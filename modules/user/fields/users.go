package fields

import (

	"github.com/graphql-go/graphql"
	"golearn/modules/user/types"
	"golearn/modules/user/resolvers"
)

// UsersType is the graphql users type
var UsersType = &graphql.Field{
	Type: graphql.NewList(types.UserType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.UsersResolver,
}
