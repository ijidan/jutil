package jutil

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"sync"
)

//验证邮箱
func CusEmail(f1 validator.FieldLevel) bool {
	value := f1.Field().String()
	//校验
	validate := validator.New()
	err := validate.Var(value, "required,email")
	if err != nil {
		return false
	}
	return true
}

//验证邮箱
func CusPassword(f1 validator.FieldLevel) bool {
	value := f1.Field().String()
	log.Println(value)
	//校验
	validate := validator.New()
	err := validate.Var(value, "required,min=6,max=18")
	if err != nil {
		return false
	}
	return true
}

//验证重复密码
func CusRptPassword(password string, rptPassword string) bool {
	if password == rptPassword {
		return true
	}
	return false
}

//验证函数
type Validate struct {
	V     *validator.Validate
	Trans ut.Translator
}

//获取实例
func NewValidate() *Validate {
	var once sync.Once
	var instance *Validate
	once.Do(func() {
		v := validator.New()
		_ = v.RegisterValidation("cusEmail", CusEmail)
		_ = v.RegisterValidation("cusPassword", CusPassword)
		//翻译注册
		uni := ut.New(zh.New())
		trans, _ := uni.GetTranslator("zh")
		_ = zhTrans.RegisterDefaultTranslations(v, trans)

		instance = &Validate{
			V:     v,
			Trans: trans,
		}
	})
	return instance
}
