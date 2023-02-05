package puddlestore

import (
	"github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/tapestry"
)

type FileType int

const (
	DIRECTORY FileType = iota
	FILE
)

const BLOCK_SIZE = uint32(4096)

type Inode struct {
	Name     string
	Type     FileType
	Size     int
	Indirect string
	AGUID    string
}

type IndirectBlock struct {
	DataBlockList []string
	LenDB         int
	AGUID         string
}

type DataBlock struct {
	AGUID string
	Bytes []byte
}

func CreateDirectoryInode(name string, absPath string) *Inode {
	inode := new(Inode)
	inode.Name = name
	inode.Type = DIRECTORY
	inode.Size = 0
	//inode.Indirect = "" to be set later
	inode.AGUID = tapestry.Hash(absPath).String()
	return inode
}

func CreateFileInode(name string, absPath string) *Inode {
	inode := new(Inode)
	inode.Name = name
	inode.Type = FILE
	inode.Size = 0
	//inode.Indirect = "" to be set later
	inode.AGUID = tapestry.Hash(absPath).String()
	return inode
}

func CreateIndirectBlock() *IndirectBlock {
	indirectBlock := new(IndirectBlock)
	indirectBlock.DataBlockList = make([]string, 0)
	indirectBlock.LenDB = 0
	indirectBlock.AGUID = makeVguidForAguid("") // generates a random 32 length AGUID
	return indirectBlock
}

func CreateDataBlock() *DataBlock {
	dataBlock := new(DataBlock)
	dataBlock.Bytes = make([]byte, BLOCK_SIZE)
	dataBlock.AGUID = makeVguidForAguid("") // generates a random 32 length AGUID
	return dataBlock
}
