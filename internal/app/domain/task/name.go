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
	// 契約内容ごとのバリデーションを追加してみる
	// これでreward.goとname.goの2箇所に契約プランごとのロジックが存在した
	switch contract.ContractType() {
	case contractDomain.FreeContract:
		if utf8.RuneCountInString(value) > 30 {
			return nil, errors.New("free契約のタスク名上限は30文字までです")
		}
	case contractDomain.LightContract:
		if utf8.RuneCountInString(value) > 50 {
			return nil, errors.New("light契約のタスク名上限は50文字までです")
		}
	}
	return name(value), nil
}

func (n name) Value() string {
	return string(n)
}
