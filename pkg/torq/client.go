package torq

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
)

const BaseUrl = "https://api.torq.io/public/v1alpha/"

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

type Client struct {
	httpClient *http.Client
	token      string
}

func NewClient(httpClient *http.Client, token string) *Client {
	return &Client{
		httpClient: httpClient,
		token:      token,
	}
}

func RequestAccessToken(ctx context.Context, clientID, clientSecret string) (string, error) {
	httpClient, err := uhttp.NewClient(ctx, uhttp.WithLogger(true, ctxzap.Extract(ctx)))
	if err != nil {
		return "", err
	}

	authUrl := "https://auth.torq.io/v1/auth/token"
	authHeader := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, authUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	var res AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res.AccessToken, nil
}

// listLocalUsers returns a list of all locally provisioned users.
func (c *Client) listLocalUsers(ctx context.Context) ([]User, error) {
	var res struct {
		Users []User `json:"users"`
	}

	usersUrl, _ := url.JoinPath(BaseUrl, "users")
	if err := c.doRequest(ctx, usersUrl, &res, nil); err != nil {
		return nil, err
	}

	return res.Users, nil
}

// ListSsoUsers returns a list of all users provisioned through SSO.
func (c *Client) listSsoUsers(ctx context.Context) ([]User, error) {
	var res struct {
		Users []User `json:"users"`
	}

	usersUrl, _ := url.JoinPath(BaseUrl, "users")
	q := url.Values{}
	q.Add("sso_provision", "true")
	if err := c.doRequest(ctx, usersUrl, &res, q); err != nil {
		return nil, err
	}

	return res.Users, nil
}

// ListUsers returns a list of all users.
func (c *Client) ListUsers(ctx context.Context) ([]User, error) {
	var allUsers []User
	localUsers, err := c.listLocalUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list local users")
	}
	allUsers = append(allUsers, localUsers...)

	ssoUsers, err := c.listSsoUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list sso users")
	}

	allUsers = append(allUsers, ssoUsers...)

	return allUsers, nil
}

// ListRoles returns a list of all roles.
func (c *Client) ListRoles(ctx context.Context, pageToken string) ([]Role, string, error) {
	var res struct {
		Roles         []Role `json:"roles"`
		NextPageToken string `json:"next_page_token,omitempty"`
	}

	q := url.Values{}
	q.Add("page_token", pageToken)

	url, _ := url.JoinPath(BaseUrl, "users/roles")
	if err := c.doRequest(ctx, url, &res, q); err != nil {
		return nil, "", err
	}

	if res.NextPageToken != "" {
		return res.Roles, res.NextPageToken, nil
	}

	return res.Roles, "", nil
}

func (c *Client) doRequest(ctx context.Context, path string, res interface{}, params url.Values) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	return nil
}
