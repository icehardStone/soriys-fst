package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	Controller
}

func NewUserController() *UserController {
	return &UserController{
		Controller: *NewController(),
	}
}

type TokenRest struct {
	Jwt string
}

// @Param   some_id     path    string     true        "Some ID"

// @Description user login
// @Accept  json
// @Produce  json
// @Tags User
// @Success 200 {object} Result{data=TokenRest} "ok"
// @Failure 400 {object} serror.APIError ""
// @Failure 404 {object} serror.APIError ""
// @Router /api/user/login [post]
func (u *UserController) Login(ctx *gin.Context) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	var hmacSampleSecret []byte = []byte("hmacSampleSecret")

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		log.Fatalf("generate a jwt string accur a error %s", tokenString)
		u.BadRequest(ctx)
	}

	ctx.JSON(http.StatusOK, ResultOk(&TokenRest{
		Jwt: tokenString,
	}))
}

func RegisterUser(c *gin.Engine) {
	g := c.Group("/api/user")

	controller := NewUserController()
	g.POST("login", controller.Login)
}
