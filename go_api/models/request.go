package models

type FileStatus string

const (
	Upload FileStatus = "Uploaded"
	Local  FileStatus = "Local Only"
)

// Requisition structure
type UploadBody struct {
	SaveFile bool `json:"save"`
}

// Response structure
type SimpleResponse struct {
	Message string `json:"message"`
}

type UploadBeastResponse struct {
	FileStatus  FileStatus   `json:"fileStatus"`
	FileContent BeastMonster `json:"fileContent"`
}

type UploadHumanResponse struct {
	FileStatus  FileStatus   `json:"fileStatus"`
	FileContent HumanMonster `json:"fileContent"`
}

// Upload Thing API
type PrepareUploadPayload struct {
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
	FileType string `json:"fileType"`
}

type PrepareUploadResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	URL     string `json:"url"`
}

// type UploadFile struct {
// 	Name string      `json:"name"`
// 	Size int         `json:"size"`
// 	Data interface{} `json:"data"`
// }

// type UploadRequest struct {
// 	Files              []File `json:"files"`
// 	ACL                string `json:"acl"`
// 	ContentDisposition string `json:"contentDisposition"`
// }
