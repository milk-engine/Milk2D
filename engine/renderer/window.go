package renderer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct{}

func (w *Window) Maximize() {
	ebiten.MaximizeWindow()
}

func (w *Window) Minimize() {
	ebiten.MinimizeWindow()
}

func (w *Window) Restore() {
	if ebiten.IsWindowMaximized() || ebiten.IsWindowMinimized() {
		ebiten.RestoreWindow()
	}
}
func (w *Window) RequestAttention() {
	ebiten.RequestAttention()
}

func (w *Window) isDecorated() bool {
	return ebiten.IsWindowDecorated()
}

func (w *Window) IsResizable() bool {
	return ebiten.WindowResizingMode() == ebiten.WindowResizingModeEnabled
}

func (w *Window) isFloating() bool {
	return ebiten.IsWindowFloating()
}

func (w *Window) isMaximized() bool {
	return ebiten.IsWindowMaximized()
}

func (w *Window) isMinimized() bool {
	return ebiten.IsWindowMinimized()
}

func (w *Window) isBeingClosed() bool {
	return ebiten.IsWindowBeingClosed()
}

func (w *Window) IsClosingHandled() bool {
	return ebiten.IsWindowClosingHandled()
}

func (w *Window) IsMousePassthrough() bool {
	return ebiten.IsWindowMousePassthrough()
}

func (w *Window) GetSize() (int, int) {
	return ebiten.WindowSize()
}

func (w *Window) GetSizeLimits() (minw, minh, maxw, maxh int) {
	return ebiten.WindowSizeLimits()
}

func (w *Window) GetResizingMode() ebiten.WindowResizingModeType {
	return ebiten.WindowResizingMode()
}

func (w *Window) GetPosition() (x, y int) {
	return ebiten.WindowPosition()
}

func (w *Window) SetTitle(title string) {
	ebiten.SetWindowTitle(title)
}

func (w *Window) SetSize(width, height int) {
	ebiten.SetWindowSize(width, height)
}

func (w *Window) SetFullscreen(fullscreen bool) {
	ebiten.SetFullscreen(fullscreen)
}

func (w *Window) SetDecorated(decorated bool) {
	ebiten.SetWindowDecorated(decorated)
}

func (w *Window) SetIcon(icons []image.Image) {
	ebiten.SetWindowIcon(icons)
}

func (w *Window) SetResizingMode(mode ebiten.WindowResizingModeType) {
	ebiten.SetWindowResizingMode(mode)
}

func (w *Window) SetSizeLimits(minw, minh, maxw, maxh int) {
	ebiten.SetWindowSizeLimits(minw, minh, maxw, maxh)
}

func (w *Window) SetFloating(float bool) {
	ebiten.SetWindowFloating(float)
}

func (w *Window) SetClosingHandled(handled bool) {
	ebiten.SetWindowClosingHandled(handled)
}

func (w *Window) SetWindowMousePassthrough(enabled bool) {
	ebiten.SetWindowMousePassthrough(enabled)
}
