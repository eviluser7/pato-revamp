//go:generate file2byteslice -package=images -input=./images/bread.png -output=./images/bread.go -var=Bread_png
//go:generate file2byteslice -package=images -input=./images/duck1.png -output=./images/duck1.go -var=Duck1_png
//go:generate file2byteslice -package=images -input=./images/duck2.png -output=./images/duck2.go -var=Duck2_png
//go:generate file2byteslice -package=images -input=./images/duckPondL.png -output=./images/duckpondl.go -var=DuckPondL_png
//go:generate file2byteslice -package=images -input=./images/duckPondR.png -output=./images/duckpondr.go -var=DuckPondR_png
//go:generate file2byteslice -package=images -input=./images/duckWalkL.png -output=./images/duckwalkl.go -var=DuckWalkL_png
//go:generate file2byteslice -package=images -input=./images/duckWalkR.png -output=./images/duckwalkr.go -var=DuckWalkR_png
//go:generate file2byteslice -package=images -input=./images/menu.png -output=./images/menu.go -var=Menu_png
//go:generate file2byteslice -package=images -input=./images/version.png -output=./images/version.go -var=Version_png
//go:generate file2byteslice -package=images -input=./images/shadow.png -output=./images/shadow.go -var=Shadow_png
//go:generate file2byteslice -package=images -input=./images/tiles.png -output=./images/tiles.go -var=Tiles_png
//go:generate file2byteslice -package=images -input=./images/water.png -output=./images/water.go -var=Water_png
//go:generate file2byteslice -package=images -input=./images/hud_bread.png -output=./images/hud_bread.go -var=BreadHud_png
//go:generate file2byteslice -package=images -input=./images/timer_hud.png -output=./images/timer_hud.go -var=TimerHud_png
//go:generate file2byteslice -package=images -input=./images/lose_text.png -output=./images/lose_text.go -var=Lose_png
//go:generate file2byteslice -package=images -input=./images/win_text.png -output=./images/win_text.go -var=Win_png
//go:generate file2byteslice -package=images -input=./images/button.png -output=./images/button.go -var=Button_png
//go:generate file2byteslice -package=images -input=./images/sign.png -output=./images/sign.go -var=Sign_png
//go:generate file2byteslice -package=images -input=./images/instructions.png -output=./images/instructions.go -var=Instructions_png
//go:generate file2byteslice -package=images -input=./images/info.png -output=./images/info.go -var=Info_png
//go:generate file2byteslice -package=images -input=./images/credits_back.png -output=./images/back.go -var=BackBtn_png
//go:generate file2byteslice -package=fonts -input=./fonts/mplus-1p-regular.ttf -output=./fonts/mplus1pregular.go -var=MPlus1pRegular_ttf
//go:generate file2byteslice -package=fonts -input=./fonts/Bevan.ttf -output=./fonts/bevan.go -var=Bevan_ttf
//go:generate file2byteslice -package=sounds -input=./sounds/ambience.wav -output=./sounds/ambience.go -var=Ambience_wav
//go:generate file2byteslice -package=sounds -input=./sounds/eat_bread.wav -output=./sounds/eat_bread.go -var=EatBread_wav
//go:generate file2byteslice -package=sounds -input=./sounds/gameover.wav -output=./sounds/gameover.go -var=GameOver_wav
//go:generate file2byteslice -package=sounds -input=./sounds/quack.wav -output=./sounds/quack.go -var=Quack_wav
//go:generate file2byteslice -package=sounds -input=./sounds/victory.wav -output=./sounds/victory.go -var=Victory_wav
//go:generate gofmt -s -w .

package resources

import (
	// Dummy imports for go.mod for some Go files with 'ignore' tags. For example, `go mod tidy` does not
	// recognize Go files with 'ignore' build tag.
	//
	// Note that this affects only importing this package, but not 'file2byteslice' commands in //go:generate.
	_ "github.com/hajimehoshi/file2byteslice"
)
