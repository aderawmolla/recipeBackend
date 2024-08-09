package services

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
)

var client *graphql.Client

func InitGraphQLClient(url string) {
	client = graphql.NewClient(url)
}

func RunGraphQLQuery(ctx context.Context, req *graphql.Request, respData interface{}) error {
	err := client.Run(ctx, req, respData)
	if err != nil {
		log.Println("GraphQL query failed:", err)
		return err
	}
	return nil
}
