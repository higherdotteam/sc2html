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
	ts := time.Now().Unix()
	tss := fmt.Sprintf("%d", ts)

	teams := strings.Split(os.Getenv("SLACK_TEAMS"), ",")
	tokens := strings.Split(os.Getenv("SLACK_TOKENS"), ",")
	for i, t := range teams {
		if t != team {
			continue
		}
		api := slack.New(tokens[i])

		stack := make([]slack.Msg, 0)
		j := 0
		for {
			j += 1000
			hp := slack.HistoryParameters{Oldest: "", Latest: tss, Count: 1000, Inclusive: false, Unreads: false}
			list, _ := api.GetIMHistory(room, hp)
			stamps := make([]string, 0)
			for _, r := range list.Messages {
				//SaveMsg(team, room, r.Msg)
				//fmt.Println(r.Msg.Timestamp)
				stack = append([]slack.Msg{r.Msg}, stack...)
				//fmt.Println(r.Msg.Text)
				//fmt.Println(r.Msg.Attachments)
				stamps = append(stamps, r.Msg.Timestamp)
			}
			if len(stamps) == 0 {
				break
			}
			tss = stamps[len(stamps)-1]
		}

		lastUser := ""
		pclass := "student"
		for _, m := range stack {
			if m.User != lastUser {
				if pclass == "student" {
					pclass = "teacher"
				} else {
					pclass = "student"
				}
				fmt.Println("</p>")
				lastUser = m.User
				fmt.Println("<p class=\"" + pclass + "\">")
			}
			fmt.Println(m.Text)
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
