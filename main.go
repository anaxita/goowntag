package main

import (
	"errors"
	"log"
	"reflect"
	"strings"
)

type User struct {
	ID   int64  `anaxita:"required"`
	Name string `anaxita:"required"`
}

func main() {
	user := User{
		ID:   0,
		Name: "",
	}

	log.Println(validate(&user, "ID", "Name"))
}

func validate(v interface{}, fields ...string) error {
	var (
		rv = reflect.ValueOf(v)
		t  reflect.Type
	)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return errors.New("v must be a struct!")
	}

	t = rv.Type()

	var b strings.Builder
	b.WriteString("VALIDATION ERROR: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i).Name

		for _, field := range fields {
			if field != f {
				continue
			}

			tag := t.Field(i).Tag.Get("anaxita")
			if tag != "required" {
				continue
			}

			switch rv.Field(i).Kind() {
			case reflect.String:
				if rv.Field(i).String() == "" {
					_, _ = b.WriteString(f)
					_, _ = b.WriteString(" ")
				}
			case reflect.Int64:
				if rv.Field(i).Int() == 0 {
					_, _ = b.WriteString(f)
					_, _ = b.WriteString(" ")
				}
			}
		}
	}

	if len(b.String()) != 18 {
		b.WriteString("are required!")
		return errors.New(b.String())
	}

	return nil
}
