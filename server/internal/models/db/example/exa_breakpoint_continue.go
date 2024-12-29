package example

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
)

// file struct, 文件结构体
type ExaFile struct {
	request.CommonModel
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	request.CommonModel
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
