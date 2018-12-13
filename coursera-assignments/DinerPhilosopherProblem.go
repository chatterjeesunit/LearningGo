package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	MAX_EAT = 3
	TOTAL_NO_PHILOSOPHERS = 5
)

var (
	wg sync.WaitGroup
	eatingDoneChannel = make(chan bool, 2)
	permissionToEatChannel = make(chan bool, 2)
)

type ChopStick struct {
	number int
	sync.Mutex
}

type Philosopher struct {
	number         int
	leftChopstick  *ChopStick
	rightChopstick *ChopStick
	numberOfTimesEaten int
}

type Host struct {
	sync.Mutex
	finishedEatingCounter int
}


func (h *Host) finish() bool {
	h.Lock()
	defer h.Unlock()
	return h.finishedEatingCounter >= MAX_EAT*TOTAL_NO_PHILOSOPHERS
}

func (h *Host) markEatingDoneEvent() {
	h.Lock()
	h.finishedEatingCounter++
	h.Unlock()
	if !h.finish() {
		permissionToEatChannel <- true
	}else {
		fmt.Println("All Philosophers have EATEN")
	}




}

func (p *Philosopher) eat() {
	// If this philosopher has eaten 3 times, then don't do anything but return
	if p.numberOfTimesEaten == MAX_EAT {
		return
	}

	p.leftChopstick.Lock()
	p.rightChopstick.Lock()

	fmt.Printf("start to eat %d\n", p.number)
	p.numberOfTimesEaten++

	time.Sleep(time.Second * 1)

	fmt.Printf("\t\tfinishing eating %d\n", p.number)

	p.leftChopstick.Unlock()
	p.rightChopstick.Unlock()

	eatingDoneChannel <- true
}



func main() {

	host, philosophers := initialize()

	var wg sync.WaitGroup
	wg.Add(1)
	go eatDinner(host, philosophers, &wg)

	wg.Wait()

	for _, p := range philosophers {
		fmt.Printf("P%d : %d, ", p.number, p.numberOfTimesEaten)
	}

}


func initialize() (*Host, []*Philosopher) {
	host := new(Host)
	permissionToEatChannel <- true
	permissionToEatChannel <- true

	chopsticks := make([]*ChopStick, 5)
	for i := 0; i < TOTAL_NO_PHILOSOPHERS; i++ {
		chopsticks[i] = new(ChopStick)
		chopsticks[i].number = i
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < TOTAL_NO_PHILOSOPHERS; i++ {
		philosophers[i] = &Philosopher{i+1, chopsticks[i], chopsticks[(i+1)%5], 0}
	}

	return host, philosophers
}

func eatDinner(host *Host, philosophers []*Philosopher, wg *sync.WaitGroup) {
	for i := 0; !host.finish(); i++ {
		select {
		case <-permissionToEatChannel:
			go philosophers[i%5].eat()
		case <-eatingDoneChannel:
			host.markEatingDoneEvent()
		}
	}
	wg.Done()
}
