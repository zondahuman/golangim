package handler

import (
	"encoding/json"
	"fmt"
	"golangim/im/model"
)

func ProcessHandler(param string) {
	fmt.Println("processHandler---------------", param)
	transferStruct(param)
}

func transferStruct(param string) {
	fmt.Println("transferStruct---------------", param)

	msgModel := &model.ChatMessageRequest{}
	json.Unmarshal([]byte(param), &msgModel)
	fmt.Println("msgModel %#v", msgModel)
	fmt.Println("transferStruct", "---------end")

}
