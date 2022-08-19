package api

import (
	"encoding/hex"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
)

func (dtr *IotdorTran) validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var nowuser string
		if token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				if user, ok := token.Claims.(jwt.MapClaims)["uname"]; ok {
					if eu, err := dtr.UserInfo(user.(string)); err == nil && eu != nil {
						nowuser = eu.Name
						dpw, _ := hex.DecodeString(eu.Passwd)
						return dpw, nil
					}
				}
				return nil, errors.New("get claims.username failed")
			}); err == nil {
			if token.Valid {
				r.Header.Add("uname", nowuser)
				next.ServeHTTP(w, r)
				return
			}
		}
		respError(http.StatusUnauthorized, w, errors.New("登录失败，或session timeout"))
		return
	})
}

func (dtr *IotdorTran) validateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if eu, err := dtr.UserInfo(r.Header.Get("uname")); err == nil && eu != nil {
			if dtr.UserAdminFlag(r.Header.Get("uname")) {
				next.ServeHTTP(w, r)
				return
			}
		}
		respError(http.StatusUnauthorized, w, errors.New("您不是管理员，木有权限"))
		return
	})
}
