package modules

import (
	user "golearn/modules/user/fields"

	"github.com/graphql-go/graphql"
)

// MutationType is the root mutation
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": user.CreateUserType,
	},
})
