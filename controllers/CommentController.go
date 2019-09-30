package controllers

import (
	"emlog-go-gin/common"
	"emlog-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//评论列表
func CommentRows(c *gin.Context) {
	gid, _ := strconv.ParseInt(c.Param("gid"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	comments := models.GetCommentList(gid, 0, page, 10)
	c.JSON(http.StatusOK, common.GetResponse(0, comments, "OK"))
}
