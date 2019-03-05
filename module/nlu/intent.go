package nlu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/countstarlight/homo/module/com"
	"io/ioutil"
	"net/http"
	"sort"
)

//API url of homo-core nlu server
const nluURL = "http://localhost:5000/parse"
const project = "rasa"
const model = "ini"

var intents = map[string]string{
	"affirm":         "表示确定",
	"ask_name":       "询问名字",
	"deny":           "表示拒绝",
	"goodbye":        "表示道别",
	"greet":          "表达问候",
	"inform_time":    "询问时间",
	"medical":        "咨询医药",
	"thanks":         "表达感谢",
	"request_search": "请求搜索",
}

var intentList = []string{
	"affirm",
	"ask_name",
	"deny",
	"goodbye",
	"greet",
	"inform_time",
	"medical",
	"thanks",
	"request_search",
}

type IntentRankingList []struct {
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

type nluReply struct {
	Intent struct {
		Name       string  `json:"name"`
		Confidence float64 `json:"confidence"`
	} `json:"intent"`
	Entities      []interface{}     `json:"entities"`
	IntentRanking IntentRankingList `json:"intent_ranking"`
	Text          string            `json:"text"`
	Project       string            `json:"project"`
	Model         string            `json:"model"`
}

type intentRequest struct {
	Query   string `json:"q"`
	Project string `json:"project"`
	Model   string `json:"model"`
}

func (l IntentRankingList) Len() int {
	return len(l)
}

func (l IntentRankingList) Less(i, j int) bool {
	return l[i].Confidence > l[j].Confidence
}

func (l IntentRankingList) Swap(i, j int) {
	l[i],
		l[j] = l[j],
		l[i]
}

func ActionLocal(text string) ([3]string, error) {
	postM := &intentRequest{
		Query:   text,
		Project: project,
		Model:   model,
	}
	var postJson, err = json.Marshal(postM)
	if err != nil {
		return [3]string{"", ""}, err
	}
	req, err := http.NewRequest("POST", nluURL, bytes.NewBuffer(postJson))
	if err != nil {
		return [3]string{"", ""}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return [3]string{"", ""}, err
	}
	defer com.IOClose("", resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return [3]string{"", ""}, err
	}
	reply := nluReply{}
	err = json.Unmarshal(body, &reply)
	if err != nil {
		return [3]string{"", ""}, err
	}
	if !com.IfStringInArray(reply.Intent.Name, actions) {
		return [3]string{"", ""}, fmt.Errorf("意图[%s]没有对应的行为", reply.Intent.Name)
	}
	var replyMessage [3]string
	replyMessage[2], err = RunActions[reply.Intent.Name]()
	//Get intent rank
	sort.Sort(reply.IntentRanking)
	rankList := reply.IntentRanking[:3]
	replyMessage[0] = "意图分析: "
	for _, r := range rankList {
		if !com.IfStringInArray(r.Name, intentList) {
			replyMessage[0] = replyMessage[0] + fmt.Sprintf("[%s]: %.4f%% ", "未知", r.Confidence*100)
		} else {
			replyMessage[0] = replyMessage[0] + fmt.Sprintf("[%s]: %.4f%% ", intents[r.Name], r.Confidence*100)
		}
	}
	replyMessage[1] = "实体分析: "
	if len(reply.Entities) > 0 {
		for _, e := range reply.Entities {
			v, ok := e.(map[string]interface{})
			if !ok {
				return [3]string{"", ""}, fmt.Errorf("获取实体失败")
			}
			replyMessage[1] = replyMessage[1] + fmt.Sprintf("[%s]: %s ", entities[v["entity"].(string)], v["value"].(string))
		}
	} else {
		replyMessage[1] = replyMessage[1] + "无实体"
	}
	return replyMessage, err
}
