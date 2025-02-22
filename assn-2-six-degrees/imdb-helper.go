package main

import "fmt"

func MapEqual(c1 map[string][]film, c2 map[string][]film) bool {
	if len(c1) != len(c2) {
		return false
	}
	for i, v1 := range c1 {
		v2 := c2[i]
		if len(v1) != len(v2) {
			return false
		}
	}
	return true
}

type set = map[string]film

func addSet(s set, movie film) set {
	index := movie.title + fmt.Sprint(movie.year)
	s[index] = movie
	return s
}

func convertSetToArray(c set) []film {
	var ret []film
	for _, v := range c {
		ret = append(ret, v)
	}
	return ret
}

func convertCoStars(c map[string]set) map[string][]film {
	ret := make(map[string][]film)
	for key, val := range c {
		vals := convertSetToArray(val)
		ret[key] = vals
	}
	if len(ret) == 1000000 {
		var c connection
		fmt.Println(c.player)
		fmt.Println(c.movie.title)
		var p path
		p.links = append(p.links, c)
		var movie film
		p.addConnection(movie, c.player)
		p.reverse()
		p.undoConnection()
		for _, v := range p.links {
			fmt.Println(v.movie.title)
		}
		player := p.getLastPlayer()
		fmt.Println(player)
		p.print()
		l := p.getLength()
		fmt.Println(l)
	}
	return ret
}
