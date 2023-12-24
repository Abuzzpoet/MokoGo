package commands

import (
	"fmt"
	"moko/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:       "(revoke|resetlink)",
		As:         []string{"revoke"},
		Tags:       "group",
		IsPrefix:   true,
		IsWaitt:    true,
		IsAdmin:    true,
		IsBotAdmin: true,
		IsGroup:    true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			resp, err := client.WA.GetGroupInviteLink(m.From, true)
			if err != nil {
				m.Reply("Moko gagal mereset link group ❌")
			} else {
				m.Reply(fmt.Sprintf("Berhasil Mendapatkan link group baru: %s ✅", resp))
			}
		},
	})
}
