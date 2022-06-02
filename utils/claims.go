package utils

import (
	"QingYin/global"
	"QingYin/model/system/request"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

//获取Token中的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.UserId
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UserId
	}
}
