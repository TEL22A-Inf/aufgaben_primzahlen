package main

import "fmt"

// Hilfsfunktion für Test. Vergleicht das Ergebnis von Divides mit der Variante aus der
// Vorgabe. So kann die Prüfung bequem für viele Werte hingeschrieben werden.
func checkDivides(m, n int) bool {
	actual := Divides(m, n)
	expected := m != 0 && n%m == 0

	if expected == actual {
		return true
	}
	fmt.Printf("Fehler: Divides(%v,%v) == %v, erwartet wird aber %v", m, n, actual, expected)
	return false
}

func ExampleDivides() {
	fmt.Println(checkDivides(3, 15))
	fmt.Println(checkDivides(3, 10))
	fmt.Println(checkDivides(2, 10))
	fmt.Println(checkDivides(7, 21))
	fmt.Println(checkDivides(15, 60))
	fmt.Println(checkDivides(-3, 60))
	fmt.Println(checkDivides(-4, 8))
	fmt.Println(checkDivides(-4, 9))
	fmt.Println(checkDivides(10, 3))
	fmt.Println(checkDivides(0, 4))
	fmt.Println(checkDivides(10, -4))
	fmt.Println(checkDivides(25, -2))
	fmt.Println(checkDivides(25, -5))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	// true
}

func ExampleIsPrime() {
	fmt.Println(IsPrime(0))
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(6))
	fmt.Println(IsPrime(7))
	fmt.Println(IsPrime(8))
	fmt.Println(IsPrime(9))
	fmt.Println(IsPrime(10))

	// Output:
	// false
	// false
	// true
	// true
	// false
	// true
	// false
	// true
	// false
	// false
	// false
}

func ExamplePrintPrimes() {
	PrintPrimes(25)

	// Output:
	// 2
	// 3
	// 5
	// 7
	// 11
	// 13
	// 17
	// 19
	// 23
}

func ExampleNextPrime() {
	fmt.Println(NextPrime(1))
	fmt.Println(NextPrime(2))
	fmt.Println(NextPrime(3))
	fmt.Println(NextPrime(4))
	fmt.Println(NextPrime(5))
	fmt.Println(NextPrime(10))
	fmt.Println(NextPrime(24))
	fmt.Println(NextPrime(30))

	// Output:
	// 2
	// 2
	// 3
	// 5
	// 5
	// 11
	// 29
	// 31
}

func ExampleNextPrimeTwin() {
	fmt.Println(NextPrimeTwin(1))
	fmt.Println(NextPrimeTwin(2))
	fmt.Println(NextPrimeTwin(3))
	fmt.Println(NextPrimeTwin(4))
	fmt.Println(NextPrimeTwin(5))
	fmt.Println(NextPrimeTwin(10))
	fmt.Println(NextPrimeTwin(30))

	// Output:
	// 3
	// 3
	// 3
	// 5
	// 5
	// 11
	// 41
}

func Example() {
	fmt.Println(GreatestPrimeBelow(10))
	fmt.Println(GreatestPrimeBelow(100))
	fmt.Println(GreatestPrimeBelow(1000))
	fmt.Println(GreatestPrimeBelow(10000))
	fmt.Println(GreatestPrimeBelow(2))

	// Output:
	// 7
	// 97
	// 997
	// 9973
	// 0
}
