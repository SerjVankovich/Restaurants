package db

import (
	"../models"
	"../utils"
	"database/sql"
	"encoding/hex"
	"errors"
)

var dbErr = errors.New("dataBase argument wasn't provided")

func GetAllUsers(dataBase *sql.DB) ([]*models.User, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query("SELECT * FROM public.users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.PrefRest, &user.Token, &user.Salt)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByEmail(dataBase *sql.DB, email string) (*models.User, error) {
	if dataBase == nil {
		return nil, dbErr
	}
	row := dataBase.QueryRow("SELECT * FROM users WHERE email = $1", email)

	user := new(models.User)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.PrefRest, &user.Token, &user.Salt, &user.Confirmed)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserById(dataBase *sql.DB, id int32) (*models.User, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	row := dataBase.QueryRow("SELECT * FROM users WHERE id = $1", id)

	user := new(models.User)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.PrefRest, &user.Token, &user.Salt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func RegisterNewUser(dataBase *sql.DB, user *models.User) error {
	if dataBase == nil {
		return dbErr
	}

	if user == nil {
		return errors.New("user not provided")
	}

	us, _ := GetUserByEmail(dataBase, user.Email)

	if us != nil {
		return errors.New("user with this email exists")
	}

	salt := utils.GetRandomString(30)

	pass := hex.EncodeToString(utils.Encrypt([]byte(user.Password), salt))

	user.Password = pass
	user.Salt = salt

	_, err := dataBase.Exec("INSERT INTO users (id, name, email, password, salt, confirmed) values (DEFAULT, $1, $2, $3, $4, DEFAULT)",
		user.Name,
		user.Email,
		user.Password,
		user.Salt)
	if err != nil {
		return err
	}

	return nil
}

func ConfirmUser(dataBase *sql.DB, email string) error {
	if dataBase == nil {
		return dbErr
	}

	user, _ := GetUserByEmail(dataBase, email)

	if user == nil {
		return errors.New("user with this email doesn't exists")
	}

	if user.Confirmed {
		return nil
	}

	_, err := dataBase.Exec(`UPDATE users SET "confirmed" = $1 WHERE email = $2`, true, email)

	return err

}
