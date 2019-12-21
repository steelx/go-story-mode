package storyNode

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"strings"
)

var (
	basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
)

type Choice struct {
	cmd         string
	description string
	nextNode    *StoryNode
}

type StoryNode struct {
	Text    string
	choices []*Choice
}

func (story *StoryNode) AddChoice(cmd, description string, nextNode *StoryNode) {
	newChoice := &Choice{cmd, description, nextNode}
	story.choices = append(story.choices, newChoice)
}

func (story StoryNode) IsEmpty() bool {
	return len(story.choices) == 0
}

func (story StoryNode) Render(win *pixelgl.Window) {
	basicTxt := text.New(pixel.V(50, 250), basicAtlas)
	fmt.Fprintln(basicTxt, story.Text)
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 1))

	descriptionTxt := text.New(pixel.V(10, 50), basicAtlas)
	if story.choices != nil {
		for _, c := range story.choices {
			fmt.Fprintf(descriptionTxt, "%s : %s \n", c.cmd, c.description)
			descriptionTxt.Draw(win, pixel.IM.Scaled(descriptionTxt.Orig, 1))
		}
	}
}

func matchStrings(str1, str2 string) bool {
	return strings.ToLower(str1) == strings.ToLower(str2)
}
func (story *StoryNode) ExecuteCMD(cmd string) *StoryNode {
	for _, c := range story.choices {
		if matchStrings(c.cmd, cmd) {
			return c.nextNode
		}
	}
	return story
}

func (story *StoryNode) Play(userInput string) *StoryNode {
	if len(story.choices) != 0 {
		return story.ExecuteCMD(userInput)
	}
	return story
}
