package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Мы используем пакет "math/big"  так как встроенные числовые типы данных Go (int, uint и float64)
	// имеют ограниченную точность и размер, что может привести к ошибкам округления и переполнения
	// при работе с очень большими числами.

	a := big.NewInt(1e7 + 875621)
	b := big.NewInt(1e7 - 654321)

	// сложение
	c := new(big.Int).Add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, c)

	// вычитание
	d := new(big.Int).Sub(a, b)
	fmt.Printf("%v - %v = %v\n", a, b, d)

	// умножение
	e := new(big.Int).Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, e)

	// деление
	f := new(big.Float).Quo(
		big.NewFloat(float64(a.Int64())),
		big.NewFloat(float64(b.Int64())),
	)

	fmt.Printf("%v / %v = %.2f\n", a, b, f)
}
