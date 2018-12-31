// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	tb "github.com/nsf/termbox-go"
)

// Init initializes termui library. This function should be called before any others.
// After initialization, the library must be finalized by 'Close' function.
func Init() error {
	if err := tb.Init(); err != nil {
		return err
	}
	tb.SetInputMode(tb.InputEsc | tb.InputMouse)
	tb.SetOutputMode(tb.Output256)

	return nil
}

// Close finalizes termui library.
// It should be called after successful initialization when termui's functionality isn't required anymore.
func Close() {
	tb.Close()
}

func TerminalSize() (int, int) {
	tb.Sync()
	width, height := tb.Size()
	return width, height
}
