package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
	"github.com/hypebeast/go-osc/osc"
)

type CLIFlags struct {
	OSCServerAddr string `kong:"default='127.0.0.1',name='server-addr',help='OSC server address'"`
	OSCServerPort int    `kong:"default=8884,name='server-port',help='OSC server port'"`
	OSCClientAddr string `kong:"default='127.0.0.1',name='client-addr',help='OSC client address'"`
	OSCClientPort int    `kong:"default=8885,name='client-port',help='OSC client port'"`
	Verbose       bool   `kong:"name='verbose',help='Show extra output in the terminal'"`
}

func main() {
	var cli CLIFlags
	kong.Parse(&cli)

	d := osc.NewStandardDispatcher()

	d.AddMsgHandler("/text", func(msg *osc.Message) {
		lastValue := msg.Arguments[len(msg.Arguments)-1]
		if strValue, ok := lastValue.(string); ok {
			parsedtext := sentitext.Parse(strValue, lexicon.DefaultLexicon)
			sentiment := sentitext.PolarityScore(parsedtext)

			if cli.Verbose {
				fmt.Println("Str:", strValue)
				fmt.Println("Pos:", sentiment.Positive)
				fmt.Println("Neg:", sentiment.Negative)
				fmt.Println("Neu:", sentiment.Neutral)
				fmt.Println("Compound/Final Sentiment:", sentiment.Compound)
			}

			msg := osc.NewMessage("/sentiment")
			msg.Append(float32(sentiment.Compound))
			if cli.Verbose {
				fmt.Printf("Sending '%s' to new OSC Client on %s:%d\n", msg, cli.OSCClientAddr, cli.OSCClientPort)
			}
			client := osc.NewClient(cli.OSCClientAddr, cli.OSCClientPort)
			client.Send(msg)
		} else {
			fmt.Println("Warning: Incoming message not formatted as '/text STRING'")
		}

	})

	if cli.Verbose {
		fmt.Printf("Starting OSC Server on on %s:%d\n", cli.OSCServerAddr, cli.OSCServerPort)
	}
	saddr := fmt.Sprintf("%s:%d", cli.OSCServerAddr, cli.OSCServerPort)
	server := &osc.Server{
		Addr:       saddr,
		Dispatcher: d,
	}
	server.ListenAndServe()
}
