package validate_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gookit/validate"
	"github.com/stretchr/testify/assert"
)

// UserForm struct
type UserForm struct {
	Name     string    `validate:"required|minLen:7"`
	Email    string    `validate:"email" message:"email is invalid"`
	Age      int       `validate:"required|int|min:1|max:99" message:"int:age must int| min: age min value is 1"`
	CreateAt int       `validate:"min:1"`
	Safe     int       `validate:"-"`
	UpdateAt time.Time `validate:"required"`
	Code     string    `validate:"customValidator"`
	// nested struct
	ExtInfo struct {
		Homepage string `validate:"required"`
		CityName string
	}
}

// CustomValidator custom validator in the source struct.
func (f UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}

// ConfigValidation config the Validation
// eg:
// - define validate scenes
func (f UserForm) ConfigValidation(v *validate.Validation) {
	v.WithScenes(validate.SValues{
		"add":    []string{"ExtInfo.Homepage", "Name", "Code"},
		"update": []string{"ExtInfo.CityName", "Name"},
	})
}

// Messages you can custom validator error messages.
func (f UserForm) Messages() map[string]string {
	return validate.MS{
		"required":      "oh! the {field} is required",
		"Name.required": "message for special field",
	}
}

// Translates you can custom field translates.
func (f UserForm) Translates() map[string]string {
	return validate.MS{
		"Name":             "User Name",
		"Email":            "User Email",
		"ExtInfo.Homepage": "Home Page",
	}
}

func Test(t *testing.T) {
	u := &UserForm{
		Name: "inhere",
	}

	v := validate.Struct(u)
	// v := validate.New(u)
	assert.NotEqual(t, false, v.Validate(), "should not equal")
	if v.Validate() { // validate ok
		// do something ...

	} else {
		fmt.Println(v.Errors)               // all error messages
		fmt.Println(v.Errors.One())         // returns a random error message text
		fmt.Println(v.Errors.Field("Name")) // returns error messages of the field
	}
}
