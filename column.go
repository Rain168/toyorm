package toyorm

import (
	"reflect"
)

type Column interface {
	Column() string
}

type StrColumn string

func (d StrColumn) Column() string {
	return string(d)
}

const StarColumn StrColumn = "*"

type ColumnValue interface {
	Column
	Value() reflect.Value
}

type modelFieldValue struct {
	Field
	value reflect.Value
}

func (s *modelFieldValue) Value() reflect.Value {
	return s.value
}