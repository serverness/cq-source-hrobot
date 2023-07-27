package order

import (
	"context"
	"log"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/floshodan/hrobot-go/hrobot"
	"github.com/serverness/cq-source-hrobot/client"
)

func MarketServers() *schema.Table {
	return &schema.Table{
		Name:      "hrobot_order_market_servers",
		Resolver:  fetchOrderMarketServers,
		Transform: transformers.TransformWithStruct(&hrobot.ServerMarketProduct{}, transformers.WithPrimaryKeys("ID")),
		Columns:   schema.ColumnList{},
	}
}

func fetchOrderMarketServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	products, _, err := cl.Client.Order.ListMarket(ctx, &hrobot.OrderMarketListOpts{})

	if err != nil {
		log.Println(err)
		return nil
	}

	res <- products

	return nil
}
