package main

// Task - single unit of work (container)
// Job - group of tasks
// (pick the machine with greatest score)Scheduler assign task to worker
// Manager - brain
// Worker - handle tasks and gather statistics (kubelet)
// Cluster - all this components together
// CLI:

// 1. Start and stop tasks
// 2. Get the status of tasks
// 3. See the state of machines (i.e. the workers)
// 4. Start the manager
// 5. Start the worker

// Manager: API for client, job current state, worker metrics, scheduling api
// Scheduler: feasability, scoring, picking
// Worker: API for manager, task state, metrics, task list

func main() {

}
