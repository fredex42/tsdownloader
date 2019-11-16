package main

import (
  "net/http"
  _ "net/url"
  "log"
  "errors"
  "os"
    "io"
)

func DownloadMediaChunk(downloadUrl string, fp *os.File) error {
  //splitUrl, _ := url.Parse(downloadUrl)

  response, dlErr := http.Get(downloadUrl)

  if dlErr != nil {
    log.Printf("could not download %s: %s", downloadUrl, dlErr)
    return dlErr
  }

  if response.StatusCode != 200 {
      log.Printf("Server returned %d trying to download %s", response.StatusCode, downloadUrl)
      return errors.New("Server error")
  }

  defer response.Body.Close()

  // filename := path.Base(splitUrl.Path)
  //
  // log.Printf("Downloading %s to %s...", downloadUrl, filename)
  // fp, err := os.Create(filename)
  // if err != nil {
  //   log.Printf("Could not create %s: %s", filename, err)
  //   return err
  // }
  //
  // defer fp.Close()

  io.Copy(fp, response.Body)
  return nil
}
