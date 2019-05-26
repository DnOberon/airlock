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
    "introductionText": "Mankind has colonized our the solar system yet still stands divided. The Inner Planets, consisting of Earth and Mars, have continually oppressed and held sway over The Colonies on Jupiter's moon and surrounding space.\n\nNeeding Earth and Mars for supplies of food and medicine The Colonies dare not openly defy them. Needing the raw minerals and hydrogen mined around Jupiter, The Inner Planets must rule space with an iron fist or risk becoming helpless and defenseless when they are unable to manufacture ships and weapons. \n\nAboard the independent space freighter Gently Weeping, you and your crew of five have managed to avoid the majority of the conflict. Recently however, you stumbled across information regarding a future attack by The Colonies on one of The Inner Planet's most lucrative, and most populated space stations. \n\nFinding the information has left your ship stranded in space and operating on emergency power. You've contacted help and a rescue ship is on the way - but you only have enough oxygen for five people. \n\nIn order to survive, and act on the intelligence you've gathered, you must jettison one of your crew into space. \n\nChoose wisely - the fate of the solar system is in your hands.",
    "initialArea": "main"
  },
  "items": [],
  "locations": [
    {
      "id": "captains_quarters",
      "name": "Captain's Quarters",
      "entryPoint": true,
      "description": "This is your cabin, and it has seen better days. Off to the left is your bunk, covered in clothes and data pads. To your right sits your desk - which is also a disaster. A single yellow light bar runs through the room lengthwise and gives off only a faint yellow light. On your desk sits a single picture and what looks like a newspaper.",
      "directions": ["_", "-", "north_hallway", "-"],
      "characters": [""],
      "pointsOfInterest": [
        {
          "name": "picture",
          "description": "A picture of your cat. Unfortunately cats and space travel don't mix so you had to leave her home."
        },
        {
          "name": "handbook",
          "description": "This handbook contains all the laws independent freighters have to follow. It looks like it has never been opened."
        },
        {
          "name": "newspaper",
          "description": "It may be 2150, but newspapers have refused to disappear. The headline reads \"Earth Raises Hydrogen Tariffs\""
        }
      ]
    },
    {
      "id": "north_hallway",
      "name": "North Hallway",
      "description": "The metal floor of the hallway has a high shine from the frequent passage of boots. The walls were white, once, but now are a non-uniform shade of grey. There are dark smears of what we hope is some kind of grease",
      "directions": ["cockpit", "hallway", "airlock", "captains_quarters"]
    },
    {
      "id": "airlock",
      "name": "Airlock",
      "description": "The only part of the ship that's still mostly white. There is a transparent steel wall in the middle of the room with a large bulkhead door set in the middle. Red warning lines flank the door on either end and on the wall to your right you see a button, also flanked with warning lines. The airlock use to be inconspicuous, now it is a cloying presence in the ship - a constant reminder of the choice you must make. \n\nYou must come back to this room in order to use the \"jettison\" command and end the game.",
      "directions": ["-", "-", "north_hallway", "-"]
    },
    {
      "id": "cockpit",
      "name": "Cockpit",
      "description": "Movies make spaceship's cockpit appear spacious and well kept. The cockpit of your ship however, is cramped and a soft cacophony of beeps, clicks, and small fans. There is no light except what the switches and readouts display. The cockpit slopes forward, metal molding into transparent steel which continues towards the ground and out of your field of vision.",
      "directions": ["-", "north_hallway", "-", "-"],
      "pointsOfInterest": [
        {
          "name": "instrument panels",
          "description": "The panels are a mix of buttons, switches, readouts, and the occasional food wrapper. You didn't go to school for this, so you understand none of it. Except the candy bar wrappers. \nYou are now hungry."
        },
        {
          "name": "forward display",
          "description": "The transparent steel display shows nothing but an inky void with faint bits of sparkling light. You're a long way from home."
        }
      ]
    },
    {
      "id": "hallway",
      "name": "Hallway",
      "description": "This part of the hallway is as dirty as the others - except for the door to the first mate's quarters which is a spotless white.",
      "directions": [
        "north_hallway",
        "south_hallway",
        "-",
        "first_mates_quarters"
      ],
      "pointsOfInterest": [
        {
          "name": "plaque",
          "description": "The ship's identification plaque. The only legible part of fading bronze plaque reads: \n\nGeently Weeps\n\nBuilt 2112 by OmesCorp"
        }
      ]
    },
    {
      "id": "first_mates_quarters",
      "name": "First Mate's Quarters",
      "description": "The room is a pristine white and extremely clean - in stark contrast to your own cabin. There are no personal affects that you can see, you assume whatever there might be is stored in the lockers beneath their bunk.",
      "pointsOfInterest": [
        {
          "name": "handbook",
          "description": "This handbook contains all the laws independent freighters have to follow. This copy is well used, dog eared pages and bookmarks stick out everywhere."
        }
      ],
      "directions": ["-", "-", "hallway", "-"]
    },
    {
      "id": "south_hallway",
      "name": "South Hallway",
      "description": "There is nothing out of the ordinary about this hallway.",
      "directions": ["hallway", "machine_shop", "common_area", "cargo_hold"]
    },
    {
      "id": "common_area",
      "name": "The Common Area",
      "description": "Next to the cargo hold, this is the largest area in the ship. A large table takes up one half of the room and is bolted the floor. What can only be described as a \"space garage sale\" fills the other half. Random placements of furniture and personal items strewn about complete the effect. There are a few doors around the room, each leading to personal or shared quarters. Soft emergency lighting lends the room an abandoned feel.",
      "pointsOfInterest": [
        {
          "name": "tv",
          "description": "TVs haven't changed much. This one is powered off and you can see your reflection in it. Not that there is enough power in the ship to run it anyway."
        },
        {
          "name": "paintings",
          "description": "On one wall sit large paintings in-between doors. Each painting depicts the same scene but in different styles. The scene is of a young girl holding a wilting flower. In the background there is a market full of people. The painting is set far before space travel was obtained."
        },
        {
          "name": "large table",
          "description": "placeholder describing table and the two characters that will be sitting at it."
        }
      ],
      "directions": ["-", "shared_quarters", "-", "south_hallway"]
    },
    {
      "id": "cargo_hold",
      "name": "Cargo Hold",
      "description": "This space is large enough that you can hear your steps echo. There are a few boxes along the walls, most holding spare parts or food for your ship. You wish you could jettison what cargo you had rather than a crew member. Unfortunately cargo doesn't breathe.",
      "pointsOfInterest": [
        {
          "name": "crates",
          "description": "Inspecting the crates more closely you see that the majority of them hold supplies for your ship. Food and water make up the bulk of these supplies with whatever room left being packed with spare parts."
        },
        {
          "name": "loader",
          "description": "The electric loader is basically a forklift and run on ship's power via a very long extension cable. It is strapped to the side of the hold."
        }
      ],
      "directions": ["-", "-", "south_hallway", "-"]
    },
    {
      "id": "machine_shop",
      "name": "Machine Shop",
      "description": "Parts, tools, and electronics cover almost every surface of this room. Hung on the wall next to the door are a few hardhats, shiny and clean from disuse. A winch is embedded in the ceiling and is currently holding what can only be described as a giant ball of electronic yarn.",
      "directions": ["south_hallway", "engine_room", "-", "-"]
    },
    {
      "id": "engine_room",
      "name": "Engine Room",
      "description": "There isn't much room here. A large, inverted pyramid of pitch black metal hands suspended from the ceiling. The tip disappearing into the floor. There are no tools or electronics in the area. Light is provided by a bank of candles and mirrors to either side of the door you just entered.",
      "pointsOfInterest": [
        {
          "name":"pyramid",
          "description": "The pyramid is twice your height and could easily fit three or full grown adults inside. It is pitch black and it seems to suck in the light given off by the candles. There is no hum or any indication that the engine is doing anything. This could be normal, you don't know. You don't come down here often."
        }
      ],
      "directions": ["machine_shop", "-", "-", "-"]
    }
  ]
}`
