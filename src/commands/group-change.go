package commands

import (
	"fmt"
	"moko/src/libs"
	"regexp"
	"strings"

	"github.com/nyaruka/phonenumbers"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:       "(demote|dm|promote|pm)",
		As:         []string{"demote", "promote"},
		Tags:       "group",
		IsPrefix:   true,
		IsGroup:    true,
		IsAdmin:    true,
		IsBotAdmin: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			var ujid []types.JID
			var ok error
			// apalah ini gw bingung
			if m.QuotedMsg != nil {
				if m.QuotedMsg.MentionedJid != nil {
					ajid := m.QuotedMsg.MentionedJid
					ujid = make([]types.JID, len(ajid))
					for i, a := range ajid {
						ujid[i], ok = types.ParseJID(a)
						if ok != nil {
							return
						}
					}
				} else {
					ujid = make([]types.JID, 0)
					jid, _ := types.ParseJID(*m.QuotedMsg.Participant)
					ujid = append(ujid, jid)
				}
			} else if len(m.Querry) > 0 {
				ajid := strings.Split(strings.Trim(m.Querry, " "), ",")
				ujid = make([]types.JID, len(ajid))
				for i, a := range ajid {
					num, err := phonenumbers.Parse(a, "ID")
					if err != nil {
						return
					}
					num_formatted := phonenumbers.Format(num, phonenumbers.E164)
					ujid[i], ok = types.ParseJID(fmt.Sprintf("%s@whatsapp.net", num_formatted[1:]))
					if ok != nil {
						return
					}
				}
			}
			if ujid == nil || len(ujid) == 0 {
				m.Reply("Tag atau balas pesan seseorang yang mau dijadikan admin/dijatuhkan dari admin ❗")
				return
			}

			if regexp.MustCompile(`demote|dm`).MatchString(m.Command) {
				resp, err := client.WA.UpdateGroupParticipants(m.From, ujid, whatsmeow.ParticipantChangeDemote)
				if err != nil {
					m.Reply("Gagal menurunkan admin ❌")
					return
				}

				for _, item := range resp {
					if item.Error == 404 {
						m.Reply("Mungkin user tersebut sudah tidak ada di grup ini ❗")
					} else if !item.IsAdmin {
						m.Reply("Sukses menurunkan admin ✅")
					}
				}
			} else {
				resp, err := client.WA.UpdateGroupParticipants(m.From, ujid, whatsmeow.ParticipantChangePromote)
				if err != nil {
					m.Reply("Gagal menjadikan admin ❌")
					return
				}

				for _, item := range resp {
					if item.IsAdmin {
						m.Reply("Sukses menjadikan admin ✅")
					}
					if item.Error == 404 {
						m.Reply("Mungkin user tersebut sudah tidak ada di grup ini ❗")
					}
				}
			}
		},
	})
}
