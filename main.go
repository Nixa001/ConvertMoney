package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	
	fmt.Print("Entrez le montant de l'argent à convertir: ")
	var amount float64
	fmt.Scanf("%f", &amount)
	fmt.Print("Entrez la devise à convertir (par exemple, EUR): ")
	var fromCurrency string
	fmt.Scanf("%s", &fromCurrency)
	
	fmt.Print("Entrez la devise à convertir (par exemple, USD): ")
	var toCurrency string
	fmt.Scanf("%s", &toCurrency)

	// URL de l'API avec les devises de départ et de destination
	apiURL := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/%s", fromCurrency)
	// Envoie du requête HTTP GET à l'API
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Erreur dans l'obtention des taux de change: %s\n", err)
		return
	}
	defer response.Body.Close()

	// Lecture du corps de la réponse
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Erreur de lecture du corps de la réponse: %s\n", err)
		return
	}

	// Parsez le taux de change de la devise de destination à partir de la réponse
	responseString := string(responseData)
	toCurrencyRateStartIndex := strings.Index(responseString, fmt.Sprintf(`"%s":`, toCurrency)) + len(toCurrency) + 3
	toCurrencyRateEndIndex := strings.Index(responseString[toCurrencyRateStartIndex:], ",")
	toCurrencyRateString := responseString[toCurrencyRateStartIndex : toCurrencyRateStartIndex+toCurrencyRateEndIndex]
	var toCurrencyRate float64
	fmt.Sscanf(toCurrencyRateString, "%f", &toCurrencyRate)

			
	convertedAmount := amount * toCurrencyRate

	
	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}

