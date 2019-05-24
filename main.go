package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dnoberon/airlock/characters"
	"github.com/dnoberon/airlock/engine"
	"github.com/dnoberon/airlock/locations"
)

// Configuration is a placeholder
type Configuration struct {
	Entry struct {
		Title            string
		Author           string
		IntroductionText string
		InitialArea      string
	} `json:"entry"`

	Characters []characters.Character
	Locations  []locations.Location
}

var version string

func main() {
	var configuration Configuration

	_, err := os.Stat("game.json")
	if !os.IsNotExist(err) {
		file, err := os.Open("game.json")
		if err != nil {
			log.Fatal(err)
		}

		err = json.NewDecoder(file).Decode(&configuration)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := json.Unmarshal([]byte(engine.BaseGame), &configuration)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf(`%s version %s 
Created By %s	
%s
`, configuration.Entry.Title, version, configuration.Entry.Author, configuration.Entry.IntroductionText)

	configuration.Characters = characters.InitCharacters(configuration.Characters)

	state := &engine.State{}
	var l []*locations.Location
	var c []*characters.Character

	for i := range configuration.Locations {
		l = append(l, &configuration.Locations[i])

		if configuration.Locations[i].ID == configuration.Entry.InitialArea {
			state.CurrentLocation = &configuration.Locations[i]
		}
	}

	for i := range configuration.Characters {
		c = append(c, &configuration.Characters[i])
	}

	locations.InitLocations(l, c)

	if state.CurrentLocation == nil {
		state.CurrentLocation = l[0]
	}

	fmt.Println("")
	fmt.Println(fmt.Sprintf("You are in the %s", state.CurrentLocation.Name))

	for {
		state.Decide()
		fmt.Println()
	}
}
