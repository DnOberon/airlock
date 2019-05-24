package convoengine

// ConversationNode is a multiple linked list node.
type ConversationNode struct {
	ID               string // slug like identifier allows for flattened input data
	Trigger          string
	Text             string
	AfterVisitedText string
	EntryPoint       bool
	ExitPoint        bool

	IgnoreAfterVisit bool
	visited          bool

	CanBeRecalled bool
	parent        *ConversationNode
	Choices       []string // slug identifiers of other conversation nodes
	choices       []*ConversationNode
}

// InitConversations is the entry point
func InitConversations(conversationNodes []ConversationNode) *ConversationNode {
	// not terribly concerned with performance here
	for i, node := range conversationNodes {
		for j, connection := range conversationNodes {
			if in(connection.ID, node.Choices) {
				conversationNodes[i].choices = append(conversationNodes[i].choices, &conversationNodes[j])
			}
		}

	}

	return &conversationNodes[0]
}
