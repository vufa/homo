package nlu

import "time"

var actions = []string{
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

func askName() (string, error) {
	return "我叫homo丫", nil
}

func affirm() (string, error) {
	return "明白", nil
}

func deny() (string, error) {
	return "明白了", nil
}

func goodbye() (string, error) {
	return "再见，和你聊天很开心", nil
}

func greet() (string, error) {
	return "你好!我叫homo", nil
}

func informTime() (string, error) {
	return "现在是" + time.Now().Format("2006-01-02 15:04:05"), nil
}
func thanks() (string, error) {
	return "不用谢", nil
}

var RunActions = map[string]func() (string, error){
	"affirm":   affirm,
	"ask_name": askName,
	"deny":     deny,
	"goodbye":  goodbye,
	"greet":    greet,
	"inform_time": informTime,
	"thanks":   thanks,
}
