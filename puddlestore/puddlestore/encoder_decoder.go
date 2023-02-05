package puddlestore

import (
	"bytes"
	"encoding/gob"
)

// file for encoding and decoding inodes, indirect blocks and data blocks to and from bytes

func GetInodeBytes(inode *Inode) ([]byte, error) {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(*inode)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func GetInodeFromBytes(b []byte) (*Inode, error) {
	buff := bytes.NewBuffer(b)
	var inode Inode
	dataDecoder := gob.NewDecoder(buff)
	err := dataDecoder.Decode(&inode)
	if err != nil {
		return nil, err
	}

	return &inode, nil
}

func GetIndirectBytes(ib *IndirectBlock) ([]byte, error) {
	ib.LenDB = len(ib.DataBlockList)
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(ib.AGUID)
	if err != nil {
		return nil, err
	}
	err = e.Encode(ib.LenDB)
	if err != nil {
		return nil, err
	}
	for _, i := range ib.DataBlockList {
		err = e.Encode(i)
		if err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil
}

func GetIndirectFromBytes(b []byte) (*IndirectBlock, error) {
	buff := bytes.NewBuffer(b)
	var ib IndirectBlock
	dataDecoder := gob.NewDecoder(buff)
	err := dataDecoder.Decode(&ib.AGUID)
	if err != nil {
		return nil, err
	}
	err = dataDecoder.Decode(&ib.LenDB)
	if err != nil {
		return nil, err
	}
	ib.DataBlockList = make([]string, ib.LenDB)
	var s string
	for i := 0; i < ib.LenDB; i++ {
		err = dataDecoder.Decode(&s)
		ib.DataBlockList[i] = s
		if err != nil {
			return nil, err
		}
	}
	return &ib, nil
}

func GetDataBytes(db *DataBlock) ([]byte, error) {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(*db)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func GetDataFromBytes(b []byte) (*DataBlock, error) {
	buff := bytes.NewBuffer(b)
	var db DataBlock
	dataDecoder := gob.NewDecoder(buff)
	err := dataDecoder.Decode(&db)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
