package connector

import (
	"context"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/conductorone/baton-torq/pkg/torq"
)

type userBuilder struct {
	resourceType *v2.ResourceType
	client       *torq.Client
}

func (u *userBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return u.resourceType
}

// Create a new connector resource for a Torq user.
func userResource(user *torq.User) (*v2.Resource, error) {
	profile := map[string]interface{}{
		// there is no name available from the API response.
		"first_name":      user.Email,
		"login":           user.Email,
		"user_id":         user.ID,
		"sso_provisioned": user.SsoProvision,
	}

	var status v2.UserTrait_Status_Status
	switch user.Status {
	case "VERIFIED":
		status = v2.UserTrait_Status_STATUS_ENABLED
	case "STATUS_UNSPECIFIED":
		status = v2.UserTrait_Status_STATUS_UNSPECIFIED
	default:
		// there is no state for 'INVITATION_SENT' status.
		status = v2.UserTrait_Status_STATUS_UNSPECIFIED
	}

	userTraitOptions := []rs.UserTraitOption{
		rs.WithUserProfile(profile),
		rs.WithEmail(user.Email, true),
		rs.WithStatus(status),
	}

	ret, err := rs.NewUserResource(
		user.Email,
		userResourceType,
		user.ID,
		userTraitOptions,
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// List returns all the users from the database as resource objects.
// Users include a UserTrait because they are the 'shape' of a standard user.
func (u *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	users, err := u.client.ListUsers(ctx)
	if err != nil {
		return nil, "", nil, err
	}

	var rv []*v2.Resource
	for _, user := range users {
		userCopy := user
		ur, err := userResource(&userCopy)
		if err != nil {
			return nil, "", nil, err
		}
		rv = append(rv, ur)
	}

	return rv, "", nil, nil
}

// Entitlements always returns an empty slice for users.
func (u *userBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

// Grants always returns an empty slice for users since they don't have any entitlements.
func (u *userBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func newUserBuilder(client *torq.Client) *userBuilder {
	return &userBuilder{
		resourceType: userResourceType,
		client:       client,
	}
}
