package main

import (
	go_status "development/go/recipes/lib"
	"development/go/recipes/lib/helpers/recipes"

	log "github.com/sirupsen/logrus"
)

func RegisterFewUsers() {
	log.Info("Objective: Sign up a few new users")
	_, err := recipes.RegisterUser(go_status.DefaultEmail)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(go_status.ActiveUser)
	if err != nil {
		return
	}
}
