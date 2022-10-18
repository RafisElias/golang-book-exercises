package main

import "testing"

func TestAvarege(t *testing.T) {
	result := Avarege([]float64{98,93,77,82,83})
	expect := 86.60

	if result != expect {
		t.Errorf("resultado '%.2f', experado '%.2f'", result, expect)
	}
}

func TestHalf(t *testing.T) {
	resultHalf, resultIsOdd := Half(1)
	expectHalf, expectIsOdd := 0, false

	if resultHalf != expectHalf {
		t.Errorf("resultado do metade '%d', medate experada '%d'", resultHalf, expectHalf)
	}
	if resultIsOdd != expectIsOdd {
		t.Errorf("resultado do número '%v', resultado experado '%v'", resultIsOdd, expectIsOdd)
	}
}

func TestGetBiggestNumber(t *testing.T) {
	result := GetBiggestNumber(0,154,8875,251,144,12,68,20)
	expect := 8875

	if result != expect {
		t.Errorf("resultado '%d', experado '%d'", result, expect)
	}
}

func TestMakeOddGenerator(t *testing.T) {
	result := MakeOddGenerator()
	var expect uint = 1

	if result() != expect {
		t.Errorf("resultado '%d', experado '%d'", result(), expect)
	}
}

func TestFibonacci(t *testing.T) {
	result := Fibonacci(12)
	expect := 144

	if result != expect {
		t.Errorf("resultado '%d', experado '%d'", result, expect)
	}
}
