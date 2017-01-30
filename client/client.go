package client

import "github.com/nlopes/slack"
import "fmt"
import "os"
import "strings"
import "time"

var links = make(map[string]int)

func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}
func rightPad(s string, padStr string, pLen int) string {
	return s + strings.Repeat(padStr, pLen)
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

func SaveHTML(team, room string) {
	ts := time.Now().Unix() - int64(31536000*5)
	tss := fmt.Sprintf("%d", ts)

	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		api := slack.New(tokens[i])

		j := 0
		for {
			fmt.Println("syncing ", j)
			j += 1000
			hp := slack.HistoryParameters{Oldest: tss, Latest: "", Count: 1000, Inclusive: false, Unreads: false}
			list, _ := api.GetGroupHistory(room, hp)
			stamps := make([]string, 0)
			for _, r := range list.Messages {
				//SaveMsg(team, room, r.Msg)
				//fmt.Println(r.Msg.Timestamp)
				//fmt.Println(r.Msg.Text)
				//fmt.Println(r.Msg.Attachments)
				stamps = append(stamps, r.Msg.Timestamp)
			}
			if len(stamps) == 0 {
				break
			}
			tss = stamps[0]
			//fmt.Println("-----")
			//time.Sleep(time.Second)
		}
	}

	for k, v := range links {
		if v > 2 {
			fmt.Println(k, v)
		}
	}
}

func ListRooms(team string) {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		api := slack.New(tokens[i])
		list, _ := api.GetChannels(false)
		for _, r := range list {
			fmt.Println(r.ID, r.Name)
		}
		list2, _ := api.GetGroups(false)
		for _, r := range list2 {
			fmt.Println(r.ID, r.Name)
		}
		list3, _ := api.GetIMChannels()
		for _, r := range list3 {
			u, _ := api.GetUserInfo(r.User)
			fmt.Println(r.ID, u.Name)
		}
	}
}

func ListTeams() {
	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	for _, team := range teams {
		fmt.Println(team)
	}
}