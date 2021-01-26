package translations

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) (err error) {

	translations := []struct {
		tag             string
		translation     string
		override        bool
		customRegisFunc validator.RegisterTranslationsFunc
		customTransFunc validator.TranslationFunc
	}{
		{
			tag:         "required",
			translation: "Ce champ est obligatoire",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("len-string", "Ce champ doit faire une taille de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} caractère", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-string-character", "{0} caractères", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "Ce champ doit être égal à {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-elements", "Ce champ doit contenir {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("len-elements-element", "{0} element", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("len-elements-element", "{0} elements", locales.PluralRuleOther, false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("len-elements-element", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-elements", fe.Field(), c)

				default:
					t, err = ut.T("len-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("min-string", "Ce champ doit faire une taille minimum de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} caractère", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-string-character", "{0} caractères", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "Ce champ doit être égal à {1} ou plus", false); err != nil {
					return
				}

				if err = ut.Add("min-elements", "Ce champ doit contenir au moins {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("min-elements-element", "{0} element", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("min-elements-element", "{0} elements", locales.PluralRuleOther, false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("min-elements-element", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-elements", fe.Field(), c)

				default:
					t, err = ut.T("min-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("max-string", "Ce champ doit faire une taille maximum de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} caractère", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-string-character", "{0} caractères", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "Ce champ doit être égal à {1} ou moins", false); err != nil {
					return
				}

				if err = ut.Add("max-elements", "Ce champ doit contenir au maximum {1}", false); err != nil {
					return
				}
				if err = ut.AddCardinal("max-elements-element", "{0} element", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("max-elements-element", "{0} elements", locales.PluralRuleOther, false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("max-elements-element", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-elements", fe.Field(), c)

				default:
					t, err = ut.T("max-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "Ce champ n'est pas égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "Ce champ ne doit pas être égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lt-string", "Ce champ doit avoir une taille inférieure à {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} caractère", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-string-character", "{0} caractères", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "Ce champ doit être inférieur à {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-elements", "Ce champ doit contenir mois de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-elements-element", "{0} element", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lt-elements-element", "{0} elements", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "Ce champ doit être avant la date et l'heure actuelle", false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-elements-element", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-elements", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lt-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lte-string", "Ce champ doit faire une taille maximum de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} caractère", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-string-character", "{0} caractères", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "Ce champ doit faire {1} ou moins", false); err != nil {
					return
				}

				if err = ut.Add("lte-elements", "Ce champ doit contenir un maximum de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-elements-element", "{0} element", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("lte-elements-element", "{0} elements", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "Ce champ doit être avant ou pendant la date et l'heure actuelle", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-elements-element", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-elements", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("lte-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gt-string", "Ce champ doit avoir une taille supérieur à {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} caractère", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-string-character", "{0} caractères", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "Ce champ doit être supérieur à {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-elements", "Ce champ doit contenir plus de {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-elements-element", "{0} element", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gt-elements-element", "{0} elements", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "Ce champ doit être après la date et l'heure actuelle", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-elements-element", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-elements", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gt-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gte-string", "Ce champ doit faire une taille d'au moins {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} caractère", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-string-character", "{0} caractères", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "Ce champ doit être {1} ou plus", false); err != nil {
					return
				}

				if err = ut.Add("gte-elements", "Ce champ doit contenir au moins {1}", false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-elements-element", "{0} element", locales.PluralRuleOne, false); err != nil {
					return
				}

				if err = ut.AddCardinal("gte-elements-element", "{0} elements", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "Ce champ doit être après ou pendant la date et l'heure actuelle", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-string", fe.Field(), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-elements-element", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-elements", fe.Field(), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
						goto END
					}

					t, err = ut.T("gte-datetime", fe.Field())

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-number", fe.Field(), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "{0} doit être égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "{0} doit être égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0} ne doit pas être égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "{0} doit être supérieur à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "{0} doit être supérieur ou égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "{0} doit être inférieur à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "{0} doit être inférieur ou égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "{0} ne doit pas être égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "{0} doit être supérieur à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "{0} doit être supérieur ou égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "{0} doit être inférieur à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "{0} doit être inférieur ou égal à {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "{0} ne doit contenir que des caractères alphabétiques",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} ne doit contenir que des caractères alphanumériques",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "Ce champ doit être une valeur numérique valide",
			override:    false,
		},
		{
			tag:         "number",
			translation: "Ce champ doit être un nombre valid",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "Ce champ doit être une chaîne de caractères au format hexadécimal valide",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "Ce champ doit être une couleur au format HEX valide",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "Ce champ doit être une couleur au format RGB valide",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "Ce champ doit être une couleur au format RGBA valide",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "Ce champ doit être une couleur au format HSL valide",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "Ce champ doit être une couleur au format HSLA valide",
			override:    false,
		},
		{
			tag:         "email",
			translation: "Ce champ doit être une adresse email valide",
			override:    false,
		},
		{
			tag:         "url",
			translation: "Ce champ doit être une URL valide",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "Ce champ doit être une URI valide",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "Ce champ doit être une chaîne de caractères au format Base64 valide",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "Ce champ doit contenir le texte '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "Ce champ doit contenir au moins l' un des caractères suivants '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "{0} ne doit pas contenir le texte '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "{0} ne doit pas contenir l'un des caractères suivants '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "{0} ne doit pas contenir ce qui suit '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "Ce champ doit être un numéro ISBN valid",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "Ce champ doit être un numéro ISBN-10 valid",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "Ce champ doit être un numéro ISBN-13 valid",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "Ce champ doit être un UUID valid",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "Ce champ doit être un UUID version 3 valid",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "Ce champ doit être un UUID version 4 valid",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "Ce champ doit être un UUID version 5 valid",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} ne doit contenir que des caractères ascii",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} ne doit contenir que des caractères ascii affichables",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "Ce champ doit contenir des caractères multioctets",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "Ce champ doit contenir une URI data valide",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "Ce champ doit contenir des coordonnées latitude valides",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "Ce champ doit contenir des coordonnées longitudes valides",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "Ce champ doit être un numéro SSN valide",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "Ce champ doit être une adressse IPv4 valide",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "Ce champ doit être une adressse IPv6 valide",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "Ce champ doit être une adressse IP valide",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "Ce champ doit contenir une notation CIDR valide",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "Ce champ doit contenir une notation CIDR valide pour une adresse IPv4",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "Ce champ doit contenir une notation CIDR valide pour une adresse IPv6",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "Ce champ doit être une adressse TCP valide",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "Ce champ doit être une adressse IPv4 TCP valide",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "Ce champ doit être une adressse IPv6 TCP valide",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "Ce champ doit être une adressse UDP valide",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "Ce champ doit être une adressse IPv4 UDP valide",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "Ce champ doit être une adressse IPv6 UDP valide",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "Ce champ doit être une adresse IP résolvable",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "Ce champ doit être une adresse IPv4 résolvable",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "Ce champ doit être une adresse IPv6 résolvable",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "Ce champ doit être une adresse UNIX résolvable",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "Ce champ doit contenir une adresse MAC valide",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "Ce champ doit être une couleur valide",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "Ce champ doit être l'un des choix suivants [{1}]",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				s, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
		},
	}

	for _, t := range translations {

		if t.customTransFunc != nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)

		} else if t.customTransFunc != nil && t.customRegisFunc == nil {

			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)

		} else if t.customTransFunc == nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)

		} else {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {

	return func(ut ut.Translator) (err error) {

		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return

	}

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {

	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
