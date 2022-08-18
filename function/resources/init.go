package resources

import (
	"embed"

	"github.com/golang/freetype/truetype"
)

var (
	//go:embed PretendardVariable.ttf
	pretendardedVariable []byte
	Pretendard, _        = truetype.Parse(pretendardedVariable)

	//go:embed graphic
	GraphicImages embed.FS
)
