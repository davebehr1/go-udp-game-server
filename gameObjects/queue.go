package gameObjects

import "sync"

type ConcurrentQueue struct {
	lock *sync.Mutex

	notEmpty *sync.Cond
	notFull  *sync.Cond

	backend *QueueBackend
}

func (c *ConcurrentQueue) enqueue(data interface{}) error {
	c.lock.Lock()

	for c.backend.isFull() {
		c.notFull.Wait()
	}
	err := c.backend.put(data)

	c.notEmpty.Signal()

	c.lock.Unlock()

	return err
}

func (c *ConcurrentQueue) dequeue() (interface{}, error) {
	c.lock.Lock()

	for c.backend.isEmpty() {
		c.notEmpty.Wait()
	}

	data, err := c.backend.pop()

	c.notFull.Signal()

	c.lock.Unlock()

	return data, err
}

func (c *ConcurrentQueue) getSize() uint32 {
	c.lock.Lock()

	size := c.backend.size

	c.lock.Unlock()

	return size
}

func NewConcurrentQueue(maxSize uint32) *ConcurrentQueue {
	queue := ConcurrentQueue{}

	queue.lock = &sync.Mutex{}
	queue.notFull = sync.NewCond(queue.lock)
	queue.notEmpty = sync.NewCond(queue.lock)

	queue.backend = &QueueBackend{}
	queue.backend.size = 0
	queue.backend.head = nil
	queue.backend.tail = nil
	queue.backend.maxSize = maxSize
	return &queue
}
