package main

type film struct {
	title string
	year  int
}

func determinePathToData(userSelectedPath *string) string {
	if userSelectedPath == nil {
		return "./data/updated/little-endian"
	}
	return "error"
}
