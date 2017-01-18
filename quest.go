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
	PAUSE time.Duration = time.Second / 2
)

func init() {
	fmt.Println(
		`Welcome to Gopher Quest!
If you don't know how to play, type 'info' now; otherwise, press enter.`)
	input("")
}

func main() {
	fmt.Println("THIS IS THE GAME LOOP")
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
Press enter to start the game.`)
	time.Sleep(PAUSE)
}

func wait() {
	time.Sleep(PAUSE)
	fmt.Scanln()
}
