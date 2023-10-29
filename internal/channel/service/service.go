package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Service ...
type Service struct {
	DiscordGo *discordgo.Session
	DataBase  *sql.DB
}

// New ..
func New(dg *discordgo.Session) Service {
	return Service{DiscordGo: dg}
}

var (
	//Ex: Nada é proibido, até que alguém te impeça. - Cid Cidoso.
	FRASE_PARA_ETERNIDADE_REGEX        = os.Getenv("FRASE_PARA_ETERNIDADE_REGEX")
	FRASE_PARA_ETERNIDADE_CHANNEL_NAME = os.Getenv("FRASE_PARA_ETERNIDADE_CHANNEL_NAME")
	FRASE_PARA_ETERNIDADE_CHANNEL_ID   = os.Getenv("FRASE_PARA_ETERNIDADE_CHANNEL_ID")
	MESSAGE_AMOUNT_TO_GET, _           = strconv.Atoi(os.Getenv("MESSAGE_AMOUNT_TO_GET"))
	MANY_TIMES_TO_GET_MESSAGE, _       = strconv.Atoi(os.Getenv("MANY_TIMES_TO_GET_MESSAGE"))
)

// GetChannelMessagesContent ...
func (s Service) GetChannelMessagesContent(channelID string) ([]ChannelMessage, error) {
	return getChannelMessagesFromJSON(channelID)
}

// GetRandomChannelMessage ...
func (s Service) GetRandomChannelMessage(channelID string) (*ChannelMessage, error) {
	msgs, err := s.GetChannelMessagesContent(channelID)
	if err != nil {
		return nil, err
	}

	randomIndex := rand.Intn(len(msgs))
	return &ChannelMessage{
		ID:             msgs[randomIndex].ID,
		Content:        msgs[randomIndex].Content,
		Author:         msgs[randomIndex].Author,
		CreatedAtMonth: msgs[randomIndex].CreatedAtMonth,
		CreatedAtYear:  msgs[randomIndex].CreatedAtYear,
		SaveApproved:   msgs[randomIndex].SaveApproved,
	}, nil
}

// UpdateData ...
func (s Service) UpdateData(channelID string) error {
	msgs, err := s.getChannelMessagesOnLimit(channelID)
	if err != nil {
		return err
	}

	channelMsgs, err := buildChannelMessagesResponse(msgs)
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(channelMsgs, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("internal/channel/service/public-data/frases-para-a-eternidade.json", file, 0644)
	if err != nil {
		fmt.Println("Error updating channel messages, ", err)
		return err
	}

	for i := range channelMsgs {
		err := s.DiscordGo.MessageReactionAdd(channelID, channelMsgs[i].ID, "✔")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) getChannelMessagesOnLimit(channelID string) ([]*discordgo.Message, error) {
	var (
		emptyString = ""
		msgs        = []*discordgo.Message{}
		err         error
	)
	for i := 1; i <= MANY_TIMES_TO_GET_MESSAGE; i++ {
		msgs, err = s.DiscordGo.ChannelMessages(channelID, MESSAGE_AMOUNT_TO_GET, emptyString, emptyString, emptyString)
		if err != nil {
			fmt.Println("Error retrieving channel messages, ", err)
			return nil, err
		}
	}

	return msgs, nil
}

func buildChannelMessagesResponse(msgs []*discordgo.Message) ([]ChannelMessage, error) {
	channelMsgs := []ChannelMessage{}
	for i := range msgs {
		matched, err := regexp.MatchString(FRASE_PARA_ETERNIDADE_REGEX, msgs[i].Content)
		if err != nil {
			return nil, err
		}

		t, err := time.Parse(time.RFC3339, string(msgs[i].Timestamp))
		if err != nil {
			return nil, err
		}

		if matched {
			cm := ChannelMessage{
				ID:             msgs[i].ID,
				Author:         msgs[i].Author.Username,
				Content:        strings.Replace(msgs[i].Content, `"`, "", 2),
				CreatedAtMonth: translateMonthFromEnglishToPortuguese(t.Month()),
				CreatedAtYear:  strconv.Itoa(t.Year()),
				SaveApproved:   matched,
			}

			channelMsgs = append(channelMsgs, cm)
		}
	}

	return channelMsgs, nil
}

// func buildPhraseForEternitiesResponse(msgs []*discordgo.Message) []PhraseForEternity {
// 	pfes := []PhraseForEternity{}
// 	for i := range msgs {
// 		pfe := buildPhraseForEternityResponse(msgs[i].Content)

// 		pfes = append(pfes, pfe)
// 	}

// 	return pfes
// }

func buildPhraseForEternityResponse(content string) PhraseForEternity {
	i := strings.Index(content, `"`)
	i2 := strings.Index(content[i+1:], `"`)
	phrase := content[i+1 : i2-1]
	i3 := strings.Index(content, `.`)
	author := content[i2+2 : i3]
	author = strings.Trim(author, ",")
	date := content[i3+1:]

	return PhraseForEternity{
		Phrase: phrase,
		Author: author,
		Date:   date,
	}
}

func translateMonthFromEnglishToPortuguese(month time.Month) string {
	switch month.String() {
	case "January":
		return "Janeiro"
	case "February":
		return "Fevereiro"
	case "March":
		return "Março"
	case "April":
		return "Abril"
	case "May":
		return "Maio"
	case "June":
		return "Junho"
	case "July":
		return "Julho"
	case "August":
		return "Agosto"
	case "September":
		return "Setembro"
	case "October":
		return "Outubro"
	case "November":
		return "Novembro"
	case "December":
		return "Dezembro"
	default:
		return "Abril"
	}
}

func getChannelMessagesFromJSON(channelID string) ([]ChannelMessage, error) {
	channelName := ""
	switch channelID {
	case FRASE_PARA_ETERNIDADE_CHANNEL_ID:
		channelName = FRASE_PARA_ETERNIDADE_CHANNEL_NAME
	default:
		channelName = FRASE_PARA_ETERNIDADE_CHANNEL_NAME
	}

	// Open our jsonFile
	jsonFile, err := os.Open("internal/channel/service/public-data/" + channelName + ".json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened" + channelName + ".json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	msgs := []ChannelMessage{}
	json.Unmarshal([]byte(byteValue), &msgs)

	return msgs, nil
}
