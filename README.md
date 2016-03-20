Samples to demonstrate usage of WorkerGo package.

## sample.go

Creates dispatcher, sends jobs and waits with an infinite loop.

## sample_waitgroup.go

Uses an existing sync.WaitGroup to wait for goroutines. Creates a sync.WaitGroup, and passes the pointer for it into the dispatcher.