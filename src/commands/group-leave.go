package commands

import (
	"fmt"
	"moko/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(leave)",
		As:       []string{"leave"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsOwner:  true,
		IsGroup:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			err := client.WA.LeaveGroup(m.From)
			if err != nil {
				m.Reply("Moko gagal keluar dari group ini ❌")
				fmt.Println(err.Error())
			}
		},
	})
}
