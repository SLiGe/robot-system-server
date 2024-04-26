package service

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"mime"
	"mime/multipart"
	"os"
	v1 "robot-system-server/api/v1"
)

type FileService interface {
	UploadFile(c *gin.Context) (*minio.UploadInfo, error)
	Img(c *gin.Context) error
}

type fileService struct {
	*Service
	minioClient *minio.Client
}

func NewFileService(service *Service, minioClient *minio.Client) FileService {
	return &fileService{
		Service:     service,
		minioClient: minioClient,
	}
}

func (s *fileService) UploadFile(c *gin.Context) (*minio.UploadInfo, error) {
	bucketName := c.PostForm("bucket")
	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}
	open, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func(open multipart.File) {
		_ = open.Close()
	}(open)
	// 获取文件扩展名
	fileExt := getFileExtension(file.Filename)
	// 通过扩展名获取ContentType
	contentType := mime.TypeByExtension(fileExt)
	info, err := s.minioClient.PutObject(c, bucketName, file.Filename, open, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (s *fileService) Img(c *gin.Context) error {
	bucket := c.Param("bucket")
	file := c.Param("file")
	if bucket == "" || file == "" {
		return v1.ErrNotFound
	}
	imageObject, err := s.minioClient.GetObject(c, bucket, file, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	objInfo, err := imageObject.Stat()
	if err != nil {
		return err
	}
	// 设置响应头部为图像类型
	c.Header("Content-Type", objInfo.ContentType)

	// 将图像内容输出到响应
	_, err = io.Copy(c.Writer, imageObject)
	if err != nil {
		return err
	}
	return nil
}

// 获取文件扩展名
func getFileExtension(fileName string) string {
	ext := ""
	if dotIndex := lastDotIndex(fileName); dotIndex >= 0 {
		ext = fileName[dotIndex:]
	}
	return ext
}

// 获取最后一个点的索引
func lastDotIndex(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if os.IsPathSeparator(s[i]) {
			break
		}
		if s[i] == '.' {
			return i
		}
	}
	return -1
}
