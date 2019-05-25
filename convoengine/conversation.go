package convoengine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dnoberon/airlock/items"
	"github.com/mitchellh/go-wordwrap"
)

func in(needle string, haystack []string) bool {
	for _, bale := range haystack {
		if bale == needle {
			return true
		}
	}

	return false
}

// FindEntryNode to being conversation on
func FindEntryNode(start *ConversationNode) *ConversationNode {
	if start.EntryPoint && (!start.visited || start.visited && !start.IgnoreAfterVisit) {
		return start
	}

	for _, choice := range start.choices {
		if c := FindEntryNode(choice); c != nil {
			return c
		}
	}

	return nil
}

// Talk starts the conversation
func (root *ConversationNode) Talk(itemList []*items.Item) {
	choiceMap := map[string]*ConversationNode{}

	fmt.Println()

	// text after visited
	if root.visited && root.AfterVisitedText != "" {
		fmt.Println(wordwrap.WrapString(root.AfterVisitedText, 80))
	} else {
		fmt.Println(wordwrap.WrapString(root.Text, 80))
	}

	// make sure we print any output before exit
	if root.ExitPoint {
		return
	}

	root.visited = true
	// build and print choice map
	for i := range root.choices {
		if root.choices[i].visited && root.choices[i].IgnoreAfterVisit {
			continue
		}

		if root.choices[i].MustHaveItem != "" {
			for _, item := range itemList {
				if item.ID == root.choices[i].MustHaveItem {
					fmt.Println(wordwrap.WrapString(fmt.Sprintf("%d. %s", i+1, root.choices[i].Trigger), 80))
					choiceMap[fmt.Sprintf("%d", i+1)] = root.choices[i]
				}
			}

			continue
		}

		fmt.Println(wordwrap.WrapString(fmt.Sprintf("%d. %s", i+1, root.choices[i].Trigger), 80))
		choiceMap[fmt.Sprintf("%d", i+1)] = root.choices[i]
	}

	if root.parent != nil && root.parent.CanBeRecalled {
		fmt.Println(fmt.Sprintf("%d. %s", len(choiceMap)+1, "Go back"))
		choiceMap[fmt.Sprintf("%d", len(choiceMap)+1)] = root.parent
	}

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")

	in, err := buf.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	node, ok := choiceMap[strings.TrimSpace(strings.TrimSuffix(in, "\n"))]

	if !ok {
		fmt.Print("No option selected")
		return
	}

	node.parent = root
	node.Talk(itemList)
}
