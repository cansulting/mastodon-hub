package main

import (
	"social/backend/constants"

	"github.com/cansulting/elabox-system-tools/foundation/app"
	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

func main() {
	con, err := app.NewController(&Activity{}, nil)
	if err != nil {
		logger.GetInstance().Panic().Err(err).Msg("Failed initializing ela.social")
	}
	constants.AppController = con
	app.RunApp(con)
}
