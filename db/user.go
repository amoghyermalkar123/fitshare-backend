package db

import (
	"context"
	"errors"
	"fitshare/api/types"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (db *DB) AddUserDetails(userDetails *types.UserDetails) error {
	coll := db.client.Database("fitshare").Collection("user")
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userDetails.Password = string(hashedPwd)

	_, err = coll.InsertOne(context.TODO(), userDetails)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetUserDetails(userLogin *types.UserLogin) (*types.UserDetailsResponse, error) {
	coll := db.client.Database("fitshare").Collection("user")

	userResponse := &types.UserDetails{}
	err := coll.FindOne(context.TODO(), bson.M{"email": userLogin.Email}).Decode(userResponse)
	if err != nil {
		return nil, err
	}

	if !db.ValidatePassword(userLogin.Password, userResponse.Password) {
		return nil, errors.New("user not authenticated")
	}

	response := &types.UserDetailsResponse{
		Name:     userResponse.Name,
		Username: userResponse.Username,
		Email:    userResponse.Email,
		Height:   userResponse.Height,
		Weight:   userResponse.Weight,
		Age:      userResponse.Age,
		UserType: userResponse.UserType,
	}

	return response, err
}

func (db *DB) ValidatePassword(inputPassword, dbPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword)); err != nil {
		return false
	}
	return true
}
