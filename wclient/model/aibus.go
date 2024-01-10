package model

import (
	"strings"

	"github.com/opentdp/wechat-rest/args"
)

func AiChat(id, msg string) string {

	var err error
	var res string

	if len(args.LLM.Models) == 0 {
		return "未配置大语言模型"
	}

	// 预设模型参数
	if _, exists := UserModels[id]; !exists {
		UserModels[id] = args.LLM.Models[0]
	}
	if _, exists := MsgHistory[id]; !exists {
		MsgHistory[id] = []*HistoryItem{}
	}

	// 调用接口生成文本
	text := strings.TrimSpace(strings.TrimPrefix(msg, "/ai"))
	switch GetUserModel(id).Provider {
	case "google":
		res, err = GoogleChat(id, text)
	case "openai":
		res, err = OpenaiChat(id, text)
	default:
		res = "暂不支持此模型"
	}

	// 返回结果
	if err != nil {
		return err.Error()
	}
	return res

}

// User Models

var UserModels = make(map[string]*args.LLModel)

func SetUserModel(id string, m *args.LLModel) string {

	UserModels[id] = m
	MsgHistory[id] = []*HistoryItem{}

	return "对话模型已切换为 " + m.Name + " [" + m.Model + "]"

}

func GetUserModel(id string) *args.LLModel {

	return UserModels[id]

}

// Message History

type HistoryItem struct {
	Content string
	Role    string
}

var MsgHistory = make(map[string][]*HistoryItem)

func CountHistory(id string) int {

	return len(MsgHistory[id])

}

func ClearHistory(id string) string {

	MsgHistory[id] = []*HistoryItem{}
	return "已清空上下文"

}

func AppendHistory(id string, items ...*HistoryItem) {

	if len(MsgHistory[id]) >= args.LLM.HistoryNum {
		MsgHistory[id] = MsgHistory[id][len(items):]
	}

	MsgHistory[id] = append(MsgHistory[id], items...)

}
