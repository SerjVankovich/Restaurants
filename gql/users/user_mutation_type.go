package users

import (
	"../../db"
	"../../models"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func register(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        UserType,
		Description: "Register one user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			email, emailOk := p.Args["email"].(string)
			password, passwordOk := p.Args["password"].(string)
			name, nameOk := p.Args["name"].(string)

			if !emailOk {
				return nil, errors.New("email not provided")
			}
			if !passwordOk {
				return nil, errors.New("password not provided")
			}
			if !nameOk {
				return nil, errors.New("name not provided")
			}

			user := &models.User{Email: email, Password: password, Name: name}

			err := db.RegisterNewUser(dataBase, user)

			if err != nil {
				return nil, err
			}

			us, err := db.GetUserByEmail(dataBase, email)

			if err != nil {
				return nil, err
			}

			return us, nil
		},
	}
}

func UserMutation(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"register": register(dataBase),
			},
		})
}
