package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var url = flag.String("url", "", "m3u8 index to download")
	var out = flag.String("out", "media.mp4", "mp4 file to output")
	flag.Parse()

	if url == nil || *url == "" {
		log.Fatal("You must specify the -url argument")
	}

	mediaList, downloadErr := DownloadIndex(*url)

	if downloadErr != nil {
		log.Fatal("Could not download")
	}

	spew.Dump(mediaList)

	totalChunks := len(mediaList)
	log.Printf("Got a total of %d chunks to download\n", totalChunks)

	pwd, _ := os.Getwd()
	outfile, temperr := ioutil.TempFile(pwd, "mediats-")
	if temperr != nil {
		log.Fatalf("Could not create tempfile for TS: %s", temperr)
	}

	filename := outfile.Name()

	log.Printf("Downloading %s to %s...", *url, filename)
	fp, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Could not create %s: %s", filename, err)
	}

	defer fp.Close()

	for ctr, chunkUrl := range mediaList {
		log.Printf("Chunk %d / %d...", ctr, totalChunks)
		DownloadMediaChunk(*chunkUrl, fp)
	}

	convertErr := RunConverter(filename, *out)
	if convertErr != nil {
		log.Fatal("Could not convert")
	}
	os.Remove(filename)
}
