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
const BaseGame = `{
  "entry": {
    "title": "Airlock",
    "author": "John Darrington",
    "introductionText": "Mankind has colonized our the solar system yet still stands divided. The Inner Planets, consisting of Earth and Mars, have continually oppressed and held sway over The Colonies on Jupiter's moon and surrounding space.\n\nNeeding Earth and Mars for supplies of food and medicine The Colonies dare not openly defy them. Needing the raw minerals and hydrogen mined around Jupiter, The Inner Planets must rule space with an iron fist or risk becoming helpless and defenseless when they are unable to manufacture ships and weapons. \n\nAboard the independent space freighter Gently Weeping, you and your crew of five have managed to avoid the majority of the conflict. Recently however, a supply run of medical supplies was interrupted by an attack from The Colonies on one of The Inner Planet's most lucrative, and most populated space stations. \n\nThe battle has left your ship stranded in space and operating on emergency power. You've contacted help and a rescue ship is on the way - but you only have enough oxygen for five people. \n\nAs the captain it is your duty to insure the safety of both crew and cargo. The duty of choosing who lives and who dies falls to you. \n\nChoose Wisely...",
    "initialArea": "main"
  },
  "items": [
    {
      "id": "syringe",
      "name": "syringe",
      "description": "The syringe is encased in glass and there is no sign of the needle. There is a rose-tinted liquid in the syringe that moves slightly when you move it."
    },
    {
      "id": "letter",
      "name": "letter",
      "description": "The letter looks very formal and is very creased, as if it'd been opened and closed numerous times. You open it slightly and see the official seal of The Colonies. The letter is addressed to an ethics board. You can't see anything else unless you open the letter further - you don't want to pry."
    }
  ],
  "locations": [
    {
      "id": "captains_quarters",
      "name": "Captain's Quarters",
      "entryPoint": true,
      "description": "This is your cabin, and it has seen better days. Off to the left is your bunk, covered in clothes and data pads. To your right sits your desk - which is also a disaster. A single yellow light bar runs through the room lengthwise and gives off only a faint yellow light. On your desk sits a single picture and what looks like a newspaper.",
      "directions": ["_", "-", "north_hallway", "-"],
      "characters": ["twins"],
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
      "directions": ["cockpit", "south_hallway", "airlock", "captains_quarters"]
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
      "characters": ["manuel"],
      "items": ["syringe"],
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
      "id": "south_hallway",
      "name": "South Hallway",
      "description": "This part of the hallway is as dirty as the others - except for the door to the first mate's quarters which is a spotless white.",
      "directions": [
        "north_hallway",
        "machine_shop",
        "common_area",
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
      "items": ["letter"],
      "pointsOfInterest": [
        {
          "name": "handbook",
          "description": "This handbook contains all the laws independent freighters have to follow. This copy is well used, dog eared pages and bookmarks stick out everywhere."
        }
      ],
      "directions": ["-", "-", "south_hallway", "-"]
    },
    {
      "id": "common_area",
      "name": "The Common Area",
      "description": "Next to the cargo hold, this is the largest area in the ship. A large table takes up one half of the room and is bolted the floor. What can only be described as a \"space garage sale\" fills the other half. Random placements of furniture and personal items strewn about complete the effect. There are a few doors around the room, each leading to personal or shared quarters. Soft emergency lighting lends the room an abandoned feel.",
      "items": ["letter"],
      "characters": ["twins"],
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
      "characters": ["bear"],
      "pointsOfInterest": [
        {
          "name": "crates",
          "description": "Inspecting the crates more closely you see that the majority of them hold supplies for your ship. Food and water make up the bulk of these supplies with whatever room left being packed with spare parts."
        },
        {
          "name": "loader",
          "description": "The electric loader is basically a forklift and run on ship's power via a very long extension cable. It is strapped to the side of the hold."
        },
        {
          "name": "stuffed bear",
          "description": "This stuffed blue bear is in very good shape. You can tell it's used frequently, but other than that it is clean and all the stitches are in place. It's sitting right next to Bear."
        }
      ],
      "directions": ["-", "-", "machine_shop", "-"]
    },
    {
      "id": "machine_shop",
      "name": "Machine Shop",
      "description": "Parts, tools, and electronics cover almost every surface of this room. Hung on the wall next to the door are a few hardhats, shiny and clean from disuse. A winch is embedded in the ceiling and is currently holding what can only be described as a giant ball of electronic yarn.",
      "directions": ["south_hallway", "engine_room", "-", "cargo_hold"],
      "characters": ["rebecca"]
    },
    {
      "id": "engine_room",
      "name": "Engine Room",
      "description": "There isn't much room here. A large, inverted pyramid of pitch black metal hands suspended from the ceiling. The tip disappearing into the floor. There are no tools or electronics in the area. Light is provided by a bank of candles and mirrors to either side of the door you just entered.",
      "pointsOfInterest": [
        {
          "name": "pyramid",
          "description": "The pyramid is twice your height and could easily fit three or full grown adults inside. It is pitch black and it seems to suck in the light given off by the candles. There is no hum or any indication that the engine is doing anything. This could be normal, you don't know. You don't come down here often."
        }
      ],
      "directions": ["machine_shop", "-", "-", "-"]
    }
  ],
  "characters": [
    {
      "id": "bear",
      "name": "Bear",
      "description": "Bear stands at least 6'5\" and weighs more than you'd like to think about. You found him on the streets and you put him to work hauling and providing general protection for the crew. He's simple. Though he can understand you, he has the mind of a five year old. To him, today is just another work day and he smiles when you talk to him.",
      "afterDeath": "You couldn't look Bear in the eyes as you tried to explain to him how he was about to help the ship. He was all smiles and wanted nothing else than to help. You led him to the airlock, the bear which was his namesake clutched in one hand with yours in the other. When the end came...you couldn't do it. You couldn't live with yourself afterwards. \n\nYou told Bear to look away and stepped into the airlock yourself.",
      "conversations": [
        {
          "id": "start",
          "entryPoint": true,
          "trigger": "Hey big guy",
          "text": "Bear looks at you and smiles.",
          "ignoreAfterVisit": true,
          "choices": ["big_guy", "end"]
        },
        {
          "id": "big_guy",
          "trigger": "Doing ok big guy?",
          "text": "\"Yes!\" then he grins as wide as he can.",
          "choices": ["ship_broken"],
          "ignoreAfterVisit": true
        },
        {
          "id": "ship_broken",
          "trigger": "You know the ship is broken?",
          "text": "Yeah boss - let's fix it!",
          "choices": ["choke_up"],
          "ignoreAfterVisit": true
        },
        {
          "id": "choke_up",
          "trigger": "Continue",
          "text": "You choke up, you can't continue the conversation. You smile, pat Bear on the arm and leave.",
          "exitPoint": true
        },
        {
          "id": "end",
          "trigger": "Cya big guy",
          "text": "Bear smiles",
          "exitPoint": true
        }
      ]
    },
    {
      "id": "manuel",
      "name": "Manuel",
      "description": "Manuel looks to be in his early forties. What hair he has on his head is reddish gray, clashing with the yellow jumpsuit he has on. He is your pilot and is currently strapped into the single seat in the room. He's a little lethargic and barely stirred when you came into the room. He seems just as ready to forget you came as talk.",
      "afterDeath": "Manuel took the order stoically, nodding his head once while he unstrapped from the chair and made his way to the airlock. \n\n\"No sense putting things off Cap'n\" he said as he entered the airlock and sealed the door himself. \n\nBefore he left he'd passed you a sealed envelope addressed to a house in Mars. You can only assume it's for someone he'll never meet again.",
      "conversations": [
        {
          "id": "start",
          "trigger": "I have some more questions",
          "text": "Cap'n",
          "entryPoint": true,
          "canBeRecalled": true,
          "choices": ["wrong_ship", "syringe", "have_to_do", "end"]
        },

        {
          "id": "wrong_ship",
          "trigger": "Do you know what's wrong with the ship?",
          "text": "All I know is that she don't fly"
        },
        {
          "id": "syringe",
          "mustHaveItem": "syringe",
          "trigger": "[syringe] Found this - you know I don't allow drugs on my ship.",
          "afterVisitedText": "Already told you Cap'n",
          "canBeRecalled": false,
          "text": "Of course - it's a memento.",
          "choices": ["memento"]
        },
        {
          "id": "memento",
          "trigger": "Oh yeah? Memento of what exactly.",
          "text": "It's from a run out to The Colonies a while back. The one where we got little less returns than we was hope'n for.",
          "choices": ["memento_end"]
        },
        {
          "id": "memento_end",
          "trigger": "And?",
          "text": "And what we was carrying was medical supplies Cap'n. Earth had decided to skip that particular station in order to \"send a message\" - we were the only ones dumb enough to ignore it and go anyways. I like to remember that.",
          "choices": ["start", "end"]
        },
        {
          "id": "have_to_do",
          "trigger": "You know what I have to do, right?",
          "text": "Yep, and I'm glad you're makin the call",
          "canBeRecalled": true,
          "choices": ["volunteer", "who", "start"]
        },
        {
          "id": "volunteer",
          "trigger": "You going to volunteer?",
          "text": "No offense Cap'n, but death ain't something I'm ready for yet."
        },
        {
          "id": "who",
          "trigger": "Who would you..?",
          "text": "That's not a question you should be asking me."
        },
        {
          "id": "what_do",
          "trigger": "What do you do?",
          "text": ".....I'm the pilot Cap'n",
          "choices": ["why_cargo"]
        },
        {
          "id": "why_cargo",
          "trigger": "Why pilot a cargo ship? There's more exciting things to fly.",
          "text": "Like what? Some warship? Nah - here there ain't any pressure and I generally get paid on the regular. Get to do my own thing.",
          "canBeRecalled": true,
          "choices": ["own_thing", "lonely", "start"]
        },
        {
          "id": "own_thing",
          "trigger": "And what is your own thing?",
          "text": "Getting paid and not taking any risks. Though I guess that last part doesn't apply anymore."
        },
        {
          "id": "lonely",
          "trigger": "Isn't it lonely out there?",
          "text": "You tell me, I don't see anyone warming your bunk at night."
        },
        {
          "id": "end",
          "trigger": "I need to go",
          "text": "Sounds good to me",
          "exitPoint": true
        }
      ]
    },
    {
      "id": "rebecca",
      "name": "Rebecca",
      "description": "Much like her quarters your first mate Rebecca is precise in movement and in dress. She salutes you, though you've repeatedly told her this isn't a military ship. She looks tired and it's a good bet she hasn't slept since the ship broke down. There isn't a lot she can do, but she'll try until she physically can't. She'd make a better captain than you would, maybe.",
      "afterDeath": "After it was over, and she was gone, you started to go through Rebecca's personal items. You were surprised how much she had stored in a single locker, and that she never kept any of it out in the open. The items spoke of a deliberate life - each choice taken with purpose and with a full idea of the consequences. \n\nAll but one. \n\nAt the bottom of the locker was a small bag with a lock of hair and what looked like a miniature shoe. The hair was the same color as Rebecca's. ",
      "conversations": [
        {
          "id": "start",
          "trigger": "Let's talk some more.",
          "text": "Captain. How can I help?",
          "entryPoint": true,
          "choices": ["fix", "colonist", "air", "end"]
        },
        {
          "id": "fix",
          "trigger": "You're an engineer now? Trying to fix the ship?",
          "text": "We're all trying to help Captain. If anything, it keeps our minds of what has to happen next.",
          "choices": ["help"]
        },
        {
          "id": "help",
          "trigger": "How can I help?",
          "text": "Make the call no one else wants to think about.",
          "choices": ["start"]
        },
        {
          "id": "air",
          "trigger": "You know we're running out of air?",
          "text": "Of course Captain. It's your job to fix it... \n\n\nCaptain?",
          "ignoreAfterVisit": true,
          "choices": ["yes"]
        },
        {
          "id": "yes",
          "trigger": "Yes?",
          "text": "Choose me.",
          "choices": ["what"]
        },
        {
          "id": "what",
          "trigger": "What?",
          "text": "Choose me. I'm the best choice. I'm mostly redundant. Everyone else has a job, I just help you carry out yours. You can do it yourself. That way you can go on working after I'm....after I'm gone sir.",
          "choices": ["not_again"]
        },
        {
          "id": "not_again",
          "trigger": "We won't talk about this again.",
          "text": "Yes sir.",
          "exitPoint": true
        },
        {
          "id": "colonist",
          "trigger": "You're a colonist, right?",
          "text": "No sir. I was born on Mars. I choose to serve in The Colonies as part of their militia.",
          "choices": ["why"]
        },
        {
          "id": "why",
          "trigger": "Then why choose to work with The Colonies?",
          "text": "Just wanted to show them we aren't all bad sir",
          "choices": ["and"]
        },
        {
          "id": "and",
          "trigger": "And how did that go?",
          "text": "Not well.",
          "choices": ["letter"]
        },
        {
          "id": "letter",
          "trigger": "[letter] Is that why you resigned?",
          "mustHaveItem": "letter",
          "text": "They wouldn't listen. Earth and Mars are just trying to help. Trying to make us all be a family again."
        },
        {
          "id": "end",
          "trigger": "I'll see what else I can do.",
          "text": "We'll make it Captain.",
          "exitPoint": true
        }
      ]
    },
    {
      "id": "twins",
      "name": "The twins, Alexi and Gregor",
      "description": "The twins are two skinny men, both around their late 50s or 60s. Alexi and Gregor serve as cook and accountant respectively. Having been born on Earth, in Russia, they've only spent a few years in space and haven't yet become fully adapted to life in low gravity. They appear uncomfortable, a mix of the low gravity, their heated conversation, and the choice that's on everyone's mind.",
      "aliases": ["alexi", "gregor"],
      "afterDeath": "Right before the airlock closed the twins decided they couldn't bear to be parted. They didn't give you time to refuse or rethink your choice. When the end came and the hatch to dark space slid open the twins didn't look at it, or you. They could only hold on to each other, their tears freezing on smiling faces as they left for the void, together.",
      "conversations": [
        {
          "id": "start",
          "trigger": "Let's talk about something else",
          "text": "Alexi: \"Hello Captain\"",
          "entryPoint": true,
          "canBeRecalled": true,
          "choices": ["talking_about", "photo", "wish", "end"]
        },
        {
          "id": "talking_about",
          "trigger": "What are you two talking about?",
          "text": "Alexi: \"Nothing Captain\" \nGregor: \"Not true - we're talking about the war\"",
          "choices": ["war"]
        },
        {
          "id": "war",
          "trigger": "War? You can hardly call a few border skirmishes a war.",
          "text": "Gregor: \"You don't get to decide what is and isn't war. People are dying - that has to be explained somehow. The Colonies have been itching to get their hands dirty, show the inner planets that they've got teeth. \n\nAlexi: \"It's not like Earth and Mars have been pacifists, this is just as much their fault as it is The Colonies.\"",
          "canBeRecalled": true,
          "choices": ["colonies", "earth_mars", "start"]
        },
        {
          "id": "colonies",
          "trigger": "The Colonies can't be expected to put up with Earth and Mar's oppression forever.",
          "text": "Alexi: \"Yes, but does that justify violence?\""
        },
        {
          "id": "earth_mars",
          "trigger": "Earth and Mars want a peaceful end to this, they stand to lose just as much as The Colonies do.",
          "text": "Gregor: \"How incredibly naive Captain. You can't seriously think that Earth and Mars want anything but control?\" "
        },
        {
          "id": "photo",
          "trigger": "[photo] You guys looked a lot younger here, who is this girl in the middle?",
          "mustHaveItem": "photo",
          "text": "Gregor: \"That's our sister, we must have been in our thirties when that was taken - and she must have been around twenty or so.\" \n\nAlexi: \"We haven't seen her in a long time.\"",
          "choices": ["where"]
        },
        {
          "id": "where",
          "trigger": "Where is she now?",
          "text": "Gregor: \"Can't say. Last I heard she'd taken a shuttle out to work on one of the hydrogen mines around Jupiter.\" \n\nAlexi: \"Space is a big place, hard to keep in contact.\"",
          "choices": ["start", "end"]
        },
        {
          "id": "end",
          "trigger": "Let's talk later",
          "text": "In unison: \"Until next time.\"",
          "exitPoint": true
        }
      ]
    }
  ]
}
`
