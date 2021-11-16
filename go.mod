module iotdor

go 1.17

replace github.com/yjiong/iotdor => ./

require (
	entgo.io/ent v0.9.1
	github.com/bitly/go-simplejson v0.5.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.4
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/lib/pq v1.10.4
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/goconvey v1.7.2
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/yjiong/iotdor v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa
)
