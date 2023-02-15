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

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

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

/*func addThousandsSeparators(f float64) string {
	s := fmt.Sprintf("%.0f", f)
	if len(s) < 4 {
		return s
	}
	result := make([]byte, 0, len(s)+int(len(s)/3))
	start := len(s) % 3
	if start == 0 {
		start = 3
	}
	result = append(result, []byte(s[:start])...)
	for i := start; i < len(s); i += 3 {
		result = append(result, ' ')
		result = append(result, []byte(s[i:i+3])...)
	}
	return string(result)
}
*/

func addThousandsSeparators(f float64) string {
	s := fmt.Sprintf("%.2f", f)
	parts := strings.Split(s, ".")
	integerPart := parts[0]
	decimalPart := parts[1]
	if len(integerPart) < 4 {
		return s
	}
	var result string
	for i := len(integerPart) - 1; i >= 0; i-- {
		result = string(integerPart[i]) + result
		if (len(integerPart)-i)%3 == 0 && i != 0 {
			result = " " + result
		}
	}
	return result + "." + decimalPart
}

func main() {

	flag.Parse()

	// Her er noen eksempler du kan bruke i den manuelle testingen
	//fmt.Println(fahr, out, funfacts)

	//fmt.Println("len(flag.Args())", len(flag.Args()))
	//fmt.Println("flag.NFlag()", flag.NFlag())

	//fmt.Println(isFlagPassed("out"))

	if out == "C" && isFlagPassed("F") {
		cel := conv.FarhenheitToCelsius(fahr)
		fmt.Printf("%.0f°F is %s°C\n", fahr, addThousandsSeparators(cel))
	}

	if out == "F" && isFlagPassed("C") {
		fahr := conv.CelsiusToFarhenheit(cel)
		fmt.Printf("%.0f°C is %s°F\n", cel, addThousandsSeparators(fahr))
	}

	if out == "K" && isFlagPassed("C") {
		kel := conv.CelsiusToKelvin(cel)
		fmt.Printf("%.0f°C is %s°K\n", cel, addThousandsSeparators(kel))
	}

	if out == "C" && isFlagPassed("K") {
		cel := conv.KelvinToCelsius(kel)
		fmt.Printf("%.0f°K is %s°C\n", kel, addThousandsSeparators(cel))
	}

	if out == "F" && isFlagPassed("K") {
		fahr := conv.KelvinToFarhenheit(kel)
		fmt.Printf("%.0f°K is %s°F\n", kel, addThousandsSeparators(fahr))
	}

	if out == "K" && isFlagPassed("F") {
		kelvin := conv.FarhenheitToKelvin(fahr)
		fmt.Printf("%.0f°F is %s°K\n", fahr, addThousandsSeparators(kelvin))
	}
	/*if out == "K" && isFlagPassed("F") {
		kel := conv.FarhenheitToKelvin(fahr)
		fmt.Printf("%.0f°F is %s°K\n", fahr, addThousandsSeparators(kel))
	}
	*/
	/*if out == "K" && isFlagPassed("F") {
		kelvin := conv.FarhenheitToKelvin(fahr)
		fmt.Printf("%.2f°F is %.0f°K\n", fahr, kelvin)
	}
	*/

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/
	// Eksempel på enkel logikk
	//if out == "C" && isFlagPassed("F") {
	// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
	// skal returnere °C
	//	fmt.Println("0°F er -17.78°C")
	//}

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
