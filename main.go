package main

import (
  "flag"
  "log"
  "github.com/davecgh/go-spew/spew"
  "os"
)

func main() {
  var url = flag.String("url","","m3u8 index to download")
  flag.Parse()

  if url==nil || *url=="" {
    log.Fatal("You must specify the -url argument")
  }

  mediaList, downloadErr := DownloadIndex(*url)

  if downloadErr != nil {
    log.Fatal("Could not download")
  }

  spew.Dump(mediaList)

  totalChunks := len(mediaList)
  log.Printf("Got a total of %d chunks to download\n", totalChunks)

  filename := "media.ts"

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

  convertErr := RunConverter(filename,"media.mp4")
  if convertErr != nil {
    log.Fatal("Could not convert")
  }
  os.Remove(filename)
}
