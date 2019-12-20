package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/steelx/go-story-mode/storyNode"
)

var (
	storyStart = storyNode.StoryNode{Text: `
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
)

func init() {
	storyStart.AddChoice("N", "Go North", &darkRoom)
	storyStart.AddChoice("S", "Go South", &darkRoom)
	storyStart.AddChoice("E", "Go East", &trapRoom)

	darkRoom.AddChoice("S", "Go back South", &lion)
	darkRoom.AddChoice("O", "Turn on Lantern", &darkRoomLit)

	darkRoomLit.AddChoice("N", "Go North", &treasureRoom)
	darkRoomLit.AddChoice("S", "Go South", &storyStart)
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

	storyStart.Play(win)

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
