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
	var list = flag.String("list", "", "Download m3u8 indices from list")
	var tsout = flag.Bool("ts",false,"Don't convert to mp4, keep as transport string")
	flag.Parse()

	if (url == nil || *url == "") && (list==nil || *list=="") {
		log.Fatal("You must specify either the -url or -list argument")
	}

	var indices *[]string
	if *list == "" {
		indices = &[]string{*url,}
	} else {
		var err error
		indices, err = ReadUrlList(list)
		if err != nil {
			log.Fatal("Could not read specified list: %s", err)
		}
	}

	spew.Dump(indices)
	pwd, _ := os.Getwd()
	outfile, temperr := ioutil.TempFile(pwd, "mediats-")
	if temperr != nil {
		log.Fatalf("Could not create tempfile for TS: %s", temperr)
	}

	filename := outfile.Name()

	fp, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Could not create %s: %s", filename, err)
	}

	defer fp.Close()

	for _, toDownload := range(*indices){
		log.Printf("Downloading %s", toDownload)
		mediaList, downloadErr := DownloadIndex(toDownload)

		if downloadErr != nil {
			log.Fatal("Could not download")
		}

		//spew.Dump(mediaList)

		totalChunks := len(mediaList)
		log.Printf("Got a total of %d chunks to download\n", totalChunks)

		log.Printf("Downloading %s to %s...", *url, filename)

		for ctr, chunkUrl := range mediaList {
			log.Printf("Chunk %d / %d...", ctr, totalChunks)
			DownloadMediaChunk(*chunkUrl, fp)
		}
	}

	if *tsout == false {
		convertErr := RunConverter(filename, *out)
		if convertErr != nil {
			log.Fatal("Could not convert")
		}
		os.Remove(filename)
	}
}
