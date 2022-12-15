package main

func main() {}

// Erwartet zwei Zahlen m und n.
// Liefert true, falls m ein Teiler von n ist.
func Divides(m, n int) bool {
	// Hinweis: Es ist eine Lösung vorgegeben, die man auch in der Praxis verwenden würde.
	// D.h. eigentlich würde man diese Funktion nicht schreiben
	// oder sie so einfach lassen, wie hier in der Vorgabe.
	//
	// Als Übungsaufgabe ersetzen Sie diese Lösung dennoch durch eine,
	// die den Modulo-Operator nicht verwendet.

	return n%m == 0
	// Erläuterung:
	// Der Modulo-Operator % liefert den Rest der ganzzahligen Division n/m.
	// Wenn dieser Rest 0 ist, dann ist n durch m teilbar.
}

// Erwartet eine Zahl n.
// Liefert true, falls n eine Primzahl ist.
func IsPrime(n int) bool {
	// Hinweis: Prüfen Sie mit einer Schleife für jede der Zahlen i von 2 bis n-1,
	// ob n durch i teilbar ist. Wenn Sie einen Teiler finden, können Sie sofort
	// (noch innerhalb der Schleife) false liefern.
	// Läuft die Schleife durch, ohne einen Teiler zu finden, ist n eine Primzahl.
	// Die Schleife ist als Gerüst vorgegeben, die Prüfung für en vorzeitigen Abbruch
	// müssen Sie noch einbauen.
	// Außerdem müssen Sie vor der Schleife evtl. noch Sonderfälle betrachten.
	for i := 2; i < n; i++ {
		// TODO: Prüfen, ob n durch i teilbar ist und false liefern, falls ja.
	}
	return true
}

// Erwartet eine Zahl n.
// Gibt alle Primzahlen auf der Konsole aus, die kleiner als n sind.
func PrintPrimes(n int) {
	// Hinweis: Laufen Sie in einer Schleife über alle Zahlen von 2 bis n und geben
	// Sie sie aus, falls es Primzahlen sind.
	// Sie können das Schleifengerüst aus IsPrime() kopieren.
	// TODO
}

// Erwartet eine Zahl n.
// Liefert die nächstgrößere Primzahl.
// Liefert n, falls n selbst eine Primzahl ist.
func NextPrime(n int) int {
	// Hinweis: Laufen Sie mit einer Schleife ab n und zählen Sie den Schleifenzähler
	// hoch, bis Sie eine Primzahl gefunden haben.
	result := n
	// TODO: result hochzählen, bis es eine Primzahl ist.
	return result
}

// Erwartet eine Zahl n.
// Liefert den nächsten Primzahlzwilling, der größer ist als n.
// Genauer: Liefert die kleinste Zahl k, so dass:
// * k >= n
// * k ist eine Primzahl
// * k + 2 ist eine Primzahl
func NextPrimeTwin(n int) int {
	// Hinweis: Schreiben Sie eine Schleife, die bei n = NextPrime(n) beginnt und n
	// so lange hochzählt, bis sowohl n als auch n+2 Primzahlen sind.
	// TODO
	return n
}

// Erwartet eine Zahl n.
// Liefert die größte Primzahl, die echt kleiner als n ist.
// Falls es keine solche Zahl gibt, wird 0 geliefert.
func GreatestPrimeBelow(n int) int {
	// Hinweis: Schreiben Sie eine Schleife, die n so lange herunterzählt,
	// bis es eine Primzahl ist.
	// TODO
	return n
}
