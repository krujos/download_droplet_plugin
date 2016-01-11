// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/krujos/download_droplet_plugin/droplet"
)

type FakeDroplet struct {
	SaveDropletStub        func(name string, path string) error
	saveDropletMutex       sync.RWMutex
	saveDropletArgsForCall []struct {
		name string
		path string
	}
	saveDropletReturns struct {
		result1 error
	}
}

func (fake *FakeDroplet) SaveDroplet(name string, path string) error {
	fake.saveDropletMutex.Lock()
	fake.saveDropletArgsForCall = append(fake.saveDropletArgsForCall, struct {
		name string
		path string
	}{name, path})
	fake.saveDropletMutex.Unlock()
	if fake.SaveDropletStub != nil {
		return fake.SaveDropletStub(name, path)
	} else {
		return fake.saveDropletReturns.result1
	}
}

func (fake *FakeDroplet) SaveDropletCallCount() int {
	fake.saveDropletMutex.RLock()
	defer fake.saveDropletMutex.RUnlock()
	return len(fake.saveDropletArgsForCall)
}

func (fake *FakeDroplet) SaveDropletArgsForCall(i int) (string, string) {
	fake.saveDropletMutex.RLock()
	defer fake.saveDropletMutex.RUnlock()
	return fake.saveDropletArgsForCall[i].name, fake.saveDropletArgsForCall[i].path
}

func (fake *FakeDroplet) SaveDropletReturns(result1 error) {
	fake.SaveDropletStub = nil
	fake.saveDropletReturns = struct {
		result1 error
	}{result1}
}

var _ droplet.Droplet = new(FakeDroplet)
