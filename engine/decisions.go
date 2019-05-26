package engine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dnoberon/airlock/items"

	"github.com/dnoberon/airlock/convoengine"
	"github.com/mitchellh/go-wordwrap"
)

var commands map[string]func(state *State, arguments ...string)

var articles = []string{"and", "to", "a"}

// Decide drives command line interaction
func (s *State) Decide() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")

	input, err := buf.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.ToLower(strings.TrimSpace(strings.TrimSuffix(input, "\n")))

	parsed := strings.Split(input, " ")

	// strip out articles
	i := 0
	for _, x := range parsed {
		if !in(x, articles) {
			// copy and increment index
			parsed[i] = x
			i++
		}
	}
	parsed = parsed[:i]

	for i := range parsed {
		command := strings.Join(parsed[0:i+1], " ")

		choice, ok := commands[command]
		if !ok {
			continue
		}

		fmt.Println("")
		choice(s, parsed[i+1:]...)
		return
	}

	fmt.Println()
	fmt.Println("I do not understand that command")
	fmt.Print("use command 'help' for information on how to play")
	return
}

func listCharactersAtLocation(state *State, arguments ...string) {
	if len(state.CurrentLocation.ActiveCharacters) == 0 {
		fmt.Print("You see no one here")
	}

	for _, character := range state.CurrentLocation.ActiveCharacters {
		fmt.Print(fmt.Sprintf("You see %s here.", character.Name))
	}
}

// move doesn't work on windows
func move(state *State, arguments ...string) {
	if len(arguments) != 1 {
		fmt.Print("You must provide a single cardinal direction e.g north, south, east, or west.")
		return
	}

	state.PreviousLocation = state.CurrentLocation

	switch arguments[0] {
	case "north":
		if state.CurrentLocation.North != nil {
			state.CurrentLocation = state.CurrentLocation.North
		} else {
			fmt.Print("You cannot go this way.")
			return
		}

	case "south":
		if state.CurrentLocation.South != nil {
			state.CurrentLocation = state.CurrentLocation.South
		} else {
			fmt.Print("You cannot go this way.")
			return
		}

	case "east":
		if state.CurrentLocation.East != nil {
			state.CurrentLocation = state.CurrentLocation.East
		} else {
			fmt.Print("You cannot go this way.")
			return
		}

	case "west":
		if state.CurrentLocation.West != nil {
			state.CurrentLocation = state.CurrentLocation.West
		} else {
			fmt.Print("You cannot go this way.")
			return
		}

	default:
		fmt.Print("You must provide a valid cardinal direction.")
		return
	}

	fmt.Print(fmt.Sprintf("You've arrived at %s.", state.CurrentLocation.Name))
}

func goBack(state *State, arguments ...string) {
	if state.PreviousLocation == nil {
		fmt.Print("You have nowhere to go back to.")
		return
	}

	state.CurrentLocation, state.PreviousLocation = state.PreviousLocation, state.CurrentLocation

	fmt.Println(fmt.Sprintf("You return to %s", state.CurrentLocation.Name))
	fmt.Print(state.CurrentLocation.Description)
}

func talkTo(state *State, arguments ...string) {
	if len(state.CurrentLocation.ActiveCharacters) == 0 {
		fmt.Print("There is no one here to talk to.")
		return
	}

	if len(arguments) < 1 {
		fmt.Print("You must provide a character name.")
		return
	}

	for _, character := range state.CurrentLocation.ActiveCharacters {
		if strings.ToLower(character.Name) == strings.ToLower(arguments[0]) {
			state.Characters = append(state.Characters, character)

			entry := convoengine.FindEntryNode(character.RootConversationNode)
			if entry == nil {
				fmt.Print("They ignore you.")
				return
			}

			entry.Talk(state.Inventory)
		}
	}

	fmt.Println()
	fmt.Print("You are in " + state.CurrentLocation.Name)
}

func examinePointOfInterest(state *State, arguments ...string) {
	if len(state.CurrentLocation.PointsOfInterest) == 0 {
		fmt.Print("There is nothing here to look at.")
		return
	}

	if len(arguments) < 1 {
		fmt.Print("You must provide the name of what you want to look at.")
		return
	}

	for _, point := range state.CurrentLocation.PointsOfInterest {
		if strings.ToLower(point.Name) == strings.ToLower(strings.Join(arguments, " ")) {
			fmt.Print(wordwrap.WrapString(point.Description, 80))
			return
		}
	}

	for _, character := range state.CurrentLocation.ActiveCharacters {
		if strings.ToLower(character.Name) == strings.ToLower(strings.Join(arguments, " ")) {
			fmt.Print(wordwrap.WrapString(character.Description, 80))
			return
		}
	}

	for _, item := range state.Inventory {
		if strings.ToLower(item.Name) == strings.ToLower(strings.Join(arguments, " ")) {
			fmt.Print(wordwrap.WrapString(item.Description, 80))
			return
		}
	}

	for _, item := range state.CurrentLocation.ActiveItems {
		if strings.ToLower(item.Name) == strings.ToLower(strings.Join(arguments, " ")) {
			fmt.Print(wordwrap.WrapString(item.Description, 80))
			return
		}
	}

	fmt.Println("There is no point of interest called: " + arguments[0])
}

func exit(state *State, arguments ...string) {
	os.Exit(1)
}

func lookAround(state *State, arguments ...string) {

	fmt.Println(wordwrap.WrapString(state.CurrentLocation.Description, 80))
	fmt.Println()

	for _, point := range state.CurrentLocation.PointsOfInterest {
		fmt.Println("You see " + point.Name + " here.")
	}

	for _, item := range state.CurrentLocation.ActiveItems {
		fmt.Println("There is " + item.Name + " here.")
	}

	for _, character := range state.CurrentLocation.ActiveCharacters {
		fmt.Println(fmt.Sprintf("You see %s here.", character.Name))
	}

	fmt.Println()

	if state.CurrentLocation.North != nil {
		fmt.Println("To the North is the " + state.CurrentLocation.North.Name)
	}
	if state.CurrentLocation.South != nil {
		fmt.Println("To the South is the " + state.CurrentLocation.South.Name)
	}
	if state.CurrentLocation.East != nil {
		fmt.Println("To the East is the " + state.CurrentLocation.East.Name)
	}
	if state.CurrentLocation.West != nil {
		fmt.Print("To the West is the " + state.CurrentLocation.West.Name)
	}
}

func inventory(state *State, arguments ...string) {
	if len(state.Inventory) == 0 {
		fmt.Print("You have no items in your inventory.")
		return
	}

	fmt.Println("You have the following items:")
	for _, item := range state.Inventory {
		fmt.Println(item.Name)
	}
}

func whereAmI(state *State, arguments ...string) {
	fmt.Println("You are in " + state.CurrentLocation.Name)
	fmt.Println()

	if state.CurrentLocation.North != nil {
		fmt.Println("To the North is the " + state.CurrentLocation.North.Name)
	} else {
		fmt.Println("There is nothing to the North")
	}

	if state.CurrentLocation.South != nil {
		fmt.Println("To the South is the " + state.CurrentLocation.South.Name)
	} else {
		fmt.Println("There is nothing to the South")
	}

	if state.CurrentLocation.East != nil {
		fmt.Println("To the East is the " + state.CurrentLocation.East.Name)
	} else {
		fmt.Println("There is nothing to the East")
	}

	if state.CurrentLocation.West != nil {
		fmt.Print("To the West is the " + state.CurrentLocation.West.Name)
	} else {
		fmt.Print("There is nothing to the West")
	}
}

func jettisonCharacter(state *State, arguments ...string) {
	if len(arguments) < 1 {
		fmt.Print("You must provide a character name")
		return
	}

	for _, character := range state.Characters {
		buf := bufio.NewReader(os.Stdin)

		if strings.ToLower(character.Name) == strings.ToLower(arguments[0]) {
			fmt.Println(wordwrap.WrapString(character.AfterDeath, 80))
			fmt.Println()

			if character.CorrectChoice {
				fmt.Println(wordwrap.WrapString("You've successfully found the saboteur and saved your crew and the intelligence. Now you can use that intelligence to threaten the enemy's hidden, peaceful worlds and end the war", 80))
			} else {
				fmt.Println(wordwrap.WrapString(`Because you did not find the saboteur you and your crew were killed when they struck again, completely destroying the ship instead of simply disabling it. 
You are dead. Good job.`, 80))
			}

			fmt.Println()
			fmt.Print("Press any key to exit the game...")
			buf.ReadString('\n')
			os.Exit(1)
		}
	}

	fmt.Print("You must first talk to a character in order to jettison them")
}

func takeItem(state *State, arguments ...string) {
	var wantedItem *items.Item
	i := 0 // output index
	for _, item := range state.CurrentLocation.ActiveItems {
		if strings.ToLower(item.Name) != strings.ToLower(strings.Join(arguments, " ")) {
			state.CurrentLocation.ActiveItems[i] = item
			i++
			continue
		}
		wantedItem = item
	}

	state.CurrentLocation.ActiveItems = state.CurrentLocation.ActiveItems[:i]
	if wantedItem != nil {
		state.Inventory = append(state.Inventory, wantedItem)
		fmt.Print(wantedItem.Name + " is now in your inventory.")
		return
	}

	fmt.Print("You can't take " + strings.Join(arguments, " "))
}

func dropItem(state *State, arguments ...string) {
	for _, item := range state.Inventory {
		if strings.ToLower(item.Name) == strings.ToLower(strings.Join(arguments, " ")) {
			state.CurrentLocation.ActiveItems = append(state.CurrentLocation.ActiveItems, item)

			i := 0
			for _, item := range state.Inventory {
				if strings.ToLower(item.Name) != strings.ToLower(strings.Join(arguments, " ")) {
					state.Inventory[i] = item
					i++
					continue
				}
			}

			fmt.Print("You've dropped " + item.Name)
			return
		}
	}

	fmt.Print("You don't have " + strings.Join(arguments, " "))
}

func init() {
	commands = make(map[string]func(state *State, arguments ...string))

	commands["help"] = help
	commands["who is here"] = listCharactersAtLocation
	commands["move"] = move
	commands["go back"] = goBack
	commands["talk"] = talkTo
	commands["exit"] = exit

	commands["inventory"] = inventory
	commands["take"] = takeItem
	commands["drop"] = dropItem
	commands["look around"] = lookAround
	commands["examine"] = examinePointOfInterest
	commands["look at"] = examinePointOfInterest
	commands["where am i"] = whereAmI

	commands["jettison"] = jettisonCharacter
}

func in(needle string, haystack []string) bool {
	for _, bale := range haystack {
		if bale == needle {
			return true
		}
	}

	return false
}

func help(state *State, arguments ...string) {
	fmt.Print(`Airlock is a text adventure game. Below is a list of commands and a description of what they do. This game is currently in Beta.

help                      - repeats this message
exit                      - ends the game
move [cardinal direction] - moves the player to the north, south, east, or west
go back                   - returns the player to their previous location
look around               - lists points of interest and any people present in the player's current location
where am I                - prints player's location and the surrounding areas
who is here               - lists any people present in the player's location
talk to [person's name]   - initiates conversation with desired person
examine [name or thing]   - provides a description of the character or point of interest selected
take [thing]              - adds item to inventory
inventory                 - list items in inventory 
drop [thing]              - drop item from inventory into current location

jettison [name]           - jettisons the selected crew member to space and ends the game
`)
}
