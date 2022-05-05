package core

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type errors map[string]string

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func NewValidator(s interface{}) errors {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return parse(trans, s)
}

func parse(trans ut.Translator, s interface{}) errors {
	errors := errors{}
	err := validate.Struct(s)
	
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			errors[e.Field()] = e.Translate(trans) 
		}
		return errors
	}

	return nil
}