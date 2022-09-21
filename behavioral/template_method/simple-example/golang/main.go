package main

type ProcessInterface interface {
	start()
	execute()
	stop()
}

type Process struct {
	ProcessInterface
}

func newProcess(p ProcessInterface) *Process {
	return &Process{p}
}

func (p *Process) StartProcessLifecycle() {
	p.start()
	p.execute()
	p.stop()
}

type CleanupProcess struct {
}

func (c *CleanupProcess) start() {
	println("Cleanup process started")
}

func (c *CleanupProcess) stop() {
	println("Cleanup process stopped")
}

func (c *CleanupProcess) execute() {
	println("Cleanup process executing")
}

func main() {
	process := newProcess(&CleanupProcess{})
	process.StartProcessLifecycle()
}
