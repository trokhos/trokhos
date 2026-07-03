package trokhos

import "embed"

//go:embed all:frontend/.output/public migrations/*.go
var FS embed.FS
