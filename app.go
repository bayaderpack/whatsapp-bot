package main

import (
	"context"
	"fmt"
	"go.mau.fi/whatsmeow/types"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) WhatsAppStart() {

	wac, err := WAConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer wac.Disconnect()

}

func (a *App) SendMessage(phone string, message string) string {
	_, err := cli.SendMessage(context.Background(),types.JID{
		User:   phone,
		Server: types.DefaultUserServer,
	},  &waProto.Message{
		Conversation: proto.String(message),
	})
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("Message: \"%s\" successfully sended to the number %s!", message, phone)
}