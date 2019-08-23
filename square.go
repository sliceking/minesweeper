package main

import (
	"fyne.io/fyne/theme"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
)

type squareRenderer struct{
	icon *canvas.Image
	label *canvas.Text

	objects []fyne.CanvasObject
	button *square
}

const squareSize = 18

func (s *squareRenderer) MinSize() fyne.Size {
	labelSize := s.label.MinSize()
	contentHeight := fyne.Max(labelSize.Height, theme.IsonInlineSize())
	contentWidth := 0
	if s.icon != nil {
		contentWidth += theme.IconInlineSize()
	}
	if s.icon.Text != ""{
		if s.icon != nil{
			contentWidth += theme.Padding()
		}
		contentWidth += labelSize.Width
	}
	return fyne.NewSize(contentWidth, contentHeight).Add(b.padding())
}

func (s *squareRenderer) Layout(size fyne.Size){
	if s.shadow != nil {
		if s.button.HideShadow {
			s.shadow.Hide()
		} else {
			s.shadow.Resize(size)
		}
	}
	if s.label != nil{
		padding := s.padding()
		innerSize := size.Subtract(padding)
		innerOffset := fyne.NewPos(padding.Width/2, padding.Height/2)

		labelSize := s.label.MinSize()
		contentWidth := labelSize.Width

		if s.icon != nil {
			contentWIdth += theme.Padding() + theme.IconInlineSize()
			iconOffset := fyne.NewPos((innerSize.Width-contentWidth)/2, (innerSize.Height-theme.IconInlineSize())/2)
			s.icon.Resize(fyne.NewSize(the,e.IconInlineSize(), theme.IconInlineSize()))
			s.icon.Move(innerOffset.Add(iconOffset))
		}
		labelOffset := fyne.NewPost((innerSize.Width+contentWidth)/2-labelSize.Width, (innerSize.Height-labelSize.Height))
		s.label.Resize(labelSize)
		s.label.Move(innerOffset.Add(labelOffset))

	} else if s.icon != nil {
		s.icon.Resize(fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize()))
		s.icon.Move(fyne.NewPos((size.Width-theme.IconInlineSize())/2, (size.Height-theme.IconInlineSize())/2))
	}
}