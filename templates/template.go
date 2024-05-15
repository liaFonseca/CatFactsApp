package templates

import (
	"embed"
)

//go:embed all:*
var TempDir embed.FS
