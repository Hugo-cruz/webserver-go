package dependencies

import (
	"fmt"
	services "webserver/src/app/repository"
	common "webserver/src/commom"
	config "webserver/src/config"
)

type Dependencies struct {
	DB *services.DatabaseRepository
}

func BuildDependencies() (Dependencies, error) {
	cnfPath := common.ConfigPath
	cnf, err := config.LoadConfigFromFile(cnfPath)
	if err != nil {
		fmt.Println("error on config")
	}

	db, err := services.NewDatabaseConnection(cnf.DBPath)
	if err != nil {
		return Dependencies{}, err
	}
	dataBaseRepository := services.NewDatabaseRepository(db)
	return Dependencies{
		DB: dataBaseRepository,
	}, nil
}
