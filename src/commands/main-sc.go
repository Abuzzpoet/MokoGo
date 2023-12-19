package commands

import (
	"moko/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(sc|source)",
		As:       []string{"sc"},
		Tags:     "main",
		IsPrefix: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			m.Reply("*Sc ini Private*\n Base yg dipake : https://github.com/fckvania/MaoGo")
		},
	})
}
