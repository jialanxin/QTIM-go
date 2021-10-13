package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/http"
	"strconv"
	"time"
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
func (info *Info) update() {
	topItemText, err := ask(info.channelID)
	if err != nil {
		fmt.Println(err)
	} else {
		if topItemText != info.currentTopText {
			err := info.informMe(topItemText)
			if err != nil {
				fmt.Println(err)
			}
			info.currentTopText = topItemText
			fmt.Println(info.channelTitle + "|" + info.currentTopText)
		}
	}
}
func (info *Info) informMe(messageTitle string) error {
	_, err := http.Get("https://sctapi.ftqq.com/SCT81476TcdRF4VBPsELPhLtP1uRkfd3X.send?title=" + info.channelTitle + "|" + messageTitle)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func ask(channelID string) (string, error) {
	url := "https://www.qingting.fm/channels/" + channelID
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return "", err
	}
	nodes, err := htmlquery.QueryAll(doc, "/html/body/div/div/div[3]/div[2]/div[1]/div[2]/div/ul/li[1]/span[1]/a/p/text()")
	if err != nil {
		return "", err
	}
	text := nodes[0]
	return text.Data, nil
}

func main() {
	GuanQi := newInfo(387255, "观棋有语")
	ChaHuaHui := newInfo(418553, "察话会")
	for {
		time.Sleep(10 * time.Minute)
		GuanQi.update()
		ChaHuaHui.update()
	}
}
