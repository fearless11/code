package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed 包含要数据的数据源信息
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化数据文件
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}
