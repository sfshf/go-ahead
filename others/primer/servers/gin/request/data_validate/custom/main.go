package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator"
)

type Booking struct {
    CheckIn time.Time   `uri:"check_in" binding:"required,bookable" time_format:"2006-01-02"`
    CheckOut time.Time  `uri:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// 自定义一个验证器
var bookable validator.Func = func(fl validator.FieldLevel) bool {
    if date, ok := fl.Field().Interface().(time.Time); ok {
        today := time.Now()
        return today.Before(date)
    }
    return false
}

func main() {

    r := gin.Default()

    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("bookable", bookable)
    }

    r.GET("/bookable/:check_in/:check_out", func(c *gin.Context) {
        var b Booking
        if err := c.ShouldBindUri(&b); err != nil {
            c.String(http.StatusBadRequest, "%s\n", err.Error())
        } else {
            c.JSON(http.StatusOK, b)
        }
    })

    r.Run(":8080")

}

// Test it with:
// curl -X GET 'http://localhost:8080/bookable/2020-03-05/2020-03-06'
// curl -X GET 'http://localhost:8080/bookable/2020-04-05/2020-03-06'
// curl -X GET 'http://localhost:8080/bookable/2020-04-05/2020-04-06'
