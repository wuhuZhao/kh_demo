package handler

import (
	"context"
	"fmt"
	"hertz_demo/kitex_gen/biz"
	"hertz_demo/kitex_gen/biz/userservice"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
)

type UserImpl struct {
	client userservice.Client
}

type data map[string]interface{}

func NewUserImpl() *UserImpl {
	c, err := userservice.NewClient("kitex_demo", client.WithHostPorts("127.0.0.1:9999"))
	if err != nil {
		panic(fmt.Sprintf("create user client error: %v", err))
	}
	return &UserImpl{client: c} // 指定下游的ip，高级用法可以使用resolver去调用服务注册中心
}

// GetUsers: Get请求，获取内存中的user信息
func (u *UserImpl) GetUsers(ctx context.Context, c *app.RequestContext) {
	r, err := u.client.GetUsers(ctx)
	if err != nil {
		c.JSON(200, data{
			"msg":  err.Error(),
			"data": r,
			"code": -1,
		})
		return
	}
	c.JSON(200, data{
		"msg":  "success",
		"data": r,
		"code": 1,
	})

}

// Login: Post请求
func (u *UserImpl) Login(ctx context.Context, c *app.RequestContext) {
	username, password := c.PostForm("username"), c.PostForm("password")
	lr, err := u.client.Login(ctx, &biz.LoginRequest{Username: username, Password: password})
	if err != nil {
		c.JSON(200, data{
			"msg":  err.Error(),
			"data": lr.GetUserToken(),
			"code": -1,
		})
		return
	}
	if lr.GetBase().GetCode() == -1 {
		c.JSON(200, data{
			"msg":  lr.GetBase().GetMsg(),
			"data": lr,
			"code": -1,
		})
		return
	}
	c.JSON(200, data{
		"msg":  lr.GetBase().GetMsg(),
		"data": lr.GetUserToken(),
		"code": lr.GetBase().GetCode(),
	})
}

// LogOut: Post请求, 同上
func (u *UserImpl) LogOut(ctx context.Context, c *app.RequestContext) {
	username := c.PostForm("username")
	lor, err := u.client.LogOut(ctx, &biz.LogoutRequest{UserToken: username})
	if err != nil {
		c.JSON(200, data{
			"msg":  err.Error(),
			"data": "",
			"code": -1,
		})
		return
	}
	if lor.GetBase().GetCode() == -1 {
		c.JSON(200, data{
			"msg":  lor.GetBase().GetMsg(),
			"data": "",
			"code": lor.GetBase().GetCode(),
		})
		return
	}
	c.JSON(200, data{
		"msg":  lor.GetBase().GetMsg(),
		"data": "",
		"code": lor.GetBase().GetCode(),
	})
}
