package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "タスク一覧を取得します",
		"data":    []string{},
	})
}

func CreateTask(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "タスクを作成します",
		"data":    gin.H{},
	})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "タスクID " + id + " を更新します",
		"data":    gin.H{},
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "タスクID " + id + " を削除します",
		"data":    gin.H{},
	})
}
