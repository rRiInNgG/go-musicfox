//go:build windows

package player

import (
	"io"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
)

// decodeMinimp3 在 Windows 下回退到标准 mp3 解码
func decodeMinimp3(r io.ReadSeekCloser) (beep.StreamSeekCloser, beep.Format, error) {
	return mp3.Decode(r)
}
