package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dnoberon/airlock/items"

	"github.com/mitchellh/go-wordwrap"

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
	Items      []items.Item
}

var version string
var baseConfigurationSite string

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

		// Load from Web
	} else {
		resp, err := http.Get(baseConfigurationSite)
		// any errors we go straight to loading the encoded game
		if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
			err := json.Unmarshal([]byte(engine.BaseGame), &configuration)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(&configuration)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// print basic information about the game
	fmt.Printf(`%s %s 
Created By %s	

%s
`,
		configuration.Entry.Title,
		version,
		configuration.Entry.Author,
		wordwrap.WrapString(configuration.Entry.IntroductionText, 80))

	configuration.Characters = characters.InitCharacters(configuration.Characters)

	state := &engine.State{}
	var locationList []*locations.Location
	var characterList []*characters.Character
	var itemList []*items.Item

	for i := range configuration.Locations {
		locationList = append(locationList, &configuration.Locations[i])

		if configuration.Locations[i].ID == configuration.Entry.InitialArea {
			state.CurrentLocation = &configuration.Locations[i]
		}
	}

	for i := range configuration.Characters {
		characterList = append(characterList, &configuration.Characters[i])
	}

	for i := range configuration.Items {
		itemList = append(itemList, &configuration.Items[i])
	}

	locations.InitLocations(locationList, characterList, itemList)

	if state.CurrentLocation == nil {
		state.CurrentLocation = locationList[0]
	}

	fmt.Println("")
	fmt.Println(fmt.Sprintf("You are in the %s", state.CurrentLocation.Name))

	for {
		state.Decide()
		fmt.Println()
	}
}
