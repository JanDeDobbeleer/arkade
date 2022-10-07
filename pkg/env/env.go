// Copyright (c) arkade author(s) 2022. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package env

import (
	"os"
	"path"
)

func GetUserHome() string {
	home := os.Getenv("HOME")
	if home != "" {
		return home
	}
	home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}
	return home
}

func LocalBinary(name, subdir string) string {
	home := GetUserHome()
	val := path.Join(home, ".arkade/bin/")
	if len(subdir) > 0 {
		val = path.Join(val, subdir)
	}

	return path.Join(val, name)
}
