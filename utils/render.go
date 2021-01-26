package utils

import (
	"fmt"

	"github.com/alcjohn/go_fullstack/models"

	"github.com/gin-gonic/gin"
)

func merge(mp1, mp2 map[string]interface{}) map[string]interface{} {

	mp3 := make(map[string]interface{})
	for k, v := range mp1 {
		if _, ok := mp1[k]; ok {
			mp3[k] = v
		}
	}

	for k, v := range mp2 {
		if _, ok := mp2[k]; ok {
			mp3[k] = v
		}
	}
	return mp3
}
func Render(c *gin.Context, code int, name string, obj gin.H) {
	user, ok := c.Get("user")

	tugudu := gin.H{
		"user": user,
	}
	if ok {

		fmt.Println(user.(models.User))
	}
	c.HTML(code, name, merge(tugudu, obj))
}
