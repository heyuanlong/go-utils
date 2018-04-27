package ownrun

import (
	"time"
	"sync"
)

type agent interface {
	DoWork(v interface{})
	TimeOut()
}

type RunServer struct {
	a 						agent
	runNums 				int
	vchan 					chan interface{}
	isRun 					bool
	msec 					int64
	mutex 					sync.Mutex
}

func NewRunServer(a agent , runNums int, chanLens int,msec int64) *RunServer {
	return &RunServer{
		a:a,
		runNums:runNums,
		isRun:false,
		msec:msec,
		vchan :make(chan interface{},chanLens),
	}
}

func (ts *RunServer) Send(v interface{})  {
	if ts.isRun == false {
		return
	}
	ts.vchan <- v
}

func (ts *RunServer) Run()  {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	if ts.isRun {
		return
	}
	ts.isRun = true
	for i:=0 ; i<ts.runNums ; i++ {
		go ts.work()
	}
}


func (ts *RunServer) work()  {
	var v interface{}
	if ts.msec > 0 {
		for {
			select {
			case v = <- ts.vchan:
				ts.a.DoWork(v)
			case <-time.After(time.Millisecond * time.Duration(ts.msec)):
				ts.a.TimeOut()
			}
		}
	}else{
		for {
			select {
			case v = <- ts.vchan:
				ts.a.DoWork(v)
			}
		}
	}
}