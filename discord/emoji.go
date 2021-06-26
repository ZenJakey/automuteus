package discord

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zenjakey/utils/pkg/game"

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
	topMap[true] = make([]Emoji, 27) // 18 colors for alive/dead
	topMap[false] = make([]Emoji, 27)
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
			ID:   "858185424362405888",
		},
		game.Blue: {
			Name: "aublue",
			ID:   "858185425386078238",
		},
		game.Green: {
			Name: "augreen",
			ID:   "858185426094653460",
		},
		game.Pink: {
			Name: "aupink",
			ID:   "858185426794709022",
		},
		game.Orange: {
			Name: "auorange",
			ID:   "858185427408126012",
		},
		game.Yellow: {
			Name: "auyellow",
			ID:   "858185427932151839",
		},
		game.Black: {
			Name: "aublack",
			ID:   "858185511805779978",
		},
		game.White: {
			Name: "auwhite",
			ID:   "858185512460353586",
		},
		game.Purple: {
			Name: "aupurple",
			ID:   "858185513106538506",
		},
		game.Brown: {
			Name: "aubrown",
			ID:   "858185513718775858",
		},
		game.Cyan: {
			Name: "aucyan",
			ID:   "858185514053009460",
		},
		game.Lime: {
			Name: "aulime",
			ID:   "858185515086118982",
		},
		game.Maroon: {
			Name: "aumaroon",
			ID:   "858185515743445032",
		},
		game.Rose: {
			Name: "aurose",
			ID:   "858185516452151296",
		},
		game.Banana: {
			Name: "aubanana",
			ID:   "858185516763971585",
		},
		game.Gray: {
			Name: "augray",
			ID:   "858185517911638036",
		},
		game.Tan: {
			Name: "autan",
			ID:   "858185518675525643",
		},
		game.Coral: {
			Name: "aucoral",
			ID:   "858185519439806475",
		},
		game.Watermelon: {
			Name: "auwatermelon",
			ID:   "858204098402582558",
		},
		game.Chocolate: {
			Name: "auchocolate",
			ID:   "858204098319220776",
		},
		game.Skyblue: {
			Name: "auskyblue",
			ID:   "858204098364964884",
		},
		game.Beige: {
			Name: "aubeige",
			ID:   "858204098342813696",
		},
		game.Hotpink: {
			Name: "auhotpink",
			ID:   "858204098053799978",
		},
		game.Turquoise: {
			Name: "auturquoise",
			ID:   "858204098570485760",
		},
		game.Lilac: {
			Name: "aulilac",
			ID:   "858204098342682634",
		},
		game.Rainbow: {
			Name: "aurainbow",
			ID:   "858204098432204800",
		},
		game.Azure: {
			Name: "auazure",
			ID:   "858204098309652540",
		},
	},
	false: []Emoji{
		game.Red: {
			Name: "aureddead",
			ID:   "858185527840604201",
		},
		game.Blue: {
			Name: "aubluedead",
			ID:   "858185528318230600",
		},
		game.Green: {
			Name: "augreendead",
			ID:   "858185529069928469",
		},
		game.Pink: {
			Name: "aupinkdead",
			ID:   "858185530076692500",
		},
		game.Orange: {
			Name: "auorangedead",
			ID:   "858185530688274432",
		},
		game.Yellow: {
			Name: "auyellowdead",
			ID:   "858185531178090507",
		},
		game.Black: {
			Name: "aublackdead",
			ID:   "858185532165324830",
		},
		game.White: {
			Name: "auwhitedead",
			ID:   "858185532781363210",
		},
		game.Purple: {
			Name: "aupurpledead",
			ID:   "858185533658103848",
		},
		game.Brown: {
			Name: "aubrowndead",
			ID:   "858185534954536980",
		},
		game.Cyan: {
			Name: "aucyandead",
			ID:   "858185535524831232",
		},
		game.Lime: {
			Name: "aulimedead",
			ID:   "858185536039944213",
		},
		game.Maroon: {
			Name: "aumaroondead",
			ID:   "858185536967409704",
		},
		game.Rose: {
			Name: "aurosedead",
			ID:   "858185537617526784",
		},
		game.Banana: {
			Name: "aubananadead",
			ID:   "858185538255585310",
		},
		game.Gray: {
			Name: "augraydead",
			ID:   "858185539039526912",
		},
		game.Tan: {
			Name: "autandead",
			ID:   "858185539692527667",
		},
		game.Coral: {
			Name: "aucoraldead",
			ID:   "858185540063199263",
		},
		game.Watermelon: {
			Name: "auwatermelondead",
			ID:   "858204098436137001",
		},
		game.Chocolate: {
			Name: "auchocolatedead",
			ID:   "858204098411102208",
		},
		game.Skyblue: {
			Name: "auskybluedead",
			ID:   "858204098587000853",
		},
		game.Beige: {
			Name: "aubeigedead",
			ID:   "858204098340192256",
		},
		game.Hotpink: {
			Name: "auhotpinkdead",
			ID:   "858204098364440576",
		},
		game.Turquoise: {
			Name: "auturquoisedead",
			ID:   "858204098373746738",
		},
		game.Lilac: {
			Name: "aulilacdead",
			ID:   "858204098343075870",
		},
		game.Rainbow: {
			Name: "aurainbowdead",
			ID:   "858204098376630274",
		},
		game.Azure: {
			Name: "auazuredead",
			ID:   "858204098422374420",
		},
	},
}
