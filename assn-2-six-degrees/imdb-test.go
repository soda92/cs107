package main

import "fmt"

func stall() string {
	var dummy string
	fmt.Printf("[Press enter to continue]")
	// dummy = readline()
	return dummy
}

func printFill() {
	for range 5 {
		fmt.Println()
	}
	fmt.Println("    .... skipping one or more records... ")
	fmt.Println()
}

/**
 * Function: listMovies
 * --------------------
 * Assumes coming in that the specified actor/actress has appeared in
 * all of the movies populating the specified credits vector.  This routine
 * prints out the first 10 and the last 10 movies, unless there are 20 or fewer
 * movies on the specified actor's/actress's resume (in which case
 * it just prints all of them.)
 *
 * @param player the actor/actress of interest.
 * @param credits the specified actor's/actress's list of movie
 *                credits.
 */
func listMovies(player string, credits []film) {
	kNumFilesToPrint := 10

	fmt.Printf("%s has starred in %d films.\n", player, len(credits))
	if len(credits) == 0 {
		return
	}
	fmt.Println("These films are:")

	numMovies := 1

	for _, curr := range credits {
		if numMovies >= kNumFilesToPrint {
			break
		}
		movie := curr
		fmt.Printf("     %d.) %s (%d)\n", numMovies, movie.title, movie.year)
		numMovies += 1
	}

	if len(credits) > 2*kNumFilesToPrint {
		printFill()
	}
	numMovies = len(credits) - kNumFilesToPrint + 1
	for _, curr := range credits[len(credits)-kNumFilesToPrint:] {
		movie := curr
		fmt.Printf("     %d.) %s (%d)\n", numMovies, movie.title, movie.year)
		numMovies += 1
	}
	stall()
}

/**
 * Function: listCostars
 * ---------------------
 * Builds up the list of costars and then prints all these
 * costars in a format similar to that used by listMovies.
 * The STL set is used to collect actor/actress names without
 * storing duplicates.
 *
 * @param player the actor/actress of interest.
 * @param credits the list of movies that the specified actor/actress has appeared in.
 *                (No integrity checks are done, so it's the client's responsibility to make sure
 *                 the specified actor/actress really has appeared in these movies.)
 * @param db the imdb housing the specified plater, list of movies, etc.  This is passed in
 *           so that each member of each cast of each movie can be added to the specified player's
 *           set of costars.
 */
func listCostars(player string, credits []film, db imdb) {
	kNumCoStarsToPrint := 10
	costars := make(map[string][]film)
	for i := range credits {
		movie := credits[i]
		cast, _ := db.getCast(movie)

		for j := range cast {
			costar := cast[j]
			if costar != player {
				costars[costar] = append(costars[costar], movie)
			}
		}
	}

	fmt.Printf("%s has worked with %d other people.\n", player, len(costars))
	if len(costars) == 0 {
		return
	}
	fmt.Println("Those other people are:")

	numCostars := 0
	startindex := 0
	for costar, films := range costars {
		fmt.Printf("     %d.) %s", numCostars, costar)
		numCostars += 1
		if len(films) > 1 {
			fmt.Printf(" (in %d different films)", len(films))
		}
		fmt.Println()
		startindex += 1
	}

	if startindex < len(costars) {
		if len(costars) > 2*kNumCoStarsToPrint {
			printFill()
		}
		for {
			if numCostars >= len(costars)-kNumCoStarsToPrint {
				break
			}
			numCostars += 1
			startindex += 1
		}

		i := 0
		for costar, films := range costars {
			if i < startindex {
				i += 1
				continue
			}
			i += 1

			fmt.Printf("     %d.) %s", numCostars, costar)
			if len(films) > 1 {
				fmt.Printf(" (in %d different films)", len(films))
			}
			fmt.Println()
		}
	}

	stall()
}

/**
 * Function: listAllMoviesAndCostars
 * ---------------------------------
 * Pings the specified imbfile to see if the specified
 * actor/actress appears in the database (and if so, has
 * appeared in a non-zero number of films.)  If the specified
 * actor/actress is missing (or if there are no films to speak
 * of), then a polite message is printed and we return immediately.
 * Otherwise, we assume that the local vector<film> has been populated
 * with real data, and we pass the buck onto the listMovies and the
 * listCostars routines.  See the documentation for each of those functions
 * on what they do and how they work.
 *
 * @param player the name of the actor/actress of interest.  No error
 *               checking is done on the string itself.
 * @param db the imdb being queried.  The assumption is that
 *           the imdb is legitimate and has already passed its own
 *           good test.
 */
func listAllMoviesAndCostars(player string, db *imdb) {
	credits, ret := db.getCredits(&player)
	if !ret {
		fmt.Printf("We're sorry, but %s doesn't appear to be in our database.\nPerhaps someone else?\n", player)
		return
	}
	listMovies(player, credits)
	listCostars(player, credits, *db)
}

/**
 * Function: queryForActors
 * ------------------------
 * Loops indefinitely, and with each iteration prompts
 * the user for the name of an actor or actress.  An
 * empty string response will terminate the program, but
 * any other responses will prompt an attempt to list all
 * of the movie credits and the costars of the specified
 * actor/actresses.  It's possible that the actor/actresses
 * doesn't exist, but the listAllmoviesAndCostrars handles
 * that situation.
 *
 * @param db a const reference to the imdb that should
 *           queried.
 */
func queryForActors(db *imdb) {
	for {
		fmt.Print("Please enter the name of an actor or actress (or [enter] to quit): ")
		response := readline()
		if response == "" {
			return
		}
		listAllMoviesAndCostars(response, db)
	}
}

/**
 * Function: imdb_test_main
 * --------------
 * Defines the entry point for the unit testing
 * program that exercises the imdb class.  Notice
 * that the imdb constructor is called,
 */
func imdb_test_main(name string) int {
	db := NewImdb(determinePathToData(nil))
	if !db.good() {
		fmt.Println("Data directory not found!  Aborting...")
		return 1
	}
	if name != "nil" {
		listAllMoviesAndCostars(name, db)
	}
	queryForActors(db)
	return 0
}
