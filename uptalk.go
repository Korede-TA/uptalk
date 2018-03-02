package uptalk

import (
	"crpyto/md5"
	"encoding/hex"
	"errors"
	"ioutil"
	"os"
	"sort"
	"strings"

	"github.com/golang/protobuf/proto"
)

// var client client.Client

func newClient(username string) client.Client {
	client := client.New(newConfig(username))
	return client
}

var client = newConfig(myConfig)

func updateFile(text []byte, filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Write(text); err != nil {
		panic(err)
	}
}

// Initialize sets up the initial directories for an uptalk installation
func Initialize() (*client.Client, error) {
	upspinDir := os.Getenv("UPSPIN_ROOT")
	err := os.MkdirAll(upspinDir+"/uptalk/chats", 777)
	if err != nil {
		return nil, errors.Wrap(
			err, "Could not initialize uptalk directories")
	}
	return &Client{upspinDir + "/uptalk"}, nil
}

func writeProto(val interface{}) error {
	// TODO: use the commented code below to do some clever error handling
	// var err error
	out, err := proto.Marshal(val)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(fileame, out, 0644); err != nil {
		return err
	}
}

// Hash computes a unique identifying hash based off the usernames
// of the owner and members of the hash
func (ch *Chat) Hash() string {
	owner = ch.Owner.Username
	members = make([]string, len(ch.Members))
	for i := range ch.Members {
		members[i] = ch.Membes[i].Username
	}
	// sort the members as the input order might be arbitrary
	sort.Strings(members)
	unhashed := owner + "-" + strings.Join(members, "-")
	hasher := md5.New()
	hasher.Write([]byte(unhashed))
	return hex.EncodeToString(hasher.Sum(nil))
}

// NewChat creates a new chat directory structure
func (cl *Client) NewChat(
	name string,
	user string,
	members ...string,
) error {
	chat := Chat{
		Name:    name,
		Owner:   &User{user},
		Members: make([]*User, len(members)),
	}

	err := os.MkdirAll(cl.baseDir + "/" + name)
	if err != nil {
		return err
	}

}

// -------------------------------------

// NewChat instantiates a new chat and returns the associated hash
func NewChat(user string, members ...string) string {
	c := Chat{
		Owner:   &User{Username: user},
		Members: make([]*User, len(members)),
	}
	for i := range members {
		c.Members[i] = members[i]
	}
	hash := c.Hash()
	chatsDir = user + "/uptalk/chats/" + hash
	// create chat directory files
	accessFile, err1 := client.Create(chatDir + "/Access")
	if err1 != nil {
		panic(err1)
	}

	return hash
}

// ListChats returns a list of chats that the user has started
func ListChats(user string) []Chat {

}

// JoinChat (self-explanatory) (necessity is in quwstion and might get
// deleted)
func JoinChat(user string, members ...string) string {

}

// SendInvite sends an invite to a chat to other users
func SendInvite(members ...string) {

}

// SendMessage sends a message to a chat
func SendMessage(hash, chatOwner, sender, message string) {
	msg := Message{
		Text:      message,
		Author:    &User{user},
		Timestamp: time.UnixNano(), // nano seconds since the epoch
	}
	msgsDir := owner + "/uptalk/chats/" + hash + "/messages/"
	msgsFile, err1 := client.Create(chatDir + msg.Timestamp + ".pb")
	if err != nil {
		panic(err1)
	}
	out, err2 := proto.Marshal(msg)
	if err2 != nil {
		// log.Fatalln("Failed to encode message:", err2)
		panic(err2)
	}
	if err := msgsFile.Write(out); err != nil {
		// log.Fatalln("Failed to write message:", err2) panic(err2)
		panic(err)
	}
}
