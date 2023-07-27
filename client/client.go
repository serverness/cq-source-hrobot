package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/floshodan/hrobot-go/hrobot"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger  zerolog.Logger
	Client  *hrobot.Client
	Backend backend.Backend

	Spec       Spec
	sourceSpec specs.Source
}

func (c *Client) ID() string {
	return c.sourceSpec.Name
}

func New(_ context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec
	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
	err := pluginSpec.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate plugin spec: %w", err)
	}
	pluginSpec.SetDefaults()

	client := hrobot.NewClient(hrobot.WithToken(os.Getenv("HETZNER_TOKEN")))

	return &Client{
		Logger:     logger,
		Backend:    opts.Backend,
		Spec:       pluginSpec,
		Client:     client,
		sourceSpec: s,
	}, nil
}
