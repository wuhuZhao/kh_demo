package main

import (
	"context"
	biz "kitex_demo/kitex_gen/biz"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

// TableName: 指定user的表名, 具体看gorm的doc
func (*User) TableName() string {
	return "user"
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	online map[string]struct{}
	mu     sync.Mutex
	db     *gorm.DB
}

func NewUserServiceImpl() *UserServiceImpl {
	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	d.AutoMigrate(&User{})                                                      // create table
	d.Create(&User{Username: "test", Password: "123456", Email: "test@qq.com"}) // insert one
	d.Create(&User{Username: "btdc", Password: "tcc", Email: "btdc@qq.com"})    // insert one
	if err != nil {
		panic("failed to connect database")
	}
	return &UserServiceImpl{online: map[string]struct{}{}, mu: sync.Mutex{}, db: d}
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *biz.LoginRequest) (resp *biz.LoginResponse, err error) {
	// TODO: Your code here...
	resp = &biz.LoginResponse{}
	if request.GetUsername() == "" || request.GetPassword() == "" {
		resp.Base = &biz.BaseResponse{Code: -1, Msg: "username or password is empty"}
		return
	}
	user := &User{}
	s.db.Where("username = ?", request.GetUsername()).First(user)
	if user.Password == request.GetPassword() && user.Username == request.GetUsername() {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.online[request.GetUsername()] = struct{}{}
		resp.Base = &biz.BaseResponse{Code: 1, Msg: "success"}
		resp.UserToken = request.GetUsername()
		return
	}
	resp.Base = &biz.BaseResponse{Code: -1, Msg: "user not found"}
	return
}

// LogOut implements the UserServiceImpl interface.
func (s *UserServiceImpl) LogOut(ctx context.Context, request *biz.LogoutRequest) (resp *biz.LogOutResponse, err error) {
	// TODO: Your code here...
	resp = &biz.LogOutResponse{}
	if request.GetUserToken() == "" {
		resp.Base = &biz.BaseResponse{Code: -1, Msg: "username is empty"}
		return
	}
	if _, ok := s.online[request.GetUserToken()]; ok {
		s.mu.Lock()
		defer s.mu.Unlock()
		delete(s.online, request.GetUserToken())
		resp.Base = &biz.BaseResponse{Code: 1, Msg: "success"}
		return
	}
	resp.Base = &biz.BaseResponse{Code: -1, Msg: "user is not online"}
	return
}

// GetUsers implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUsers(ctx context.Context) (resp []*biz.User, err error) {
	// TODO: Your code here...
	resp = []*biz.User{}
	var users []User
	d := s.db.Find(&users)
	if d.RowsAffected == 0 {
		return
	}
	for i := int64(0); i < d.RowsAffected; i++ {
		resp = append(resp, &biz.User{Username: users[i].Username, Password: users[i].Password, Email: users[i].Email})
	}
	return
}
