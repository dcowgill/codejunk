#!/usr/bin/env python

import collections
import heapq
import itertools
import math
import random


INTERVAL = 30


class EmptyQueueException(Exception): pass


class Dist(object):
    def __init__(self):
        self.n = 0.0
        self.mean = 0.0
        self.M2 = 0.0

    def add(self, x):
        self.n += 1
        delta = x - self.mean
        self.mean += delta / self.n
        self.M2 += delta * (x - self.mean)

    @property
    def variance(self):
        return self.M2/(self.n - 1)

    @property
    def stddev(self):
        return math.sqrt(self.variance)

    def __repr__(self):
        return "({0:.3f}, {1:.3f})".format(self.mean, self.stddev)


class Task(object):
    def __init__(self, priority, duration):
        self.priority = priority
        self.duration = duration
        self.enqueued_at = None
        self.dequeued_at = None

    @property
    def waited_for(self):
        if self.dequeued_at is not None:
            return self.dequeued_at - self.enqueued_at
        return None

    def __cmp__(self, other):
        if not isinstance(other, type(self)):
            return NotImplemented
        return self.priority - other.priority

    def __repr__(self):
        return "Task(%d, %d)" % (self.priority, self.duration)


class Worker(object):
    def __init__(self, id):
        self.id = id
        self._task = None

    def is_idle(self):
        return not self._task

    def accept(self, task):
        self._task = task

    def result(self, time):
        if self.is_idle():
            return None
        if (time - self._task.dequeued_at) < self._task.duration:
            return None
        retval = self._task
        self._task = None
        return retval


class GaussianTaskGenerator(object):
    def __init__(self, rate, cost):
        self.rate_mu, self.rate_sigma = rate
        self.cost_mu, self.cost_sigma = cost

    def reset(self): pass

    def generate_tasks(self, priority, interval):
        n = int(interval * random.gauss(self.rate_mu, self.rate_sigma))
        return (Task(priority, self._random_duration()) for _ in xrange(n))

    def _random_duration(self):
        return random.gauss(self.cost_mu, self.cost_sigma)


class UpfrontTaskGenerator(object):
    def __init__(self, n, cost):
        self.n = n
        self.cost_mu, self.cost_sigma = cost

    def reset(self):
        self.first = True

    def generate_tasks(self, priority, interval):
        if not self.first:
            return []
        self.first = False
        return (Task(priority, self._random_duration()) for _ in xrange(self.n))

    def _random_duration(self):
        return random.gauss(self.cost_mu, self.cost_sigma)


class Workload(object):
    def __init__(self, task_generators):
        self._generators = list(enumerate(task_generators))

    def reset(self):
        for _, g in self._generators:
            g.reset()

    def num_priorities(self):
        return len(self._generators)

    def generate_tasks(self, interval):
        for priority, g in self._generators:
            for task in g.generate_tasks(priority, interval):
                yield task


class BasicStrategy(object):
    def pop(self, queue, worker):
        return queue.dequeue()


class ProbabilisticStrategy(object):
    def __init__(self, p):
        self.p = p

    def pop(self, queue, worker):
        if random.random() >= self.p:
            return queue.dequeue()
        return queue.dequeue_random()

    def __str__(self):
        return "ProbabilisticStrategy(p={:0.2f})".format(self.p)


class PriorityAffinityStrategy(object):
    def pop(self, queue, worker):
        level = worker.id % queue.num_levels()
        return queue.dequeue_preferred(level)


class MultiLevelQueue(object):
    def __init__(self, num_levels):
        self._levels = range(num_levels)
        self._queues = [collections.deque() for _ in self._levels]

    def enqueue(self, task):
        self._queues[task.priority].appendleft(task)

    def dequeue(self):
        for q in self._queues:
            if q:
                return q.pop()
        raise EmptyQueueException()

    def dequeue_random(self):
        random.shuffle(self._levels)
        for level in self._levels:
            q = self._queues[level]
            if q:
                return q.pop()
        raise EmptyQueueException()

    def dequeue_preferred(self, level):
        q = self._queues[level]
        if q:
            return q.pop()
        return self.dequeue()

    def dequeue_all(self):
        for q in self._queues:
            while q:
                yield q.pop()

    def sizes(self):
        for i in self._levels:
            yield i, len(self._queues[i])

    def empty(self):
        return not any(self._queues)

    def num_levels(self):
        return len(self._levels)


class Simulation(object):
    def __init__(self, workload, num_workers=10, strategy=None):
        nlevels = workload.num_priorities()
        self._strategy = strategy or BasicStrategy()
        self._workload = workload
        self._workers = [Worker(i) for i in xrange(num_workers)]
        self._queue = MultiLevelQueue(nlevels)
        self._queue_size = dict((i, Dist()) for i in xrange(nlevels))
        self._wait_times = dict((i, Dist()) for i in xrange(nlevels))
        self._completed = dict((i, 0) for i in xrange(nlevels))
        self._idle_workers = Dist()
        self._time = 0
        self._workload.reset()

    def run(self, duration):
        for _ in xrange(duration):
            self._step()

        for task in self._queue.dequeue_all():
            task.dequeued_at = self._time
            self._wait_times[task.priority].add(task.waited_for)

        return {
            'tasks_completed': self._completed,
            'wait_times': self._wait_times,
            'queue_size': self._queue_size,
            'idle_workers': self._idle_workers,
        }

    def _step(self):
        if self._time % INTERVAL == 0:
            for task in self._workload.generate_tasks(INTERVAL):
                task.enqueued_at = self._time
                self._queue.enqueue(task)

        for worker in self._workers:
            task = worker.result(self._time)
            if task:
                self._completed[task.priority] += 1

        for worker in self._workers:
            if self._queue.empty():
                break
            if worker.is_idle():
                task = self._strategy.pop(self._queue, worker)
                task.dequeued_at = self._time
                self._wait_times[task.priority].add(task.waited_for)
                worker.accept(task)

        for i, n in self._queue.sizes():
            self._queue_size[i].add(n)

        self._idle_workers.add(sum(1 for w in self._workers if w.is_idle()))
        self._time += 1


def report(header, results):
    tasks_completed = sorted(results['tasks_completed'].iteritems())
    queue_size = sorted(results['queue_size'].iteritems())
    wait_times = sorted(results['wait_times'].iteritems())
    total_tasks_completed = sum(n for _, n in tasks_completed)

    print header
    print '-' * 80
    print "Tasks completed: {} = {} total".format(
        tasks_completed, total_tasks_completed)
    print "Queue size:      {}".format(queue_size)
    print "Wait times:      {}".format(wait_times)
    print "Idle workers:    {}".format(results['idle_workers'])
    print ""


if __name__ == '__main__':
    workloads = collections.OrderedDict([
        ("Low throughput, fast tasks", Workload([
            GaussianTaskGenerator(rate=(0.5, 0.5), cost=(2, 1)),
            GaussianTaskGenerator(rate=(0.5, 0.5), cost=(2, 1)),
            GaussianTaskGenerator(rate=(0.5, 0.5), cost=(2, 1)),
        ])),
        ("All priorities equal, fast tasks", Workload([
            GaussianTaskGenerator(rate=(2.0, 0.1), cost=(2, 1)),
            GaussianTaskGenerator(rate=(2.0, 0.1), cost=(2, 1)),
            GaussianTaskGenerator(rate=(2.0, 0.1), cost=(2, 1)),
        ])),
        ("All priorities equal, slow tasks", Workload([
            GaussianTaskGenerator(rate=(2.0, 0.1), cost=(10, 0)),
            GaussianTaskGenerator(rate=(2.0, 0.1), cost=(10, 0)),
            GaussianTaskGenerator(rate=(2.0, 0.1), cost=(10, 0)),
        ])),
        ("Too many high priority", Workload([
            GaussianTaskGenerator(rate=(5.0, 0.1), cost=(2, 1)),
            GaussianTaskGenerator(rate=(0.5, 0.1), cost=(2, 1)),
            GaussianTaskGenerator(rate=(0.5, 0.1), cost=(2, 1)),
        ])),
        ("Too many low priority", Workload([
            GaussianTaskGenerator(rate=(0.5, 0.1), cost=(2, 1)),
            GaussianTaskGenerator(rate=(0.5, 0.1), cost=(2, 1)),
            GaussianTaskGenerator(rate=(5.0, 0.1), cost=(2, 1)),
        ])),
        ("Lots of upfront low priority tasks", Workload([
            GaussianTaskGenerator(rate=(3.0, 0.1), cost=(2, 1)),
            GaussianTaskGenerator(rate=(2.0, 0.1), cost=(2, 1)),
            UpfrontTaskGenerator(n=10000, cost=(2, 0)),
        ])),
        ("High variability", Workload([
            GaussianTaskGenerator(rate=(0, 5), cost=(1, 5)),
            GaussianTaskGenerator(rate=(0, 5), cost=(1, 5)),
            GaussianTaskGenerator(rate=(0, 5), cost=(1, 5)),
        ])),
        ("Low and high", Workload([
            GaussianTaskGenerator(rate=(5, 1), cost=(1, 2)),
            GaussianTaskGenerator(rate=(5, 1), cost=(1, 2)),
        ])),
    ])

    strategies = [
        BasicStrategy(),
        ProbabilisticStrategy(p=0.1),
        ProbabilisticStrategy(p=0.5),
        PriorityAffinityStrategy(),
    ]

    for name, workload in workloads.iteritems():
        print "*" * 80
        for strategy in strategies:
            sim = Simulation(workload, num_workers=10, strategy=strategy)
            report("{0}; {1!s}".format(name, strategy), sim.run(3600))
