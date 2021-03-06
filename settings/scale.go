package settings

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type scaleItems struct {
	scale float32
	name  string
	//preview *canvas.Text
	//button *widget.Button
}

func (s *Settings) appliedScale(value float32) {
	//for _, scale := range scales {
	//scale.preview.TextSize = int(float32(theme.TextSize()) * scale.scale / value)
	//}
}

func (s *Settings) chooseScale(value string) {
	for _, scale := range s.scaleItems {
		if scale.name == value {
			s.config.System.Setting.Scale = scale.scale
		}
	}
}
func (s *Settings) makeScaleSelect(sc float32) *widget.Select {
	var scaleNames []string
	selected := ""
	for _, scale := range s.scaleItems {
		scaleNames = append(scaleNames, scale.name)
		if sc == scale.scale {
			selected = scale.name
		}
	}
	scaleSelect := widget.NewSelect(scaleNames, func(v string) {
		s.chooseScale(v)
	})
	scaleSelect.SetSelected(selected)
	return scaleSelect
}

func (s *Settings) makeScaleSetting(scale float32) fyne.CanvasObject {
	scaleLabel := widget.NewLabel(s.lang.System.ScaleLabel)
	scaleSelect := s.makeScaleSelect(scale)
	return fyne.NewContainerWithLayout(layout.NewGridLayout(2), scaleLabel, scaleSelect)
}

// refreshMonitor is a simple widget that updates canvas components when the UI is asked to refresh.
// Captures theme and scale changes without the settings monitoring code.
type refreshMonitor struct {
	widget.Label
	settings *Settings
}

func (r *refreshMonitor) Refresh() {
	//r.settings.refreshScalePreviews()
	r.Label.Refresh()
}

func newRefreshMonitor(s *Settings) *refreshMonitor {
	r := &refreshMonitor{settings: s}
	r.Hide()
	return r
}
