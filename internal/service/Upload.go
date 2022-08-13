package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
