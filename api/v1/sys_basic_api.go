package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type basicApi struct{}

//相当于controller层,调用service层方法实现业务逻辑
func (b *basicApi) Feed(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/feed接口")
}

func (b *basicApi) UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/user/接口")
}

func (b *basicApi) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/user/login/接口")
}

func (b *basicApi) Register(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/user/register/接口")
}

func (b *basicApi) Publish(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/publish/action/接口")
}

func (b *basicApi) PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, "测试/publish/list/接口")
}
