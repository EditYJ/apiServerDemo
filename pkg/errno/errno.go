package errno

import "fmt"

// 普通的错误信息结构
type Errno struct {
	Code    int
	Message string
}

// 实现error的Error()接口
func (err Errno) Error() string {
	return err.Message
}

// 返回调试信息[Err]的结构体
type Err struct {
	Code    int
	Message string
	Err     error
}

// 返回Err对象
// 用来新建定制的错误
func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

// 添加错误信息[Message]
// 如果想对外展示更多的信息可以调用此函数
func (err *Err) Add(message string) *Err {
	err.Message += " "+message
	return err
}

// 添加格式化错误信息[Message]
// 如果想对外展示更多的信息可以调用此函数
func (err *Err) Addf(format string, args ...interface{}) *Err {
	err.Message += " "+ fmt.Sprintf(format, args...)
	return err
}

// 实现error的Error()接口
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// 判断是否是用户无法找到的错误
func IsErrUserNotFound(err error) bool {
	code, _:=DecodeErr(err)
	return code == ErrUserNotFound.Code
}

// 解码错误信息
// 用来解析定制的错误
func DecodeErr(err error) (int, string) {
	if err == nil{
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}


