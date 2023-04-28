package main

import (
	"testing"
)

type UserEntity struct {
	ID    int
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
	Seven string
	Eight string
	Nine  string
}

type UserModel struct {
	ID    int
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
	Seven string
	Eight string
	Nine  string
}

type UserDto struct {
	ID    int
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
	Seven string
	Eight string
	Nine  string
}

func pretendThatThisIsDBQuery() UserEntity {
	return UserEntity{
		ID:    1,
		One:   "One",
		Two:   "Two",
		Three: "Three",
		Four:  "Four",
		Five:  "Five",
		Six:   "Six",
		Seven: "Seven",
		Eight: "Eight",
		Nine:  "Nine",
	}
}

func mapperEntityToModelValue(e UserEntity) UserModel {
	return UserModel{
		ID:    e.ID,
		One:   e.One,
		Two:   e.Two,
		Three: e.Three,
		Four:  e.Four,
		Five:  e.Five,
		Six:   e.Six,
		Seven: e.Seven,
		Eight: e.Eight,
		Nine:  e.Nine,
	}
}

func mapperEntityToModelPointer(e *UserEntity) *UserModel {
	return &UserModel{
		ID:    e.ID,
		One:   e.One,
		Two:   e.Two,
		Three: e.Three,
		Four:  e.Four,
		Five:  e.Five,
		Six:   e.Six,
		Seven: e.Seven,
		Eight: e.Eight,
		Nine:  e.Nine,
	}
}

func mapperModelToDtoValue(m UserModel) UserDto {
	return UserDto{
		ID:    m.ID,
		One:   m.One,
		Two:   m.Two,
		Three: m.Three,
		Four:  m.Four,
		Five:  m.Five,
		Six:   m.Six,
		Seven: m.Seven,
		Eight: m.Eight,
		Nine:  m.Nine,
	}
}

func mapperModelToDtoPointer(m *UserModel) *UserDto {
	return &UserDto{
		ID:    m.ID,
		One:   m.One,
		Two:   m.Two,
		Three: m.Three,
		Four:  m.Four,
		Five:  m.Five,
		Six:   m.Six,
		Seven: m.Seven,
		Eight: m.Eight,
		Nine:  m.Nine,
	}
}

//goos: windows
//goarch: amd64
//pkg: some-benchmark/get_object_like_api_server/simple
//cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
//BenchmarkByValue
//BenchmarkByValue-12             70588234                15.88 ns/op            0B/op          0 allocs/op
//BenchmarkByPointer
//BenchmarkByPointer-12           141153536                8.477 ns/op           0B/op          0 allocs/op

func BenchmarkByValue(b *testing.B) {
	e := pretendThatThisIsDBQuery()
	for i := 0; i < b.N; i++ {
		m := mapperEntityToModelValue(e)
		mapperModelToDtoValue(m)
	}
}

func BenchmarkByPointer(b *testing.B) {
	e := pretendThatThisIsDBQuery()
	for i := 0; i < b.N; i++ {
		m := mapperEntityToModelPointer(&e)
		mapperModelToDtoPointer(m)
	}
}
