package uptalk

import (
	"os"
	"testing"
	"upspin.io/upbox"
)

// Upbox package provides utils to set up testing environment for Upspin
const testUpboxConfig = "upbox.yml"

var sch *upbox.Schema

const testFileName = "test.pb"

func setup() {
	var err error
	sch, err = upbox.SchemaFromFile(testUpboxConfig)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	err := sch.Stop()
	if err != nil {
		panic(err)
	}
}

var testFile = os.Create(testFileName)

// TestNewChat tests the instantiation of a new upspin chat room
// chets are unique based on the uniqueness of the owner (who's dir
// it is stored) and the participants
// creating a new chat dispatches a unique hash based off the owner and
// members of the chat. This hash can be used as a reference.
func TestNewChat(t *testing.T) {
	c = NewClient(file)
	c.NewChat("recepient@example.com")
}

func TestNewChatWithMultipleMembers(t *testing.T) {

}

func TestListChatsOwnedByMe(t *testing.T) {

}

func TestListChats(t *testing.T) {

}

// TestJoinChat tests the ability for a user to join a chat created by
// another user
func TestJoinChat(t *testing.T) {

}

// TestSendMessage tests sending a message to a specific chat room based
// off the owner and participants
func TestSendMessage(t *testing.T) {

}

func TestSendMedia(t *testing.T) {

}

func TestListInvites(t *testing.T) {

}

func TestAcceptInvite(t *testing.T) {

}
