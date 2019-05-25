# Airlock - a customizable text adventure game

Airlock is a text adventure with a simple engine and end condition. You start the game in a certain location and can interact with items, characters, and other points of interest in that location. You can also move to other locations, each having their own set of items, characters etc. The game ends when you select a character to remove from the game (or in this case, jettison out an airlock).

The game engine itself is static. Commands and interaction types do not change. There will always be locations, characters, and items. You can always interact with them in certain ways and the game will always end when a character is removed from the game.

The game's characters (and their conversation trees), items, and locations can easily be changed by providing your own, custom `game.json` file. 

*Disclaimer: There be dragons in the codebase.* 

<br>

## Available commands
```
help                      - displays this list of commands help text
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
```

<br>

## The `game.json` file
------

The `game.json` file consists of three high level object collections: `items`, `characters`, and `locations`. It also consists of a single high level object called `entry`. Each of these are required in order for a custom `game.json` file to work. See the end of the file for an example.

<br>

## The `entry` object
-----
The `entry` object contains information on the game itself, an introduction text block, and the start location of the player.

Sample:
```
  "entry": {
    "title": "Airlock",
    "author": "John Darrington",
    "introductionText": "The galaxy is at war. Posing as an independent freighter, you and your crew of six are attempting to return home with intelligence that could turn the tide of war in your favor.\n\nSomeone aboard has purposely disabled the ship and you must now await rescue. You will run out of air before rescue arrives. Your only chance is to jettison a single crew member in order to make the air in the ship last long enough. \n\nIf you jettison the spy the intel is safe. If you jettison someone else - let's just say the war will only get bloodier. \n \nChoose wisely...",
    "initialArea": "main"
  }
  ```

<br>

## The `item` object
----
The `item` object represents an in-game item. ID should be one word, or multiple words in snake case. e.g `leather_wallet`. The ID is how you will tell locations what items it contains.

Sample:
```
{
  "id": "wallet",
  "name": "Leather wallet",
  "description": "The leather wallet has the initials BYB on it."
}
```

<br>
 
## The `location` object
-----
The `location` object represents a physical location in the game. The location _can_ contain items, characters, and points of interest. The location _must_ contain a name, ID, and description. 

#### `location` Characters
The `characters` key in the location object should be an array of `string`. Each entry is the ID of a character inside the `characters` high level object of the configuration file. When the locations are initialized on game start, it will attempt to find every character listed in this array by ID and include that in the in-game location.

### `location` Items
Like the `characters` key, `items` is an array of `string` with each entry being the ID of an item in the `items` high level object of the configuration file.

### `location` Directions
The `directions` key is how you connect locations to each other. `directions` is an array of type `string`. Each entry should be the ID of another location in the `locations` high level object of the configuration file. Movement in the game is accomplished by using the `move` command with a cardinal direction as the argument. In order to make that work the `directions` the location ID entries must be follow this order: `[north, south, east, west]`. Leave a blank string, `""`, if you do not want to specify a location for one of the cardinal directions.

Sample:
```
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
    }
]
```

### `location` Points of Interest
The `pointsOfInterest` is a collection of the `pointOfInterest` type. This allows the player to focus on certain parts of a location. A point of interest must have a name and description to be considered valid

Sample:
```
"pointsOfInterest": [
        {
          "name": "plaque",
          "description": "The ship's information plaque: \n Built 2051 by OmesCorp"
        }
      ]
```

<br>

## The `character` object
-----
The `character` object represents an in-game character contains a conversation tree unique to that character. This allows the player to examine and interact a character through scripted conversation paths. Must contain a unique id, name, and description. Must contain at least one conversation object in the `conversations` list. 

_Note:_ The `afterDeath` key allows you to display text if that character is chosen to be removed as part of the end game functionality.

### 

Sample:
```
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
          "choices": [],
          "canBeRecalled": true
        }
       ]
    }
]

```

<br>

## The `conversation` object
----
Being able to talk to a character is a large part of this engine and the original game. As such, this is probably the most complex concept of the `game.json` file. Each conversation object in the `conversations` list of a character represents a single node of a conversation tree. Each node contains a trigger (the text when the node is listed as a choice in the conversation itself), the conversation text itself, conversation text that displays after the node has been visited once. 

The object also contains information on how the game should treat this node such as whether or not it's considered an entry point to a conversation tree or if it will end the conversation and return the user to the main play screen. 

A node can do the following:

* terminate a conversation
* return the player to the previous conversation node
* can be ignored when a parent node is listing possible conversation choices
* can be the first node a player visits when beginning the conversation with a character
* can contain different text after having been visited once
* can be available only if the player has a certain item in their inventory.

Sample:
```
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
```

<br>

## Full sample `game.json` file
----
```
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
```