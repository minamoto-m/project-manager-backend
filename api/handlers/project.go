package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProjects(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "プロジェクト一覧を取得します",
		"data":    []string{},
	})
}

func CreateProject(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "プロジェクトを作成します",
		"data":    gin.H{},
	})
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "プロジェクトID " + id + " を更新します",
		"data":    gin.H{},
	})
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "プロジェクトID " + id + " を削除します",
		"data":    gin.H{},
	})
}
