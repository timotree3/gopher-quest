/*
Copyright (C) 2017 (timotree3, et al.)

This file is part of gopher-quest.

gopher-quest is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

gopher-quest is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with gopher-quest.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"
	"time"
)

const (
	PAUSE      time.Duration = time.Second / 2
	GRIDSIZE   int           = 12   // the width and height of both territory grids
	UNFOUND    rune          = ' '  // for spots not checked yet
	EMPTY      rune          = '\'' // not part of a base/burrow
	HORIZONTAL rune          = '-'  // part of a base/burrow going side-to-side
	VERTICAL   rune          = '|'  // part of a base/burrow going up and down
	TOP        rune          = '^'  // an entrance at the top of a base/burrow
	BOTTOM     rune          = 'v'  // an entrance at the bottom of a base/burrow
	RIGHT      rune          = '>'  // an entrance at the right end of a base/burrow
	LEFT       rune          = '<'  // an entrance at the left end of a base/burrow
	YAXIS      bool          = false
	XAXIS      bool          = true
)

type cell struct {
	found    bool // whether it should be rendered
	rendered rune // one of (O, -, |, ^, v, >, <)
}

type grid [GRIDSIZE][GRIDSIZE]cell

// the territory grids
var gopherland grid
var cowlair grid

func init() {
	for y := 0; y < GRIDSIZE; y++ {
		for x := 0; x < GRIDSIZE; x++ {
			gopherland[y][x] = cell{true, EMPTY} // should be {false, EMPTY} may be true for testing purposes
			cowlair[y][x] = cell{true, EMPTY}    // ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
		}

	}
	fmt.Println(
		`Welcome to Gopher Quest!
If you don't know how to play, type 'info' now; otherwise, press enter.`)
	input("") // 'info' can be typed at any time.
}

func main() {
	// for testing features until I implement game flow
	gopherland.make_base(0, 0, YAXIS, GRIDSIZE)
	cowlair.make_base(GRIDSIZE-12, GRIDSIZE-12, YAXIS, 12)
	gopherland.print()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	cowlair.print()
}

func (plot *grid) make_base(y_start int, x_start int, axis bool, distance int) (err error) {
	if axis == YAXIS { // if this is a vertical base
		plot[y_start][x_start].rendered = TOP
		for y := y_start + 1; y < y_start+distance-1; y++ {
			plot[y][x_start].rendered = VERTICAL
		}
		plot[y_start+distance-1][x_start].rendered = BOTTOM
	} else { // if this is a horizontal base
		plot[y_start][x_start].rendered = LEFT
		for x := x_start + 1; x < x_start+distance-1; x++ {
			plot[y_start][x].rendered = HORIZONTAL
		}
		plot[y_start][x_start+distance-1].rendered = RIGHT
	}
	return nil
}

func (plot grid) print() {
	for _, row := range plot {
		for _, spot := range row {
			if spot.found { // if the player should be able to see the spot
				fmt.Print(string(spot.rendered)) // print the spot's real value
			} else {
				fmt.Print(string(UNFOUND)) // print a blank spot
			}
		}
		fmt.Println() // make a new line for each row
	}
}

func input(prompt string) (response string) {
	for {
		response = ""
		fmt.Print(prompt)
		fmt.Scanln(&response)
		if response != "info" {
			return
		}
		info()
	}
}

func info() {
	fmt.Println(
		`You control a nation of Gophers.
Press enter to continue.`)
	wait()
	fmt.Println(
		`Recently, the evil Cows declared war on your nation.
Not only did they capture all of your citizens, but they are also attacking your burrows!
You need to rescue your Gophers before all of your burrows are destroyed!`)
	wait()
	fmt.Println(
		`Each turn, you choose a place to send a rescue party.
If you find nothing, it will be marked on your map as empty.
If you find an entrance to a Cow base, all Gophers being held captive there will be rescued.
If you find a part of a Cow base that they can't enter, it will be marked on your map as part of a base.`)
	wait()
	fmt.Println(
		`After your turn, the evil Cows will try to invade your burrows.
They can only invade a burrow if they find an entrance.
You will be notified by your guards when a Cow squad is near a burrow.`)
	wait()
	fmt.Println(
		`That's how the game works.
If you forget, just type 'info' again at any time.
Press enter to start the game.`)
	time.Sleep(PAUSE)
}

func wait() {
	time.Sleep(PAUSE)
	fmt.Scanln()
}
