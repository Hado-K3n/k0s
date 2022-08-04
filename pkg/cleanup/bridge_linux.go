/*
Copyright 2022 k0s authors

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

package cleanup

import (
	"fmt"
	"runtime"

	"github.com/vishvananda/netlink"
)

type bridge struct{}

// Name returns the name of the step
func (b *bridge) Name() string {
	return "kube-bridge leftovers cleanup step"
}

// Run removes found kube-bridge leftovers
func (b *bridge) Run() error {
	if runtime.GOOS == "windows" {
		return nil
	}

	lnks, err := netlink.LinkList()
	if err != nil {
		return fmt.Errorf("failed to get link list from netlink: %v", err)
	}

	for _, l := range lnks {
		if l.Attrs().Name == "kube-bridge" {
			err := netlink.LinkDel(l)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
