package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"strings"
)

func (db *imdb) DecodeActor(index int) (string, []int32) {
	name := GetRecord(db.actorFile, index)
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
		reader2.Seek(2, io.SeekCurrent)
		totalLen += 2
	}

	offsets := make([]int32, numMovies)
	for i := 0; i < int(numMovies); i++ {
		binary.Read(reader2, binary.NativeEndian, &offsets[i])
	}
	// fmt.Println(name1)
	return name1, offsets
}

func GetRecord(db []byte, index int) string {
	index = index * 4
	var firstAddr int32
	binary.Decode(db[index:index+4], binary.LittleEndian, &firstAddr)
	var nextAddr int32
	binary.Decode(db[index+4:index+8], binary.LittleEndian, &nextAddr)
	if firstAddr >= nextAddr {
		log.Fatal("index error")
	}

	record := string(db[firstAddr:nextAddr])
	return record
}

func (db *imdb) DecodeMovie(index int) film {
	name := GetRecord(db.movieFile, index)
	len1 := strings.IndexByte(name, 0x00)
	name1 := name[:len1]
	rest := name[len1+1:]
	// totalLen := len1+1
	year := int(rest[0])
	year += 1900
	var movie film
	movie.title = name1
	movie.year = year
	return movie
}

func (db *imdb) getCastFromMovie(index int) []string {
	var ret []string
	return ret
}

func GetOffsetByIndex(db []byte, index int) int {
	var offset int32
	binary.Decode(db[index*4:(index+1)*4], binary.NativeEndian, &offset)
	return int(offset)
}
