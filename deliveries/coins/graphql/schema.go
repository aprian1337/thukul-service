package graphql

type Schema struct {
	CoinsResolver Resolver
}

func NewSchema(resolver Resolver) Schema {
	return Schema{
		CoinsResolver: resolver,
	}
}
