package main

import (
	"encoding/binary"
	"testing"
)

func TestDecodeRecord(
	t *testing.T) {
	db := NewImdb(determinePathToData(nil))
	var num int32
	binary.Decode(db.actorFile[num:num+4], binary.LittleEndian, &num)

	player := "Ewan McGregor"
	actorIndex, found := db.BinarySearch(player, 1, 1+int(num))
	if !found {
		t.Fatalf(`%s not found in db`, player)
	}
	actorName, _ := db.DecodeActor(actorIndex)
	if actorName != player {
		t.Fatal("actor name not equal")
	}
}

func TestIndexOf(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	index := IndexOf(data, 0x02)
	if index != 1 {
		t.Fatal("index error")
	}
}

func TestCasts(t *testing.T) {
	var movie film
	movie.title = "Anno Domini"
	movie.year = 2000
	movie.offsetInMovieFile = 1841336
	db := NewImdb(determinePathToData(nil))

	casts, found := db.getCast(movie)
	if len(casts) != 6 || !found {
		t.Fatal("casts wrong")
	}
}
