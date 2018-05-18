package ipc

import "fmt"

func CardCallBack(card IData) {
	fmt.Println("callback with", card.projectType)
}
