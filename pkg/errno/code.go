//错误代码说明：
//
// 服务级别错误：1 为系统级错误；2 为普通错误，通常是由账号非法操作引起的
// 服务模块为两位数：一个大型系统的服务模块通常不超过两位数，如果超过，说明这个系统该拆分了
// 错误码为两位数：防止一个模块定制过多的错误码，后期不好维护
// code = 0 说明是正确返回，code > 0 说明是错误返回
// 错误通常包括系统级错误码和服务级错误码
// 建议代码中按服务模块将错误分类
// 错误码均为 >= 0 的数
// 在 apiserver 中 HTTP Code 固定为 http.StatusOK，错误码通过 code 来表示。
package errno

// 定义错误代码
var (
	// 一般错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "内部服务出错！"}
	ErrBind             = &Errno{Code: 10002, Message: "将请求体绑定到结构体时发生错误。请求数据结构有误"}

	ErrValidation = &Errno{Code: 20001, Message: "验证失败！"}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库发生错误！"}
	ErrToken      = &Errno{Code: 20003, Message: "在签名JSON web token的时候发生错误！"}

	// 用户操作错误
	ErrEncrypt           = &Errno{Code: 20101, Message: "加密用户密码的时候发生错误！"}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "无法找到此用户！"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "token无效！"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码错误！"}
)
