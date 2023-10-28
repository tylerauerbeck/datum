package schema

import (
	"entgo.io/ent/schema/mixin"
)

// AuthMixin for all schemas in the graph.
type AuthMixin struct {
	mixin.Schema
}

// // Policy defines the privacy policy of the AuthMixin.
// func (AuthMixin) Policy() ent.Policy {
// 	return privacy.Policy{
// 		Query: privacy.QueryPolicy{
// 			// Deny any query operation in case
// 			// there is no "viewer context".
// 			rule.DenyIfNoViewer(),
// 			// Allow admins to query any information.
// 			rule.AllowIfAdmin(),
// 		},
// 		Mutation: privacy.MutationPolicy{
// 			// Deny any mutation operation in case
// 			// there is no "viewer context".
// 			rule.DenyIfNoViewer(),
// 		},
// 	}
// }
