package main

import (
	"fmt"

	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
	"github.com/hypebeast/go-osc/osc"
)

func main() {
	addr := "172.23.225.228:8765"
	client := osc.NewClient("192.168.4.124", 8766)

	d := osc.NewStandardDispatcher()

	d.AddMsgHandler("/text", func(msg *osc.Message) {
		// osc.PrintMessage(msg)
		lastValue := msg.Arguments[len(msg.Arguments)-1]
		if strValue, ok := lastValue.(string); ok {
			parsedtext := sentitext.Parse(strValue, lexicon.DefaultLexicon)
			sentiment := sentitext.PolarityScore(parsedtext)
			fmt.Println("Str:", strValue)
			fmt.Println("Pos:", sentiment.Positive)
			fmt.Println("Neg:", sentiment.Negative)
			fmt.Println("Neu:", sentiment.Neutral)
			fmt.Println("Compound/Final Sentiment:", sentiment.Compound)

			msg := osc.NewMessage("/sentiment")
			msg.Append(float32(sentiment.Compound))
			client.Send(msg)
		}

	})

	server := &osc.Server{
		Addr:       addr,
		Dispatcher: d,
	}
	server.ListenAndServe()
}
