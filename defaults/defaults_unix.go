//go:build !windows
// +build !windows

/*
   Copyright The containerd Authors.

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

package defaults

const (
	// DefaultRootDir is the default location used by containerd to store
	// persistent data
	DefaultRootDir = "/mnt/user/0/emulated/0/k3s/var/lib/containerd"
	// DefaultStateDir is the default location used by containerd to store
	// transient data
	DefaultStateDir = "/mnt/user/0/emulated/0/k3s/run/containerd"
	// DefaultAddress is the default unix socket address
	DefaultAddress = "/mnt/user/0/emulated/0/k3s/run/containerd/containerd.sock"
	// DefaultDebugAddress is the default unix socket address for pprof data
	DefaultDebugAddress = "/mnt/user/0/emulated/0/k3s/run/containerd/debug.sock"
	// DefaultFIFODir is the default location used by client-side cio library
	// to store FIFOs.
	DefaultFIFODir = "/mnt/user/0/emulated/0/k3s/run/containerd/fifo"
	// DefaultRuntime is the default linux runtime
	DefaultRuntime = "io.containerd.runc.v2"
	// DefaultConfigDir is the default location for config files.
	DefaultConfigDir = "/mnt/user/0/emulated/0/k3s/etc/containerd"
)
