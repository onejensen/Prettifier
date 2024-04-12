package handlers

import (
	"bufio"
	"fmt"
	"itinerary/pkg/checkers"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

// Process_inputs takes in a string input_database and a slice of Airport
// structures csv_database. It replaces any IATA, ICAO or time codes
// within the input_database with their corresponding values from the
// csv_database. It also adds colour to any IATA or ICAO codes.
// It returns a tuple with the processed input_database and a coloured version
// of the input_database.
// Parameters:
// input_database (string): The string of user input to process
// csv_database ([]checkers.Airport): A slice of Airport structures
func Process_inputs(input_database string, csv_database []checkers.Airport) (string, string) {

	coloured_inputs := input_database
	re_city := regexp.MustCompile(`\*#+[A-Z]{3,4}\b`)
	re_icao := regexp.MustCompile(`\##[A-Z]{4}\b`)
	re_iata := regexp.MustCompile(`\#[A-Z]{3}\b`)
	re_time := regexp.MustCompile(`(D|T12|T24)(\([^)]+\))`)

	matches_city := re_city.FindAllString(input_database, -1)
	for _, match := range matches_city {
		city_name := replace_Municipality(match, csv_database)
		input_database = strings.ReplaceAll(input_database, match, city_name)
		coloured_inputs = strings.ReplaceAll(coloured_inputs, match, ColorBlue+city_name+ColorReset)
	}
	matches_iata := re_iata.FindAllString(input_database, -1)
	matches_icao := re_icao.FindAllString(input_database, -1)
	matches_time := re_time.FindAllString(input_database, -1)

	for _, match := range matches_iata {
		airport_name := Convert_IATAcodes(match, csv_database)
		input_database = strings.ReplaceAll(input_database, match, airport_name)
		coloured_inputs = strings.ReplaceAll(coloured_inputs, match, ColorBlue+airport_name+ColorReset)
	}

	for _, match := range matches_icao {
		airport_name := Convert_ICAOcodes(match, csv_database)
		input_database = strings.ReplaceAll(input_database, match, airport_name)
		coloured_inputs = strings.ReplaceAll(coloured_inputs, match, ColorBlue+airport_name+ColorReset)
	}

	for _, match := range matches_time {
		entrytime := Convert_times(match, input_database)
		input_database = strings.ReplaceAll(input_database, match, entrytime)
		coloured_inputs = strings.ReplaceAll(coloured_inputs, match, ColorRed+entrytime+ColorReset)
	}
	return trim_spaces(input_database), trim_spaces(coloured_inputs)
}

// Convert_IATAcodes takes a string that matches the pattern "#IATACODE"
// and returns the name of the airport. It searches the
// csv_database for a match and returns the airport's name if found.
// If no match is found, it returns the original string.
// This function is responsible for converting IATA codes into airport names
// based on a csv_database of airports.
// Parameters:
// match (string): The string to check for a match
// csv_database ([]checkers.Airport): A slice of Airport structs containing airport data.
// Returns:
// (string): The name of the airport if a match is found, the original string otherwise.
func Convert_IATAcodes(match string, csv_database []checkers.Airport) string {
	for _, airport := range csv_database {
		if airport.Iata_code == match[1:] {
			return airport.Name
		}
	}
	return match
}

// Convert_ICAOcodes takes a string that matches the pattern "##ICAOCODE"
// and returns the name of the airport. It searches the
// csv_database for a match and returns the airport's name if found.
// If no match is found, it returns the original string.
func Convert_ICAOcodes(match string, csv_database []checkers.Airport) string {
	for _, airport := range csv_database {
		if airport.Icao_code == match[2:] {
			return airport.Name
		}
	}
	return match
}

// replace_Municipality takes a string that matches the pattern "*##icao_code"
// or "*#iata_code" and returns the name of the airport. It searches the
// csv_database for a match and returns the airport's municipality if found.
// If no match is found, it returns the original string.
func replace_Municipality(match string, csv_database []checkers.Airport) string {
	for _, airport := range csv_database {
		if strings.TrimPrefix(match, "*##") == airport.Icao_code || strings.TrimPrefix(match, "*#") == airport.Iata_code {
			return airport.Municipality
		}
	}
	return match
}

// Convert_times takes a string containing a time in the format
// (D(2006-01-02T15:04Z07:00) or T12(2006-01-02T15:04Z07:00) or
// T24(2006-01-02T15:04Z07:00)) and returns a string with the
// formatted time.
func Convert_times(match string, input_database string) string {
	time_prefix := strings.Split(match, "(")[0]
	cleanTime := strings.TrimSuffix(strings.Split(match, "(")[1], ")")

	const layout = "2006-01-02T15:04Z07:00"
	parsed_time, err := time.Parse(layout, cleanTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return match
	}
	formattedTime := ""
	switch {
	case time_prefix == "D":
		formattedTime = parsed_time.Format("02 Jan 2006")
	case time_prefix == "T12":
		formattedTime = parsed_time.Format("03:04PM (-07:00)")
	case time_prefix == "T24":
		formattedTime = parsed_time.Format("15:04 (-07:00)")
	}
	return formattedTime
}

// Ok_Output displays a nice message to the user after the output file is
// successfully created, and asks the user if they want to print the result
// to the command line. It takes three arguments:
//
//	output_database: the result of the prettification
//	output_file: the name of the output file
//	coloured_output: the prettified version of the input, with ANSI colours
func Ok_Output(output_database string, output_file string, coloured_output string) {

	good_bye := "\n" + "Thank you for using Anywhere Holidays Prettifier Tool" + "\n\n" + "See you soon!" + "\n\n"
	output_done := "\n" + "-= Output succesfully written to ->" + output_file + "=-" + "\n\n" + "Do you want to print the result in the command line? (Y/N)" + "\n"

	err := os.WriteFile(output_file, []byte(output_database), 0644)
	for _, letter := range output_done {
		time.Sleep(25 * time.Millisecond)
		fmt.Print(string(letter))
	}
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(0)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		time.Sleep(250 * time.Millisecond)
		fmt.Print("\n", "> ")
		if !scanner.Scan() {
			break
		}
		answer := scanner.Text()

		if answer == "n" || answer == "N" {
			break
		}

		if answer == "y" || answer == "Y" {
			fmt.Println()
			for _, letter := range coloured_output {
				time.Sleep(20 * time.Millisecond)
				fmt.Print(string(letter))
			}
			fmt.Println()
			time.Sleep(500 * time.Millisecond)
			break
		}
		fmt.Println("Invalid answer. Please enter Y or N.")
	}
	for _, letter := range good_bye {
		time.Sleep(20 * time.Millisecond)
		fmt.Print(string(letter))
	}
}

// trim_spaces removes any unnecessary spaces and replaces consecutive blank
// lines with at most two blank lines.
// It replaces:
// - 2 or more spaces with a single space
// - 3 or more consecutive blank lines with 2 blank lines
func trim_spaces(text string) string {
	text = strings.ReplaceAll(text, "\r", "\n")
	text = strings.ReplaceAll(text, "\v", "\n")
	text = strings.ReplaceAll(text, "\f", "\n")

	reSpaces := regexp.MustCompile(`[ ]{2,}`)
	text = reSpaces.ReplaceAllString(text, " ")

	re := regexp.MustCompile(`\n{3,}`)
	text = re.ReplaceAllString(text, "\n\n")

	return text
}
