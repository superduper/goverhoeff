package verhoeff

import (
	"fmt"
	"strconv"
)

type row [10]int

// multiplication table
type mTable [10]row

var d mTable = mTable{
	row{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	row{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
	row{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
	row{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
	row{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
	row{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
	row{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
	row{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
	row{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
	row{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
}

// permutation table
type pTable [8]row

var p pTable = pTable{
	row{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	row{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
	row{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
	row{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
	row{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
	row{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
	row{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
	row{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
}

var inv row = row{0, 4, 3, 2, 1, 5, 6, 7, 8, 9}

const (
	maxValidateNumDigits = len(p) + 1
	maxGenerateNumDigits = len(p)
)

func validateNum(num string, max int) error {
	if len(num) > max {
		return fmt.Errorf("Too many digits(%d), %d is maximum allowed digit count", len(num), max)
	}
	_, err := strconv.Atoi(num)
	return err
}

func Generate(num string) (string, error) {
	if err := validateNum(num, maxGenerateNumDigits); err != nil {
		return "", err
	}
	c := 0
	for pos, el := range reversedStringToIntArray(num) {
		c = d[c][p[((pos + 1) % 8)][el]]
	}
	return strconv.Itoa(inv[c]), nil
}

func Validate(num string) (bool, error) {
	if err := validateNum(num, maxValidateNumDigits); err != nil {
		return false, err
	}
	c := 0
	for pos, el := range reversedStringToIntArray(num) {
		c = d[c][p[(pos % 8)][el]]
	}
	return c == 0, nil
}

func reversedStringToIntArray(num string) []int {
	i := len(num)
	o := make([]int, i)
	for _, el := range num {
		i--
		o[i] = int(el - '0')
	}
	return o
}

func main() {
	num := "999999199"
	chk, err := Generate(num)
	if err != nil {
		panic(err.Error())
	}
	signed := fmt.Sprintf("%s%s", num, chk)
	ok, _ := Validate(signed)
	fmt.Printf("Generate(\"%s\") is %v and Validate(\"%s\") is %v \n", num, chk, signed, ok)

	ok, _ = Validate("123451")
	fmt.Println(`Validate("123451") is `, ok, `should be`, true)
	ok, _ = Validate("122451")
	fmt.Println(`Validate("122451") is `, ok, `should be`, false)
	ok, _ = Validate("128451")
	fmt.Println(`Validate("128451") is `, ok, `should be`, false)
}
