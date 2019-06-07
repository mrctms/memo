/*
Copyright (C) MarckTomack <marcktomack@tutanota.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	m "./management"
	"flag"
	"fmt"
	"os"
)

var (
	add         = flag.Bool("a", false, "[memo] | To add a memo")
	addShort    = flag.Bool("ash", false, "[long memo] [shorted memo] | Add a shorted memo")
	show        = flag.Bool("s", false, "To show all memo")
	delete      = flag.Bool("d", false, "[position number] | To delete a memo")
	deleteAll   = flag.Bool("da", false, "To delete all memo")
	reveal      = flag.Bool("r", false, "[position number] | Show the complete memo")
	modify      = flag.Bool("m", false, "[position number] [memo] | To edit a memo")
	modifyShort = flag.Bool("msh", false, "[position number] [memo] | To edit the memo behind the shorted memo")
)

func main() {
	m.GetUserHome()
	flag.Parse()
	if *add {
		m.CreateMemo(os.Args[2])
	} else if *delete {
		m.DeleteMemos(os.Args[1:])
	} else if *deleteAll {
		m.DeleteAllMemos()
	} else if *show {
		m.SelectMemo()
	} else if *addShort {
		m.CreateShortMemo(os.Args[2], os.Args[3])
	} else if *reveal {
		m.SelectShortMemo(os.Args[2])
	} else if *modify {
		m.ModifyMemo(os.Args[2], os.Args[3])
	} else if *modifyShort {
		m.ModifyMemoShort(os.Args[2], os.Args[3])
	} else {
		fmt.Println("Something went wrong")
	}
}
