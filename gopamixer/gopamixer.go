package gopamixer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Somestuff() {
	//fmt.Println("Hello, World!")

	box := tview.NewBox().SetBorder(true).SetTitle(" GoPa mixer! ")
	box.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {

		ypos := y + 2
		for cx := x + 2; cx < width-2; cx++ {
			screen.SetContent(cx, ypos, tview.BlockLowerOneEighthBlock, nil, tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorBlack))
		}
		ypos++

		third := (width - 4) / 3
		firstthird := third + ((width - 4) % 3)
		cx := x + 2
		for ; cx < firstthird; cx++ {
			screen.SetContent(cx, ypos, tview.BlockDarkShade, nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGreen))
		}
		for ; cx < firstthird+third; cx++ {
			screen.SetContent(cx, ypos, tview.BlockDarkShade, nil, tcell.StyleDefault.Foreground(tcell.ColorDarkGoldenrod))
		}
		for ; cx < width-2; cx++ {
			screen.SetContent(cx, ypos, tview.BlockDarkShade, nil, tcell.StyleDefault.Foreground(tcell.ColorDarkRed))
		}
		ypos++
		for cx := x + 2; cx < width-2; cx++ {
			screen.SetContent(cx, ypos, tview.BlockUpperOneEighthBlock, nil, tcell.StyleDefault.Foreground(tcell.ColorGray).Background(tcell.ColorBlack))
		}
		ypos++

		return x + 1, ypos, width - 2, height - 5
	})
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
