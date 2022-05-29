package graphql

type Schema struct {
	PaymentResolver Resolver
}

func NewSchema(paymentResolver Resolver) Schema {
	return Schema{
		PaymentResolver: paymentResolver,
	}
}
