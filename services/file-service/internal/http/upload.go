package http

import (
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/config"
	"github.com/Zhoangp/Api_Gateway-Courses/services/file-service/pb"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
	"time"
)

type fileHandler struct {
	cf     *config.Config
	client pb.FileServiceClient
}

func NewFileHandler(cf *config.Config, client pb.FileServiceClient) *fileHandler {
	return &fileHandler{cf: cf, client: client}
}

func (hdl fileHandler) UploadAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		res, err := hdl.client.UploadAvatar(c, &pb.UploadAvatarRequest{
			File: &pb.File{
				FileName: fileName,
				Size:     fileHeader.Size,
				Content:  dataBytes,
				Folder:   folder,
			},
		})
		if res.Error != nil {
			panic(res.Error)
		}
		c.JSON(200, &res.Url)
	}
}
func (hdl fileHandler) UploadAsset() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
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

		res, err := hdl.client.UploadAsset(c)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		go func() {
			res.Send(&pb.UploadAssetRequest{
				File: &pb.File{
					FileName: strconv.Itoa(int(time.Now().UnixMicro())),
					Size:     fileHeader.Size,
					Folder:   folder,
				},
			})
			chunkSize := 1000000
			for i := 0; i < len(dataBytes); i += chunkSize {
				end := i + chunkSize
				if end > len(dataBytes) {
					end = len(dataBytes)
				}
				if err = res.Send(&pb.UploadAssetRequest{
					File: &pb.File{
						Content: dataBytes[i:end],
					},
				}); err != nil {
					log.Fatal(err)
				}
			}

			res.CloseSend()
		}()
		reply, err := res.Recv()
		if err != nil {
			if err == io.EOF {
			}
		}

		c.JSON(200, reply)
	}
}
