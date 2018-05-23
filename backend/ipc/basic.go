package ipc

/*
#cgo LDFLAGS: -L ../cplus/lib -lipc -lIPCReading-S-r0.0.1 -lstdc++ -lDNN -lopenblas -lopencv_highgui -lopencv_imgcodecs -llibjasper -llibjpeg -llibtiff -llibpng -llibwebp -lopencv_imgproc -lopencv_core  -lpthread -lstdc++ -fopenmp -ldl -lz -lm
#include "../cplus/ipc.h"
#include <stdlib.h>
*/
import "C"
import (
	"identify/backend/common"
	"sync"
)

type IData C.struct_data

var (
	mu  sync.Mutex
	fns = make(map[common.ProjectType]func(IData))
)

func (data *IData) SetParams(modelPath, imgPath, app string, ptype common.ProjectType) {
	data.modelPath = C.CString(modelPath)
	data.imgPath = C.CString(imgPath)
	data.app = C.CString(app)
	data.projectType = C.int(ptype)
}

func (data *IData) Check() int {
	var result C.int
	switch common.ProjectType(data.projectType) {
	case common.ProjectCard:
		result = C.CardIdentify(C.struct_data(*data))
	}
	return int(result)
}

func lookup(ptype common.ProjectType) func(IData) {
	mu.Lock()
	defer mu.Unlock()
	return fns[ptype]
}

func Register(fn func(IData), ptype common.ProjectType) {
	mu.Lock()
	defer mu.Unlock()
	fns[ptype] = fn
}

func UnRegister(ptype common.ProjectType) {
	mu.Lock()
	defer mu.Unlock()
	delete(fns, ptype)
}

//export CBGateway
func CBGateway(data IData) {
	fn := lookup(common.ProjectType(data.projectType))
	fn(data)
}
