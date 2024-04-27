package dependencies

import (
	services "webserver/src/app/services"
	common "webserver/src/commom"
)

func buildDependencies() {
	cnfPath := common.ConfigPath
	config, err := LoadConfigFromFile(cnfPath)
	if err != nil {
		panic("error on config")
	}
	dbService := services.NewDatabaseService(config)
}
