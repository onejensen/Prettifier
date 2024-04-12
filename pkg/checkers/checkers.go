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

// Check_args checks if the user has provided the correct number of arguments,
// and if they have provided the -h flag, it prints the usage instructions and exits.
// The function returns the input, output and lookup files as strings.
func Check_args(args []string) (string, string, string) {

	if len(args) != 4 || len(args) == 2 && args[1] == "-h" {

		// Print the usage instructions.
		fmt.Println("Itinerary usage:")
		fmt.Println("go run ./cmd/itinerary/. ./inandout/input.txt ./inandout/output.txt ./airports_lookup.csv")

		os.Exit(0)
	}

	return args[1], args[2], args[3]
}

// Read_txt reads the input file and returns its contents as a string.
// If the input file cannot be found or cannot be read, the program exits with
// a non-zero status and prints an error message to stderr.
func Read_txt(input_file string) string {

	input, err := os.ReadFile(input_file)

	if err != nil {
		fmt.Println("Input not found:", err)
		os.Exit(0)
	}

	return string(input)
}

// Read_csv reads the csv file containing airport data and returns
// it as a slice of Airport structs. If the file cannot be found or
// cannot be read, the program exits with a non-zero status and prints
// an error message to stderr.
func Read_csv(csv_file string) []Airport {

	lookup, err := os.Open(csv_file)

	if err != nil {
		fmt.Println("Airport lookup not found:", err)
		os.Exit(0)
	}

	defer lookup.Close()

	reader := csv.NewReader(lookup)

	airport_data := []Airport{}

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("csv file corrupted:", err)
			os.Exit(0)
		}

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

	return airport_data
}
