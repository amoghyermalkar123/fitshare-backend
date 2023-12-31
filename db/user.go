package db

import (
	"context"
	"errors"
	"fitshare/api/types"
	dbtypes "fitshare/db/dbTypes"
	"time"

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

func (db *DB) GetGymEventsAndPeople(username string) (*types.UserHomePage, error) {
	memColl := db.client.Database("fitshare").Collection("gym_members")
	gymMembersDetails := &dbtypes.DbGymMembers{}

	err := memColl.FindOne(context.TODO(), bson.D{{"members", username}}).Decode(gymMembersDetails)
	if err != nil {
		return nil, err
	}

	coll := db.client.Database("fitshare").Collection("gym_schedule")

	filter := bson.M{
		"gym_id": gymMembersDetails.GymID,
	}

	gymSchedule := &dbtypes.GymWeeklySchedule{}
	err = coll.FindOne(context.Background(), filter).Decode(gymSchedule)
	if err != nil {
		return nil, err
	}

	homeFeed := &types.UserHomePage{}
	for _, user := range gymMembersDetails.MemberUsernames {
		if user == username {
			continue
		}
		homeFeed.DiscoverPeople = append(homeFeed.DiscoverPeople, types.DiscoverPeople{
			Username: user,
		})
	}

	for _, event := range gymSchedule.Schedule {
		if isDateToday(event.DateTime) {
			homeFeed.TodaysEvent = append(homeFeed.TodaysEvent, types.TodaysEvent{
				EventName: event.EventName,
				EventTime: event.DateTime,
			})
		}
	}

	homeFeed.UserName = username
	return homeFeed, nil
}

func isDateToday(dateToCheck time.Time) bool {
	// Get the current date
	currentDate := time.Now().Truncate(24 * time.Hour)

	// Truncate the provided date to remove the time part
	dateToCheck = dateToCheck.Truncate(24 * time.Hour)

	// Compare the two dates
	return dateToCheck.Equal(currentDate)
}
