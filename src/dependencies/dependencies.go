package dependencies

import (
	"fmt"
	services "webserver/src/app/services"
	common "webserver/src/commom"
	config "webserver/src/config"
)

type Dependencies struct {
	db *services.DatabaseService
}

func BuildDependencies() (Dependencies, error) {
	cnfPath := common.ConfigPath
	config, err := config.LoadConfigFromFile(cnfPath)
	if err != nil {
		fmt.Println("error on config")
	}
	dbService, err := services.NewDatabaseService(config)

	if err != nil {
		return Dependencies{}, err
	}

	return Dependencies{
		db: dbService,
	}, nil
}
