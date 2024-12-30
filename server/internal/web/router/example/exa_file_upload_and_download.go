package example

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/example"
	"github.com/gin-gonic/gin"
)

func NewFileUploadAndDownload(exaFileUploadAndDownloadApi *example.FileUploadAndDownloadApi) *FileUploadAndDownloadRouter {
	return &FileUploadAndDownloadRouter{
		exaFileUploadAndDownloadApi: exaFileUploadAndDownloadApi,
	}
}

type FileUploadAndDownloadRouter struct {
	exaFileUploadAndDownloadApi *example.FileUploadAndDownloadApi
}

func (r *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := router.Group("fileUploadAndDownload")
	{
		fileUploadAndDownloadRouter.POST("upload", r.exaFileUploadAndDownloadApi.UploadFile)                                 // 上传文件
		fileUploadAndDownloadRouter.POST("getFileList", r.exaFileUploadAndDownloadApi.GetFileList)                           // 获取上传文件列表
		fileUploadAndDownloadRouter.POST("deleteFile", r.exaFileUploadAndDownloadApi.DeleteFile)                             // 删除指定文件
		fileUploadAndDownloadRouter.POST("editFileName", r.exaFileUploadAndDownloadApi.EditFileName)                         // 编辑文件名或者备注
		fileUploadAndDownloadRouter.POST("breakpointContinue", r.exaFileUploadAndDownloadApi.BreakpointContinue)             // 断点续传
		fileUploadAndDownloadRouter.GET("findFile", r.exaFileUploadAndDownloadApi.FindFile)                                  // 查询当前文件成功的切片
		fileUploadAndDownloadRouter.POST("breakpointContinueFinish", r.exaFileUploadAndDownloadApi.BreakpointContinueFinish) // 切片传输完成
		fileUploadAndDownloadRouter.POST("removeChunk", r.exaFileUploadAndDownloadApi.RemoveChunk)                           // 删除切片
		fileUploadAndDownloadRouter.POST("importURL", r.exaFileUploadAndDownloadApi.ImportURL)                               // 导入URL
	}
}
