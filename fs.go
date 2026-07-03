package trokhos

import "embed"

//go:embed all:frontend/dist migrations/*.go
var FS embed.FS
