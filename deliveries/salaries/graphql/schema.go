package graphql

type Schema struct {
	SalaryResolver Resolver
}

func NewSchema(salaryResolver Resolver) Schema {
	return Schema{
		SalaryResolver: salaryResolver,
	}
}
