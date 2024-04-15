run:
	go run ./cmd/itinerary/. ./inandout/input.txt ./inandout/output.txt ./airports_lookup.csv     
	
run2:
	go run ./cmd/itinerary/. ./inandout/input2.txt ./inandout/output.txt ./airports_lookup.csv
	
missing:
	go run ./cmd/itinerary/. ./inandout/input2.txt ./inandout/output.txt ./testers/missing.csv

malformed:
	go run ./cmd/itinerary/. ./inandout/input2.txt ./inandout/output.txt ./testers/malformed.csv

malformed2:
	go run ./cmd/itinerary/. ./inandout/input2.txt ./inandout/output.txt ./testers/malformed2.csv