## wrap gin handler response


gin HandlerFunc 很容易出现忘记`return`问题：
```go
func(c *gin.Context) {
    var json struct {
    Value string `json:"value" binding:"required"`
    }
    
    if err := c.Bind(&json); err != nil {
        c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
    }
    
    c.JSON(200, gin.H{
        "message": "pong",
    })
}

```

可以通过封装一个`wrap`统一处理响应和错误。附件响应可以不走`wrap`或者通过解析`Metadata`处理。

```go
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
```


## 感谢
[How to wrap route handler function gin.HandlerFunc](https://stackoverflow.com/questions/63968197/how-to-wrap-route-handler-function-gin-handlerfunc)
<br/>
[gin](https://github.com/gin-gonic/gin)
<br/>
[go-kratos](https://github.com/go-kratos/kratos)
