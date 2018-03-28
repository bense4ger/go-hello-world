package model

// FileReference describes the shape of types which can process file content
type FileReference interface {
	ProcessContent() (string, error)
}
