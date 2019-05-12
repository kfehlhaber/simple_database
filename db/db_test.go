package db

import (
	"testing"
)

func TestFetch(t *testing.T) {
	Serve()
	result := Fetch("hello")

	if result != "" {
		t.Fail()
	}
}

func TestPersist(t *testing.T) {
	Serve()

	if Persist("key", "value") != "SUCCESS" {
		t.Fail()
	}
}

func TestSetFetch(t *testing.T) {
	Serve()
	Persist("key", "value")

	if Fetch("key") != "value" {
		t.Fail()
	}

	if Fetch("wrongkey") != "" {
		t.Fail()
	}
}

func BenchmarkPersistFetch(b *testing.B) {
	Serve()
	for i := 1; i < 100000; i++ {
		go Persist("again", "and again")
		go Fetch("again")
	}
}