package v1

import "github.com/gin-gonic/gin"

type extraApi_2 struct{}

func (*extraApi_2) Relation(c *gin.Context) {
	// 获取关系
}

func (*extraApi_2) FollowList(c *gin.Context) {
	// 关注列表
}

func (*extraApi_2) FollowerList(c *gin.Context) {
	// 粉丝列表
}
