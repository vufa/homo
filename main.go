//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package main

import (
	"github.com/aiicy/aiicy/cmd"
	_ "github.com/aiicy/aiicy/master/engine/native"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: use pgx instead of sqlite3
func main() {
	cmd.Execute()
}
