package graphql

type Schema struct {
	UsersResolver Resolver
}

func NewSchema(usersResolver Resolver) Schema {
	return Schema{
		UsersResolver: usersResolver,
	}
}
