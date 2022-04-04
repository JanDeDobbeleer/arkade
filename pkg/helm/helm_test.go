// Copyright (c) arkade author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package helm

import "testing"

func Test_GetHelmURL_GitBash(t *testing.T) {
	arch := "amd64"
	os := "windows"

	got := GetHelmURL(arch, os, helmVersion)
	want := "https://get.helm.sh/helm-v3.1.2-windows-amd64.tar.gz"
	if got != want {
		t.Fatalf("want: %s, but got: %s", want, got)
	}
}
