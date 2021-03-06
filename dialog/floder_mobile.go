// +build ios android

package dialog

import (
	"fyne.io/fyne"
	"fyne.io/fyne/internal/driver/gomobile"
)

func (f *folderDialog) loadPlaces() []fyne.CanvasObject {
	return nil
}

func isHidden(file, _ string) bool {
	return false
}

func folderOpenOSOverride(f *FileDialog) bool {
	gomobile.ShowFileOpenPicker(f.callback.(func(string)), f.filter)
	return true
}

func folderSaveOSOverride(f *FileDialog) bool {
	ShowInformation("File Save", "File save not available on mobile", f.parent)

	callback := f.callback.(func(string))
	if callback != nil {
		callback(nil, nil)
	}

	return true
}
