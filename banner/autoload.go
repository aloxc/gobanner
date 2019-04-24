// Copyright 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package autoload configure the banner loader with defaults
// Import the package. Thats it.
package banner

import (
	"github.com/mattn/go-colorable"
)

func init() {
	var (
		isEnabled      bool = true
		isColorEnabled bool = true
	)

	Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled)
}
