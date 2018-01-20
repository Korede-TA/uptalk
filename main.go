package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/korede-ta/uptalk"
	// "github.com/urfave/cli"
)

func main() {
	fmt.Println("vim-go")

	command := os.Args[1]
	user := os.Getenv("UPSPINUSR")
	usage := `
	Uptalk Usage:
	`

	switch command {
	case "new-chat":
		var members []string
		for i := 2; i < len(os.Args); i++ {
			members = append(members, os.Args[i])
		}

		hash := uptalk.NewChat(user, members...)
		fmt.Printf("New Chat Hash: %v\n", hash)
		break

	case "list-chats":
		chats := uptalk.ListChats(user)

		var members []string
		for _, member := range chat.Members {
			members = append(members, member.Username)
		}

		for _, chat := range chats {
			fmt.Printf(
				"Owner: %s, Members: %v, Hash: %v\n",
				user,
				members,
				chat.GetHash().String(),
			)
		}
		break

	// not sure if this functionality will end up existing
	// just seems extraneous
	case "join-chat":
		var members []string
		for i := 1; i < len(os.Args); i++ {
			members = append(members, os.Args[i])
		}

		hash := uptalk.JoinChat(user, members...)
		fmt.Printf("Joined Chat Hash: %v\n", hash)
		break

	case "send-invite":
		var invitees []string
		for i := 1; i < len(os.Args); i++ {
			invitees = append(invitees, os.Args[i])
		}
		uptalk.SendInvite(invitees...)
		break

	case "send-message":
		chatHash := os.Args[2]
		// messages are plain text. Could potentially take more standard
		// formats in the future (markdown?)
		message := os.Args[3]
		uptalk.SendMessage(chatHash, message)
		break

	case "send-media":
		chatHash := os.Args[2]
		// file path to media to be sent
		mediaPath := os.Args[3]
		uptalk.sendMedia(chatHash, mediaPath)
		break

	case "help":
	default:
		fmt.Printf(usage)
		break
	}

}
