package connector

import (
	"context"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"github.com/conductorone/baton-torq/pkg/torq"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
)

type Connector struct {
	client *torq.Client
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (c *Connector) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(c.client),
		newRoleBuilder(c.client),
	}
}

// Metadata returns metadata about the connector.
func (c *Connector) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "Torq Connector",
		Description: "Connector sycing users and roles from Torq to Baton.",
	}, nil
}

// Validate is called to ensure that the connector is properly configured. It should exercise any API credentials
// to be sure that they are valid.
func (c *Connector) Validate(ctx context.Context) (annotations.Annotations, error) {
	_, err := c.client.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error validating Torq connector: %w", err)
	}
	return nil, nil
}

// New returns a new instance of the connector.
func New(ctx context.Context, clientId string, clientSecret string) (*Connector, error) {
	httpClient, err := uhttp.NewClient(ctx, uhttp.WithLogger(true, ctxzap.Extract(ctx)))
	if err != nil {
		return nil, err
	}

	token, err := torq.RequestAccessToken(ctx, clientId, clientSecret)
	if err != nil {
		return nil, fmt.Errorf("torq-connector: failed to get token: %w", err)
	}

	return &Connector{
		client: torq.NewClient(httpClient, token),
	}, nil
}
