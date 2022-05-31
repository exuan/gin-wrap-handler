package main

import (
	"github.com/exuan/gin-wrap-handler/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ret struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func wrap(f func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := f(c)

		// if nil, then return ''
		if res == nil {
			res = ``
		}

		ret := Ret{
			Data: res,
		}
		code := http.StatusOK

		if err != nil {
			e := errors.FromError(err)
			// wrap http status code
			if e.Code > 0 && e.Code <= 600 {
				code = e.Code
			}

			ret.Msg = e.Msg
			ret.Code = e.Code
		}

		c.JSON(code, ret)
	}
}

func main() {
	r := gin.Default()
	r.GET("/", wrap(welcome))
	r.GET("/err", wrap(err))
	r.Run()
}

func welcome(c *gin.Context) (interface{}, error) {
	return nil, nil
}

func err(c *gin.Context) (interface{}, error) {
	return nil, errors.New(http.StatusBadRequest, "bad request")
}
