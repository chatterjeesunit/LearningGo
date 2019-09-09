package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"math/rand"
	"os"
	"time"
)

func main() {

	fileName := "/Users/in-sunit.chatterjee/learn/GitHub/go/src/LearningGo/examples/fileio/demo-watch.txt"

	watchFile(fileName)


	file ,_ := os.Create(fileName)
	fmt.Println("file Created...")
	file.Close()
	writeToFile(fileName, false, 1)



	writeToFile(fileName, true, 10)

	fmt.Println("Done Main...")
}

func writeToFile(fileName string, sleep bool, n int) {

	for i:=0; i < n; i++{
		if sleep {
			rand.Seed(time.Now().UnixNano())
			sleepTime := rand.Int31n(10)
			fmt.Printf("Sleeping for %d seconds \n", sleepTime)
			time.Sleep(time.Duration(sleepTime) * time.Second)

		}

		file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		fmt.Println("Writing to file")
		file.Write([]byte("some random text"))
	}
}

func watchFile(fileName string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Starting Watcher : ", err)
		os.Exit(1)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				fmt.Println("\tevent:", event)
				if !ok {
					fmt.Println("\tEvents... Not OK")
					//return
				}
				fmt.Println("\tOperation : ", event.Op.String())

			case err, ok := <-watcher.Errors:
				fmt.Println("\terror:", err)
				if !ok {
					fmt.Println("\tErrors... Not OK")
					//return
				}

			}
		}
	}()

	err = watcher.Add(fileName)
	if err != nil {
		fmt.Println("\tError Adding file", err)
		os.Exit(1)
	}
	fmt.Println("watching file : ", fileName)
}