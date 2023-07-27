package order

import (
	"context"
	"log"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/floshodan/hrobot-go/hrobot"
	"github.com/serverness/cq-source-hrobot/client"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:      "hrobot_order_servers",
		Resolver:  fetchOrderServers,
		Transform: transformers.TransformWithStruct(&hrobot.ServerProduct{}, transformers.WithPrimaryKeys("ID")),
		Columns:   schema.ColumnList{},
	}
}

func fetchOrderServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	products, _, err := cl.Client.Order.ProductList(ctx, &hrobot.OrderServerListOpts{})

	if err != nil {
		log.Println(err)
		return nil
	}

	res <- products

	return nil
}
