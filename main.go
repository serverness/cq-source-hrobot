package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/serverness/cq-source-hrobot/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
