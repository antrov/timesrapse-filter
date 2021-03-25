package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kelvins/sunrisesunset"
)

func main() {
	name := flag.String("name", "", "file with unix timestamp in name")
	offset := flag.Int("offset", 60, "offset befure sunrise and after sunset")
	flag.Parse()

	if name == nil || len(*name) == 0 {
		log.Fatalln("filename not passed")
	}

	filename := filepath.Base(*name)

	value := strings.TrimSuffix(filename, filepath.Ext(filename))
	epoch, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	date := time.Unix(epoch, 0)

	p := sunrisesunset.Parameters{
		Latitude:  51.570786,
		Longitude: 19.265566,
		UtcOffset: 1.0,
		Date:      date,
	}

	// Calculate the sunrise and sunset times
	sunrise, sunset, err := p.GetSunriseSunset()

	// If no error has occurred, print the results
	if err != nil {
		fmt.Println(err)
	}

	// log.Println("date", date)
	// log.Println("sunrise", sunrise, "->", sunrise.Add(time.Duration(-1*offset)*time.Minute))
	// log.Println("sunset", sunset, "->", sunset.Add(time.Duration(offset)*time.Minute))

	if date.Before(sunrise.Add(time.Duration(-1**offset) * time.Minute)) {
		// log.Println("przed wschodem")
	} else if date.After(sunset.Add(time.Duration(*offset) * time.Minute)) {
		// log.Println("po zachodzie")
	} else {
		fmt.Printf("file '%s'\n", *name)
	}
}
