package recipes

import (
	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers/recipes"

	log "github.com/sirupsen/logrus"
)

func RegisterFewUsers() {
	log.Info("Objective: Sign up a few new users")
	_, err := recipes.RegisterUser(lib.DefaultEmail)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(lib.ActiveUser)
	if err != nil {
		return
	}
}
