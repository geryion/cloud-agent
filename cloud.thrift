service Cloud{
    /*
    * 用户登录接口
    * user: 用户名
    * passwd: 用户密码
    * */
    string Login(1:string user, 2:string passwd)

    /*
    * 用户注册接口
    * user: 用户名
    * passwd: 用户密码
    * passwd: 用户密码确认
    * */
    string Register(1:string user, 2:string passwd, 3:string passwd2)

    /*
    * 用户注销接口
    * user: 用户名
    * passwd: 用户密码
    * */
    string CancelLation(1:string user, 2:string passwd)

    /*
    * 用户修改密码
    * user: 用户名
    * passwd: 用户密码
    * */
    string ChangePasswd(1:string user, 2:string passwd)

    /*
    * 忘记密码接口
    * type: 验证方式
    * phone/mail/question/: 验证信息
    * confirCode: 验证码
    * */
    string ResetPasswd(1:string type, 2:string info, 3:string code)
}