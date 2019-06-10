//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, March 2019
//

package nlu

import (
	"github.com/countstarlight/homo/cmd/webview/config"
	"time"
)

var RunActions = map[string]func(entitiesList map[string]string) (string, error){
	"confirm":     confirm,
	"ask_name":    askName,
	"deny":        deny,
	"goodbye":     goodbye,
	"greet":       greet,
	"inform_time": informTime,
	"thanks":      thanks,
	"medical":     medical,
	"switch_mode": switchMode,
}

var actionList []string

func init() {
	actionList = make([]string, 0, len(RunActions))
	for k := range RunActions {
		actionList = append(actionList, k)
	}
}

func askName(entitiesList map[string]string) (string, error) {
	return "莫非我忘记自我介绍了？我是你的homo助理，你好呀", nil
}

func confirm(entitiesList map[string]string) (string, error) {
	return "明白", nil
}

func deny(entitiesList map[string]string) (string, error) {
	return "明白了", nil
}

func goodbye(entitiesList map[string]string) (string, error) {
	return "回头见", nil
}

func greet(entitiesList map[string]string) (string, error) {
	return "你好，我是homo，你的智能助理", nil
}

func informTime(entitiesList map[string]string) (string, error) {
	return "现在是" + time.Now().Format("2006-01-02 15:04:05"), nil
}

func medical(entitiesList map[string]string) (string, error) {
	return "[伤心]哎...希望你早日康复", nil
}

func thanks(entitiesList map[string]string) (string, error) {
	return "应该的，有事随时找我", nil
}

func switchMode(entitiesList map[string]string) (string, error) {
	if entitiesList["mode"] == "分析" || entitiesList["mode"] == "调试" {
		if config.AnalyticalMode {
			return "已经处于[分析模式]", nil
		} else {
			config.AnalyticalMode = true
			return "已进入[分析模式]", nil
		}
	} else if entitiesList["mode"] == "交互" {
		if config.SilenceMode {
			config.SilenceMode = false
			config.AnalyticalMode = false
			return "[勿扰模式]已关闭", nil
		}
		if !config.AnalyticalMode {
			return "已经处于[交互模式]", nil
		} else {
			config.AnalyticalMode = false
			return "已进入[交互模式]", nil
		}
	} else if entitiesList["mode"] == "勿扰" {
		config.SilenceMode = true
		return "已进入[勿扰模式]，再次启用语音识别请切换到[交互模式]", nil
	}
	return "无效的模式", nil
}
