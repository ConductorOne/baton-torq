package main

import (
	cfg "github.com/conductorone/baton-torq/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("torq", cfg.Config)
}
