package task

import (
	"errors"
	contractDomain "go-ca/internal/app/domain/contract"
	"unicode/utf8"
)

type Name interface {
	Value() string
}
type name string

func NewName(value string, contract contractDomain.Contract) (Name, error) {
	if value == "" || len(value) > 50 {
		return nil, errors.New("タスク名が不正です")
	}

	if utf8.RuneCountInString(value) > contract.GetMaxTaskNameLength() {
		return nil, errors.New("契約プランのタスク名上限を超えています")
	}
	return name(value), nil
}

func (n name) Value() string {
	return string(n)
}
