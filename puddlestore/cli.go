package main

import (
	//"flag"
	"errors"
	"fmt"
	"github.com/abiosoft/ishell"
	core "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore"
	raftCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/client"
	kv "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/raft/kvstatemachine"
	tapCli "github.com/brown-csci1380/mkohn-smeeks-s19/puddlestore/puddlestore/tapestry/client"
	"github.com/samuel/go-zookeeper/zk"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// check server state w ruok
	ok := zk.FLWRuok([]string{"127.0.0.1:2181"}, time.Second*2)

	if !ok[0] {
		panic("zk server is not okay")
	}

	// establish a connection to zookeeper server w address 127.0.0.1 at port 2181
	//  that times out after 3s
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*3)
	if err != nil {
		handleError(err)
	}

	// remove any existing raftlogs
	//removeRaftLogs()

	// start raft cluster and tapestry network of given sizes
	rchildren, _, _ := conn.Children("/Raft")
	tchildren, _, _ := conn.Children("/Tapestry")
	if len(rchildren) == 0 || len(tchildren) == 0 {
		removeRaftLogs()
		core.StartRaftAndTapestry(3, 5)
	}

	tapestryClient := tapCli.Connect(GetNodeAddress(conn, "/Tapestry"))

	raftClient, err := raftCli.Connect(GetNodeAddress(conn, "/Raft"))
	if err != nil {
		handleError(err)
	}

	// initialize kv state machine
	err = raftClient.SendRequest(kv.KV_STATE_INIT, []byte{})
	handleError(err)

	// create root directory
	core.CreateRootDir(tapestryClient, raftClient)

	// Kick off shell
	shell := ishell.New()

	// add 'open' command

	shell.AddCmd(&ishell.Cmd{
		Name: "open",
		Help: "obtains a reference to a particular file",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Usage: open <absPath>")
				return
			}
			path := c.Args[0]
			var fn string
			var parentPath string
			// change to add / at beginning if not there
			if string(path[0]) != "/" {
				path = "/" + path
			}

			segmentedPath := strings.Split(path, "/")
			if len(segmentedPath) == 2 {
				parentPath = "/"
				fn = segmentedPath[len(segmentedPath)-1]
			} else {
				fn = segmentedPath[len(segmentedPath)-1]
				parentPath = path[:len(path)-len(fn)]
			}

			//fmt.Printf("\n\n parent path: %s, fn: %s\n\n", parentPath, fn)
			inode, err := core.Open(parentPath, fn, tapestryClient, raftClient)
			if err != nil {
				fmt.Printf("The file %v does not exist\n", c.Args[0])
			} else {
				fmt.Printf("The ID for %v is %v\n", c.Args[0], inode.AGUID)
			}
			return
		},
	})

	// add 'read' command
	shell.AddCmd(&ishell.Cmd{
		Name: "read",
		Help: "examine the data contained in a file",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 2 {
				shell.Println("Usage: read <absPath> <startIndex>")
				return
			}
			var buffer []byte
			startIndex, _ := strconv.Atoi(c.Args[1])
			argPath := c.Args[0]
			var parentPath string
			var fn string

			// change to add / at beginning if not there
			if string(argPath[0]) != "/" {
				argPath = "/" + argPath
			}

			segmentedPath := strings.Split(argPath, "/")
			if len(segmentedPath) == 1 || len(segmentedPath) == 2 {
				parentPath = "/"
				fn = segmentedPath[len(segmentedPath)-1]
			} else {
				fn = segmentedPath[len(segmentedPath)-1]
				parentPath = argPath[:len(argPath)-len(fn)]
			}
			//fmt.Printf("\n\n parent path: %s, fn: %s\n\n", parentPath, fn)
			data := core.Read(parentPath, fn, uint32(startIndex), buffer, raftClient, tapestryClient)
			fmt.Println(string(data[:]))

		},
	})

	// add 'write' command
	shell.AddCmd(&ishell.Cmd{
		Name: "write",
		Help: "change the data contained in a file",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 3 {
				shell.Println("Usage: write <absPath> <startIndex> <data>")
				return
			}
			path := c.Args[0]
			var parentPath string
			var fn string
			// change to add / at beginning if not there
			if string(path[0]) != "/" {
				path = "/" + path
			}

			//fmt.Printf("argPath: %v\n", path)

			segmentedPath := strings.Split(path, "/")
			if len(segmentedPath) == 2 {
				parentPath = "/"
				fn = segmentedPath[len(segmentedPath)-1]
			} else {
				fn = segmentedPath[len(segmentedPath)-1]
				parentPath = path[:len(path)-len(fn)]
			}

			buffer := []byte(c.Args[2])
			//fmt.Printf("buf daddy: %v\n", buffer)
			startIndex, _ := strconv.Atoi(c.Args[1])
			err := core.Write(parentPath, fn, uint32(startIndex), buffer, tapestryClient, raftClient)
			handleError(err)

		},
	})

	// add 'create' command
	shell.AddCmd(&ishell.Cmd{
		Name: "mkfile",
		Help: "creates a new file",
		Func: func(c *ishell.Context) {
			//shell.Println(len(c.Args))
			if len(c.Args) != 1 {
				shell.Println("Usage: mkfile <absPath>")
				return
			}
			path := c.Args[0]
			var parentPath string
			var fn string
			// change to add / at beginning if not there
			if string(path[0]) != "/" {
				path = "/" + path
			}

			//fmt.Printf("argPath: %v\n", path)

			segmentedPath := strings.Split(path, "/")
			if len(segmentedPath) == 2 {
				parentPath = "/"
				fn = segmentedPath[len(segmentedPath)-1]
			} else {
				fn = segmentedPath[len(segmentedPath)-1]
				parentPath = path[:len(path)-len(fn)]
			}
			//fmt.Printf("\n\n parent path: %s, fn: %s\n\n", parentPath, fn)
			_, err := core.Mkfileordir(parentPath, fn, tapestryClient, raftClient, true)
			handleError(err)
		},
	})

	// add 'mkdir' command
	shell.AddCmd(&ishell.Cmd{
		Name: "mkdir",
		Help: "creates a new directory within an existing directory",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Usage: mkdir <absPath>")
				return
			}
			var argPath string
			path := c.Args[0]
			var parentPath string
			var fn string
			// change to add / at beginning if not there
			if string(path[0]) != "/" {
				path = "/" + path
			}
			// check if ends in / or not
			if string(path[len(path)-1]) == "/" {
				argPath = path[:len(path)-1]
			} else {
				argPath = path
			}

			//fmt.Printf("argPath: %v\n", argPath)

			segmentedPath := strings.Split(argPath, "/")
			if len(segmentedPath) == 2 {
				parentPath = "/"
				fn = segmentedPath[len(segmentedPath)-1]
			} else {
				fn = segmentedPath[len(segmentedPath)-1]
				parentPath = argPath[:len(argPath)-len(fn)]
			}
			//fmt.Printf("\n\n parent path: %s, fn: %s\n\n", parentPath, fn)
			_, err := core.Mkfileordir(parentPath, fn, tapestryClient, raftClient, false)
			handleError(err)

		},
	})

	// add 'list' command
	shell.AddCmd(&ishell.Cmd{
		Name: "ls",
		Help: "obtains a list of the contents of a particular directory",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Usage: ls <absPath>")
				return
			}
			var argPath string
			path := c.Args[0]
			// change to add / at beginning if not there
			if string(path[0]) != "/" {
				path = "/" + path
			}
			// check if ends in / or not
			if string(path[len(path)-1]) == "/" {
				argPath = path
			} else {
				argPath = path + "/"
			}

			_, err := core.Ls(argPath, tapestryClient, raftClient)
			handleError(err)
		},
	})

	// add 'remove' command
	shell.AddCmd(&ishell.Cmd{
		Name: "rm",
		Help: "deletes an entry from a directory",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Println("Usage: remove <absPath>")
				return
			}
			path := c.Args[0]
			var argPath string
			var parentPath string
			var fn string

			// change to add / at beginning if not there
			if string(path[0]) != "/" {
				path = "/" + path
			}

			if string(path[len(path)-1]) == "/" {
				argPath = path[:len(path)-1]
			} else {
				argPath = path
			}

			segmentedPath := strings.Split(argPath, "/")
			if len(segmentedPath) == 2 {
				parentPath = "/"
				fn = segmentedPath[1]
			} else {
				fn = segmentedPath[len(segmentedPath)-1]
				parentPath = argPath[:len(argPath)-len(fn)]
			}
			//fmt.Printf("\n\n parent path: %s, fn: %s\n\n", parentPath, fn)
			err := core.Rm(parentPath, fn, tapestryClient, raftClient)
			//err := core.Rmfile(parentPath, fn, tapestryClient, raftClient)
			handleError(err)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "quit",
		Help: "quit puddlestore client",
		Func: func(c *ishell.Context) {
			core.Cleanup()
			os.Exit(0)
		},
	})

	shell.Println(shell.HelpText())
	shell.Run()
}

// acquires address of random node in children at path
func GetNodeAddress(conn *zk.Conn, path string) (addr string) {
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)

	children, _, err := conn.Children(path)
	handleError(err)

	if len(children) == 0 {
		err = errors.New("No children in path" + path)
		handleError(err)
	}

	index := rand.Intn(len(children))
	startOfAddr := strings.Index(children[index], "/")

	return children[index][startOfAddr+1:]
}

// handles errors, panics if found
func handleError(err error) {
	if err != nil {
		//core.Cleanup()
		fmt.Println(err)
	}
}

// rmeove raft logs
func removeRaftLogs() {
	os.RemoveAll("raftlogs/")
}
