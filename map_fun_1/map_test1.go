package main

// Run it with: map_test1 [ minute | hour | day | month | year | whatever ]

import (
	"os"
	"fmt"
	"time"
)

func main() {
	args_without_prog := os.Args[1:]

	var now = time.Now()
	
	for _, element := range args_without_prog {

		fmt.Println(timeToBucket(now,element))
	}
}

var time_to_bucket_format = map[string]string {
	"minute" : "2006-01-02T15:04:05",
	"hour" : "2006-01-02T15",
	"day" : "2006-01-02",
	"month" : "2006-01",
	"year" : "2006",
}

func timeToBucket(tm time.Time,trunc string)(string,error) {
	var bucket string
	var response error
	
	bucket, ok := time_to_bucket_format[trunc]
	
	if ok {
		return bucket, response
	} else {
		response = fmt.Errorf("Invalid truncation period: %s", trunc)
		return bucket, response;
	}
}
