package main

import (
  "log"
  "os/exec"
  "errors"
)

func RunConverter(incomingFile string, outgoingFile string) error {
  if incomingFile==outgoingFile {
    log.Printf("Incoming and outgoing file names must be different")
    return errors.New("Incoming and outgoing names must be different")
  }


  cmd := exec.Command("/usr/bin/ffmpeg", "-i", incomingFile, "-vcodec", "copy", "-acodec", "copy", "-bsf:a", "aac_adtstoasc", "-f", "mp4", outgoingFile)

  outputBytes, cmdErr := cmd.Output()

  outputString := string(outputBytes)
  log.Printf("Output from converter: \n%s", outputString)

  if cmdErr != nil {
    log.Printf("Could not run converter: %s", cmdErr)
    return cmdErr
  }
  return nil
}
