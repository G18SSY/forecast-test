package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/joefitzgerald/forecast"
)

func main() {
	api := createApi("local.settings.json")
	today := time.Now().Local().Format("2006-01-02")

	user, _ := api.WhoAmI()

	projects, _ := api.Projects()

	filter := forecast.AssignmentFilter{
		PersonID:  user.ID,
		StartDate: today,
		EndDate:   today,
	}

	assignments, _ := api.AssignmentsWithFilter(filter)

	for _, assignment := range assignments {
		project := findProject(projects, assignment.ProjectID)
		fmt.Print(project.Name)

		if len(assignment.Notes) > 0 {
			fmt.Print(": ")
			fmt.Println(assignment.Notes)
		}
	}
}

func findProject(projects []forecast.Project, id int) forecast.Project {
	for _, project := range projects {
		if project.ID == id {
			return project
		}
	}

	return forecast.Project{}
}

func createApi(config string) *forecast.API {
	file, err := os.Open(config)
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	var configData map[string]string
	json.Unmarshal(bytes, &configData)

	api := forecast.New(
		"https://api.forecastapp.com",
		configData["accountId"],
		configData["accessToken"],
	)

	return api
}
