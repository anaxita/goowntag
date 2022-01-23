package main

import "testing"

func BenchmarkValidate(b *testing.B) {
	b.ReportAllocs()

	user := User{
		ID:   1,
		Name: "2",
	}

	for i := 0; i < b.N; i++ {

		_ = validate(user, "ID", "Name")
	}
}

func BenchmarkValidateParallel(b *testing.B) {
	b.ReportAllocs()

	user := User{
		ID:   0,
		Name: "",
	}

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				_ = validate(user, "ID", "Name")
			}
		},
	)
}
