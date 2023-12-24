package commands

import (
	"moko/src/libs"
	"time"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(ping|speed)",
		As:       []string{"ping"},
		Tags:     "main",
		IsPrefix: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			start := time.Now()
			m.Reply("⚡ Speed: " + time.Since(start).String())
		},
	})
}
