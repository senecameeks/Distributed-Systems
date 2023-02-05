package puddlestore

import (
	"errors"
	"fmt"
	raftCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/client"
	//kv "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/kvstatemachine"
	tapCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/client"
	//	"strconv"
	//	"strings"
)

// contains functionality for 6 main functions

// returns reference to file with filename fn, in dir with path path
func Open(parentPath string, fn string, tapClient *tapCli.Client, raftClient *raftCli.Client) (*Inode, error) {
	//fmt.Printf("mf path: %v\n", parentPath+fn)
	inode, found := traversePath(parentPath, tapClient, raftClient)
	//fmt.Printf("am i found: %v\n", found)
	if found {
		inode = getInodeFromPath(parentPath+fn, raftClient, tapClient)
		return inode, nil
	} else {
		return nil, fmt.Errorf("invalid path")
	}
}

// returns contents of a file with filename fn, in dir with path path, starting at index start
func Read(path string, fn string, start uint32, buffer []byte, raftClient *raftCli.Client, tapClient *tapCli.Client) []byte {
	destInode, err := Open(path, fn, tapClient, raftClient)
	//fmt.Printf("d1: %v\n", destInode)
	handleError(err)

	/*if destInode.Type != FILE {
		fmt.Println("cannot read from directory")
	}*/
	//fmt.Println("im here")
	destIndirectBlock := getIndirectBlock(destInode.Indirect, raftClient, tapClient)

	if uint32(destInode.Size) < start {
		handleError(errors.New("file smaller than start value"))
	}
	//fmt.Println("im there")
	blockIndex := start / BLOCK_SIZE
	startIndex := start % BLOCK_SIZE
	//fmt.Printf("im gonna access this vguid: %v\n", getVguidFromAguid(destIndirectBlock.AGUID, raftClient))
	//fmt.Printf("len of db list: %v with vguid: %v\n", len(destIndirectBlock.DataBlockList), getVguidFromAguid(destIndirectBlock.AGUID, raftClient))
	//fmt.Printf("what should was stored: %v\n", destIndirectBlock)
	db := getDataBlock(destIndirectBlock.DataBlockList[blockIndex], raftClient, tapClient)
	buffer = db.Bytes[startIndex:]

	for i := blockIndex + 1; i < uint32(len(destIndirectBlock.DataBlockList)); i++ {
		db = getDataBlock(destIndirectBlock.DataBlockList[i], raftClient, tapClient)
		for b := 0; b < len(db.Bytes); b++ {
			buffer = append(buffer, db.Bytes[b])
		}
	}

	return buffer
}

func Write(parentPath string, fn string, start uint32, buffer []byte, tapClient *tapCli.Client, raftClient *raftCli.Client) (err error) {
	//fmt.Printf("\n\n\n INSIDE WRITE \n\n\n")
	//destInode, err := Open(absPath, tapClient, raftClient)
	destInode, err := Open(parentPath, fn, tapClient, raftClient)
	if err != nil {
		return fmt.Errorf("could not find file")
	}
	//fmt.Printf("d1: %v\n", destInode)
	newbie := new(Inode)
	newbie.AGUID = destInode.AGUID
	newbie.Indirect = destInode.Indirect
	newbie.Name = destInode.Name
	newbie.Size = destInode.Size
	newbie.Type = destInode.Type
	destIndirectBlock := getIndirectBlock(destInode.Indirect, raftClient, tapClient)
	//var bufferAguid string
	//fmt.Printf("\n\n buffer size: %d\n\n", len(buffer))
	startIndex := start / BLOCK_SIZE
	endIndex := (start + uint32(len(buffer))) / BLOCK_SIZE

	// create new datablocks if necessary
	if endIndex+1 > uint32(len(destIndirectBlock.DataBlockList)) || len(destIndirectBlock.DataBlockList) == 0 {
		// create new data blocks
		//fmt.Printf("\n\n\n Creating new Data blocks \n\n\n")
		for i := len(destIndirectBlock.DataBlockList); i <= int(endIndex); i++ {
			newDataBlock := CreateDataBlock()
			destIndirectBlock.DataBlockList = append(destIndirectBlock.DataBlockList, newDataBlock.AGUID)
			//Tell Raft and Tapestry about the new datablock
			putDataIntoTap(newDataBlock, tapClient, raftClient)

		}
		//Tell Raft nd Tapestry about the updated Indirect Block
		//fmt.Printf("im gonna insert this vguid: %v\n", getVguidFromAguid(destIndirectBlock.AGUID, raftClient))
		//fmt.Printf("len of db list: %v with vguid: %v\n", len(destIndirectBlock.DataBlockList), getVguidFromAguid(destIndirectBlock.AGUID, raftClient))
		putIndirectIntoTap(destIndirectBlock, tapClient, raftClient)
		//fmt.Printf("what should be stored: %v\n", destIndirectBlock)
		//fmt.Printf("this is my changed vguid: %v\n", getVguidFromAguid(destIndirectBlock.AGUID, raftClient))
	}
	//fmt.Println("i am here to tell the tale")
	var db *DataBlock
	var almostDone uint32
	almostDone = uint32(len(buffer))
	sliceUntil := BLOCK_SIZE - start
	startFromHere := sliceUntil

	for j := startIndex; j <= endIndex; j++ {
		dbVguid := getVguidFromAguid(destIndirectBlock.DataBlockList[j], raftClient)
		db = getDataBlockFromVguid(dbVguid, tapClient)
		//fmt.Printf("\n\n\n Indexing into first data block startIndex: %d\n\n\n", startIndex)
		// insert bytes
		if j == startIndex {
			endIndex := sliceUntil
			if uint32(len(buffer)) < sliceUntil {
				endIndex = uint32(len(buffer))
			}
			for byteIndex := 0; uint32(byteIndex) < endIndex; byteIndex++ {
				db.Bytes[start] = buffer[byteIndex]
				start++
			}
			//db.Bytes[start:] = buffer[:sliceUntil]
			almostDone -= sliceUntil

		} else {
			if almostDone > BLOCK_SIZE {
				db.Bytes = buffer[startFromHere : startFromHere+BLOCK_SIZE]
				startFromHere = startFromHere + BLOCK_SIZE
				almostDone -= BLOCK_SIZE
			} else {
				//fmt.Printf("start: %v stop: %v\n", startFromHere%BLOCK_SIZE, almostDone)
				db.Bytes = buffer[(startFromHere % BLOCK_SIZE):almostDone]
			}
		}
		putDataIntoTap(db, tapClient, raftClient)
	}
	//fmt.Printf("d2: %v\n", destInode)
	//fmt.Printf("n: %v\n", newbie)
	putInodeIntoTap(newbie, tapClient, raftClient)
	return nil
}

func Mkfileordir(parentPath string, fnDir string, tapClient *tapCli.Client, raftClient *raftCli.Client, isFile bool) (*Inode, error) {
	//parentPath := getParentPath(absPath)
	//fmt.Printf("parent path %v and fn %v\n", parentPath, fnDir)
	parentInode, found := traversePath(parentPath, tapClient, raftClient)
	//fmt.Printf(parentInode.Indirect)
	//fmt.Printf("\n\n\n\n\n")
	if found {
		//fmt.Printf("root found")
		//fmt.Printf("in mkfile absPath is " + parentPath)
		//fmt.Printf("\n\n\n\n")
		var fileDirInode *Inode
		if isFile {
			//fmt.Printf("mf fn %v\n", parentPath+fnDir)
			fileDirInode = CreateFileInode(fnDir, parentPath+fnDir)
		} else {
			fileDirInode = CreateDirectoryInode(fnDir, parentPath+fnDir+"/")
		}
		fileDirIndirect := CreateIndirectBlock()
		fileDirInode.Indirect = fileDirIndirect.AGUID
		// if isFile {
		// 	newDataBlock := CreateDataBlock()
		// 	fileDirIndirect.DataBlockList = append(fileDirIndirect.DataBlockList, newDataBlock.AGUID)
		// 	putDataIntoTap(newDataBlock, tapClient, raftClient)
		// 	//fmt.Printf("\n\n MKFILE len of data block is %d for inode %v\n", len(fileDirIndirect.DataBlockList), fileDirInode)
		// 	//fmt.Printf("\nunderneath\n")
		// }
		//fmt.Printf("\nabove\n")
		//mt.Printf("\n\n not MKFILE but is len of data block is %d for inode %v", len(fileDirIndirect.DataBlockList), fileDirInode)
		putIndirectIntoTap(fileDirIndirect, tapClient, raftClient)
		//fmt.Printf("\n\n not MKFILE but is len of data block is %d for inode %v", len(fileDirIndirect.DataBlockList), fileDirInode)
		//fmt.Printf("after putting into tap")
		//fmt.Printf("haha ok dad")
		//parentInode := getInodeFromPath(parentPath, raftClient, tapClient)
		parentIndirectBlock := getIndirectBlock(parentInode.Indirect, raftClient, tapClient)
		parentIndirectBlock.DataBlockList = append(parentIndirectBlock.DataBlockList, fileDirInode.AGUID)
		// put updated parent indirect block in tap
		putIndirectIntoTap(parentIndirectBlock, tapClient, raftClient)
		putInodeIntoTap(parentInode, tapClient, raftClient)
		//fmt.Printf("\n\nlen of indirect inodes %d \n\n", len(parentIndirectBlock.DataBlockList))
		putInodeIntoTap(fileDirInode, tapClient, raftClient)
		return fileDirInode, nil
	} else {
		//fmt.Printf("root not found")
		return nil, fmt.Errorf("invalid path")
	}
}

func Ls(absPath string, tapClient *tapCli.Client, raftClient *raftCli.Client) ([]string, error) {
	//fmt.Printf("\n ls absolute path" + absPath)
	//fmt.Printf("\n\n\n\n")
	_, found := traversePath(absPath, tapClient, raftClient)
	var items []string
	if found {
		destInode := getInodeFromPath(absPath, raftClient, tapClient)
		destIndirectBlock := getIndirectBlock(destInode.Indirect, raftClient, tapClient)
		//fmt.Printf("\n\n in ls %v and length of indirblock inodes is %d \n\n", absPath, len(destIndirectBlock.DataBlockList))
		for _, aguid := range destIndirectBlock.DataBlockList {
			curr_vguid := getVguidFromAguid(aguid, raftClient)
			item := getInodeFromVguid(curr_vguid, tapClient)
			items = append(items, item.Name)

			fmt.Println(item.Name)
		}
		return items, nil
	} else {
		return nil, fmt.Errorf("invalid path")
	}
}

// determines whether fn is a file or directory and removes appropriately
func Rm(parentPath string, fn string, tapClient *tapCli.Client, raftClient *raftCli.Client) error {
	inode, found := traversePath(parentPath, tapClient, raftClient)
	if found {
		// tapestry does not handle deletion so no need to worry about this
		fnInode := getInodeFromPath(parentPath+fn, raftClient, tapClient)
		dirInode := getInodeFromPath(parentPath+fn+"/", raftClient, tapClient)
		if fnInode != nil {
			// issa file
			parIndirect := getIndirectBlock(inode.Indirect, raftClient, tapClient)
			for i, aguid := range parIndirect.DataBlockList {
				if aguid == fnInode.AGUID {
					if i == len(parIndirect.DataBlockList)-1 {
						parIndirect.DataBlockList = parIndirect.DataBlockList[:i]
					} else {
						parIndirect.DataBlockList = append(parIndirect.DataBlockList[:i], parIndirect.DataBlockList[i+1:]...)
					}
					putIndirectIntoTap(parIndirect, tapClient, raftClient)
					putInodeIntoTap(inode, tapClient, raftClient)
				}
			}
		} else if dirInode != nil {
			// issa directory
			// check if directory empty (can be deleted)
			dirIndirect := getIndirectBlock(dirInode.Indirect, raftClient, tapClient)
			if len(dirIndirect.DataBlockList) > 0 {
				return fmt.Errorf("cannot delete non-empty directory")
			}

			// now delete empty directory
			parIndirect := getIndirectBlock(inode.Indirect, raftClient, tapClient)
			for i, aguid := range parIndirect.DataBlockList {
				if aguid == dirInode.AGUID {
					if i == len(parIndirect.DataBlockList)-1 {
						parIndirect.DataBlockList = parIndirect.DataBlockList[:i]
					} else {
						parIndirect.DataBlockList = append(parIndirect.DataBlockList[:i], parIndirect.DataBlockList[i+1:]...)
					}
					putIndirectIntoTap(parIndirect, tapClient, raftClient)
					putInodeIntoTap(inode, tapClient, raftClient)
				}
			}
		} else {
			//fmt.Printf("cannot find %v%v\n", parentPath, fn)
			return fmt.Errorf("file does not exist")
		}
	}
	return nil
}

/*func Rmfile(absPath string, fn string, tapClient *tapCli.Client, raftClient *raftCli.Client) error {
	inode, found := traversePath(absPath, tapClient, raftClient)
	if found {
		//parentPath := getParentPath(absPath)
		destInode := getInodeFromPath(absPath, raftClient, tapClient)
		destIndirectBlock := getIndirectBlock(destInode.Indirect, raftClient, tapClient)
		for i, aguid := range destIndirectBlock.DataBlockList {
			if aguid == inode.AGUID {
				//remove file
				destIndirectBlock.DataBlockList[i] = ""
				putIndirectIntoTap(destIndirectBlock, tapClient, raftClient)
				putInodeIntoTap(inode, tapClient, raftClient)
				break
			}
		}
		return nil
	} else {
		return fmt.Errorf("cannot remove file")
	}
}

func Rmdir(absPath string, tapClient *tapCli.Client, raftClient *raftCli.Client) error {
	inode, found := traversePath(absPath, tapClient, raftClient)
	if found {
		parentPath := getParentPath(absPath)
		destInode := getInodeFromPath(parentPath, raftClient, tapClient)
		destIndirectBlock := getIndirectBlock(destInode.Indirect, raftClient, tapClient)
		for i, aguid := range destIndirectBlock.DataBlockList {
			if aguid == inode.AGUID {
				//check if dir is empty
				vguid := getVguidFromAguid(aguid, raftClient)
				dirInode := getInodeFromVguid(vguid, tapClient)
				destIndirectBlock = getIndirectBlock(dirInode.AGUID, raftClient, tapClient)
				canDeleteDir := true
				for i := 0; i < len(destIndirectBlock.DataBlockList); i++ {
					if destIndirectBlock.DataBlockList[i] != "" {
						canDeleteDir = false
						break
					}
				}
				// remove if empty
				if canDeleteDir {
					destIndirectBlock.DataBlockList[i] = ""
					putInodeIntoTap(inode, tapClient, raftClient)
					break
				} else {
					return fmt.Errorf("cannot remove non empty dir")
				}
			}
		}
		return nil
	} else {
		return fmt.Errorf("cannot remove dir")
	}
}*/
