package task

import "errors"

type Name interface {
	Value() string
}
type name string

func NewName(value string) (Name, error) {
	if value == "" || len(value) > 50 {
		return nil, errors.New("タスク名が不正です")
	}
	return name(value), nil
}

func (n name) Value() string {
	return string(n)
}
