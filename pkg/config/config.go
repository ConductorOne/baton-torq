package config

import (
	"fmt"

	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	TorqClientId = field.StringField(
		"torq-client-id",
		field.WithDescription("Client ID used to authenticate to the Torq API."),
		field.WithRequired(true),
		field.WithDisplayName("Client ID"),
	)

	TorqClientSecret = field.StringField(
		"torq-client-secret",
		field.WithDescription("Client Secret used to authenticate to the Torq API."),
		field.WithRequired(true),
		field.WithIsSecret(true),
		field.WithDisplayName("Client Secret"),
	)

	BaseURLField = field.StringField(
		"base-url",
		field.WithDescription("Override the Torq API URL (for testing)"),
		field.WithHidden(true),
		field.WithExportTarget(field.ExportTargetCLIOnly),
	)

	// FieldRelationships defines relationships between the fields listed in
	// Config that can be automatically validated.
	FieldRelationships = []field.SchemaFieldRelationship{}
)

//go:generate go run ./gen
var Config = field.NewConfiguration([]field.SchemaField{
	TorqClientId,
	TorqClientSecret,
	BaseURLField,
})

// ValidateConfig is run after the configuration is loaded, and should return an
// error if it isn't valid. Implementing this function is optional, it only
// needs to perform extra validations that cannot be encoded with configuration
// parameters.
func ValidateConfig(cfg *Torq) error {
	if cfg.TorqClientId == "" {
		return fmt.Errorf("torq client id is missing, please provide it via --torq-client-id flag")
	}
	if cfg.TorqClientSecret == "" {
		return fmt.Errorf("torq client secret is missing, please provide it via --torq-client-secret flag")
	}
	return nil
}
