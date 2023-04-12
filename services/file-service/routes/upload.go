package routes

import (
	"github.com/Zhoangp/Api_Gateway-Courses/services/file-service/pb"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Upload(c *gin.Context, client pb.FileServiceClient) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	folder := c.DefaultPostForm("folder", "img")
	file, err := fileHeader.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dataBytes := make([]byte, fileHeader.Size)
	if _, err := file.Read(dataBytes); err != nil {
		panic(err)
	}
	fileName := "thumbnails-" + strconv.Itoa(int(time.Now().UnixMicro()))
	res, err := client.UploadFile(c, &pb.UploadFileRequest{
		FileName: fileName,
		Size:     fileHeader.Size,
		Content:  dataBytes,
		Folder:   folder,
	})
	if res.Error != nil {
		panic(res.Error)
	}
	c.JSON(200, &res.Url)
}
