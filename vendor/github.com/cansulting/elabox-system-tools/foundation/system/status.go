// Copyright 2021 The Elabox Authors
// This file is part of the elabox-system-tools library.

// The elabox-system-tools library is under open source LGPL license.
// If you simply compile or link an LGPL-licensed library with your own code,
// you can release your application under any license you want, even a proprietary license.
// But if you modify the library or copy parts of it into your code,
// you’ll have to release your application under similar terms as the LGPL.
// Please check license description @ https://www.gnu.org/licenses/lgpl-3.0.txt

// This file is used for setting and getting the current system status

package system

import "os"

type Status string

const (
	STOPPED     = "inactive"
	RUNNING     = "active"
	BOOTING     = "booting"
	INIT_UPDATE = "init_update"
	UPDATING    = "updating"
)

func GetStatus() string {
	return os.Getenv("elastatus")
}

func SetStatus(status string) {
	os.Setenv("elastatus", status)
}
