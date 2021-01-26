package config

import (
	"reflect"
	"strings"
	"sync"

	"github.com/alcjohn/go_fullstack/translations"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/fr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var Trans ut.Translator

var _ binding.StructValidator = &DefaultValidator{}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		fr := fr.New()
		uni := ut.New(fr, fr)
		trans, _ := uni.GetTranslator("fr")
		Trans = trans
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		translations.RegisterDefaultTranslations(v.validate, trans)
		// Print JSON name on validator.FieldError.Field()
		v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
			return name
		})
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
