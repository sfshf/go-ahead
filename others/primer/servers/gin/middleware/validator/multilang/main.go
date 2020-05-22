package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    en2 "github.com/go-playground/locales/en"
    zh2 "github.com/go-playground/locales/zh"
    "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator"
    en_translations "github.com/go-playground/validator/translations/en"
    zh_translations "github.com/go-playground/validator/translations/zh"
)

type Person struct {
    Name string     `form:"name" validate:"required"`
    Age int         `form:"age" validate:"required,gt=10"`
    Address string  `form:"address" validate:"required"`
}

var (
    Uni *ut.UniversalTranslator
    Validate *validator.Validate
)

// 验证信息多语言化
func main() {
    Validate = validator.New()  //验证器

    zh := zh2.New()
    en := en2.New()
    Uni = ut.New(zh, en)

    r := gin.Default()
    r.GET("/testing", func(c *gin.Context) {
        locale := c.DefaultQuery("locale", "zh")
        trans, _ := Uni.GetTranslator(locale)
        switch locale {
        case "zh":
            zh_translations.RegisterDefaultTranslations(Validate, trans)
        case "en":
            en_translations.RegisterDefaultTranslations(Validate, trans)
        default:
            zh_translations.RegisterDefaultTranslations(Validate, trans)
        }
        var person Person
        if err := c.ShouldBind(&person); err != nil {
            c.String(http.StatusInternalServerError, "%v\n", err)
            c.Abort()
            return
        }
        if err := Validate.Struct(person); err != nil {
            errs := err.(validator.ValidationErrors)
            sliceErrs := []string{}
            for _, e := range errs {
                sliceErrs = append(sliceErrs, e.Translate(trans))
            }
            c.String(http.StatusInternalServerError, "%v\n", sliceErrs)
            c.Abort()
            return
        }
    })
    r.Run()

}
