package commands

import (
	"fmt"
	"moko/src/libs"
	"moko/src/libs/api"
	"regexp"
	"time"

	"github.com/Pauloo27/searchtube"
)

type method string

const (
	audio string = "Audio"
	video string = "Video"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(ytmp4|ytmp3|play)",
		As:       []string{"ytmp4", "ytmp3", "play"},
		Tags:     "downloader",
		IsPrefix: true,
		IsQuerry: true,
		IsWaitt:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			var url string
			if api.IsYoutubeURL(m.Querry) {
				url = m.Querry
			} else {
				ser, _ := searchtube.Search(m.Querry, 10)
				if len(ser) == 0 {
					m.Reply("Not Found")
					return
				}
				for _, v := range ser {
					d, _ := v.GetDuration()
					if d < 8*time.Minute {
						url = v.URL
						break
					}
				}
			}
			yt, err := api.YoutubeDL(url)
			if err != nil {
				m.Reply(err.Error())
				return
			}

			caption := fmt.Sprintf("🏷️ *Title*: %s\n👤 *Author*: %s\n📝 *Duration*: %s\n", yt.Info.Title, yt.Info.Author, yt.Info.Duration)

			if reg, _ := regexp.MatchString(`(ytmp3|play)`, m.Command); reg {
				build, err := yt.Link.Audio[0].Url()
				if err != nil {
					m.Reply(err.Error())
					return
				}

				bytes, err := client.GetBytes(build)
				if err != nil {
					m.Reply(err.Error())
					return
				}
				caption += "🔖 *Quality:* " + yt.Link.Audio[0].Quality
				caption += "\n📥 *Size:* " + yt.Link.Audio[0].Size
				caption += "\n🔎 *Format:* " + yt.Link.Audio[0].Format
				client.SendDocument(m.From, bytes, fmt.Sprintf("%s.mp3", yt.Info.Title), caption, m.ID)
			} else {
				build, err := yt.Link.Video[0].Url()
				if err != nil {
					m.Reply(err.Error())
					return
				}
				bytes, err := client.GetBytes(build)
				if err != nil {
					m.Reply(err.Error())
					return
				}
				caption += "🔖 *Quality:* " + yt.Link.Video[0].Quality
				caption += "\n📥 *Size:* " + yt.Link.Video[0].Size
				caption += "\n🔎 *Format:* " + yt.Link.Video[0].Format
				client.SendVideo(m.From, bytes, caption, m.ID)
			}
		},
	})
}
