package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/steelx/go-story-mode/storyNode"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"time"
)

var (
	basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
	storyStart = &storyNode.StoryNode{Text: `
You enter a big cave.

			[ N ]
	[ E ]    
			[ S ]

You see three passages leading out.
North leads into darkness, south has some up hills,
east appears to have some foot trails..
	`}

	darkRoom = storyNode.StoryNode{Text: `
Its pitch dark you cannot see shit 
@_@
	`}
	darkRoomLit = storyNode.StoryNode{Text: `
The dark room is now lit with your lantern. 
			^_-
You can go back south or head out to North
	`}
	lion = storyNode.StoryNode{Text: `
While stumbling around, you get eaten by Lion -O-
	`}
	trapRoom = storyNode.StoryNode{Text: `
You head to East where it seems well traveled.
It's a trap!!
You fall into a pit \m/
	`}
	treasureRoom = storyNode.StoryNode{Text: `
You arrive at a room fill with gold treasures!!
$_$
	`}
	frameRate = 33 * time.Millisecond
)

func init() {
	storyStart.AddChoice("N", "Go North", &darkRoom)
	storyStart.AddChoice("S", "Go South", &darkRoom)
	storyStart.AddChoice("E", "Go East", &trapRoom)

	darkRoom.AddChoice("W", "Go back South", &lion)
	darkRoom.AddChoice("O", "Turn on Lantern", &darkRoomLit)

	darkRoomLit.AddChoice("E", "Go North", &treasureRoom)
	darkRoomLit.AddChoice("S", "Go South", storyStart)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tick := time.Tick(frameRate)
	for !win.Closed() {
		win.Clear(colornames.Firebrick)
		var userInput string
		if win.Pressed(pixelgl.KeyN) {
			userInput = "n"
		}
		if win.Pressed(pixelgl.KeyS) {
			userInput = "s"
		}
		if win.Pressed(pixelgl.KeyO) {
			userInput = "o"
		}
		if win.Pressed(pixelgl.KeyE) {
			userInput = "e"
		}

		select {
		case <-tick:
			storyStart.Render(win)
			storyStart = storyStart.Play(userInput)
		}

		if storyStart.IsEmpty() {
			endTxt := text.New(pixel.V(100, 100), basicAtlas)
			fmt.Fprintln(endTxt, "THE END.")
			endTxt.Draw(win, pixel.IM.Scaled(endTxt.Bounds().Center(), 3))
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
