package engine

import (
	"github.com/dnoberon/airlock/characters"
	"github.com/dnoberon/airlock/items"
	"github.com/dnoberon/airlock/locations"
)

// State represents the current state of the game
type State struct {
	PlayerName       string
	CurrentLocation  *locations.Location
	PreviousLocation *locations.Location

	Inventory  []*items.Item
	Characters []*characters.Character
	Locations  []*locations.Location
}

// BaseGame is to be used if user opts out of newest game release
const BaseGame = `
{
  "entry": {
    "title": "Airlock",
    "author": "John Darrington",
    "introductionText": "The galaxy is at war. Posing as an independent freighter, you and your crew of six are attempting to return home with intelligence that could turn the tide of war in your favor.\n\nSomeone aboard has purposely disabled the ship and you must now await rescue. You will run out of air before rescue arrives. Your only chance is to jettison a single crew member in order to make the air in the ship last long enough. \n\nIf you jettison the spy the intel is safe. If you jettison someone else - let's just say the war will only get bloodier. \n \nChoose wisely...",
    "initialArea": "main"
  },
  "items":[
{
  "id": "wallet",
  "name": "Leather wallet",
  "description": "The leather wallet has the initials BYB on it."
}
  ],
  "locations": [
    {
      "id": "main_room",
      "name": "Main Room",
      "entryPoint": true,
      "description": "The Main Playing Area",
      "directions": ["hallway", "-", "-", "-"],
      "characters": ["james"],
      "pointsOfInterest": [
        {
          "name": "plaque",
          "description": "The ship's information plaque: \n Built 2051 by OmesCorp"
        }
      ],
      "items":["wallet"]
    },
    {
      "id": "hallway",
      "name": "East Hallway",
      "description": "Hallway that is ugly",
      "characters": [],
      "directions": ["", "main_room"]
    }
  ],
  "characters": [
    {
      "id": "james",
      "name": "James",
      "afterDeath": "James did not sabotage the ship. He has a family at home you monster.",
      "description": "James is a ruggedly handsome asian man.",
      "conversations": [
        {
          "id": "root",
          "text": "Hello, my name is James - what's yours?",
          "entryPoint": true,
          "choices": ["say_name", "say_goodbye", "call_name"],
          "canBeRecalled": true
        },
        {
          "id": "say_name",
          "trigger": "My name is Bob",
          "text": "That's great. I had a cousin named Bob",
          "afterVisitedText": "Yeah I know, you told me.",
          "choices": ["say_goodbye", "call_name"]
        },
        {
          "id": "say_goodbye",
          "trigger": "I have to go.",
          "text": "Oh, ok.",
          "exitPoint": true
        },
        {
          "id": "call_name",
          "mustHaveItem": "wallet",
          "trigger": "You're a turd.",
          "text": "That's not very nice.",
          "choices": ["say_goodbye"]
        }
      ]
    },
    {
      "id": "robert",
      "name": "Robert",
      "afterDeath": "Robert worked hard on the ship - then worked hard on stopping the ship. After you jettisoned him you found all the evidence you needed.",
      "correctChoice": true,
      "conversations": [
        {
          "id": "root",
          "text": "Hello, my name is James - what's yours?",
          "entryPoint": true,
          "choices": ["say_name", "say_goodbye", "call_name"],
          "canBeRecalled": true
        }
      ]
    }
  ]
}
`
