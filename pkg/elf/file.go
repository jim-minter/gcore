package elf

type File struct {
	Count    uint64
	PageSize uint64
}

type FileElement struct {
	Start   uint64
	End     uint64
	FileOfs uint64
}
