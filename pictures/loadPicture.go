package pictures

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"os"
)

/*
e.g. usage :

	heroFrames := pictures.LoadAsFrames(imgSprite, 16)
	heroSprite := pixel.NewSprite(imgSprite, heroFrames[10])
	scaledMatrix := pixel.IM.Scaled(pixel.ZV, 16)
	heroSprite.Draw(win, scaledMatrix.Moved(win.Bounds().Center()))

	*******************************OR********************************
	*** below will render sprite everytime you click on Window    ***
	*****************************************************************
	heroFrames := pictures.LoadAsFrames(imgSprite, 16)
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		tree := pixel.NewSprite(imgSprite, heroFrames[rand.Intn(len(heroFrames))])
		trees = append(trees, tree)
		matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(win.MousePosition()))
	}
	for i, tree := range trees {
		tree.Draw(win, matrices[i])
	}
*/

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func LoadAsFrames(imgSprite pixel.Picture, pixelSize float64) []pixel.Rect {
	var spriteFrames []pixel.Rect
	for x := imgSprite.Bounds().Min.X; x < imgSprite.Bounds().Max.X; x += pixelSize {
		for y := imgSprite.Bounds().Min.Y; y < imgSprite.Bounds().Max.Y; y += pixelSize {
			spriteFrames = append(spriteFrames, pixel.R(x, y, x+pixelSize, y+pixelSize))
		}
	}
	//e.g. pixel.NewSprite(imgSprite, spriteFrames[frameIndex])
	return spriteFrames
}
