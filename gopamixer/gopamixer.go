package gopamixer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var focusIndex = 0
var App *tview.Application

func Somestuff() {
	//fmt.Println("Hello, World!")

	App = tview.NewApplication()

	headerBox := tview.NewBox().SetBorder(true).SetTitle(" GoPa mixer! ")

	vols := []*tview.Box{DrawVolumeBar("one"), DrawVolumeBar("two"), DrawVolumeBar("three")}

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(headerBox, 3, 1, false)

	for _, vol := range vols {
		flex.AddItem(vol, 7, 1, true)
	}
	flex.AddItem(tview.NewBox(), 0, 1, false) // fills the rest of the space

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyUp {
			if focusIndex > 0 {
				focusIndex--
				App.SetFocus(vols[focusIndex])
			}
			return nil
		} else if event.Key() == tcell.KeyDown {
			if focusIndex < len(vols)-1 {
				focusIndex++
				App.SetFocus(vols[focusIndex])
			}
			return nil
		} else if event.Key() == tcell.KeyRune {
			if event.Rune() == 'q' {
				App.Stop()
				return nil
			}
		}
		return event
	})

	if err := App.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func DrawVolumeBar(name string) *tview.Box {
	box := tview.NewBox().SetBorder(true).SetTitle(" Some Volume: " + name + " ")
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
	/*box.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		panic("ohnoo")
		//return event
	})*/
	return box
}
