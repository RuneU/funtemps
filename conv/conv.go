package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(fahr float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	var Celsius = (fahr - 32) * 5 / 9

	return Celsius
}

// De andre konverteringsfunksjonene implementere her
// ...
func CelsiusToFahrenheit(cel float64) float64 {
	//Farhrenheit = Celsius*(9/5) + 32
	var fahr = cel*9/5 + 32
	return fahr
}

func FarhenheitToKelvin(fahr float64) float64 {
	var kel = (fahr-32)*5/9 + 273.15
	return kel
}

func KelvinToFarhenheit(kel float64) float64 {
	//Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
	var fahr = (kel-273.15)*9/5 + 32
	return fahr
}

func CelsiusToKelvin(cel float64) float64 {
	//Kelvin = Celsius + 273.15
	var kel = cel + 273.15
	return kel
}

func KelvinToCelsius(kel float64) float64 {
	//Celsius = Kelvin - 273.15
	var cel = kel - 273.15
	return cel
}
