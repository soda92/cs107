package main

import (
	"log"
	"os"

	"github.com/edsrzf/mmap-go"
)

func mmap_(f *os.File) (mmap.MMap, *byte) {
	m, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		log.Fatal("error mapping file")
	}
	return m, &m[0]
}

func unmap_(m mmap.MMap) {
	m.Unmap()
}
