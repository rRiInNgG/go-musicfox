//go:build linux || darwin

package player

import (
	"io"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/minimp3"
	minimp3pkg "github.com/tosone/minimp3"
)

// decodeMinimp3 使用 minimp3 进行解码 (仅 Unix 系统)
func decodeMinimp3(r io.ReadSeekCloser) (beep.StreamSeekCloser, beep.Format, error) {
	minimp3pkg.BufferSize = 1024 * 50
	return minimp3.Decode(r)
}
