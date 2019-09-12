package users

import (
	"../../db"
	"../../models"
	"../../utils"
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/graphql-go/graphql"
	"os"
)

func register(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type:        HashType,
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

			path, _ := os.Getwd()

			hmac := utils.ParseHMACSecret(path + "\\keys.json")

			hash := hex.EncodeToString(utils.Encrypt([]byte(user.Email), hmac))

			return &models.Hash{Hash: hash}, nil
		},
	}
}

func login(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: ConfirmedType,
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve:     loginResolver(dataBase),
		Description: "User login",
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

		user, err := db.GetUserByEmail(dataBase, email)

		if err != nil {
			return nil, err
		}

		bytePass, err := hex.DecodeString(user.Password)
		userPassword := string(utils.Decrypt(bytePass, user.Salt))

		if userPassword != password {
			return nil, errors.New("authentication failed: wrong password")
		}

		path, _ := os.Getwd()

		conf, isConfirm := CheckUserConfirm(*user)

		if !isConfirm {
			return conf, nil
		}

		jwtSecret := utils.ParseJwtSecret(path + "\\keys.json")

		token, err := utils.CreateToken(jwtSecret, user.Email, "user")

		if err != nil {
			return nil, err
		}

		regConfirm := models.Confirm{
			IsOk:        true,
			AccessToken: token,
			ConfirmHash: "",
		}

		err = db.ConfirmUser(dataBase, email, token)

		if err != nil {
			return nil, err
		}

		return regConfirm, nil
	}
}

func CheckUserConfirm(user models.User) (*models.Confirm, bool) {
	path, _ := os.Getwd()

	if !user.Confirmed {
		hmac := utils.ParseHMACSecret(path + "\\keys.json")

		hash := hex.EncodeToString(utils.Encrypt([]byte(user.Email), hmac))
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
		Type:        ConfirmedType,
		Description: "Confirm user registration",
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

		token, err := utils.CreateToken(jwtSecret, email, "user")

		if err != nil {
			return nil, err
		}

		err = db.ConfirmUser(dataBase, email, token)

		if err != nil {
			return nil, err
		}

		regConfirm.AccessToken = token
		regConfirm.IsOk = true

		return &regConfirm, nil
	}
}

func UserMutation(dataBase *sql.DB) *graphql.Object {
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
