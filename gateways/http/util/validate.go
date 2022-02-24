package util

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/mitchellh/mapstructure"
)

type StructValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

type ErrorsMap map[string]interface{}

func NewValidate() *StructValidator {
	v := validator.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	return &StructValidator{
		validator: v,
		trans:     trans,
	}
}

// Validates a struct and return an interface with errors messages
// Is important to notice that each field returns only one error
func (s StructValidator) Validate(payload interface{}, responsePayload interface{}) error {
	err := s.validator.Struct(payload)
	if err != nil {
		errs := s.errorMap(err.(validator.ValidationErrors))

		_ = mapstructure.Decode(errs, responsePayload)

		return errors.New("validation error")
	}
	return nil
}

func (s StructValidator) errorMap(errs validator.ValidationErrors) ErrorsMap {
	errsMap := make(ErrorsMap, len(errs))

	for _, e := range errs {
		// Only saves if the field hasn't an error registered in the map
		if _, ok := errsMap[e.Field()]; !ok {
			errsMap[e.Field()] = strings.ToLower(e.Translate(s.trans))
		}
	}

	return errsMap
}
