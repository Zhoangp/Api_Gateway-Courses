package http

import (
	"bytes"
	"fmt"
	"github.com/Zhoangp/Api_Gateway-Courses/services/user/pb"
	"github.com/gin-gonic/gin"
	"image"
	"strings"
)

func (hdl UserHandler) UpdateAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.MustGet("emailUser").(string)
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
		res, err := hdl.client.UpdateAvatar(c, &pb.UpdateAvatarRequest{
			FileName: "avatar-" + strings.Split(email, "@")[0], //strconv.Itoa(int(time.Now().UnixMicro())),
			Size:     fileHeader.Size,
			Content:  dataBytes,
			Folder:   folder,
			Email:    email,
			Width:    "250px",
			Height:   "250px",
		})
		if err != nil {

			fmt.Println(err)
			panic(err)
		}

		if res.Error != nil {
			panic(res.Error)
		}

		c.JSON(200, gin.H{"Url": &res.Url, "Message": "Update avatar success"})
	}
}

func getImageDemension(dataBytes []byte) (int, int, error) {
	fileBytes := bytes.NewBuffer(dataBytes)
	img, _, err := image.Decode(fileBytes)
	fmt.Println(fileBytes)
	if err != nil {
		return 0, 0, err
	}

	return img.Bounds().Dx(), img.Bounds().Dy(), nil
}
