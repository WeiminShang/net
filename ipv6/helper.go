// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv6

import (
	"errors"
	"net"
)

var (
	errMissingAddress  = errors.New("missing address")
	errHeaderTooShort  = errors.New("header too short")
	errInvalidConnType = errors.New("invalid conn type")
	errOpNoSupport     = errors.New("operation not supported")
	errNoSuchInterface = errors.New("no such interface")
)

func boolint(b bool) int {
	if b {
		return 1
	}
	return 0
}

func netAddrToIP16(a net.Addr) net.IP {
	switch v := a.(type) {
	case *net.UDPAddr:
		if ip := v.IP.To16(); ip != nil && ip.To4() == nil {
			return ip
		}
	case *net.IPAddr:
		if ip := v.IP.To16(); ip != nil && ip.To4() == nil {
			return ip
		}
	}
	return nil
}
