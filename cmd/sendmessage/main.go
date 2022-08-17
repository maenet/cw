package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/maenet/go-chatwork"
)

var token = flag.String("token", os.Getenv("CHATWORK_API_TOKEN"),
	"The Chatwork API token. If not specified, the CHATWORK_API_TOKEN environment variable will be read.")
var selfUnread = flag.Bool("unread", false, "Make the message you send unread.")

func sendmessage() error {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 2 {
		return fmt.Errorf("invalid arguments")
	}
	args := flag.Args()
	roomID := args[0]
	message := args[1]

	form := &chatwork.PostMessageForm{}
	form.Body = message
	form.SelfUnread = *selfUnread

	// TODO logging
	client, err := chatwork.NewClient(chatwork.BaseURLV2, *token, nil)
	if err != nil {
		return err
	}

	if _, err := client.PostMessage(context.Background(), roomID, form); err != nil {
		return err
	}

	return nil
}

func main() {
	// TODO handle signal
	if err := sendmessage(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: sendmessage [flags] <room id> <message>\n")
	flag.PrintDefaults()
}
