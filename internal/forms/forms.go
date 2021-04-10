package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

// Form creates custom form struct
type Form struct {
	url.Values
	Errors errors
}

// Valid validates form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "this field cannot be blank")
		return false
	}
	return true
}

//Required check required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "this field cant be blank")
		}
	}
}

// MinLength of field
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("this field must be at least %d chars long", length))
		return false
	}
	return true
}

//IsEmail checks email validity
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}