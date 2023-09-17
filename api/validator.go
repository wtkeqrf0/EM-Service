package api

import (
	"github.com/go-playground/locales/en"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

var (
	vr  *validator.Validate
	tr  *translator.UniversalTranslator
	eng translator.Translator
)

// newValidator returns a new instance of 'validate' with sane defaults.
// Validate is designed to be thread-safe and used as a singleton instance.
// It caches information about your struct and validations,
// in essence only parsing your validation tags once per struct type.
// Using multiple instances neglects the benefit of caching.
//
// This alternative constructor adds a new tags to validate specific struct fields
// and translations.
func newValidator(itf interface{ ValidateName(string) bool }) *validator.Validate {
	if vr != nil {
		return vr
	}

	vr = validator.New()
	tr = translator.New(en.New(), en.New())
	eng, _ = tr.GetTranslator("en")

	// Is done to FieldError.Field() returned the tag.
	vr.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	_ = en_translations.RegisterDefaultTranslations(vr, eng)

	// Abbreviation for validation.
	registerValidation := func(tag string, fn func(fl validator.FieldLevel) bool) {
		if err := vr.RegisterValidation(tag, fn); err != nil {
			panic(err)
		}
	}

	// Abbreviation for translations.
	registerTranslation := func(tag string, translationTmpl string) {
		if err := vr.RegisterTranslation(tag, eng,
			func(ut translator.Translator) error { return ut.Add(tag, translationTmpl, true) },
			func(ut translator.Translator, fe validator.FieldError) (translated string) {
				translated, _ = ut.T(tag, fe.Field())
				return
			},
		); err != nil {
			panic(err)
		}
	}

	// ----------------------Validations and translations----------------------

	registerValidation("name", func(fl validator.FieldLevel) bool {
		return itf.ValidateName(fl.Field().String())
	})

	registerTranslation("name", "field {0} is not correct name")

	// -------------------------------------------------------------------------

	registerValidation("order", func(fl validator.FieldLevel) bool {
		switch fl.Field().String() {
		case "ASC", "DESC":
			return true
		}

		return false
	})

	registerTranslation("order", "field {0} can have `ASC` or `DESC` value")

	// -------------------------------------------------------------------------

	return vr
}
