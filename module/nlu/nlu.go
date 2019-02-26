package nlu

import (
	"bytes"
	"encoding/json"
	"github.com/countstarlight/homo/module/com"
	"io/ioutil"
	"net/http"
)

//API url of homo-core
const coreURL = "http://localhost:5005/conversations/default/respond"

type coreReply struct {
	RecipientID string `json:"recipient_id"`
	Text        string `json:"text"`
}

func ChatWithCore(text string) (string, error) {
	var postJson = []byte(`{"query":"` + text + `"}`)
	req, err := http.NewRequest("POST", coreURL, bytes.NewBuffer(postJson))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer com.IOClose("", resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	reply := []coreReply{}
	err = json.Unmarshal(body, &reply)
	if err != nil {
		return "", err
	}
	return reply[0].Text, err
}
