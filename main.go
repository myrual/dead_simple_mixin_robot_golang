package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/MooooonStar/mixin-sdk-go/messenger"
)

type Listener struct {
	*messenger.Messenger
}

// interface to implement if you want to handle the message
func (l *Listener) OnMessage(ctx context.Context, msg messenger.MessageView, userId string) error {
	data, err := base64.StdEncoding.DecodeString(msg.Data)
	if err != nil {
		return err
	}
	if msg.Category == "SYSTEM_ACCOUNT_SNAPSHOT" {
		var transfer messenger.TransferView
		if err := json.Unmarshal(data, &transfer); err != nil {
			return err
		}
		log.Println("I got a snapshot: ", transfer)
		return l.SendPlainText(ctx, msg.ConversationId, msg.UserId, string(data))
	} else {
		log.Printf("I got a message, it said: %s", string(data))
		return l.SendPlainText(ctx, msg.ConversationId, msg.UserId, string(data))
	}
}

const (
	UserId    = "595053f7-c441-4266-a8e7-ca6eb6123d63"
	PinCode   = "518443"
	SessionId = "0b5ab266-0da3-43e1-b24a-b4d11671c134"
	PinToken  = "O2D8XVoPIkCPDCVSy7wmRLxV/AqELZl//Is2ZY0SLHtYtAHsEDlL8ie7EcvjoPeJMIPsOWdIXhMDuJdL2Q2kUAxutGA4//dw1w/EEqCUonWVGAGOEw2GPK2q8yYkkiWuS+ygD6ze+U7dd8SYrmAFeWoqIO2KnXaUBY5w9sA1b8A="
	//please delele the blank of PrivateKey the before each line
	PrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCpA3wZQI0txHzbXGh7PmWKUAtIj3ddS1XynqFforCC+0DQkI17
dI5eROKruJL88+fI3MMFz/z4v7n8slWVcA7wDtGmaRM/d0fwNmVpwfBD7s7KXujg
DG/JYNL3Jasl/oyX7d/8/2qIeK5wRo26NM9TE11yJWnUJkeaHWqvwd1CKQIDAQAB
AoGAcDNLCCyNTKvbQ3j2N5L1r7txsirj0AB6P0XF1YpYaavInGzKMAHxnW1XmS5A
DkMRRYwgv9QSmUQWgUYWgM1SbD4uwwmjoZW6SJXT6nFzDfNVgiOIOw65qwCwodpI
5yedcJWU2/SgMg6koesl/oG2h3kUVufa1cZLLYc7B+ozP7ECQQDwNlz+sKopwTy8
b6yfdKGMU2ilxCMZnxTdXGNT6e1EqChup9qqupjLvHI5ww7tUUfyMibbxFyXFps+
aAGuzRR9AkEAtB8vB2MKdZF9nlUhYtKg586sprfa4xm0e2RwfFqut5gSBBX0vmZ5
mByJfi9UoC1uGMMHGfmNmzmETv/cUnywHQJASt/nJMxUfidpkyMMH7fvExp8qA73
EelwS3+cxm9IMfpof/V6R3VaY0ceI7sVUAvh7TCMxDv00HOPeGha/e2W0QJBAIlu
UZ+9Pbm+rK4mHjccppC37Ju4JFaqwj5zwC/hsPV8HiLn7bloztX00CNrUZJ1l09l
XwG+mwEYvY7ZQe3eHTECQDo6eldTLlUW7Q6ALjwUzNTBaQ0Q/19zhXkd3SfYPXzt
g07CDzMD0ynLxqpChz5oCThooLMIOhU+EMzceJfMmj0=
-----END RSA PRIVATE KEY-----`
)
func main() {
	ctx := context.Background()
	m := messenger.NewMessenger(UserId, SessionId, PrivateKey)
	//replace with your own listener
	//go m.Run(ctx, messenger.DefaultBlazeListener{})
	l := &Listener{m}
	go m.Run(ctx, l)

	//your mixin user id, can get from  "Search User"
	snow := "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"

	//must create conversation first. If have created before, skip this step.
	if _, err := m.CreateConversation(ctx, messenger.CategoryContact, messenger.Participant{UserID: snow}); err != nil {
		log.Println("create conversation error", err)
	}

	conversation, err := m.CreateConversation(ctx, messenger.CategoryContact,
		messenger.Participant{UserID: snow},
	)
	if err != nil {
		log.Println("create error", err)
	}

	if err := m.SendImage(ctx, conversation.ID, snow, "../donate.png"); err != nil {
		log.Println("send image error:", err)
	}

	if err := m.SendVideo(ctx, conversation.ID, snow, "../sample.mp4"); err != nil {
		log.Println("send video error", err)
	}

	if err := m.SendFile(ctx, conversation.ID, snow, "../demo.pdf", "application/pdf"); err != nil {
		log.Println("send video error", err)
	}

	if err := m.SendPlainText(ctx, conversation.ID, snow, "please send me a message and transfer some CNB to me."); err != nil {
		log.Println("send text error:", err)
	}

	select {}
}
