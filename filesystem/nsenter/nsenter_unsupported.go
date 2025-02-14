//go:build !linux
// +build !linux

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nsenter

import (
	"context"
	"fmt"

	"github.com/Valdenirmezadri/go-utils/exec"
)

const (
	// DefaultHostRootFsPath is path to host's filesystem mounted into container
	// with kubelet.
	DefaultHostRootFsPath = "/rootfs"
)

// Nsenter is a type alias for backward compatibility
type Nsenter = NSEnter

// NSEnter is part of experimental support for running the kubelet
// in a container.
type NSEnter struct {
	// a map of commands to their paths on the host filesystem
	Paths map[string]string
}

// NewNsenter constructs a new instance of NSEnter
func NewNsenter(hostRootFsPath string, executor exec.Interface) (*Nsenter, error) {
	return &Nsenter{}, nil
}

// Exec executes nsenter commands in hostProcMountNsPath mount namespace
func (ne *NSEnter) Exec(cmd string, args []string) exec.Cmd {
	return nil
}

// AbsHostPath returns the absolute runnable path for a specified command
func (ne *NSEnter) AbsHostPath(command string) string {
	return ""
}

// SupportsSystemd checks whether command systemd-run exists
func (ne *NSEnter) SupportsSystemd() (string, bool) {
	return "", false
}

// Command returns a command wrapped with nenter
func (ne *NSEnter) Command(cmd string, args ...string) exec.Cmd {
	return nil
}

// CommandContext returns a CommandContext wrapped with nsenter
func (ne *NSEnter) CommandContext(ctx context.Context, cmd string, args ...string) exec.Cmd {
	return nil
}

// LookPath returns a LookPath wrapped with nsenter
func (ne *NSEnter) LookPath(file string) (string, error) {
	return "", fmt.Errorf("not implemented, error looking up : %s", file)
}

var _ exec.Interface = &NSEnter{}
