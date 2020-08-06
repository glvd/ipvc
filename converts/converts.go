package converts

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/glvd/ipcv/config"
	"github.com/glvd/ipcv/dialog"
	"github.com/glvd/ipcv/i18n"
)

// Settings gives access to user interfaces to control Fyne settings
type Converts struct {
	//fyneSettings app.SettingsSchema
	config     config.Config
	lang       i18n.Converts
	inputPath  string
	outputPath string
}

// NewConverts returns a new settings instance with the current configuration loaded
func NewConverts(language i18n.Converts) *Converts {
	s := &Converts{
		config: config.Mirror(),
		lang:   language,
	}
	return s
}

// LoadConvertScreen returns the icon for converts
func (c *Converts) ConvertIcon() fyne.Resource {
	return theme.NewThemedResource(convertIcon, nil)
}

// LoadConvertScreen creates a new convert screen to handle appearance configuration
func (c *Converts) LoadConvertScreen(w fyne.Window) fyne.CanvasObject {
	//------------------------------SettingSystem------------------------------//
	input := c.makeInputConvert(w)
	output := c.makeOutputConvert(w)
	//themes := c.makeThemeSetting(c.config.SettingSystem.Setting.ThemeLabel)
	system := widget.NewGroup(c.lang.Input.Title, input, output)

	bottom := widget.NewHBox(layout.NewSpacer(),
		&widget.Button{Text: "Run", Style: widget.PrimaryButton, OnTapped: func() {
			//_, err := config.Update(func(config *config.Config) {
			//	*config = c.config
			//})
			//if err != nil {
			//	fyne.LogError("failed on update", err)
			//}
			//err = c.save()
			//if err != nil {
			//	fyne.LogError("failed on saving", err)
			//}
		}})

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(system, bottom, nil, nil),
		system, bottom)
}

func (c *Converts) makeInputConvert(w fyne.Window) fyne.CanvasObject {
	//label := widget.NewLabel(c.lang.Input.Label)
	text := widget.NewEntry()
	text.Disable()
	button := widget.NewButton(c.lang.Input.Button, func() {
		dialog.ShowFloderOpen(func(s string, err error) {
			if len(s) > 60 {
				c.inputPath = s
				s = s[0:60] + "..."
			}
			text.SetText(s)
		}, w)
	})
	inputItem := widget.NewFormItem(c.lang.Input.Label, text)
	box := widget.NewHBox(layout.NewSpacer(), button)
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewForm(inputItem), box)
}
func (c *Converts) makeOutputConvert(w fyne.Window) fyne.CanvasObject {
	text := widget.NewEntry()
	text.Disable()
	button := widget.NewButton(c.lang.Output.Button, func() {
		dialog.ShowFloderOpen(func(s string, err error) {
			if len(s) > 60 {
				c.outputPath = s
				s = s[0:60] + "..."
			}
			text.SetText(s)
		}, w)
	})
	inputItem := widget.NewFormItem(c.lang.Output.Label, text)
	box := widget.NewHBox(layout.NewSpacer(), button)
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewForm(inputItem), box)
}
