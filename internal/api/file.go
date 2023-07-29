package api

import (
	"TSM-Server/internal/server"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// 多文件
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// 上传文件到指定的路径
		if err := c.SaveUploadedFile(file, "/home/user/Downloads/"); err != nil {
			log.Println(err)
		}

	}
	msg := fmt.Sprintf("%d files uploaded!", len(files))
	c.JSON(200, gin.H{"msg": msg})
}

func DownloadFile(c *gin.Context) {
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "test.txt"))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("/home/user/Downloads/test.txt")
}

func GetFileList(c *gin.Context) {
	var fileService server.FileService
	c.ShouldBindJSON(&fileService)
	response := fileService.GetFileListService()
	c.JSON(200, response)
}

//func GetFile(c *gin.Context) {
//	fileService := server.FileService{}
//	c.ShouldBind(&fileService)
//	response := fileService.GetFileService()
//	c.JSON(200, response)
//}

func DelFile(c *gin.Context) {
	fileService := server.FileService{}
	c.ShouldBind(&fileService)
	response := fileService.DelFileService()
	c.JSON(200, response)
}
