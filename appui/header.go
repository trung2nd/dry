package appui

import (
	"fmt"

	gizaktermui "github.com/gizak/termui"
	"github.com/moncho/dry/docker"
	"github.com/moncho/dry/ui"
	"github.com/moncho/dry/ui/termui"
)

//SortableColumnHeader is a column header associated to a  sort mode
type SortableColumnHeader struct {
	Title string // Title to display in the tableHeader.
	Mode  docker.SortMode
}

//WidgetHeader is a widget that renders a line with the result of
//appending the given what, count and details in a common format.
func WidgetHeader(what string, howMany int, details string) *termui.MarkupPar {
	par := termui.NewParFromMarkupText(DryTheme,
		fmt.Sprintf(
			"<b><blue>%s: </><yellow>%d</></>", what, howMany)+" "+details)

	par.SetX(0)
	par.Border = false
	par.Width = ui.ActiveScreen.Dimensions.Width
	par.TextBgColor = gizaktermui.Attribute(DryTheme.Bg)
	par.Bg = gizaktermui.Attribute(DryTheme.Bg)

	return par
}
