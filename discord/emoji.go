package discord

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ZenJakey/utils/pkg/game"

	"github.com/bwmarrin/discordgo"
)

// Emoji struct for discord
type Emoji struct {
	Name string
	ID   string
}

// FormatForReaction does what it sounds like
func (e *Emoji) FormatForReaction() string {
	return "<:" + e.Name + ":" + e.ID
}

// FormatForInline does what it sounds like
func (e *Emoji) FormatForInline() string {
	return "<:" + e.Name + ":" + e.ID + ">"
}

// GetDiscordCDNUrl does what it sounds like
func (e *Emoji) GetDiscordCDNUrl() string {
	return "https://cdn.discordapp.com/emojis/" + e.ID + ".png"
}

// DownloadAndBase64Encode does what it sounds like
func (e *Emoji) DownloadAndBase64Encode() string {
	url := e.GetDiscordCDNUrl()
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	encodedStr := base64.StdEncoding.EncodeToString(bytes)
	return "data:image/png;base64," + encodedStr
}

func emptyStatusEmojis() AlivenessEmojis {
	topMap := make(AlivenessEmojis)
	topMap[true] = make([]Emoji, 20) // 12 colors for alive/dead + 8 for Town of Us
	topMap[false] = make([]Emoji, 20)
	return topMap
}

func (bot *Bot) addAllMissingEmojis(s *discordgo.Session, guildID string, alive bool, serverEmojis []*discordgo.Emoji) {
	for i, emoji := range GlobalAlivenessEmojis[alive] {
		alreadyExists := false
		for _, v := range serverEmojis {
			if v.Name == emoji.Name {
				emoji.ID = v.ID
				bot.StatusEmojis[alive][i] = emoji
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			b64 := emoji.DownloadAndBase64Encode()
			em, err := s.GuildEmojiCreate(guildID, emoji.Name, b64, nil)
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Added emoji %s successfully!\n", emoji.Name)
				emoji.ID = em.ID
				bot.StatusEmojis[alive][i] = emoji
			}
		}
	}
}

// AlivenessEmojis map
type AlivenessEmojis map[bool][]Emoji

// GlobalAlivenessEmojis keys are IsAlive, Color
var GlobalAlivenessEmojis = AlivenessEmojis{
	true: []Emoji{
		game.Red: {
			Name: "aured",
			ID:   "762392085768175646",
		},
		game.Blue: {
			Name: "aublue",
			ID:   "762392085629632512",
		},
		game.Green: {
			Name: "augreen",
			ID:   "762392085889417226",
		},
		game.Pink: {
			Name: "aupink",
			ID:   "762392085726363648",
		},
		game.Orange: {
			Name: "auorange",
			ID:   "762392085264728095",
		},
		game.Yellow: {
			Name: "auyellow",
			ID:   "762392085541158923",
		},
		game.Black: {
			Name: "aublack",
			ID:   "762392086493790249",
		},
		game.White: {
			Name: "auwhite",
			ID:   "762392085990866974",
		},
		game.Purple: {
			Name: "aupurple",
			ID:   "762392085973303376",
		},
		game.Brown: {
			Name: "aubrown",
			ID:   "762392086023634986",
		},
		game.Cyan: {
			Name: "aucyan",
			ID:   "762392087945281557",
		},
		game.Lime: {
			Name: "aulime",
			ID:   "762392088121442334",
		},
		game.Watermelon: {
			Name: "auwatermelon",
			ID:   "846734024780546059",
		},
		game.Chocolate: {
			Name: "auchocolate",
			ID:   "846734025258565692",
		},
		game.Skyblue: {
			Name: "auskyblue",
			ID:   "846734024730345473",
		},
		game.Beige: {
			Name: "aubeige",
			ID:   "846734025070739456",
		},
		game.Hotpink: {
			Name: "auhotpink",
			ID:   "846734025140862986",
		},
		game.Turquoise: {
			Name: "auturquoise",
			ID:   "846734025073623050",
		},
		game.Lilac: {
			Name: "aulilac",
			ID:   "846734025116483614",
		},
		game.Rainbow: {
			Name: "aurainbow",
			ID:   "846734025237987368",
		},
	},
	false: []Emoji{
		game.Red: {
			Name: "aureddead",
			ID:   "762397192362393640",
		},
		game.Blue: {
			Name: "aubluedead",
			ID:   "762397192349679616",
		},
		game.Green: {
			Name: "augreendead",
			ID:   "762397192060272724",
		},
		game.Pink: {
			Name: "aupinkdead",
			ID:   "762397192643805194",
		},
		game.Orange: {
			Name: "auorangedead",
			ID:   "762397192333819904",
		},
		game.Yellow: {
			Name: "auyellowdead",
			ID:   "762397192425046016",
		},
		game.Black: {
			Name: "aublackdead",
			ID:   "762397192291090462",
		},
		game.White: {
			Name: "auwhitedead",
			ID:   "762397192409186344",
		},
		game.Purple: {
			Name: "aupurpledead",
			ID:   "762397192404860958",
		},
		game.Brown: {
			Name: "aubrowndead",
			ID:   "762397192102739989",
		},
		game.Cyan: {
			Name: "aucyandead",
			ID:   "762397192307867698",
		},
		game.Lime: {
			Name: "aulimedead",
			ID:   "762397192366325793",
		},
		game.Watermelon: {
			Name: "auwatermelondead",
			ID:   "846734099733413888",
		},
		game.Chocolate: {
			Name: "auchocolatedead",
			ID:   "846734099677970432",
		},
		game.Skyblue: {
			Name: "auskybluedead",
			ID:   "846734099875495956",
		},
		game.Beige: {
			Name: "aubeigedead",
			ID:   "846734099552796722",
		},
		game.Hotpink: {
			Name: "auhotpinkdead",
			ID:   "846734099565641749",
		},
		game.Turquoise: {
			Name: "auturquoisedead",
			ID:   "846734099501809684",
		},
		game.Lilac: {
			Name: "aulilacdead",
			ID:   "846734099687145482",
		},
		game.Rainbow: {
			Name: "aurainbowdead",
			ID:   "846734099745734706",
		},
	},
}
