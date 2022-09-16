package dto

import (
	"eCommerce/infrastructure/errs"
	"testing"

	"gotest.tools/assert"
)

func TestCreatePersonDto_IsValid(t *testing.T) {
	personFunc := func() *PersonCreateDto {
		return &PersonCreateDto{
			Name:            "John",
			Email:           "john@example.com",
			Role:            "admin",
			Password:        "password",
			PasswordConfirm: "password",
		}
	}

	p1 := personFunc()
	p2 := personFunc()
	p2.Name = ""
	p3 := personFunc()
	p3.Email = ""
	p4 := personFunc()
	p4.Role = ""
	p5 := personFunc()
	p5.Password = ""
	p5a := personFunc()
	p5a.Password = "1234567"
	p6 := personFunc()
	p6.PasswordConfirm = ""
	p6a := personFunc()
	p6a.Password = "12345678"
	p6a.PasswordConfirm = "12345679"

	tests := []struct {
		name     string
		dto      *PersonCreateDto
		expected error
	}{
		{"typical no error", p1, nil},
		{"blank name", p2, errs.E(errs.Validation, errs.Parameter("name"), "Name is required")},
		{"blank email", p3, errs.E(errs.Validation, errs.Parameter("email"), "Email is required")},
		{"blank role", p4, errs.E(errs.Validation, errs.Parameter("role"), "Role is required")},
		{"blank password", p5, errs.E(errs.Validation, errs.Parameter("password"), "Password is required")},
		{"7 char password", p5a, errs.E(errs.Validation, errs.Parameter("password"), "Password must be at least 8 characters")},
		{"blank password confirm", p6, errs.E(errs.Validation, errs.Parameter("passwordConfirm"), "Password Confirm is required")},
		{"password confirm not matched", p6a, errs.E(errs.Validation, errs.Parameter("passwordConfirm"), "Password Confirm does not match")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualError := tt.dto.IsCreateDtoValid()
			if (actualError != nil) && (tt.expected == nil) {
				t.Errorf("IsCreateDtoValid() error = %v; nil expected", actualError)
				return
			}
			assert.Equal(t, true, errs.Match(tt.expected, actualError))
		})
	}
}
