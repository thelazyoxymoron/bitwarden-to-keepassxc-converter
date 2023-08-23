package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type BitwardenLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	TOTP     string `json:"totp"`
	URLs     []struct {
		URL string `json:"uri"`
	} `json:"uris"`
}

type BitwardenCard struct {
	CardholderName string `json:"cardholderName"`
	Brand          string `json:"brand"`
	Number         string `json:"number"`
	ExpiryMonth    string `json:"expMonth"`
	ExpiryYear     string `json:"expYear"`
	Cvv            string `json:"code"`
}

type BitwardenItem struct {
	Name  string         `json:"name"`
	Type  int            `json:"type"`
	Login BitwardenLogin `json:"login"`
	Card  BitwardenCard  `json:"card"`
	Notes string         `json:"notes"`
}

type BitwardenExport struct {
	Items []BitwardenItem `json:"items"`
}

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <jsonFilePath>")
		return
	}

	jsonFilePath := os.Args[1]
	csvFilePath := "filtered_bitwarden.csv"

	// Read the JSON file
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal the JSON data into a slice of BitwardenItem
	var bwExport BitwardenExport
	err = json.Unmarshal(jsonData, &bwExport)
	if err != nil {
		fmt.Println("Error unmarshaling json: ", err)
		return
	}

	// Create a new CSV file
	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Initialize the CSV writer
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Write the header row to the CSV file
	header := []string{"Group", "Name", "Username", "Password", "URL", "Notes", "TOTP"}
	err = csvWriter.Write(header)
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	// Write each Bitwarden item to the CSV file
	for _, item := range bwExport.Items {
		// Export secure notes
		if item.Type == 2 {
			row := []string{"Secure Note", item.Name, "", "", "", item.Notes, ""}
			err = csvWriter.Write(row)
			if err != nil {
				fmt.Println("Error writing CSV row:", err)
				return
			}
		} else if item.Type == 3 {
			// Export credit/debit cards, store everything in notes column
			row := []string{
				"Card",
				item.Name,
				"",
				"",
				"",
				"Cardholder Name: " + item.Card.CardholderName + "\n" +
					"Brand: " + item.Card.Brand + "\n" +
					"Number: " + item.Card.Number + "\n" +
					"Expiry: " + item.Card.ExpiryMonth + "-" + item.Card.ExpiryYear + "\n" +
					"CVV: " + item.Card.Cvv + "\n\n\n" + item.Notes,
				"",
			}
			err = csvWriter.Write(row)
			if err != nil {
				fmt.Println("Error writing CSV row:", err)
				return
			}
		} else {
			// In Bitwarden, an item can have multiple URLs, so we'll write each URL as a separate row
			for _, url := range item.Login.URLs {
				row := []string{
					"Website Login",
					item.Name,
					item.Login.Username,
					item.Login.Password,
					url.URL,
					item.Notes,
					item.Login.TOTP,
				}
				err = csvWriter.Write(row)
				if err != nil {
					fmt.Println("Error writing CSV row:", err)
					return
				}
			}
		}
	}

	fmt.Println("Filtered data has been written to", csvFilePath)
}
