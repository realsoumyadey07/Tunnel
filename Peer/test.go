package main

import (
	"log"
	"os"
)

func testFunction() {

	parentFile := "storage/blossom.mp4"
	madeFile := "made/blossom.mp4"

	tfm, err := SegmentFile(parentFile)

	if err != nil {
		log.Panicln("Error:", err.Error())
	}

	for _, c := range tfm.AllSegments {
		log.Println(c.FileDestination, "|", c.ContentSize)
	}

	jointBLFiles(tfm)

	pb, err := os.ReadFile(parentFile)

	if err != nil {
		log.Panicln("Error:", err.Error())
	}

	mb, err := os.ReadFile(madeFile)

	if err != nil {
		log.Panicln("Error:", err.Error())
	}

	for i, b := range pb {

		if mb[i] != b {
			log.Println(mb[i], "!=", b)
		}
	}

}
