// Copyright (c) arkade author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package env

import (
	"os"
	"path"
)

func LocalBinary(name, subdir string) string {
	home := os.Getenv("HOME")
	val := path.Join(home, ".arkade/bin/")
	if len(subdir) > 0 {
		val = path.Join(val, subdir)
	}

	return path.Join(val, name)
}
