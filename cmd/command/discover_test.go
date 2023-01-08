package command

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"

	"github.com/amimof/huego"
)

type MockDiscoverer struct{}

func (dd *MockDiscoverer) Discover() ([]huego.Bridge, error) {
	bridge1 := huego.Bridge{Host: "127.0.0.1", ID: "mock-bridge-id-1"}
	bridge2 := huego.Bridge{Host: "192.168.2.59", ID: "mock-bridge-id-2"}
	bridges := []huego.Bridge{bridge1, bridge2}

	return bridges, nil
}

var _ Discoverer = (*MockDiscoverer)(nil) // Assure that when new method is added to interface, compiler will scream

func TestNewDiscoverCmd(t *testing.T) {
	md := &MockDiscoverer{}
	cmd := NewDiscoverCmd(md)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(out), "127.0.0.1")
	assert.Contains(t, string(out), "mock-bridge-id-1")
	assert.Contains(t, string(out), "192.168.2.59")
	assert.Contains(t, string(out), "mock-bridge-id-2")
}
