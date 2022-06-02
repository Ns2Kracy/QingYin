package middleware

import (
	"QingYin/model/system/response"
	"QingYin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//若请求路径为/douyin/feed则直接放行
		if ctx.Request.URL.Path == "/douyin/feed" {
			ctx.Next()
			return
		}

		//从Header请求头中获取token字符串
		token := ctx.Request.Header.Get("x-token")

		//为空则中断
		if token == "" {
			ctx.JSON(http.StatusOK, response.Status{StatusCode: ERROR, StatusMsg: "token为空"})
			ctx.Abort()
			return
		}

		j := utils.NewJWT()
		//parse Token
		claims, err := j.ParseToken(token)
		if err != nil {
			//过期中断
			if err == utils.TokenExpired {
				ctx.JSON(http.StatusOK, response.Status{StatusCode: ERROR, StatusMsg: "token过期"})
				ctx.Abort()
				return
			}
			//其他错误
			ctx.JSON(http.StatusOK, response.Status{StatusCode: ERROR, StatusMsg: err.Error()})
			ctx.Abort()
			return
		}

		//过期操作?暂不实现
		// if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
		// 	claims.ExpiresAt = time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime
		// 	newToken, _ := j.CreateTokenByOldToken(token, *claims)
		// 	newClaims, _ := j.ParseToken(newToken)
		// 	ctx.Header("new-token", newToken)
		// 	ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		// }

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
