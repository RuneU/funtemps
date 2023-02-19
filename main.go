package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/RuneU/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string
var kel float64
var cel float64

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&kel, "K", 0.0, "temperatur i grader kelvin")
	flag.Float64Var(&cel, "C", 0.0, "temperatur i grader celsius")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
}

func formatNumber(number float64) string {
	numStr := fmt.Sprintf("%.2f", number)
	numParts := strings.Split(numStr, ".")
	intPart := numParts[0]
	intPartLen := len(intPart)

	// Remove trailing zeros after decimal point
	decPart := strings.TrimRight(numParts[1], "0")
	if decPart == "" {
		return intPart
	}

	// Format integer part with thousands separators
	if intPartLen <= 3 {
		return intPart + "." + decPart
	}

	start := intPartLen % 3
	if start == 0 {
		start = 3
	}

	var result string
	for i, digit := range intPart {
		if i == start {
			result += " "
			start += 3
		}
		result += string(digit)
	}

	return result + "." + decPart
}

func main() {

	flag.Parse()

	if out == "C" && isFlagPassed("F") {
		cel := conv.FarhenheitToCelsius(fahr)
		fmt.Printf("%.2f°F is %s°C\n", fahr, formatNumber(cel))
	}

	if out == "F" && isFlagPassed("C") {
		fahr := conv.CelsiusToFarhenheit(cel)
		fmt.Printf("%.2f°C is %s°F\n", cel, formatNumber(fahr))
	}

	if out == "K" && isFlagPassed("C") {
		kel := conv.CelsiusToKelvin(cel)
		fmt.Printf("%.2f°C is %s°K\n", cel, formatNumber(kel))
	}

	if out == "C" && isFlagPassed("K") {
		cel := conv.KelvinToCelsius(kel)
		fmt.Printf("%.2f°K is %s°C\n", kel, formatNumber(cel))
	}

	if out == "F" && isFlagPassed("K") {
		fahr := conv.KelvinToFarhenheit(kel)
		fmt.Printf("%.2f°K is %s°F\n", kel, formatNumber(fahr))
	}

	if out == "K" && isFlagPassed("F") {
		kelvin := conv.FarhenheitToKelvin(fahr)
		fmt.Printf("%.v°F is %s°K\n", (fahr), formatNumber(kelvin))
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
