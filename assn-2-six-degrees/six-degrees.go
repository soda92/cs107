package main

import (
	"fmt"
	"os"
)

/**
 * Using the specified prompt, requests that the user supply
 * the name of an actor or actress.  The code returns
 * once the user has supplied a name for which some record within
 * the referenced imdb existsif (or if the user just hits return,
 * which is a signal that the empty string should just be returned.)
 *
 * @param prompt the text that should be used for the meaningful
 *               part of the user prompt.
 * @param db a reference to the imdb which can be used to confirm
 *           that a user's response is a legitimate one.
 * @return the name of the user-supplied actor or actress, or the
 *         empty string.
 */
func promptForActor(prompt string, db *imdb) string {
	var response string
	for {
		fmt.Printf("%s [or <enter> to quit]: ", prompt)
		response = readline()
		if response == "" {
			return ""
		}
		_, ret := db.getCredits(&response)
		if ret {
			return response
		}
		fmt.Printf("We coun't find %s in the movie database. Please try again.\n", response)
	}
}

/**
 * Serves as the main entry point for the six-degrees executable.
 * There are no parameters to speak of.
 *
 * @param argc the number of tokens passed to the command line to
 *             invoke this executable.  It's completely ignored
 *             here, because we don't expect any arguments.
 * @param argv the C strings making up the full command line.
 *             We expect argv[0] to be logically equivalent to
 *             "six-degrees" (or whatever absolute path was used to
 *             invoke the program), but otherwise these are ignored
 *             as well.
 * @return 0 if the program ends normally, and undefined otherwise.
 */
func six_dg_main(userSelectedPath *string) {
	db := NewImdb(
		determinePathToData(userSelectedPath))

	if !db.good() {
		fmt.Println("Failed to properly initialize the imdb database.")
		fmt.Println("Please check to make sure the source file exist and that you have permission to read them.")
		os.Exit(1)
	}

	for {
		source := promptForActor("Actor or actress", db)
		if source == "" {
			break
		}

		target := promptForActor("Another actor or actress", db)
		if target == "" {
			break
		}

		if source == target {
			fmt.Println("Good one. This is only interesting if you specify two different people.")
		} else {
			// replace the following line by a call to your generateShortestPath routine...
			fmt.Println()
			fmt.Println("No path between those two people cound be found.")
		}
	}
	fmt.Println("Thanks for playing!")
}
