package main

import (
	"github.com/snowpal/pitch-go-status-sdk/lib/config"
	"github.com/snowpal/pitch-go-status-sdk/lib/recipes"

	log "github.com/sirupsen/logrus"
	teamstatuses "github.com/snowpal/pitch-go-status-sdk/lib/recipes/recipe.2.report_team_statuses"
)

func main() {
	var err error
	if config.Files, err = config.InitConfigFiles(); err != nil {
		log.Fatal(err.Error())
		return
	}

	// We require that the first recipe be run before anything else as it registers a bunch of
	// users.
	recipeID := 1
	switch recipeID {
	case 1:
		log.Info("Run Recipe1")
		recipes.RegisterFewUsers()
	case 2:
		log.Info("Run Recipe2")
		teamstatuses.ReportTeamStatuses()
	default:
		log.Info("pick a specific recipe to run")
	}
}
