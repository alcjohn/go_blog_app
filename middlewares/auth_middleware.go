package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alcjohn/go_fullstack/global"
	"github.com/alcjohn/go_fullstack/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		if userID == nil {
			fmt.Println("user not connected")
			return
		}
		var user models.User
		err := global.DB.First(&user, userID).Error
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("user connected : %s", user.Email))
		c.Set("user", user)
	}
}
func MustBeAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get("user")
		if ok && user.(models.User).ID != 0 {
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/signin")
		return
	}
}
