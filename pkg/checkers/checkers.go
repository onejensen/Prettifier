package checkers

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
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
		fmt.Println("Input not found", err)
		os.Exit(0)
	}

	return string(input)
}

// Read_csv reads the csv file containing airport data and returns
// it as a slice of Airport structs. If the file cannot be found or
// cannot be read, the program exits with a non-zero status and prints
// an error message to stderr.
// Read_csv reads the csv file containing airport data and returns
// it as a slice of Airport structs. If the file cannot be found or
// cannot be read, the program exits with a non-zero status and prints
// an error message to stderr.
func Read_csv(csv_file string) []Airport {

	lookup, err := os.Open(csv_file)
	if err != nil {
		fmt.Println("Airport lookup not found", err)
		os.Exit(1)
	}
	defer lookup.Close()

	reader := csv.NewReader(lookup)
	header, err := reader.Read()
	if err != nil {
		return nil
	}

	airport_data := []Airport{}
	lineNumber := 1

	columnPositions := make(map[string]int)
	if len(header) < 6 {
		fmt.Println("Airport lookup malformed (Not enough columns)")
		os.Exit(0)
	}

	for i, column := range header {
		if column == "" {
			fmt.Println("Airport lookup malformed (Empty column)")
			os.Exit(0)
		}
		columnPositions[column] = i
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Airport lookup malformed", "("+err.Error()+")")
			os.Exit(0)
		}
		for _, field := range line {
			if strings.TrimSpace(field) == "" {
				fmt.Printf("Airport lookup malformed (Empty field on line %d) \n", lineNumber)
				os.Exit(0)
			}
		}
		lineNumber++

		airport := Airport{
			Name:         line[columnPositions["name"]],
			Iso_country:  line[columnPositions["iso_country"]],
			Municipality: line[columnPositions["municipality"]],
			Icao_code:    line[columnPositions["icao_code"]],
			Iata_code:    line[columnPositions["iata_code"]],
			Coordinates:  line[columnPositions["coordinates"]],
		}
		airport_data = append(airport_data, airport)
	}

	return airport_data
}
