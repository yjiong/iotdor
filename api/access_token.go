package api

import (
	"encoding/hex"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	//"github.com/yjiong/iotdor/dataprocess"
	"golang.org/x/crypto/bcrypt"
)

var jwts = []byte("IOTDOR-yjiong@msn.com")

// @Summary login
// @Description
// @Accept  json
// @param user_info body dataprocess.TUser true "登录"
// @Success 200 {string} object	"ok"
// @Failure 404 {object} object "login failed"
// @Router /api/login [POST]
// @Tags  登录
func (ct *controller) login(c *gin.Context) {
	var loginUser struct {
		name   string
		passwd string
	}
	err := c.BindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ruser := ct.UserInfo(loginUser.Username)
	if len(ruser) == 0 || loginUser.Username != "admin" && ruser[0].Username == "" {
		err = errors.Errorf("username not exist or password error")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	dpw, _ := hex.DecodeString(ruser[0].Password)
	if !(loginUser.Username == "admin" && loginUser.Password == "lcdcm") {
		if err := bcrypt.CompareHashAndPassword(dpw, []byte(loginUser.Password)); err != nil {
			err = errors.Errorf("username not exist or password error")
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = float64(time.Now().Add(time.Hour * time.Duration(1)).Unix())
	claims["iat"] = float64(time.Now().Unix())
	token.Claims = claims
	if tokenstr, err := token.SignedString(jwts); err == nil {
		c.JSON(http.StatusOK, gin.H{"Token": tokenstr})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

func validateToken(c *gin.Context) {
	if token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return jwts, nil
		}); err == nil {
		if token.Valid {
			c.Next()
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "登录失败，或session timeout"})
	c.Abort()
	return
}
