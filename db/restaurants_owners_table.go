package db

import (
	"../models"
	"../utils"
	"encoding/hex"
	"errors"
)

func GetAllOwners(dataBase dbInterface) ([]*models.RestaurantOwner, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query("SELECT * FROM public.restaurants_owners")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var owners []*models.RestaurantOwner

	for rows.Next() {
		owner := new(models.RestaurantOwner)
		err := rows.Scan(&owner.Id, &owner.Name, &owner.Email, &owner.Password, &owner.Token, &owner.Salt, &owner.Confirmed)

		if err != nil {
			return nil, err
		}
		owners = append(owners, owner)
	}

	return owners, nil
}

func GetOwnerByEmail(dataBase dbInterface, email string) (*models.RestaurantOwner, error) {
	if dataBase == nil {
		return nil, dbErr
	}
	row := dataBase.QueryRow("SELECT * FROM restaurants_owners WHERE email = $1", email)

	owner := new(models.RestaurantOwner)

	err := row.Scan(&owner.Id, &owner.Name, &owner.Email, &owner.Password, &owner.Token, &owner.Salt, &owner.Confirmed)

	if err != nil {
		return nil, err
	}

	return owner, nil
}

func GetOwnerById(dataBase dbInterface, id int32) (*models.RestaurantOwner, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	row := dataBase.QueryRow("SELECT * FROM restaurants_owners WHERE id = $1", id)

	owner := new(models.RestaurantOwner)

	err := row.Scan(&owner.Id, &owner.Name, &owner.Email, &owner.Password, &owner.Token, &owner.Salt, &owner.Confirmed)
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func RegisterNewOwner(dataBase dbInterface, owner *models.RestaurantOwner) error {
	if dataBase == nil {
		return dbErr
	}

	if owner == nil {
		return errors.New("restaurant owner not provided")
	}

	ow, _ := GetOwnerByEmail(dataBase, owner.Email)

	if ow != nil {
		return errors.New("restaurant owner with this email exists")
	}

	salt := utils.GetRandomString(30)

	pass := hex.EncodeToString(utils.Encrypt([]byte(owner.Password), salt))

	owner.Password = pass
	owner.Salt = &salt

	_, err := dataBase.Exec("INSERT INTO restaurants_owners (id, name, email, password, salt, confirmed) values (DEFAULT, $1, $2, $3, $4, DEFAULT)",
		owner.Name,
		owner.Email,
		owner.Password,
		owner.Salt)
	if err != nil {
		return err
	}

	return nil
}

func ConfirmOwner(dataBase dbInterface, email string, token string) error {
	if dataBase == nil {
		return dbErr
	}

	owner, _ := GetOwnerByEmail(dataBase, email)

	if owner == nil {
		return errors.New("restaurant owner with this email doesn't exists")
	}

	if owner.Confirmed {
		return nil
	}

	_, err := dataBase.Exec(`UPDATE restaurants_owners SET "confirmed" = $1, "token" = $3 WHERE email = $2`, true, email, token)

	return err

}
