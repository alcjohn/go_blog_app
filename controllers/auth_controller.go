package controllers

import (
	"fmt"
	"net/http"

	"github.com/alcjohn/go_fullstack/utils"

	"github.com/alcjohn/go_fullstack/config"
	"github.com/alcjohn/go_fullstack/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Auth Controller
type AuthController struct{}

func (controller AuthController) SignIn(c *gin.Context) {
	utils.Render(c, http.StatusOK, "auth/signin", gin.H{
		"title": "Se connecter",
	})
}
func (controller AuthController) SignUp(c *gin.Context) {
	e, _ := c.Get("errors")
	fmt.Println(e)
	utils.Render(c, http.StatusOK, "auth/signup", gin.H{
		"title": "S'inscrire",
	})
}

type LoginForm struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type RegisterForm struct {
	Firstname string `form:"firstname" binding:"required"`
	Lastname  string `form:"lastname" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required"`
}

func loginError(c *gin.Context) {

	utils.Render(c, http.StatusBadRequest, "auth/signin", gin.H{
		"title": "Se connecter",
		"errors": map[string]string{
			"email":    "Email ou Mot de passe incorect",
			"password": "Email ou Mot de passe incorect",
		},
	})
}

// Login Method
func (controller AuthController) Login(c *gin.Context) {
	session := sessions.Default(c)
	var input LoginForm
	if err := c.ShouldBind(&input); err != nil {
		fmt.Println(err)
		loginError(c)
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		fmt.Println(err)
		loginError(c)
		return
	}
	fmt.Println(user)
	session.Set("user_id", user.ID)
	session.Save()
	c.Redirect(http.StatusFound, "/me")
}

func (controller AuthController) Register(c *gin.Context) {
	var input RegisterForm
	if e := c.ShouldBind(&input); e != nil {
		errors := map[string]string{}
		for _, err := range e.(validator.ValidationErrors) {

			errors[err.Field()] = err.Translate(config.Trans)
		}
		fmt.Println(errors)
		utils.Render(c, http.StatusBadRequest, "auth/signup", gin.H{
			"title":  "S'inscrire",
			"errors": errors,
			"input":  input,
		})
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	newUser := models.User{
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  string(password),
	}

	err := config.DB.Create(&newUser).Error
	if err != nil {
		fmt.Println(err)
		utils.Render(c, http.StatusBadRequest, "auth/signup", gin.H{
			"title": "S'inscrire",
			"input": input,
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", newUser.ID)
	session.Save()
	c.Redirect(http.StatusPermanentRedirect, "/me")

}

// Logout Method
func (controller AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("user_id", 0)
	session.Save()
	c.Redirect(http.StatusPermanentRedirect, "/")
}
