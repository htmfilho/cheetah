package team

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Repository interface {
	GetCurrentTeam() Team
}

type JsonRepository struct {
	File string
}

func (jsonRepo *JsonRepository) GetCurrentTeam() Team {
	jsonFile, err := ioutil.ReadFile(jsonRepo.File)
	if err != nil {
		log.Fatalf("Error reading the file %s: %v", jsonRepo.File, err)
	}

	var team Team
	err = json.Unmarshal(jsonFile, &team)
	if err != nil {
		log.Fatalf("Error unmarshaling the json content: %v", err)
	}

	return team
}

func GetRepository() Repository {
	return &JsonRepository{}
}
