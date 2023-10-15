package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {

	client := connectCache()

	// set initial visits value as 0 if the key does not exist
	_, err := client.SetNX("visits", 0, 2*time.Minute).Result()
	if err != nil {
		fmt.Println("Error setting initial value: " + err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get no. of visits stored in cache
		val, err := client.Get("visits").Result()
		if err != nil {
			fmt.Println("Error retrieving results " + err.Error())
			return
		}

		// Write results as response
		w.Write([]byte("No of visits " + val))

		// update visits value
		ival, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Error converting str to int " + err.Error())
			return
		}

		// Updating visits value
		err = client.Set("visits", ival+1, 2*time.Minute).Err()
		if err != nil {
			fmt.Println("Error while setting the value to cache")
			return
		}
	})

	// starting the server at :8080
	http.ListenAndServe(":8080", nil)

}
