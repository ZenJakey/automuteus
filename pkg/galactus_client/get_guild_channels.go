package galactus_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
)

func (galactus *GalactusClient) GetGuildChannels(guildID string) ([]*discordgo.Channel, error) {
	resp, err := galactus.client.Post(galactus.Address+GetGuildChannelsPartial+guildID, "application/json", bytes.NewBufferString(""))
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading all bytes from resp body for getChannels")
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := errors.New("non-200 status code received for GetChannels:")
		return nil, err
	}

	var channels []*discordgo.Channel
	err = json.Unmarshal(respBytes, &channels)
	if err != nil {
		return nil, err
	}
	return channels, nil
}