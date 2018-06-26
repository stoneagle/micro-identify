package common

type ErrorCode int

const (
	ErrorOk     ErrorCode = 0
	ErrorParams ErrorCode = -1
	ErrorAuth   ErrorCode = -401

	ErrorCardIdentify ErrorCode = -10000
	ErrorCardDetail   ErrorCode = -10001
	ErrorFiles        ErrorCode = -10002
	ErrorMysql        ErrorCode = -10003
	ErrorRedis        ErrorCode = -10004
	ErrorServer       ErrorCode = -10500
)

type ProjectType int

const (
	ProjectCard ProjectType = iota
)

var (
	projectTypeMap = map[ProjectType]string{
		ProjectCard: "card",
	}
)
