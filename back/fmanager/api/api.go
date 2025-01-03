package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/icehardstone/fmanager/serror"
)

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} serror.APIError "We need ID!!"
// @Failure 404 {object} serror.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
func GetStringByInt(c *gin.Context) {
	err := serror.APIError{}
	fmt.Println(err)
}

// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Limit"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} serror.APIError "We need ID!!"
// @Failure 404 {object} serror.APIError "Can not find ID"
// @Router /testapi/get-struct-array-by-string/{some_id} [get]
func GetStructArrayByString(c *gin.Context) {

}

type Pet3 struct {
	ID int `json:"id"`
}

const (
	Success = "操作成功"
)

type Result struct {
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func ResultOk(data any) *Result {
	return &Result{
		Data: data,
		Msg:  Success,
	}
}

// @description page arguments
type PageArgs struct {
	PageSize  int `json:"pageSize" form:"pageSize"`
	PageIndex int `json:"pageIndex" form:"pageIndex"`
}

// offset number
func (p PageArgs) OffSet() int {
	if p.PageIndex <= 0 {
		return 0
	}
	return (p.PageIndex - 1) * p.PageSize
}

// take number
func (p PageArgs) Take() int {
	return p.PageSize
}
