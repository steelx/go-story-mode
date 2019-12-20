package storyNode

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"strings"
)

var (
	basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
)

type Choice struct {
	cmd string
	description string
	nextNode *StoryNode
	nextChoice *Choice
}

type StoryNode struct {
	Text    string
	choices *Choice
}

func (story *StoryNode) AddChoice(cmd, description string, nextNode *StoryNode) {
	newChoice := &Choice{
		cmd:         cmd,
		description: description,
		nextNode:    nextNode,
		nextChoice:  nil,
	}

	if story.choices == nil {
		story.choices = newChoice
	} else {
		currentChoice := story.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = newChoice
	}
}

func (story StoryNode) Render(win *pixelgl.Window) {
	basicTxt := text.New(pixel.V(10, 250), basicAtlas)
	fmt.Fprintln(basicTxt, story.Text)
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1))

	currentChoice := story.choices
	for currentChoice != nil {
		fmt.Fprintf(basicTxt, "%s : %s \n", currentChoice.cmd, currentChoice.description)
		currentChoice = currentChoice.nextChoice
		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1))
	}
}

func matchStrings(str1, str2 string) bool {
	return strings.ToLower(str1) == strings.ToLower(str2)
}
func (story *StoryNode) ExecuteCMD(cmd string) *StoryNode {
	basicTxt := text.New(pixel.V(10, 250), basicAtlas)
	currentChoice := story.choices
	for currentChoice != nil {
		if matchStrings(currentChoice.cmd, cmd) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Fprintln(basicTxt,"Sorry, didn't understand that CMD please enter again!")
	return story
}

func (story *StoryNode) Play(win *pixelgl.Window) {
	win.Clear(colornames.Black)
	story.Render(win)
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
	if win.Pressed(pixelgl.KeyR) {
		userInput = "r"
	}
	win.Update()
	if story.choices != nil {
		story.ExecuteCMD(userInput).Play(win)
	}

	basicTxt := text.New(pixel.V(100, 100), basicAtlas)
	fmt.Fprintln(basicTxt, "THE END.")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 3))
}

var (
	Width = 480
	Height = 320
)
func PercentWidth(pcent int) int {
	percent := (Width * pcent) / 100
	return percent
}
func PercentHeight(pcent int) int {
	percent := (Height * pcent) / 100
	return percent
}