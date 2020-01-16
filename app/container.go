/*
@Time : 2019/12/23 10:24
@Software: GoLand
@File : aside
@Author : Bingo <airplayx@gmail.com>
*/
package _app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"gormat/app/config"
	"gormat/app/sql2struct"
	"gormat/internal/Sql2struct"
	"time"
)

func Container(app fyne.App, win fyne.Window) *widget.TabContainer {
	var options = Sql2struct.Configs()
	var dbBox, iPBox = widget.NewTabContainer(), widget.NewTabContainer()
	currentDB := make(chan *widget.TabItem)
	for _, v := range options.SourceMap {
		for _, curDb := range v.Db {
			dbBox.Items = append(dbBox.Items, widget.NewTabItemWithIcon(
				curDb, config.Database,
				sql2struct.Screen(win, []interface{}{
					v.User,
					v.Password,
					v.Host,
					v.Port,
					curDb,
				})))
		}
		dbBox.SetTabLocation(widget.TabLocationLeading)
		iPBox.Items = append(iPBox.Items, widget.NewTabItem(v.Host, dbBox))
	}
	toolBar := ToolBar(win, iPBox, dbBox, options)
	s2sBox := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolBar, nil, nil, nil),
		toolBar,
		WelcomeScreen(),
	)
	go func() {
		for {
			time.Sleep(time.Microsecond * 200)
			if <-currentDB != dbBox.CurrentTab() {

			}
		}
	}()
	if len(iPBox.Items) > 0 {
		go func() {
			currentDB <- &widget.TabItem{}
			for {
				currentDB <- dbBox.CurrentTab()
			}
		}()
		iPBox.SetTabLocation(widget.TabLocationLeading)
		s2sBox.AddObject(iPBox)
	}
	c := widget.NewTabContainer(
		//widget.NewTabItemWithIcon("", config.Home, WelcomeScreen()),
		//widget.NewTabItemWithIcon("", theme.SettingsIcon(), _app.SettingScreen(app, win)),
		widget.NewTabItemWithIcon("", config.Store, s2sBox),
		widget.NewTabItemWithIcon("", config.Daily, fyne.NewContainer()),
		//widget.NewTabItemWithIcon("", config.Video, fyne.NewContainer()),
	)
	c.SetTabLocation(widget.TabLocationBottom)
	return c
}
