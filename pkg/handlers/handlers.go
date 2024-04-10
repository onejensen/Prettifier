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

func Process_inputs(input_database string, csv_database []checkers.Airport) string {

	//coloured_database := input_database

	// Create regular expressions to find IATA, ICAO codes and Times
	//re := regexp.MustCompile(`(?:\*##([A-Z]{4})|\*#([A-Z]{3})|##([A-Z]{4})|#([A-Z]{3}))\b`)
	re_city := regexp.MustCompile(`\*#+[A-Z]{3,4}\b`)
	re_icao := regexp.MustCompile(`\##[A-Z]{4}\b`)
	re_iata := regexp.MustCompile(`\#[A-Z]{3}\b`)
	re_time := regexp.MustCompile(`(D|T12|T24)(\([^)]+\))`)

	// Use the regular expressions to find all the IATA and ICAO codes and times
	matches_city := re_city.FindAllString(input_database, -1)
	for _, match := range matches_city {
		city_name := replace_Municipality(match, csv_database)
		input_database = strings.ReplaceAll(input_database, match, city_name)
	}
	matches_iata := re_iata.FindAllString(input_database, -1)
	matches_icao := re_icao.FindAllString(input_database, -1)
	matches_time := re_time.FindAllString(input_database, -1)

	for _, match := range matches_iata {
		airport_name := Convert_IATAcodes(match, csv_database)
		input_database = strings.ReplaceAll(input_database, match, airport_name)
	}

	for _, match := range matches_icao {
		airport_name := Convert_ICAOcodes(match, csv_database)
		input_database = strings.ReplaceAll(input_database, match, airport_name)
	}

	for _, match := range matches_time {
		entrytime := Convert_times(match, input_database)

		input_database = strings.ReplaceAll(input_database, match, entrytime)
	}
	return trim_spaces(input_database)
}

func Convert_IATAcodes(match string, csv_database []checkers.Airport) string {
	for _, airport := range csv_database {
		if airport.Iata_code == match[1:] {
			return airport.Name
		}
	}
	return match
}

func Convert_ICAOcodes(match string, csv_database []checkers.Airport) string {
	for _, airport := range csv_database {
		if airport.Icao_code == match[2:] {
			return airport.Name
		}
	}
	return match
}
func replace_Municipality(match string, csv_database []checkers.Airport) string {
	for _, airport := range csv_database {
		if strings.TrimPrefix(match, "*##") == airport.Icao_code || strings.TrimPrefix(match, "*#") == airport.Iata_code {
			return airport.Municipality
		}
	}
	return match
}

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

func Ok_Output(output_database string, output_file string) {
	// Write the output file with BONUS "animation"
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
			fmt.Println("\n", "See you soon!")
			fmt.Println()
			break
		}
		if answer == "y" || answer == "Y" {
			fmt.Println()
			for _, letter := range output_database {
				time.Sleep(20 * time.Millisecond)
				fmt.Print(string(letter))
			}
			fmt.Println()
			time.Sleep(500 * time.Millisecond)
			break
		}
		fmt.Println("Invalid answer. Please enter Y or N.")
	}
}

func trim_spaces(text string) string {
	text = strings.ReplaceAll(text, "\r", "\n")
	text = strings.ReplaceAll(text, "\v", "\n")
	text = strings.ReplaceAll(text, "\f", "\n")
	//Bonus:Trim extra spaces
	reSpaces := regexp.MustCompile(`[ ]{2,}`)
	text = reSpaces.ReplaceAllString(text, " ")
	// Use regular expression to reduce consecutive blank lines to at most two
	re := regexp.MustCompile(`\n{3,}`)
	text = re.ReplaceAllString(text, "\n\n")

	return text
}
