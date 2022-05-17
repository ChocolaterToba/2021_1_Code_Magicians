package domain

import "io"

type FileWithName struct {
	File     io.Reader
	Filename string
}
