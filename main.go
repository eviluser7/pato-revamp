package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/eviluser7/pato_revamped/resources/fonts"
	"github.com/eviluser7/pato_revamped/resources/images"
	"github.com/eviluser7/pato_revamped/resources/sounds"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	duckLeftImage  *ebiten.Image
	duckLeftWalk   *ebiten.Image
	duckLeftPond   *ebiten.Image
	duckRightImage *ebiten.Image
	duckRightWalk  *ebiten.Image
	duckRightPond  *ebiten.Image
	tilesImage     *ebiten.Image
	shadowImage    *ebiten.Image
	pondImage      *ebiten.Image
	breadImage     *ebiten.Image
	breadHud       *ebiten.Image
	timeHud        *ebiten.Image
	sign           *ebiten.Image
	/////////
	menuImage       *ebiten.Image
	version         *ebiten.Image
	mplusNormalFont font.Face
	bevanFont       font.Face
	button          *ebiten.Image
	info            *ebiten.Image
	/////////
	loseText     *ebiten.Image
	winText      *ebiten.Image
	instructions *ebiten.Image
	backBtn      *ebiten.Image

	// Sound related
	audioContext = audio.NewContext(44100)
	quackWav     *audio.Player
	breadGrab    *audio.Player
	victory      *audio.Player
	lose         *audio.Player
	ambience     *audio.Player

	// Window
	icon16 image.Image
	icon32 image.Image
	icon48 image.Image
)

const (
	screenWidth  = 1280
	screenHeight = 720
	tileSize     = 32
	tileXNum     = 2 // Tiles per row
	/////////
	creditText  = `Made by eviluser7 in 2020`
	restartText = `Press 'R' to retry.`
	leaveText   = `Press 'L' to leave.`
)

func loadResources() {
	var err error

	// Icon
	f16, err := os.Open("icon16.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f16.Close()

	icon16, err = png.Decode(f16)
	if err != nil {
		log.Fatal(err)
	}

	f32, err := os.Open("icon32.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f32.Close()

	icon32, err = png.Decode(f32)
	if err != nil {
		log.Fatal(err)
	}

	f48, err := os.Open("icon48.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f48.Close()

	icon48, err = png.Decode(f48)
	if err != nil {
		log.Fatal(err)
	}

	// Sounds
	quackS, err := wav.Decode(audioContext, bytes.NewReader(sounds.Quack_wav))
	quackWav, err = audio.NewPlayer(audioContext, quackS)

	breadS, err := wav.Decode(audioContext, bytes.NewReader(sounds.EatBread_wav))
	breadGrab, err = audio.NewPlayer(audioContext, breadS)

	victS, err := wav.Decode(audioContext, bytes.NewReader(sounds.Victory_wav))
	victory, err = audio.NewPlayer(audioContext, victS)

	loseS, err := wav.Decode(audioContext, bytes.NewReader(sounds.GameOver_wav))
	lose, err = audio.NewPlayer(audioContext, loseS)

	ambS, err := wav.Decode(audioContext, bytes.NewReader(sounds.Ambience_wav))
	ambience, err = audio.NewPlayer(audioContext, ambS)

	// Duck images
	imgDuckR, _, err := image.Decode(bytes.NewReader(images.Duck1_png))
	duckRightImage = ebiten.NewImageFromImage(imgDuckR)

	imgDuckL, _, err := image.Decode(bytes.NewReader(images.Duck2_png))
	duckLeftImage = ebiten.NewImageFromImage(imgDuckL)

	// Duck walking images
	duckWalkR, _, err := image.Decode(bytes.NewReader(images.DuckWalkR_png))
	duckRightWalk = ebiten.NewImageFromImage(duckWalkR)

	duckWalkL, _, err := image.Decode(bytes.NewReader(images.DuckWalkL_png))
	duckLeftWalk = ebiten.NewImageFromImage(duckWalkL)

	// Duck pond images
	duckPondR, _, err := image.Decode(bytes.NewReader(images.DuckPondR_png))
	duckRightPond = ebiten.NewImageFromImage(duckPondR)

	duckPondL, _, err := image.Decode(bytes.NewReader(images.DuckPondL_png))
	duckLeftPond = ebiten.NewImageFromImage(duckPondL)

	// World images
	imgTile, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	tilesImage = ebiten.NewImageFromImage(imgTile)

	imgShadow, _, err := image.Decode(bytes.NewReader(images.Shadow_png))
	shadowImage = ebiten.NewImageFromImage(imgShadow)

	imgPond, _, err := image.Decode(bytes.NewReader(images.Water_png))
	pondImage = ebiten.NewImageFromImage(imgPond)

	imgBread, _, err := image.Decode(bytes.NewReader(images.Bread_png))
	breadImage = ebiten.NewImageFromImage(imgBread)

	imgBreadHud, _, err := image.Decode(bytes.NewReader(images.BreadHud_png))
	breadHud = ebiten.NewImageFromImage(imgBreadHud)

	imgTimeHud, _, err := image.Decode(bytes.NewReader(images.TimerHud_png))
	timeHud = ebiten.NewImageFromImage(imgTimeHud)

	imgSign, _, err := image.Decode(bytes.NewReader(images.Sign_png))
	sign = ebiten.NewImageFromImage(imgSign)

	// Menu images
	imgMenu, _, err := image.Decode(bytes.NewReader(images.Menu_png))
	menuImage = ebiten.NewImageFromImage(imgMenu)

	imgVer, _, err := image.Decode(bytes.NewReader(images.Version_png))
	version = ebiten.NewImageFromImage(imgVer)

	imgButton, _, err := image.Decode(bytes.NewReader(images.Button_png))
	button = ebiten.NewImageFromImage(imgButton)

	imgInfo, _, err := image.Decode(bytes.NewReader(images.Info_png))
	info = ebiten.NewImageFromImage(imgInfo)

	// Lose/Win screen
	imgLose, _, err := image.Decode(bytes.NewReader(images.Lose_png))
	loseText = ebiten.NewImageFromImage(imgLose)

	imgWin, _, err := image.Decode(bytes.NewReader(images.Win_png))
	winText = ebiten.NewImageFromImage(imgWin)

	// Instructions
	imgInstructions, _, err := image.Decode(bytes.NewReader(images.Instructions_png))
	instructions = ebiten.NewImageFromImage(imgInstructions)

	imgBack, _, err := image.Decode(bytes.NewReader(images.BackBtn_png))
	backBtn = ebiten.NewImageFromImage(imgBack)

	////////
	// Fonts
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	tt, err = truetype.Parse(fonts.Bevan_ttf)

	bevanFont = truetype.NewFace(tt, &truetype.Options{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Bread code
type Bread struct {
	w       int
	h       int
	x       int
	y       int
	counter int
}

func createBread(g *Game) *Bread {
	bread := &Bread{
		x:       randomInt(64, 1219),
		y:       randomInt(45, 654),
		counter: 180,
	}
	return bread
}

// Duck sprite code
type Duck struct {
	w              int
	h              int
	x              int
	y              int
	direction      int
	isMoving       bool
	animCount      int
	insidePond     bool
	breadCollected int
}

// Update duck position
func (d *Duck) Update(g *Game) {

	// Check normal movement
	if ebiten.IsKeyPressed(ebiten.KeyD) && d.x < (screenWidth-45) {
		if d.insidePond {
			d.x++
		} else {
			d.x += 3
		}
		d.direction = 1
		d.isMoving = true
		d.animCount++
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && d.x > 45 {
		if d.insidePond {
			d.x--
		} else {
			d.x -= 3
		}
		d.direction = 0
		d.isMoving = true
		d.animCount++
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) && d.y > 12 {
		if d.insidePond {
			d.y--
		} else {
			d.y -= 3
		}
		d.isMoving = true
		d.animCount++
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) && d.y < (screenHeight-57) {
		if d.insidePond {
			d.y++
		} else {
			d.y += 3
		}
		d.isMoving = true
		d.animCount++
	}

	// Check two keys at the same time
	if ebiten.IsKeyPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyA) {
		d.x += 0
		d.isMoving = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) && ebiten.IsKeyPressed(ebiten.KeyS) {
		d.y += 0
		d.isMoving = false
	}

	// Check if no key is pressed
	if !ebiten.IsKeyPressed(ebiten.KeyW) &&
		!ebiten.IsKeyPressed(ebiten.KeyA) &&
		!ebiten.IsKeyPressed(ebiten.KeyS) &&
		!ebiten.IsKeyPressed(ebiten.KeyD) {
		d.x += 0
		d.y += 0
		d.isMoving = false
	}

	// Pond collisions
	if d.x >= 130 && d.x <= 298 && // 1st
		d.y >= 312 && d.y <= 363 {
		d.insidePond = true
	} else if d.x >= 706 && d.x <= 838 && // 2nd
		d.y >= 45 && d.y <= 87 {
		d.insidePond = true
	} else if d.x >= 1036 && d.x <= 1156 && // 3rd
		d.y >= 351 && d.y <= 396 {
		d.insidePond = true
	} else if d.x >= 640 && d.x <= 825 && // 4th
		d.y >= 300 && d.y <= 367 {
		d.insidePond = true
	} else if d.x >= 313 && d.x <= 498 && // 5th
		d.y >= 90 && d.y <= 157 {
		d.insidePond = true
	} else if d.x >= 395 && d.x <= 580 && // 6th
		d.y >= 595 && d.y <= 637 {
		d.insidePond = true
	} else if d.x >= 328 && d.x <= 513 && // 7th
		d.y >= 468 && d.y <= 510 {
		d.insidePond = true
	} else if d.x >= 599 && d.x <= 784 && // 8th
		d.y >= 199 && d.y <= 241 {
		d.insidePond = true
	} else if d.x >= 1123 && d.x <= 1308 && // 8th
		d.y >= 570 && d.y <= 612 {
		d.insidePond = true
	} else if d.x >= 115 && d.x <= 300 && // 9th
		d.y >= 450 && d.y <= 492 {
		d.insidePond = true
	} else if d.x >= 1014 && d.x <= 1200 && // 10th
		d.y >= 128 && d.y <= 170 {
		d.insidePond = true
	} else if d.x >= 798 && d.x <= 983 && // 11th
		d.y >= 470 && d.y <= 512 {
		d.insidePond = true
	} else {
		d.insidePond = false
	}

	if d.animCount > 18 {
		d.animCount = 0
	}

	// Check bread collisions
	if d.x < (g.bread.x+g.bread.w)+64 &&
		d.x+d.w > g.bread.x+64 &&
		d.y < (g.bread.y+g.bread.h)+64 &&
		d.h+d.y > g.bread.y+64 {
		g.bread.counter = 0
		d.breadCollected++
		breadGrab.Rewind()
		breadGrab.Play()
	}
}

// Draw duck and change sprite
func (d *Duck) Draw(screen *ebiten.Image) {
	s := duckRightImage

	switch {
	// Idle ground
	case d.direction == 0 && !d.isMoving && !d.insidePond:
		s = duckLeftImage
	case d.direction == 1 && !d.isMoving && !d.insidePond:
		s = duckRightImage

	// Walking animation
	// Left
	case d.direction == 0 && d.isMoving && d.animCount < 10 && !d.insidePond:
		s = duckLeftImage
	case d.direction == 0 && d.isMoving && d.animCount >= 10 && !d.insidePond:
		s = duckLeftWalk

	// Right
	case d.direction == 1 && d.isMoving && d.animCount < 10 && !d.insidePond:
		s = duckRightImage
	case d.direction == 1 && d.isMoving && d.animCount >= 10 && !d.insidePond:
		s = duckRightWalk

	// Pond
	case d.direction == 0 && d.insidePond:
		s = duckLeftPond
	case d.direction == 1 && d.insidePond:
		s = duckRightPond
	}

	op := &ebiten.DrawImageOptions{}

	// Shadow
	op.GeoM.Reset()
	op.GeoM.Translate(-float64(d.w)/2, -float64(d.h)/2) // Anchoring
	switch {                                            // Positioning
	case d.direction == 0 && !d.insidePond:
		op.GeoM.Translate(float64(d.x+28), float64(d.y+87))
	case d.direction == 1 && !d.insidePond:
		op.GeoM.Translate(float64(d.x+7), float64(d.y+87))

	case d.direction == 0 && d.insidePond:
		op.GeoM.Translate(float64(d.x+28), float64(d.y+83))
	case d.direction == 1 && d.insidePond:
		op.GeoM.Translate(float64(d.x+7), float64(d.y+83))
	}

	screen.DrawImage(shadowImage, op)

	// Duck
	op.GeoM.Reset()
	op.GeoM.Translate(-float64(d.w)/2, -float64(d.h)/2) // Anchoring
	op.GeoM.Translate(float64(d.x), float64(d.y))       // Positioning
	screen.DrawImage(s, op)

}

// Game code
type Game struct {
	inited         bool
	duck           *Duck
	bread          *Bread
	op             ebiten.DrawImageOptions
	layers         [][]int
	isFullscreen   bool
	scene          string
	textColor      color.RGBA
	timer          int
	innerTimer     int
	requiredAmount int
	mouseX         int
	mouseY         int
}

// Update game state
func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	g.mouseX = x
	g.mouseY = y

	if !g.inited {
		g.init()
	}

	if g.scene == "menu" {
		ambience.Pause()
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if x >= 565 && x <= 765 && // Button click
				y >= 408 && y <= 508 {
				g.scene = "game"
				g.timer = 100
				g.innerTimer = 60
				g.duck.breadCollected = 0
				g.duck.x = 400
				g.duck.y = 300
				g.requiredAmount = randomInt(10, 30)
				rand.Seed(time.Now().UnixNano())
			} else if x >= 1205 && x <= 1269 && // Instructions click
				y >= 600 && y <= 664 {
				g.scene = "instructions"
			}
		}
	}

	if g.scene == "instructions" {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if x >= 10 && x <= 145 &&
				y >= 620 && y <= 705 {
				g.scene = "menu"
			}
		}
	}

	if g.scene == "game" {
		if !ambience.IsPlaying() {
			ambience.Rewind()
			ambience.Play()
		}
		g.duck.Update(g)
		g.innerTimer--
		if g.innerTimer <= 0 {
			g.timer--
			g.innerTimer = 60
		}
		if g.duck.breadCollected <= g.requiredAmount && g.timer <= 0 {
			lose.Rewind()
			lose.Play()
			g.scene = "finish"
		}
		if g.duck.breadCollected >= g.requiredAmount && g.timer > 0 {
			victory.Rewind()
			victory.Play()
			g.scene = "finish"
		}

	}

	if g.scene == "finish" {
		ambience.Pause()
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.scene = "game"
			g.timer = 100
			g.innerTimer = 60
			g.duck.breadCollected = 0
			g.duck.x = 400
			g.duck.y = 300
			g.requiredAmount = randomInt(10, 30)
			rand.Seed(time.Now().UnixNano())
		} else if inpututil.IsKeyJustPressed(ebiten.KeyL) {
			g.scene = "menu"
		}
	}
	return nil
}

// Draw game
func (g *Game) Draw(screen *ebiten.Image) {

	if g.scene == "menu" {
		const creditX = 965
		const creditY = 700
		g.op.GeoM.Reset() // Title screen
		g.op.GeoM.Translate(float64(0), float64(0))
		screen.DrawImage(menuImage, &g.op)

		g.op.GeoM.Reset() // Version
		g.op.GeoM.Translate(float64(775), float64(175))
		screen.DrawImage(version, &g.op)

		g.op.GeoM.Reset() // Play button
		g.op.GeoM.Translate(float64(565), float64(400))
		screen.DrawImage(button, &g.op)

		g.op.GeoM.Reset() // Info button
		g.op.GeoM.Translate(float64(1205), float64(600))
		screen.DrawImage(info, &g.op)

		// Draw credit text
		text.Draw(screen, creditText, mplusNormalFont, creditX, creditY, color.White)
	}

	if g.scene == "instructions" {
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(0), float64(0)) // Screen
		screen.DrawImage(instructions, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(10), float64(620))
		screen.DrawImage(backBtn, &g.op)
	}

	if g.scene == "game" {

		// Map
		const xNum = screenWidth / tileSize
		for _, l := range g.layers {
			for i, t := range l {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

				sx := (t % tileXNum) * tileSize
				sy := (t / tileXNum) * tileSize
				screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
			}
		}

		// Water ponds
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(120), float64(350))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(1000), float64(390))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(640), float64(342))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(313), float64(132))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(395), float64(637))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(328), float64(510))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(599), float64(241))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(1123), float64(612))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(1014), float64(170))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(115), float64(492))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(798), float64(512))
		screen.DrawImage(pondImage, &g.op)

		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(670), float64(80))
		screen.DrawImage(pondImage, &g.op)

		// Breads
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(-float64(g.bread.w)/2, -float64(g.bread.h)/2) // Anchoring
		g.op.GeoM.Translate(float64(g.bread.x), float64(g.bread.y))       // Position
		if g.bread.counter > 0 {
			g.bread.counter--
			screen.DrawImage(breadImage, &g.op)
		} else {
			g.bread = createBread(g)
		}

		// Draw duck
		g.duck.Draw(screen)

		// Hud
		// Breads collected
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(1050), float64(20))
		screen.DrawImage(breadHud, &g.op)
		text.Draw(screen, fmt.Sprint(g.duck.breadCollected), bevanFont, 1120, 90, color.White)

		// Time
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(1050), float64(120))
		screen.DrawImage(timeHud, &g.op)
		text.Draw(screen, fmt.Sprint(g.timer), bevanFont, 1110, 190, color.White)

		// Amount required
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(float64(0), float64(20))
		screen.DrawImage(sign, &g.op)
		text.Draw(screen, fmt.Sprint(g.requiredAmount), bevanFont, 200, 70, color.White)

		// Leave
		text.Draw(screen, leaveText, mplusNormalFont, 10, 125, color.White)
		if inpututil.IsKeyJustPressed(ebiten.KeyL) {
			g.scene = "menu"
		}
	}

	if g.scene == "finish" {
		if g.duck.breadCollected < g.requiredAmount {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(float64(350), float64(320))
			screen.DrawImage(loseText, &g.op)
		} else {
			g.op.GeoM.Reset()
			g.op.GeoM.Translate(float64(350), float64(320))
			screen.DrawImage(winText, &g.op)
		}
		text.Draw(screen, restartText, bevanFont, 480, 500, color.White)
		text.Draw(screen, leaveText, bevanFont, 480, 570, color.White)
	}
}

// Layout window
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

// Init game
func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	w := 100
	h := 100
	x := 400
	y := 300
	g.duck = &Duck{
		w:              w,
		h:              h,
		x:              x,
		y:              y,
		direction:      1,
		breadCollected: 0,
	}

	g.bread = createBread(g)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	loadResources()

	game := &Game{
		layers: [][]int{
			{
				// First half
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,

				0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
				0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // Half map

				// Second half
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0,
				0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, // Half map
			},
		}, isFullscreen: false,
		scene:          "menu",
		timer:          100,
		innerTimer:     60,
		requiredAmount: randomInt(10, 30),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pato Goes For a Walk")
	ebiten.SetFullscreen(game.isFullscreen)
	ebiten.SetWindowIcon([]image.Image{icon16, icon32, icon48})
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
