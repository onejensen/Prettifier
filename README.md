# Itinerary prettifier Tool

This tool processes the input data from a txt file (replacing airport codes, handling 
timezones, etc.), and then writes the output to a txt file and prints out the info on 
your command line with a bonus "animation" effect and colors.

# Instructions:

- Main.go file is located in ./cmd/itinerary/ In case you want to test it from there.

- By default, input and output files are allocated in ./inandout.

- You can run the tool just by using 'Make run' from your command line, 'Make run' is 
a shortcut that will run automatically the usage of the tool so you don't have to type it.

- There's few more options in Makefile for testing the tool, for example:
"make malformed", "make malformed2", "make missing"...
You can also modify the Makefile so it suits your own files path.

- Feel free to modify the input txt however you want to test the functionality of the tool.

# Usage:

'go run ./cmd/itinerary/. ./inandout/input.txt ./inandout/output.txt ./airports_lookup.csv'