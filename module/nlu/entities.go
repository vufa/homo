//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package nlu

type EntitiesList struct {
	Disease string `json:"disease"`
	Food    string `json:"food"`
	Item    string `json:"item"`
	Time    string `json:"time"`
}

var entities = map[string]string{
	"disease":      "疾病",
	"food":         "食物",
	"item":         "商品",
	"time":         "时间",
	"mode":         "模式",
	"phone_number": "电话号码",
	"price":        "价格",
}
