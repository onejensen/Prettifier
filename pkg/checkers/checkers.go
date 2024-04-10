package checkers

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Airport struct {
	Name         string
	Iso_country  string
	Municipality string
	Icao_code    string
	Iata_code    string
	Coordinates  string
}

func Check_args(args []string) (string, string, string) {
	//Check if there's three arguments or two arguments and one of them is "-h"
	if len(args) != 4 || len(args) == 2 && args[1] == "-h" {
		os.Exit(0)
	}
	//Return the name of the input, output and csv
	return args[1], args[2], args[3]
}

// Read_txt reads the input file and returns its contents as a string.
// If the input file cannot be found or cannot be read, the program exits with
// a non-zero status and prints an error message to stderr.
func Read_txt(input_file string) string {
	// Open the input file for reading
	input, err := os.ReadFile(input_file)

	// If there was an error opening the file, print an error and exit
	if err != nil {
		fmt.Println("Input not found:", err)
		os.Exit(0)
	}

	// The file contents are returned as a byte slice. Convert it to a string
	return string(input)
}

// Read_csv reads the csv file containing airport data and returns
// it as a slice of Airport structs. If the file cannot be found or
// cannot be read, the program exits with a non-zero status and prints
// an error message to stderr.
func Read_csv(csv_file string) []Airport {

	// Open the csv file for reading
	lookup, err := os.Open(csv_file)

	// If there was an error opening the file, print an error and exit
	if err != nil {
		fmt.Println("Airport lookup not found:", err)
		os.Exit(0)
	}

	// Defer the closing of the file until after the function exits
	defer lookup.Close()

	// Create a new csv reader using the file
	reader := csv.NewReader(lookup)

	// Define a slice to store the airport data
	airport_data := []Airport{}

	// Read the csv file line by line
	for {
		// Read a line from the csv file
		line, err := reader.Read()

		// If we have reached the end of the file, break the loop
		if err == io.EOF {
			break
		}

		// If there was an error reading the file, print an error
		// and exit
		if err != nil {
			fmt.Println("csv file corrupted:", err)
			os.Exit(0)
		}

		// Create a new Airport struct and add it to the slice
		airport := Airport{
			Name:         line[0],
			Iso_country:  line[1],
			Municipality: line[2],
			Icao_code:    line[3],
			Iata_code:    line[4],
			Coordinates:  line[5],
		}
		airport_data = append(airport_data, airport)
	}

	// Return the slice of Airport structs
	return airport_data
}
