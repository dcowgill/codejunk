package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Work struct{ cost time.Duration }

func (w *Work) do() { time.Sleep(w.cost) }

type ProducerConsumerService struct {
	source chan chan *Work
	closed bool
	quitc  chan struct{}
	wg     sync.WaitGroup
}

func NewProducerConsumerService() *ProducerConsumerService {
	return &ProducerConsumerService{
		source: make(chan chan *Work),
		quitc:  make(chan struct{}),
	}
}

func (s *ProducerConsumerService) Stop() {
	if !s.closed {
		close(s.quitc)
		s.closed = true
	}
}

func (s *ProducerConsumerService) Start(numWorkers int) error {
	s.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go s.consume(i)
	}
	s.produce()
	s.wg.Wait()
	return nil
}

func (s *ProducerConsumerService) produce() {
	for {
		select {
		case workc := <-s.source:
			log.Print("producer: got a worker, looking for work")
			for {
				if w := getWork(); w != nil {
					log.Print("producer: got work")
					workc <- w
					break
				}
				log.Print("producer: no work available, sleeping")
				select {
				case <-time.After(time.Second):
				case <-s.quitc:
					log.Print("producer: exiting")
					return
				}
			}
		case <-s.quitc:
			log.Print("producer: exiting")
			return
		}
	}
}

func (s *ProducerConsumerService) consume(id int) {
	defer s.wg.Done()
	workerc := s.source
	workc := make(chan *Work)
	for {

		select {
		case workerc <- workc:
			workerc = nil
		case work := <-workc:
			log.Printf("worker %d: got %s of work", id, work.cost)
			work.do()
			workerc = s.source
			log.Printf("worker %d: work complete", id)
		case <-s.quitc:
			log.Printf("worker %d: exiting", id)
			return
		}
	}
}

type ringBuffer struct {
	q []*Work
	p int
	n int
}

func (b *ringBuffer) enqueue(w *Work) {
	if b.full() {
		panic("enqueue on full ring buffer")
	}
	b.q[b.p] = w
	b.p = (b.p + 1) % len(b.q)
	b.n++
}

func (b *ringBuffer) dequeue() *Work {
	if b.n == 0 {
		panic("dequeue on empty ring buffer")
	}
	i := (b.p - b.n + len(b.q)) % len(b.q)
	b.n--
	return b.q[i]
}

func (b *ringBuffer) size() int  { return b.n }
func (b *ringBuffer) full() bool { return b.size() == len(b.q) }

// Global work queue. Database stand-in.
var workQueue = struct {
	sync.Mutex
	q *ringBuffer
}{
	q: &ringBuffer{q: make([]*Work, 1000)},
}

// Periodically adds work to the queue. Loops forever.
func makeWork() {
	for {
		workQueue.Lock()
		if !workQueue.q.full() {
			workQueue.q.enqueue(&Work{cost: randomDuration(50, 200)})
		}
		workQueue.Unlock()
		time.Sleep(randomDuration(0, 50))
	}
}

func randomDuration(l, h int) time.Duration {
	return time.Duration(rand.Intn(h-l+1)+l) * time.Millisecond
}

// Returns nil if there isn't any work.
func getWork() *Work {
	var work *Work
	workQueue.Lock()
	defer workQueue.Unlock()
	if workQueue.q.size() != 0 {
		log.Printf("work queue length = %d", workQueue.q.size())
		work = workQueue.q.dequeue()
	}
	return work
}

func installSignalHandlers(service *ProducerConsumerService) {
	signalc := make(chan os.Signal)
	signal.Notify(signalc, syscall.SIGINT)
	go func() {
		for s := range signalc {
			if s == syscall.SIGINT {
				log.Print("received SIGINT")
				service.Stop()
				return
			}
		}
	}()
}

func main() {
	go makeWork()
	service := NewProducerConsumerService()
	installSignalHandlers(service)
	log.Printf("s.Start returned: %v", service.Start(3))
}
