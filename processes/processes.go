package processes

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

// Processes is a struct that wraps around the go gopsutil process library and
// provides a number of helper functions to make life easier
type Processes struct {
	ProcessList []*process.Process
}

// Add takes a pid number, creates a new process and adds it to the currently known processes
func (p *Processes) Add(pid int32) error {
	process, err := process.NewProcess(pid)

	if err != nil {
		return fmt.Errorf("unable to create process object from pid: %s", err)
	}

	p.ProcessList = append(p.ProcessList, process)
	return nil
}

// New returns a pointer to a new Processes struct
func New() (*Processes, error) {
	var processes Processes

	pids, err := process.Pids()

	if err != nil {
		return nil, fmt.Errorf("unable to get PIDs: %s", err)
	}

	for _, pid := range pids {
		err := processes.Add(pid)

		if err != nil {
			return nil, fmt.Errorf("error adding process to process list: %s", err)
		}
	}

	return &processes, nil
}
