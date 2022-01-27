package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/antchfx/htmlquery"
)

type Info struct {
	channelID      string
	channelTitle   string
	currentTopText string
}

func newInfo(channelID int, channelTitle string) *Info {
	channelIDString := strconv.Itoa(channelID)
	channelTopText, err := ask(channelIDString)
	if err != nil {
		panic(err)
	}
	fmt.Println(channelTitle + "|" + channelTopText)
	return &Info{
		channelID:      channelIDString,
		channelTitle:   channelTitle,
		currentTopText: channelTopText,
	}
}
func (info *Info) update(token string) {
	topItemText, err := ask(info.channelID)
	if err != nil {
		fmt.Println(err)
	} else {
		if topItemText != info.currentTopText {
			err := info.informMe(topItemText, token)
			if err != nil {
				fmt.Println(err)
			}
			info.currentTopText = topItemText
			fmt.Println(info.channelTitle + "|" + info.currentTopText)
		}
	}
}
func (info *Info) informMe(messageTitle string, token string) error {
	_, err := http.Get("https://sctapi.ftqq.com/" + token + ".send?title=" + info.channelTitle + "|" + messageTitle)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func ask_old(channelID string) (string, error) {
	url := "https://www.qingting.fm/channels/" + channelID
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return "", err
	}
	nodes, err := htmlquery.QueryAll(doc, "/html/body/div/div/div[2]/div[2]/div[1]/div[2]/div/ul/li[1]/span[1]/a/p/text()")
	if err != nil {
		return "", err
	}
	text := nodes[0]
	return text.Data, nil
}

func ask(channelID string) (string, error) {
	url := "https://i.qingting.fm/capi/v3/channel/" + channelID
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	s, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", err2
	}
	var jsonTemplate map[string]interface{}
	err3 := json.Unmarshal(s, &jsonTemplate)
	if err3 != nil {
		return "", err3
	}
	data := (jsonTemplate["data"]).(map[string]interface{})
	lastestProgram := data["latest_program"].(string)
	return lastestProgram, nil
}

func main() {
	token := os.Args[1]
	GuanQi := newInfo(387255, "观棋有语")
	ChaHuaHui := newInfo(418553, "察话会")
	for {
		time.Sleep(10 * time.Minute)
		GuanQi.update(token)
		ChaHuaHui.update(token)
	}
}
