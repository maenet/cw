# sendmessage

Example CLI to send a message to the Chatwork room

## Installation

```
$ go install github.com/maenet/go-chatwork/cmd/sendmessage@latest
```

## Usage

```
$ sendmessage -h
usage: sendmessage [flags] <room id> <message>
  -token string
        The Chatwork API token. If not specified, the CHATWORK_API_TOKEN environment variable will be read.
  -unread
        Make the message you send unread.
```
