package verhoeff

import (
	"fmt"
	"github.com/jmcvetta/randutil"
	"strconv"
	"testing"
)

func TestGeneration(t *testing.T) {
	// input validation
	if _, err := Generate("12345678"); err != nil {
		t.Error(`Generate("12345678") shouldn't return validation error`)
	}

	if _, err := Generate("123456789"); err == nil {
		t.Error(`Generate("123456789") should return validation error`)
	}
	// run some combinations
	for x := 0; x < 100000; x++ {
		go func() {
			num, _ := randutil.IntRange(100000000, 99999999)
			chk, err := Generate(strconv.Itoa(num))
			if err != nil {
				t.Error("Failed to generate check number for", x, "got error: ", err.Error())
			}
			signed := fmt.Sprintf("%d%s", num, chk)
			ok, err := Validate(signed)
			if err != nil {
				t.Error("Failed to validate number with check digit for", signed, "got error: ", err.Error())
			}
			if !ok {
				t.Error("Failed to validate number with check digit: ", signed, " ->  ", num)
			}
		}()
	}
}

func TestValidation(t *testing.T) {
	if _, err := Validate("123456789"); err != nil {
		t.Error(`Validate("123456789") shouldn't return validation error`)
	}

	if _, err := Validate("1234567890"); err == nil {
		t.Error(`Validate("1234567890") should return validation error`)
	}

	if ok, _ := Validate("543217"); !ok {
		t.Error(`Validate("543217") is `, ok, `should be`, true)
	}

	if ok, _ := Validate("123451"); !ok {
		t.Error(`Validate("123451") is `, ok, `should be`, true)
	}

	if ok, _ := Validate("122451"); ok {
		t.Error(`Validate("122451") is `, ok, `should be`, false)
	}

	if ok, _ := Validate("128451"); ok {
		t.Error(`Validate("128451") is `, ok, `should be`, false)
	}
}
