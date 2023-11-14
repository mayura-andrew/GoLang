package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5
const numMeals = 3
const maxConcurrentEaters = 2

var wg sync.WaitGroup

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id                int
	leftChopstick     *Chopstick
	rightChopstick    *Chopstick
	hostSemaphore     chan struct{}
	timesAte          int
	mealsBeforeFinish int
}

func (p *Philosopher) eat() {
	defer wg.Done()

	for p.timesAte < p.mealsBeforeFinish {
		<-p.hostSemaphore
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Millisecond * 100) // simulate eating
		fmt.Printf("finishing eating %d\n", p.id)

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
		p.hostSemaphore <- struct{}{}

		p.timesAte++
	}
}

func main() {
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	hostSemaphore := make(chan struct{}, maxConcurrentEaters)
	for i := 0; i < maxConcurrentEaters; i++ {
		hostSemaphore <- struct{}{}
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:                i + 1,
			leftChopstick:     chopsticks[i],
			rightChopstick:    chopsticks[(i+1)%numPhilosophers],
			hostSemaphore:     hostSemaphore,
			timesAte:          0,
			mealsBeforeFinish: numMeals,
		}
	}

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].eat()
	}

	wg.Wait()
}
