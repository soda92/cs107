package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"strings"
)

func (db *imdb) DecodeActor(index int) (string, []int32) {
	index = index * 4
	var firstAddr int32
	binary.Decode(db.actorFile[index:index+4], binary.LittleEndian, &firstAddr)
	var nextAddr int32
	binary.Decode(db.actorFile[index+4:index+8], binary.LittleEndian, &nextAddr)

	if firstAddr <= nextAddr {
		log.Fatal("index error")
	}

	name := string(db.actorFile[firstAddr:nextAddr])
	len1 := strings.IndexByte(name, 0x00)
	totalLen := len1
	name1 := name[:len1]
	rest := name[len1:]
	if len(name1)%2 != 0 {
		rest = rest[1:]
		totalLen += 1
	}

	reader2 := bytes.NewReader([]byte(rest))
	var numMovies int16
	binary.Read(reader2, binary.NativeEndian, &numMovies)

	totalLen += 2
	if totalLen%4 != 0 {
		// we already know it's a multiple of 2, so we just pad 2 bytes to get to a 4-align
		reader2.Seek(2, 1) // 1: io.SeekCurrent
		totalLen += 2
	}

	movieIndexes := make([]int32, numMovies)
	for i := 0; i < int(numMovies); i++ {
		binary.Read(reader2, binary.NativeEndian, &movieIndexes[i])
	}
	// fmt.Println(name1)
	return name1, movieIndexes
}

func (db *imdb) DecodeMovie(index int) film {
	index = index * 4
	var firstAddr int32
	binary.Decode(db.movieFile[index:index+4], binary.LittleEndian, &firstAddr)
	var nextAddr int32
	binary.Decode(db.movieFile[index+4:index+8], binary.LittleEndian, &nextAddr)
	if firstAddr <= nextAddr {
		log.Fatal("index error")
	}

	name := string(db.movieFile[firstAddr:nextAddr])
	len1 := strings.IndexByte(name, 0x00)
	totalLen := len1
	name1 := name[:len1]
	rest := name[len1:]
	if len(name1)%2 != 0 {
		rest = rest[1:]
		totalLen += 1
	}

	reader2 := bytes.NewReader([]byte(rest))
	var numMovies int16
	binary.Read(reader2, binary.NativeEndian, &numMovies)

	totalLen += 2
	if totalLen%4 != 0 {
		// we already know it's a multiple of 2, so we just pad 2 bytes to get to a 4-align
		reader2.Seek(2, 1) // 1: io.SeekCurrent
		totalLen += 2
	}

	movieIndexes := make([]int32, numMovies)
	for i := 0; i < int(numMovies); i++ {
		binary.Read(reader2, binary.NativeEndian, &movieIndexes[i])
	}
	// fmt.Println(name1)
	// return name1, movieIndexes
	var ret film
	return ret
}

func (db *imdb) getCastFromMovie(index int) []string {
	var ret []string
	return ret
}
