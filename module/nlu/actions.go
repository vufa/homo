//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package nlu

import (
	"github.com/countstarlight/homo/cmd/webview/config"
	"time"
)

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

func askName(entitiesList map[string]string) (string, error) {
	return "我叫homo丫", nil
}

func affirm(entitiesList map[string]string) (string, error) {
	return "明白", nil
}

func deny(entitiesList map[string]string) (string, error) {
	return "明白了", nil
}

func goodbye(entitiesList map[string]string) (string, error) {
	return "再见，和你聊天很开心", nil
}

func greet(entitiesList map[string]string) (string, error) {
	return "你好!我叫homo", nil
}

func informTime(entitiesList map[string]string) (string, error) {
	return "现在是" + time.Now().Format("2006-01-02 15:04:05"), nil
}

func medical(entitiesList map[string]string) (string, error) {
	return "[伤心]哎...希望你早日康复", nil
}

func thanks(entitiesList map[string]string) (string, error) {
	return "不用谢", nil
}

func switchModel(entitiesList map[string]string) (string, error) {
	if entitiesList["model"] == "分析" || entitiesList["model"] == "调试" {
		if config.AnalyticalMode {
			return "已经处于[分析模式]", nil
		} else {
			config.AnalyticalMode = true
			return "已进入[分析模式]", nil
		}
	} else if entitiesList["model"] == "交互" {
		if !config.AnalyticalMode {
			return "已经处于[交互模式]", nil
		} else {
			config.AnalyticalMode = false
			return "已进入[交互模式]", nil
		}
	}
	return "无效的模式", nil
}

var RunActions = map[string]func(entitiesList map[string]string) (string, error){
	"affirm":       affirm,
	"ask_name":     askName,
	"deny":         deny,
	"goodbye":      goodbye,
	"greet":        greet,
	"inform_time":  informTime,
	"thanks":       thanks,
	"medical":      medical,
	"switch_model": switchModel,
}
