package locations

import (
	"github.com/dnoberon/airlock/characters"
)

// Location represents a physical location in the game space
type Location struct {
	ID               string // slug identifier
	Name             string
	Description      string
	EntryPoint       bool
	PointsOfInterest []PointOfInterest
	Characters       []string
	ActiveCharacters []*characters.Character
	Directions       []string

	North *Location
	South *Location
	East  *Location
	West  *Location
}

// PointOfInterest allows the look at and look around feature to work
type PointOfInterest struct {
	Name        string
	Description string
}

// InitLocations builds map of locations and characters
func InitLocations(locations []*Location, characters []*characters.Character) {
	for _, rootLocation := range locations {
		// load characters
		for _, character := range characters {
			if in(character.ID, rootLocation.Characters) {
				rootLocation.ActiveCharacters = append(rootLocation.ActiveCharacters, character)
			}
		}

		for i, direction := range rootLocation.Directions {
			for _, location := range locations {
				if direction != location.ID {
					continue
				}

				switch i {
				case 0:
					rootLocation.North = location

				case 1:
					rootLocation.South = location

				case 2:
					rootLocation.East = location

				case 3:
					rootLocation.West = location
				}
			}
		}
	}
}

func in(needle string, haystack []string) bool {
	for _, bale := range haystack {
		if bale == needle {
			return true
		}
	}

	return false
}
