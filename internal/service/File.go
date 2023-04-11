package service

import (
	"TSM-Server/internal/serializer"
	"os"
	"strconv"
)

type FileService struct {
	Path string `json:"path"`
}

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size string `json:"size"`
	Date string `json:"date"`
	Type string `json:"type"`
}

func (s *FileService) GetFileListService() serializer.Response {
	root := s.Path // 指定目录的路径
	files, err := os.ReadDir(root)
	if err != nil {
		return serializer.HandleErr(err, "获取文件列表失败")
	}
	var fileList []File
	for _, file := range files {
		fileInfo, _ := file.Info()
		if err != nil {
			// handle the error
		}
		fileList = append(fileList, File{
			Name: fileInfo.Name(),
			Size: strconv.FormatInt(fileInfo.Size(), 10),
			Date: fileInfo.ModTime().String(),
			Type: fileInfo.Mode().String(),
			Path: fileInfo.Name(),
		})
	}

	return serializer.Response{
		Status: 200,
		Msg:    "获取文件列表",
		Data:   fileList,
	}
}

func (s *FileService) DelFileService() serializer.Response {
	if err := os.RemoveAll(s.Path); err != nil {
		return serializer.HandleErr(err, "删除文件失败")
	}
	return serializer.Response{
		Msg: "删除文件成功",
	}
}

func (s *FileService) UploadFileService() serializer.Response {
	return serializer.Response{
		Msg: "上传文件",
	}
}

func (s *FileService) DownloadFileService() serializer.Response {
	return serializer.Response{
		Msg: "下载文件",
	}
}
