package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/model"
	"go-blog/utils/errmsg"
	"log"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	fileSize := fileHeader.Size

	//fileName := fileHeader.Filename

	url, code := model.UpLoadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
