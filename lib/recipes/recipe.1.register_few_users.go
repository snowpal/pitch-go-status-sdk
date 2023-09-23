package recipes

import (
	log "github.com/sirupsen/logrus"
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"
)

func RegisterFewUsers() {
	log.Info("Objective: Sign up a few new users")
	var err error
	if err = addUserAndProfile(lib.DefaultEmail, "krish12345", "krish",
		"123-234-5566"); err != nil {
		return
	}

	if err = addUserAndProfile(lib.ActiveUser, "varun12345", "varun",
		"123-234-5566"); err != nil {
		return
	}

	if err = addUserAndProfile(lib.SecondUser, "raj1234567", "raj",
		"123-234-5566"); err != nil {
		return
	}

	if err = addUserAndProfile(lib.ThirdUser, "ben1234567", "ben",
		"123-234-5566"); err != nil {
		return
	}
}

func addUserAndProfile(emailAddress, userID, firstName, phoneNumber string) error {
	user, err := recipes.RegisterUser(emailAddress)
	if err != nil {
		return err
	}

	_, err = recipes.AddNewProfile(user, userID, firstName, phoneNumber)
	if err != nil {
		return err
	}

	return nil
}
