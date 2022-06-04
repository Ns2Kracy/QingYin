package middleware

import (
	"QingYin/model/system/response"
	"QingYin/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestPath := ctx.Request.URL.Path

		//放行规则
		if strings.Contains(requestPath, "/douyin/feed/") ||
			strings.Contains(requestPath, "/douyin/user/register/") ||
			strings.Contains(requestPath, "/douyin/user/login/") {
			ctx.Next()
			return
		}

		//请求参数中获取token
		token_GET := ctx.Query("token")
		token_POST := ctx.Request.FormValue("token")

		token := ""
		if token_GET == "" {
			token = token_POST
		} else {
			token = token_GET
		}

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
