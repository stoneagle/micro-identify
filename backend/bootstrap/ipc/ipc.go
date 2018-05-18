package ipc

import (
	"identify/backend/bootstrap"
	"identify/backend/common"
	identify "identify/backend/ipc"
	"os"
)

func Configure(b *bootstrap.Bootstrapper) {
	identify.Register(identify.CardCallBack, common.ProjectCard)
	cardPath := identify.GetImagePath(b.Config.Card.Ipc.Img, common.ProjectCard)
	checkAndCreateDir(cardPath)
}

func checkAndCreateDir(path string) {
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			panic(err)
		}
	}
	if _, err := os.Stat(path); err != nil {
		panic(err)
	}
}
