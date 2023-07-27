package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/serverness/cq-source-hrobot/client"
	"github.com/serverness/cq-source-hrobot/resources/order"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"hrobot",
		Version,
		schema.Tables{
			order.Servers(),
			order.MarketServers(),
		},
		client.New,
	)
}
