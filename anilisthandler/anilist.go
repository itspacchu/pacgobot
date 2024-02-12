package anilisthandler

import (
	"fmt"
	"pacgobot/miscutils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type BotHandler struct {
	Name        string
	Description string
	ShortHand   string
	HandlerFunc func(s *discordgo.Session, m *discordgo.MessageCreate)
}

var PREFIX = ""

func SetSelfAnilistUserHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	COMMAND := PREFIX + "aniuser"
	if strings.HasPrefix(m.Content, COMMAND) {
		username, exists := strings.CutPrefix(m.Content, COMMAND)
		if exists {
			fmt.Println(username)
		}
	}
}

func AnilistUserHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	COMMAND := PREFIX + "aniuser"
	if strings.HasPrefix(m.Content, COMMAND) {
		username, exists := strings.CutPrefix(m.Content, COMMAND)
		if exists {
			udata := GetUserData(strings.TrimSpace(username))
			// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", udata))

			embed := &discordgo.MessageEmbed{
				URL:   udata.Data.User.SiteURL,
				Title: fmt.Sprintf("%s's anilist summary", strings.TrimSpace(username)),
				Color: miscutils.GetAverageColor(udata.Data.User.Avatar.Medium),
				Description: fmt.Sprintf(`
				UID: %d
				---
				Anime Watched %d
				Anime Mean Score %.2f
				---
				Manga Watched %d
				Manga Mean Score %.2f
				---
				`, udata.Data.User.ID, udata.Data.User.Statistics.Anime.Count, udata.Data.User.Statistics.Anime.MeanScore, udata.Data.User.Statistics.Manga.Count, udata.Data.User.Statistics.Anime.MeanScore),
				Image: &discordgo.MessageEmbedImage{
					URL: udata.Data.User.Avatar.Large,
				},
			}
			s.ChannelMessageSendEmbedReply(m.ChannelID, embed, m.Message.Reference())

		}
	}
}
