package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

func IndexOf(data []byte, val byte) int {
	for i, v := range data {
		if val == v {
			return i
		}
	}
	log.Fatal("index error")
	return -1
}

func (db *imdb) DecodeActorRecord(record []byte) (string, []int32) {
	len1 := IndexOf(record, 0x00)
	totalLen := len1
	name := string(record[:len1])
	rest := record[len1:]
	if len(name)%2 != 0 {
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
	return name, offsets
}

func (db *imdb) DecodeActor(index int) (string, []int32) {
	record := GetRecordFromOffset(db.actorFile, GetOffsetByIndex(db.actorFile, index))
	return db.DecodeActorRecord(record)
}

func (db *imdb) DecodeActorFromOffset(offset int) (string, []int32) {
	record := GetRecordFromOffset(db.actorFile, offset)
	return db.DecodeActorRecord(record)
}

func GetRecord(db []byte, index int) []byte {
	index = index * 4 // the index was already plus-ed 1
	var offset int32
	binary.Decode(db[index:index+4], binary.LittleEndian, &offset)

	return GetRecordFromOffset(db, int(offset))
}

func GetRecordFromOffset(db []byte, offset int) []byte {
	return db[offset:]
}

func (db *imdb) DecodeMovie(index int) film {
	name := GetRecord(db.movieFile, index)
	len1 := IndexOf(name, 0x00)
	name1 := name[:len1]
	rest := name[len1+1:]
	// totalLen := len1+1
	year := int(rest[0])
	year += 1900
	var movie film
	movie.title = string(name1)
	movie.year = year
	return movie
}

func (db *imdb) getCastFromMovie(index int) []string {
	movie := db.DecodeMovie(index)
	record := GetRecord(db.movieFile, index)
	totalLen := len(movie.title) + 1 /*\0*/ + 1 /*year*/
	if totalLen%2 != 0 {
		totalLen += 1
	}
	record = record[totalLen:]
	bin := []byte(record)
	var numActors int16
	binary.Decode(bin[0:2], binary.NativeEndian, &numActors)
	totalLen += 2
	offset := 2
	if totalLen%4 != 0 {
		totalLen += 2
		offset += 2
	}
	record = record[offset:]
	// if len(record) != int(numActors)*4 {
	// 	log.Fatal("movie record corrupt: record doesn't have correct number of actors")
	// }
	// bin = []byte(record)
	reader := bytes.NewReader(record)
	var offsets []int
	for range numActors {
		var offset int32
		binary.Read(reader, binary.NativeEndian, &offset)
		offsets = append(offsets, int(offset))
	}
	var ret []string
	for _, offset := range offsets {
		actor, _ := db.DecodeActorFromOffset(offset)
		ret = append(ret, actor)
	}
	return ret
}

func GetOffsetByIndex(db []byte, index int) int {
	var offset int32
	binary.Decode(db[index*4:(index+1)*4], binary.NativeEndian, &offset)
	return int(offset)
}
