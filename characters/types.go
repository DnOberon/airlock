package characters

import (
	"github.com/dnoberon/airlock/convoengine"
)

// Character represents an interaction ready person
type Character struct {
	ID          string // slug like identifier
	Name        string
	Description string
	AfterDeath  string
	Aliases     []string

	RootConversationNode *convoengine.ConversationNode
	Conversations        []convoengine.ConversationNode
}

// InitCharacters loads both characters and conversations
func InitCharacters(characters []Character) []Character {

	for i := range characters {
		characters[i].RootConversationNode = convoengine.InitConversations(characters[i].Conversations)
	}

	return characters
}
