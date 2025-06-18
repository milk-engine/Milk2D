package renderer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) WindowSetTitle(title string) {
	ebiten.SetWindowTitle(title)
}

func (r *Renderer) WindowSetSize(w, h int) {
	ebiten.SetWindowSize(w, h)
}

func (r *Renderer) WindowSetFullscreen(fullscreen bool) {
	ebiten.SetFullscreen(fullscreen)
}

func (r *Renderer) WindowMaximize() {
	ebiten.MaximizeWindow()
}

func (r *Renderer) WindowMinimize() {
	ebiten.MinimizeWindow()
}

func (r *Renderer) WindowSetDecorated(decorated bool) {
	ebiten.SetWindowDecorated(decorated)
}

func (r *Renderer) WindowRestore() {
	if ebiten.IsWindowMaximized() || ebiten.IsWindowMinimized() {
		ebiten.RestoreWindow()
	}
}

func (r *Renderer) SetWindowIcon(icons []image.Image) {
	ebiten.SetWindowIcon(icons)
}

func (r *Renderer) WindowRequestAttention() {
	ebiten.RequestAttention()
}

func SetWindowResizingMode(mode ebiten.WindowResizingModeType) {
	ebiten.SetWindowResizingMode(mode)
}
