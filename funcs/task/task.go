package main

import (
    "log"
)


// State is task state.
type State byte

const (
    Ready State = iota + 1
    Working
    Done
)

// Task is a task to execute.
type Task struct {
    ID     uint
    Result any
    Err    error
    State  State

    Work     func() (any, error)
    Watchers []func(*Task)
}

// NewTask returns a new Task that will execute work.
func NewTask(id uint, work func() (any, error)) *Task {
    t := Task{
        ID:    id,
        State: Ready,
        Work:  work,
    }
    return &t
}



// Subscribe subscribes a watcher to task.
func (t *Task) Subscribe(w func(*Task)) {
    t.Watchers = append(t.Watchers, w)
}



// Execute executes the task.
func (t *Task) Execute() error {
    t.State = Working
    defer func() { t.State = Done }() 

    t.Result, t.Err = t.Work()
    if t.Err != nil {
        log.Printf("error: Task %d failed - %s", t.ID, t.Err)
    }

    // Notify watchers
    for _, s := range t.Watchers {
        s(t)
    }

    return t.Err
}



// Watcher is a task watcher.
type Watcher struct{}

func (w *Watcher) Handle(t *Task) {
    log.Printf("info: w2: from %d - %#v %v", t.ID, t.Result, t.Err)
}


func main() {
    t := NewTask(7, func() (any, error) { return "done", nil })
    t.Subscribe(func(t *Task) {
        log.Printf("info: w1: from %d - %#v %v", t.ID, t.Result, t.Err)
    })
    var w Watcher
    t.Subscribe(w.Handle) 

    t.Execute()

}

/* Output:
2023/06/21 16:17:52 info: w1: from 7 - "done" <nil>
2023/06/21 16:17:52 info: w2: from 7 - "done" <nil>
*/
