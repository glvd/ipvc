package windows

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/i18n"
	"github.com/glvd/ipcv/settings"
)

var Size = fyne.NewSize(800, 600)

type MainFrame struct {
	root     fyne.Window
	language *i18n.Language
}

func New(config config.Config) *MainFrame {
	setLanguageFontEnv(config.System.Language)

	language := i18n.Load(config.System.Language.Name)
	err := i18n.SaveTemplate(language)
	if err != nil {
		return nil
	}
	win := app.New().NewWindow(language.Title)
	s := settings.NewSettings(language.Settings)
	appearance := s.LoadAppearanceScreen(win)
	tabs := widget.NewTabContainer(
		&widget.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance},
		//&widget.TabItem{Text: "Language", Icon: s.LanguageIcon(), Content: language},
	)
	tabs.SetTabLocation(widget.TabLocationLeading)
	win.SetIcon(resourceShipPng)
	win.SetContent(tabs)
	win.Resize(Size)
	return &MainFrame{
		root:     win,
		language: language,
	}
	//app := app.New()
	//s := settings.NewSettings()
	//w := app.NewWindow(Title)
}

func (f *MainFrame) Run() {
	f.root.ShowAndRun()
}
