package main
import (
  "log"
  "os"
  "bufio"
)

func ReadUrlList(listfile *string) (*[]string, error) {
  file, err := os.Open(*listfile)
  if err != nil {
    log.Printf("Could not open list file %s: %s", listfile, err)
    return nil, err
  }

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var rtn []string

  for scanner.Scan() {
    rtn = append(rtn, scanner.Text())
  }
  return &rtn, nil
}
