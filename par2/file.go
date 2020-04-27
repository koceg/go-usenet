package par2

import "fmt"

type File struct {
	*FileDescPacket
	*IFSCPacket
}

func (f *File) ID() string {
	return fmt.Sprintf("%x", f.FileDescPacket.FileID)
}

func (f *File) Filename() string {
	return f.FileDescPacket.Filename
}

func (f *File) Valid() bool {
	return f.FileDescPacket != nil && f.IFSCPacket != nil
}
