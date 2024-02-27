package utils

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xfrr/goffmpeg/transcoder"
)

func ConvertToHLS(input string, output string, configQualities []string) {
	// validate configQualities to ensure it is not empty and they are string of numbers
	if len(configQualities) == 0 {
		log.Fatalf("configQualities cannot be empty")
	} else {
		for _, quality := range configQualities {
			if _, err := strconv.Atoi(quality); err != nil {
				log.Fatalf("configQualities must be a string of numbers")
			}
		}
	}
	for _, quality := range configQualities {
		convertToHLSWithQuality(input, output, quality)
	}
}

func convertToHLSWithQuality(input string, output string, quality string) {
	trans := new(transcoder.Transcoder)

	err := trans.Initialize(input, fmt.Sprintf("%s_%s.m3u8", output, quality))
	if err != nil {
		log.Fatalf("Could not initialize transcoder: %s", err)
	}

	trans.MediaFile().SetVideoFilter(fmt.Sprintf("scale=-2:%s", quality))

	done := trans.Run(true)

	progress := trans.Output()

	for msg := range progress {
		log.Printf("Convert video progress [%v]: %v\n", quality, msg)
	}

	err = <-done
	if err != nil {
		log.Fatalf("Transcoder error: %s", err)
	}
}
