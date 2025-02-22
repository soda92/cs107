package main

func MovieEqual(movie1, movie2 film) bool {
	return ((movie1.title == movie2.title) && movie1.year == movie2.year)
}

func MovieGreater(movie1, movie2 film) bool {
	if movie1.title > movie2.title {
		return true
	} else if movie1.title < movie2.title {
		return false
	}
	// proceed only if movie1.title == movie2.title
	if movie1.year > movie2.year {
		return true
	}
	return false
}

func (db *imdb) BinarySearch(player string, start, end int) (int, bool) {
	if start == end {
		return 0, false
	}
	middle := start + (end-start)/2
	name, _ := db.DecodeActor(middle)
	if name == player {
		return middle, true
	}
	if name > player {
		return db.BinarySearch(player, start, middle)
	} else {
		return db.BinarySearch(player, middle+1, end)
	}
}

func (db *imdb) BinarySearchMovie(movie film, start, end int) (int, bool) {
	if start == end {
		return 0, false
	}
	middle := start + (end-start)/2
	movie_ := db.DecodeMovie(middle)
	if MovieEqual(movie_, movie) {
		return middle, true
	}
	if MovieGreater(movie_, movie) {
		return db.BinarySearchMovie(movie, start, middle)
	} else {
		return db.BinarySearchMovie(movie, middle+1, end)
	}
}
