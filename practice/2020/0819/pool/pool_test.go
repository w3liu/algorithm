package main

import "testing"

func BenchmarkNewUser1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := new(User)
		u.Name = "111"
		u.Age = 111
		u.Phone = "222"
		_ = u
	}
}

func BenchmarkNewUser2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := pool.Get().(*User)
		u.Name = "111"
		u.Age = 111
		u.Phone = "222"
		_ = u
		pool.Put(u)
	}
}
