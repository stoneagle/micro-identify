package common

type ErrorCode int

const (
	ErrorOk     ErrorCode = 0
	ErrorParams ErrorCode = 1
	ErrorFiles  ErrorCode = 2
	ErrorServer ErrorCode = 500
)

type ProjectType int

const (
	ProjectCard ProjectType = iota
)
