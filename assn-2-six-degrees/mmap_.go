package main

import (
	"log"
	"os"

	"github.com/edsrzf/mmap-go"
)

type MMap = mmap.MMap

func mmap_(f *os.File) (mmap.MMap, []byte) {
	m, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		log.Fatal("error mapping file")
	}
	return m, m
}

func unmap_(m mmap.MMap) {
	m.Unmap()
}
