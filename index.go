package main

import (
	"bufio"
	"errors"
	"log"
	"net/http"
	"strings"
)

func extract_url(line string) *string {
	if strings.HasPrefix(line, "#") {
		return nil
	}
	if !strings.HasPrefix(line, "http") {
		return nil
	}
	return &line
}

func DownloadIndex(indexUrl string) ([]*string, error) {
	response, err := http.Get(indexUrl)

	if err != nil {
		log.Printf("Could not download '%s': %s\n", indexUrl, err)
		return make([]*string, 0), err
	}

	if response.StatusCode != 200 {
		log.Printf("Server returned an error code %d", response.StatusCode)
		return make([]*string, 0), errors.New("Could not download, server returned an error")
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanLines)

	var urlsList []*string

	for scanner.Scan() {
		line := scanner.Text()
		result := extract_url(line)
		if result != nil {
			urlsList = append(urlsList, result)
		}
	}

	return urlsList, nil
}
