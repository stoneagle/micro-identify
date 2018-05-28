package common

type ErrorCode int

const (
	ErrorOk     ErrorCode = 0
	ErrorParams ErrorCode = 1
	ErrorFiles  ErrorCode = 2
	ErrorMysql  ErrorCode = 3
	ErrorRedis  ErrorCode = 4
	ErrorAuth   ErrorCode = 5
	ErrorServer ErrorCode = 500

	ErrorCardIdentify ErrorCode = 1000
	ErrorCardDetail   ErrorCode = 1001
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
