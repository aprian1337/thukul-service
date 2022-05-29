package gql

import (
	_coinsSchema "aprian1337/thukul-service/deliveries/coins/graphql"
	_paymentsSchema "aprian1337/thukul-service/deliveries/payments/graphql"
	_salariesSchema "aprian1337/thukul-service/deliveries/salaries/graphql"
	_usersSchema "aprian1337/thukul-service/deliveries/users/graphql"
	"github.com/graphql-go/graphql"
)

type Schema struct {
	SalarySchema  _salariesSchema.Schema
	CoinsSchema   _coinsSchema.Schema
	PaymentSchema _paymentsSchema.Schema
	UsersSchema  _usersSchema.Schema
}

var LoginGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Login",
		Fields: graphql.Fields{
			"session_token": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: UsersGraphQL,
			},
		},
	})

var CoinsGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Coins",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"symbol": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var SalaryGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Salary",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"maximal": &graphql.Field{
				Type: graphql.Int,
			},
			"minimal": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var WalletsGraphql = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Wallets",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"total": &graphql.Field{
				Type: graphql.Float,
			},
		},
	})

var UsersWithChildGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UsersWithChild",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"salary": &graphql.Field{
				Type: SalaryGraphQL,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"is_admin": &graphql.Field{
				Type: graphql.Int,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"birthday": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"company": &graphql.Field{
				Type: graphql.String,
			},
			"is_valid": &graphql.Field{
				Type: graphql.Int,
			},
			"wallets": &graphql.Field{
				Type: WalletsGraphql,
			},
		},
	})

var UsersGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Users",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"is_admin": &graphql.Field{
				Type: graphql.Int,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"birthday": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"company": &graphql.Field{
				Type: graphql.String,
			},
			"is_valid": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

var UsersRegisterGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UsersReegister",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"salary_id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"is_admin": &graphql.Field{
				Type: graphql.Int,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"birthday": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"company": &graphql.Field{
				Type: graphql.String,
			},
			"is_valid": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

var BuySellGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BuySellCoin",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var TopUpGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TopUpCoin",
		Fields: graphql.Fields{
			"total": &graphql.Field{
				Type: graphql.Float,
			},
			"nominal_transaction": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"GetCoinBySymbol": &graphql.Field{
				Name: "GetCoinsBySymbol",
				Type:        CoinsGraphQL,
				Resolve:     s.CoinsSchema.CoinsResolver.GetBySymbol,
				Description: "Get Coin By Symbol",
				Args: graphql.FieldConfigArgument{
					"symbol": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"GetCoinsAll": &graphql.Field{
				Name: "GetAllCoins",
				Type:        graphql.NewList(CoinsGraphQL),
				Resolve:     s.CoinsSchema.CoinsResolver.GetAllSymbol,
				Description: "Get All Coins",
			},
			"GetAllSalary": &graphql.Field{
				Name: "GetAllSalary",
				Type:        graphql.NewList(SalaryGraphQL),
				Resolve:     s.SalarySchema.SalaryResolver.GetSalaries,
				Description: "Get All Salary",
			},
			"GetAllUsers": &graphql.Field{
				Name:              "GetAllUsers",
				Type:              graphql.NewList(UsersWithChildGraphQL),
				Resolve:           s.UsersSchema.UsersResolver.GetUsers,
				Description:       "Get All Users",
			},
			"GetUserById": &graphql.Field{
				Name:              "GetUserById",
				Type:              UsersWithChildGraphQL,
				Resolve:           s.UsersSchema.UsersResolver.GetUserDetail,
				Description:       "Get User By Id",
			},
		},
	}
	return graphql.NewObject(objectConfig)
}

func (s Schema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: graphql.Fields{
			"TopUpCoin": &graphql.Field{
				Type:        TopUpGraphQL,
				Resolve:     s.PaymentSchema.PaymentResolver.TopUpCoin,
				Description: "Top Up Coin",
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"nominal": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
			},
			"BuyCoin": &graphql.Field{
				Type:        BuySellGraphQL,
				Resolve:     s.PaymentSchema.PaymentResolver.BuyCoin,
				Description: "Buy Coin",
				Args: graphql.FieldConfigArgument{
					"symbol": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"qty": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
			},
			"SellCoin": &graphql.Field{
				Type:        BuySellGraphQL,
				Resolve:     s.PaymentSchema.PaymentResolver.SellCoin,
				Description: "Sell Coin",
				Args: graphql.FieldConfigArgument{
					"symbol": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"qty": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
			},
			"LoginUser": &graphql.Field{
				Type: LoginGraphQL,
				Resolve: s.UsersSchema.UsersResolver.LoginUser,
				Description: "Login User",
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"RegisterUser": &graphql.Field{
				Type: UsersRegisterGraphQL,
				Resolve: s.UsersSchema.UsersResolver.RegisterUser,
				Description: "Register User",
				Args: graphql.FieldConfigArgument{
					"salaryId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gender": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"birthday": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"company": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
		},
	}
	return graphql.NewObject(objectConfig)
}
