package cmd

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/config"
	"github.com/oltarzewskik/tibivi/pkg/keybindings"
	"github.com/oltarzewskik/tibivi/pkg/layout"
	"github.com/oltarzewskik/tibivi/pkg/layout/utils"
	"github.com/oltarzewskik/tibivi/pkg/schedule"
)

// Run runs Tibivi
func Run() error {
	if err := common.SetCommonVars(); err != nil {
		return err
	}
	defer common.G.Close()

	if err := config.CreateDotTibivi(); err != nil {
		return err
	}

	if err := schedule.SetSchedule(); err != nil {
		return err
	}

	common.G.Highlight = true
	common.G.InputEsc = true
	common.G.SelFgColor = gocui.ColorGreen
	common.G.SelectedDay = common.CurrentDay

	common.G.SetManagerFunc(layout.Layout)

	go layout_utils.UpdateLayoutOnCurrentBlockChange()

	go layout_utils.UpdateLayoutOnResize()

	if err := keybindings.Keybindings(); err != nil {
		return err
	}

	if err := common.G.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}
