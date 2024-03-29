package recipes

import (
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/profiles"
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/registration"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/common"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func RegisterUser(email string) (response.User, error) {
	signUpReqBody := request.SignupReqBody{
		Email:           email,
		Password:        lib.Password,
		ConfirmPassword: lib.Password,
	}
	user, err := registration.RegisterNewUserByEmail(signUpReqBody)
	if err != nil {
		return response.User{}, err
	}

	log.Info(".activate user ID: ", user.ID)
	err = registration.ActivateUser(user.ID)
	if err != nil {
		return response.User{}, err
	}
	return user, nil
}

func SignIn(email string, password string) (response.User, error) {
	log.Info("Sign in user with email: ", email)
	SleepBefore()
	user, err := registration.SignInByEmail(request.SignInReqBody{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return response.User{}, err
	}
	log.Info(".User successfully signed in, acquired JWT token")
	SleepAfter()
	return user, nil
}

func AddNewProfile(user response.User, userID string, firstName string, phoneNumber string) (response.Profile, error) {
	log.Info("Adding profile for ", user.Email)
	var resProfile response.Profile

	profileReqBody := request.ProfileReqBody{
		UserID: userID,
		Name: common.ProfileName{
			FirstName: firstName,
		},
		PhoneNumbers: []common.PhoneNumber{
			{Primary: true, PhoneNumber: phoneNumber},
		},
		EmailAddresses: []common.EmailAddress{
			{Primary: true, EmailAddress: user.Email},
		},
	}
	resProfile, err := profiles.AddProfile(user.JwtToken, profileReqBody)
	if err != nil {
		return resProfile, err
	}
	log.Info(".Profile successfully created for ", user.Email)

	return resProfile, nil
}
