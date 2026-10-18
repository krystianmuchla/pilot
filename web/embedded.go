package web

import "embed"

//go:embed resources/*
var Resources embed.FS

//go:embed index.html
var IndexHtml string
