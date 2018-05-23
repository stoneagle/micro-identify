package image

import (
	"identify/backend/bootstrap"
	"identify/backend/common"
	"os"
)

func Configure(b *bootstrap.Bootstrapper) {
	cardPath := common.GetImagePath(b.Config.Card.Rpc.Img, common.ProjectCard)
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
