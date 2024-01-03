package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/graphql-go/graphql"
	_ "github.com/graphql-go/graphql"
	"log"
)

func main() {
	fmt.Println("Hello GraphQL...")

	fields := graphql.Fields{"hello": response2()}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new GraphQL Schema, err %v", err)
	}
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute GraphQL Operation, err %v", err)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}

func response2() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String, Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "World", nil
		},
	}
}
