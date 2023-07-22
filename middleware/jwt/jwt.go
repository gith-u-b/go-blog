package jwt

import (
	"github.com/gin-gonic/gin"
	"go-blog-step-by-step/pkg/e"
	"go-blog-step-by-step/pkg/util"
	"log"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context){
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")

		if token == "" {
			token = c.Request.Header.Get("jwt-token")
		}

		if token == ""{
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)

			if err != nil {
				log.Printf("%+v", err)
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}