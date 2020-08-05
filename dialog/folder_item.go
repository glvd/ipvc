package dialog

import (
	"fyne.io/fyne/dialog"
	"image/color"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const (
	fileIconSize      = 64
	fileTextSize      = 24
	fileIconCellWidth = fileIconSize * 1.25
)

type folderDialogItem struct {
	widget.BaseWidget
	picker    *folderDialog
	isCurrent bool

	icon fyne.CanvasObject
	name string
	path string
	dir  bool
}

func (i *folderDialogItem) Tapped(_ *fyne.PointEvent) {
	i.picker.setSelected(i)
	i.Refresh()
}

func (i *folderDialogItem) TappedSecondary(_ *fyne.PointEvent) {
}

func (i *folderDialogItem) CreateRenderer() fyne.WidgetRenderer {
	text := widget.NewLabelWithStyle(i.name, fyne.TextAlignCenter, fyne.TextStyle{})
	text.Wrapping = fyne.TextTruncate

	return &fileItemRenderer{item: i,
		img: i.icon, text: text, objects: []fyne.CanvasObject{i.icon, text}}
}

func fileName(path string) (name string) {
	name = filepath.Base(path)
	ext := filepath.Ext(path)
	name = name[:len(name)-len(ext)]

	return
}

func (i *folderDialogItem) isDirectory() bool {
	return i.dir
}

func (f *folderDialog) newFileItem(path string, dir bool) *folderDialogItem {
	var icon fyne.CanvasObject
	if dir {
		icon = canvas.NewImageFromResource(theme.FolderIcon())
	} else {
		icon = dialog.NewFileIcon(storage.NewURI("file://" + path))
	}
	name := fileName(path)

	ret := &folderDialogItem{
		picker: f,
		icon:   icon,
		name:   name,
		path:   path,
		dir:    dir,
	}
	ret.ExtendBaseWidget(ret)
	return ret
}

type fileItemRenderer struct {
	item *folderDialogItem

	img     fyne.CanvasObject
	text    *widget.Label
	objects []fyne.CanvasObject
}

func (s fileItemRenderer) Layout(size fyne.Size) {
	iconAlign := (size.Width - fileIconSize) / 2
	s.img.Resize(fyne.NewSize(fileIconSize, fileIconSize))
	s.img.Move(fyne.NewPos(iconAlign, 0))

	s.text.Resize(fyne.NewSize(size.Width, fileTextSize))
	s.text.Move(fyne.NewPos(0, fileIconSize+theme.Padding()))
}

func (s fileItemRenderer) MinSize() fyne.Size {
	return fyne.NewSize(fileIconSize, fileIconSize+fileTextSize+theme.Padding())
}

func (s fileItemRenderer) Refresh() {
	canvas.Refresh(s.item)
}

func (s fileItemRenderer) BackgroundColor() color.Color {
	if s.item.isCurrent {
		return theme.PrimaryColor()
	}
	return theme.BackgroundColor()
}

func (s fileItemRenderer) Objects() []fyne.CanvasObject {
	return s.objects
}

func (s fileItemRenderer) Destroy() {
}