package types

import (
	"github.com/graphql-go/graphql"

	user "golearn/modules/user/fields"
	userType "golearn/modules/user/types"
)

// ViewerType is the graphql viewer type
var ViewerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Viewer",
	Fields: graphql.Fields{
		"users": user.UsersType,
	},
})

// Viewer is the viewer type
type Viewer struct {
	Users []userType.User `json:"users"`
}
