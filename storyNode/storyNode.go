package storyNode

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/steelx/go-story-mode/pictures"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"math"
	"math/rand"
	"strings"
	"time"
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

var (
	camPos       = pixel.ZV
	camSpeed     = 500.0
	camZoom      = 1.0
	camZoomSpeed = 1.2
	last         = time.Now()
	trees        []*pixel.Sprite
	matrices     []pixel.Matrix
)

func (story *StoryNode) Play(win *pixelgl.Window) {
	win.Clear(colornames.Firebrick)

	imgSprite, err := pictures.LoadPicture("./resources/old_hero_16x16.png")
	if err != nil {
		panic(err)
	}

	// Move camera with arrow keys
	// delta time for fixing inconsistencies
	dt := time.Since(last).Seconds()
	last = time.Now()
	if win.Pressed(pixelgl.KeyLeft) {
		camPos.X -= camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyRight) {
		camPos.X += camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		camPos.Y -= camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyUp) {
		camPos.Y += camSpeed * dt
	}
	// camera Zoom function
	camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)
	camScaledMatrix := pixel.IM.Scaled(camPos, camZoom)

	// Create Camera at 0,0
	camera := camScaledMatrix.Moved(win.Bounds().Center().Sub(camPos))
	win.SetMatrix(camera)

	heroFrames := pictures.LoadAsFrames(imgSprite, 16)
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		tree := pixel.NewSprite(imgSprite, heroFrames[rand.Intn(len(heroFrames))])
		trees = append(trees, tree)
		mouse := camera.Unproject(win.MousePosition()) //fix mouse clicks with new Camera view
		matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
	}
	for i, tree := range trees {
		tree.Draw(win, matrices[i])
	}

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

	win.Update()
	if len(story.choices) != 0 {
		story.ExecuteCMD(userInput).Play(win)
	}

	basicTxt := text.New(pixel.V(100, 100), basicAtlas)
	fmt.Fprintln(basicTxt, "THE END.")
	basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 3))
}
