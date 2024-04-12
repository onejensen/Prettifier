package main

import (
	"itinerary/pkg/checkers"
	"itinerary/pkg/handlers"
	"os"
)

func main() {
	//Check the arguments
	input_file, output_file, csv_file := checkers.Check_args(os.Args)
	//Read the input file
	input_database := checkers.Read_txt(input_file)
	//Read the csv file and make a database
	csv_database := checkers.Read_csv(csv_file)
	//Process the input (airport codes, whitespaces, timezones)
	output_database, coloured_database := handlers.Process_inputs(input_database, csv_database)
	//Write the output file
	handlers.Ok_Output(output_database, output_file, coloured_database)
}

/*
This function is responsible for processing the input text,
looking for IATA and ICAO codes, changing them into their corresponding
airport names, looking for timezones, changing the time, looking for
excesive whitespaces and removing extra whitespaces.
Finally, it returns the output_database.
*/
