package service

var ContentTypeMap map[string]string

type Service struct {
	FileManagement
}

func New() *Service {
	ContentTypeMap = map[string]string{
		"md":   "text/markdown",
		"pdf":  "application/pdf",
		"txt":  "text/plain",
		"png":  "image/png",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
	}
	service := &Service{}
	return service
}
