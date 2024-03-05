package main

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/spf13/cobra"
)

// config defines the external configuration required for the connector to run.
type config struct {
	cli.BaseConfig `mapstructure:",squash"` // Puts the base config options in the same place as the connector options

	ClientID     string `mapstructure:"torq-client-id"`
	ClientSecret string `mapstructure:"torq-client-secret"`
}

// validateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func validateConfig(ctx context.Context, cfg *config) error {
	if cfg.ClientID == "" {
		return fmt.Errorf("torq client id is missing, please provide it via --torq-client-id flag")
	}
	if cfg.ClientSecret == "" {
		return fmt.Errorf("torq client secret is missing, please provide it via --torq-client-secret flag")
	}

	return nil
}

// cmdFlags sets the cmdFlags required for the connector.
func cmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("torq-client-id", "", "Client ID used to authenticate to the Torq API. ($BATON_TORQ_CLIENT_ID)")
	cmd.PersistentFlags().String("torq-client-secret", "", "Client Secret used to authenticate to the Torq API. ($BATON_TORQ_CLIENT_SECRET)")
}
