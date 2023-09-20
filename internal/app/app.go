package app

import (
	"discord-oauth2/internal/app/database"
	"discord-oauth2/internal/app/server"
)

func Run() error {
	var err error

	err = database.Connect()
	if err != nil {
		return err
	}

	err = database.Migrate()
	if err != nil {
		return err
	}

	err = server.Start()
	if err != nil {
		return err
	}

	return nil
}
