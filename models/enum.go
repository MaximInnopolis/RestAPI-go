package models

import (
	"fmt"
	"strings"
)

type UserStatus string

type UserRole string

const (
	Active  UserStatus = "ACTIVE"
	Blocked UserStatus = "BLOCKED"
)

const (
	Admin  UserRole = "ADMIN"
	Reader UserRole = "READER"
)

func (st *UserStatus) UnmarshalJSON(b []byte) error {
	status := UserStatus(strings.Trim(string(b), "\""))
	if status != Active && status != Blocked {
		return fmt.Errorf("status %s is illegal", status)
	}

	*st = status
	return nil
}

func (r *UserRole) UnmarshalJSON(b []byte) error {
	role := UserRole(strings.Trim(string(b), "\""))
	if role != Admin && role != Reader {
		return fmt.Errorf("role %s is illegal", role)
	}

	*r = role
	return nil
}
