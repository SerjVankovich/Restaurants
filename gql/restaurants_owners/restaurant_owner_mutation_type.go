package restaurants_owners

import (
	"../../db"
	"../../models"
	"../../utils"
	"../users"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"os"
)

func register(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        users.HashType,
		Description: "Register one restaurant owner",
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

			owner := &models.RestaurantOwner{Email: email, Password: password, Name: name}

			err := db.RegisterNewOwner(dataBase, owner)

			if err != nil {
				return nil, err
			}

			path, _ := os.Getwd()

			hmac := utils.ParseHMACSecret(path + "\\keys.json")

			hash := hex.EncodeToString(utils.Encrypt([]byte(owner.Email), hmac))

			return &models.Hash{Hash: hash}, nil
		},
	}
}

func login(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: users.ConfirmedType,
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve:     loginResolver(dataBase),
		Description: "Restaurant owner login",
	}
}

func loginResolver(dataBase *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		email, emailOk := p.Args["email"].(string)
		password, passwordOk := p.Args["password"].(string)

		if !emailOk {
			return nil, errors.New("email not provided")
		}

		if !passwordOk {
			return nil, errors.New("password not provided")
		}

		owner, err := db.GetOwnerByEmail(dataBase, email)

		if err != nil {
			return nil, err
		}

		bytePass, err := hex.DecodeString(owner.Password)
		ownerPassword := string(utils.Decrypt(bytePass, *owner.Salt))

		if ownerPassword != password {
			return nil, errors.New("authentication failed: wrong password")
		}

		path, _ := os.Getwd()

		conf, isConfirm := CheckOwnerConfirm(*owner)

		if !isConfirm {
			return conf, nil
		}

		jwtSecret := utils.ParseJwtSecret(path + "\\keys.json")

		token, err := utils.CreateToken(jwtSecret, owner.Email, "owner")

		if err != nil {
			return nil, err
		}

		regConfirm := models.Confirm{
			IsOk:        true,
			AccessToken: token,
			ConfirmHash: "",
		}

		return regConfirm, nil
	}
}

func CheckOwnerConfirm(owner models.RestaurantOwner) (*models.Confirm, bool) {
	path, _ := os.Getwd()

	if !owner.Confirmed {
		hmac := utils.ParseHMACSecret(path + "\\keys.json")

		hash := hex.EncodeToString(utils.Encrypt([]byte(owner.Email), hmac))
		return &models.Confirm{
			IsOk:        false,
			AccessToken: "",
			ConfirmHash: hash,
		}, false
	}

	return nil, true
}

func confirmRegister(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        users.ConfirmedType,
		Description: "Confirm restaurant owner registration",
		Args: graphql.FieldConfigArgument{
			"hash": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: confirmRegisterResolver(dataBase),
	}
}

func confirmRegisterResolver(dataBase *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		hash, hashOk := p.Args["hash"].(string)

		if !hashOk {
			return nil, errors.New("hash not provided")
		}

		path, _ := os.Getwd()

		hmac := utils.ParseHMACSecret(path + "\\keys.json")

		byteValue, err := hex.DecodeString(hash)

		if err != nil {
			return nil, err
		}
		email := string(utils.Decrypt(byteValue, hmac))

		var regConfirm models.Confirm

		jwtSecret := utils.ParseJwtSecret(path + "\\keys.json")

		token, err := utils.CreateToken(jwtSecret, email, "owner")

		if err != nil {
			return nil, err
		}

		fmt.Println(token)
		err = db.ConfirmOwner(dataBase, email, token)

		if err != nil {
			return nil, err
		}

		regConfirm.AccessToken = token
		regConfirm.IsOk = true

		return &regConfirm, nil
	}
}

func RestaurantOwnerMutation(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"register":        register(dataBase),
				"login":           login(dataBase),
				"confirmRegister": confirmRegister(dataBase),
			},
		})
}
