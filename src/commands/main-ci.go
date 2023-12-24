package commands

import (
	"fmt"
	"moko/src/libs"
	"regexp"
	"strings"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(channelinfo|ci)",
		As:       []string{"ci"},
		Tags:     "main",
		IsPrefix: true,
		IsQuerry: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			pattern := regexp.MustCompile(`https?://whatsapp.com/channel/`)
			if !pattern.MatchString(m.Querry) {
				m.Reply("Url Invalid")
				return
			}

			key, err := client.WA.GetNewsletterInfoWithInvite(strings.Split(m.Querry, "/")[4])
			if err != nil {
				m.Reply("Moko Tidak Tau Ya ❗")
				return
			}

			m.Reply(fmt.Sprintf("🔎 *Channel Information*\n📌 *Link:* %s\n🏷️ *ID:* %s\n🔖 *Name:* %v\n👥 *Followers:* %v\n\n📝 *Description:* %v\n*Create At:* %v", m.Querry, key.ID, key.ThreadMeta.Name.Text, key.ThreadMeta.SubscriberCount, key.ThreadMeta.Description.Text, key.ThreadMeta.CreationTime))
		},
	})
}
