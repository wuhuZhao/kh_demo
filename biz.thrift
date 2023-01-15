namespace go biz

struct BaseResponse {
    1: i32 code; // 1成功，-1失败
    2: string msg;
}


struct LoginRequest {
    1: required string username;
    2: required string password;
}


struct LoginResponse {
    1: BaseResponse base;
    2: string userToken; // token使用username代替
}


struct LogoutRequest  {
    1: required string userToken; // token使用username代替
}

struct LogOutResponse {
    1: BaseResponse base;
}


struct User {
    1: string username;
    2: string password;
    3: string email;
}


service UserService {
    LoginResponse Login(1: LoginRequest request)
    LogOutResponse LogOut(1: LogoutRequest request)
    list<User> GetUsers()
}