package dto

import (
	"eCommerce/errs"
	"time"

	"github.com/google/uuid"
)

type PersonGetDto struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type PersonCreateDto struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type PersonUpdateDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (p PersonCreateDto) IsValid() error {
	switch {
	case p.Name == "":
		return errs.E(errs.Validation, "Name is required")
	case p.Email == "":
		return errs.E(errs.Validation, "Email is required")
	case p.Role == "":
		return errs.E(errs.Validation, "Role is required")
	case p.Password == "":
		return errs.E(errs.Validation, "Password is required")
	case len(p.Password) < 8:
		return errs.E(errs.Validation, "Password must be at least 8 characters")
	case p.PasswordConfirm == "":
		return errs.E(errs.Validation, "Password Confirm is required")
	case p.Password != p.PasswordConfirm:
		return errs.E(errs.Validation, "Password Confirm does not match")
	}
	return nil
}
