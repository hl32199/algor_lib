package algor_lib

import (
	"testing"
)

func TestFib1(t *testing.T) {
	var result uint
	result = Fib1(0)
	if result != 0 {
		t.Fatalf("error,result:%d,should:%d", result, 0)
	}

	result = Fib1(1)
	if result != 1 {
		t.Fatalf("error,result:%d,should:%d", result, 1)
	}

	result = Fib1(2)
	if result != 1 {
		t.Fatalf("error,result:%d,should:%d", result, 1)
	}

	result = Fib1(3)
	if result != 2 {
		t.Fatalf("error,result:%d,should:%d", result, 2)
	}

	result = Fib1(10)
	if result != 55 {
		t.Fatalf("error,result:%d,should:%d", result, 55)
	}

	result = Fib1(13)
	if result != 233 {
		t.Fatalf("error,result:%d,should:%d", result, 233)
	}
}

func TestFib2(t *testing.T) {
	var result uint
	result = Fib2(0)
	if result != 0 {
		t.Fatalf("error,result:%d,should:%d", result, 0)
	}

	result = Fib2(1)
	if result != 1 {
		t.Fatalf("error,result:%d,should:%d", result, 1)
	}

	result = Fib2(2)
	if result != 1 {
		t.Fatalf("error,result:%d,should:%d", result, 1)
	}

	result = Fib2(3)
	if result != 2 {
		t.Fatalf("error,result:%d,should:%d", result, 2)
	}

	result = Fib2(10)
	if result != 55 {
		t.Fatalf("error,result:%d,should:%d", result, 55)
	}

	result = Fib2(13)
	if result != 233 {
		t.Fatalf("error,result:%d,should:%d", result, 233)
	}
}
