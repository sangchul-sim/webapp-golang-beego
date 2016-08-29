package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
)

const (
	ErrInputData    = "데이터 입력 오류"
	ErrDatabase     = "데이터베이스 동작 오류"
	ErrDupUser      = "사용자 정보가 이미 존재합니다"
	ErrNoUser       = "사용자 정보가 존재하지 않습니다"
	ErrNoDealList   = "딜리스트 정보가 존재하지 않습니다"
	ErrPass         = "잘못된 암호"
	ErrNoUserPass   = "사용자 정보가 존재하지 않거나 암호가 잘못되었습니다"
	ErrNoUserChange = "사용자 정보가 존재하지 않거나, 데이터가 변경되지 않았습니다"
	ErrInvalidUser  = "사용자 정보가 올바르지 않습니다"
	ErrOpenFile     = "파일 열기 오류"
	ErrWriteFile    = "파일 쓰기 오류"
	ErrSystem       = "운영체제 오류"
	ErrNoPage       = "잘못된 페이지 호출"
	PagePer         = 3
)

type ControllerError struct {
	StatusCode int    `json:"status"`
	AppCode    int    `json:"code"`
	StatusText string `json:"message"`
}

var (
	err404          = &ControllerError{http.StatusNotFound, 404, http.StatusText(http.StatusNotFound)}
	errInputData    = &ControllerError{http.StatusBadRequest, 10001, ErrInputData}
	errDatabase     = &ControllerError{http.StatusInternalServerError, 10002, ErrDatabase}
	errDupUser      = &ControllerError{http.StatusBadRequest, 10003, ErrDupUser}
	errNoUser       = &ControllerError{http.StatusBadRequest, 10004, ErrNoUser}
	errNoDealList   = &ControllerError{http.StatusBadRequest, 10004, ErrNoDealList}
	errPass         = &ControllerError{http.StatusBadRequest, 10005, ErrPass}
	errNoUserPass   = &ControllerError{http.StatusBadRequest, 10006, ErrNoUserPass}
	errNoUserChange = &ControllerError{http.StatusBadRequest, 10007, ErrNoUserChange}
	errInvalidUser  = &ControllerError{http.StatusBadRequest, 10008, ErrInvalidUser}
	errOpenFile     = &ControllerError{http.StatusInternalServerError, 10009, ErrOpenFile}
	errWriteFile    = &ControllerError{http.StatusInternalServerError, 10010, ErrWriteFile}
	errSystem       = &ControllerError{http.StatusInternalServerError, 10011, ErrSystem}
	errExpired      = &ControllerError{http.StatusBadRequest, 10012, "로그인 만료"}
	errPermission   = &ControllerError{http.StatusBadRequest, 10013, "권한이 없습니다"}
	errNoPage       = &ControllerError{http.StatusBadRequest, 10014, ErrNoPage}
)

type BaseController struct {
	beego.Controller
}

// 가변인자 전달
// 함수의 매개변수 갯수가 정해지지 않고 유동적으로 변하는 형태를 가변인자
// func 함수명(매개변수명 ..자료형) 리턴값_자료형 {}
// 매개변수의 자료형 앞에 ...을 붙여 가변인자로 지정한다.
// 가변인자로 받은 변수는 슬라이스 타입이므로 range 키워드를 사용하여 값을 꺼내면 된다.
// 가변인자 매개변수로 슬라이스를 보낼때는 슬라이스만 ㄴ허지 않고 뒤에 ...을 붙인다.
// ...을 붙이면 슬라이스에 들어 있는 요소를 각각 넘겨준다.
func (c *BaseController) RetError(e *ControllerError, args ...string) {
	if len(args) > 0 && args[0] == "view" {
		c.Ctx.Output.SetStatus(e.StatusCode)
		c.TplName = "error.html"
		c.Data["StatusCode"] = e.StatusCode
		c.Data["AppCode"] = e.AppCode
		c.Data["StatusText"] = e.StatusText
		c.Render()
		c.StopRun()
	} else {
		c.Ctx.Output.SetStatus(e.StatusCode)
		c.Data["json"] = e
		c.ServeJSON()
		c.StopRun()
	}
}
