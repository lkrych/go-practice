package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

//collect the weights of each animal on a frequent basis

var animaltags = []string{"Monarch Butterfly;Rex", "Caterpillar brood;Squirmy", "Tarantula;Harry"}

const myDB = "dino"

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	queryDB(c, "", "Create Database "+myDB) //create db if it does not exist
	//create a batch points object
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  myDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{} //shuts down gracefully by releasing all memory and ensuring all pending tasks get finished
	//A wait group is a way for go routines to communicate they are currently processing something
	detectSignal := checkStopOSSignal(&wg)
	rand.Seed(time.Now().UnixNano())
	//routine for adding random weight data to the batchpoints object
	for !(*detectSignal) { //runs as long no one exits the program
		animaltag := animaltags[rand.Intn(len(animaltags))]
		split := strings.Split(animaltag, ";")
		tags := map[string]string{
			"animal_type": split[0],
			"nickname":    split[1],
		}
		fields := map[string]interface{}{
			"weight": rand.Intn(300) + 1,
		}
		fmt.Println(animaltag, fields["weight"])
		pt, err := client.NewPoint("weightmeasures", tags, fields, time.Now())
		if err != nil {
			log.Println(err)
			continue
		}
		bp.AddPoint(pt)
		time.Sleep(1 * time.Second)
	}
	log.Println("Exit signal triggered, writing data... ")
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	wg.Wait() //waits for wait group to finish
	log.Println("Exiting program...")
}

func queryDB(c client.Client, database, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: database,
	}
	response, err := c.Query(q)
	if err != nil {
		return res, err
	}
	if response.Error() != nil {
		return res, response.Error() //res was initialized in function declaration
	}
	res = response.Results
	return res, nil
}

func checkStopOSSignal(wg *sync.WaitGroup) *bool {
	Signal := false //will return true when OS reports that user wants to exit the program
	go func(s *bool) {
		wg.Add(1)
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch //listen and block on the channel until signal is received
		log.Println("Exit signals received... ")
		*s = true
		wg.Done() //decrements wg.Add, allows anything waiting on wg to continue
	}(&Signal)
	return &Signal
}
