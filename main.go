package main

func main() {
	config, err := LoadConfig("./")
	if err != nil {
		panic(err)
	}

	// TODO: Initialize and start the bot using the loaded configuration
}
