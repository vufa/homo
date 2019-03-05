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
	"phone_number": "电话号码",
	"price":        "价格",
}
