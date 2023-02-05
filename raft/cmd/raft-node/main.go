package main

import (
	"flag"
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/brown-csci1380/mkohn-smeeks-s19/raft/raft"
)

type RaftNode struct {
	*raft.RaftNode
}

func (node RaftNode) findNode(id string) *raft.RemoteNode {
	nodeList := node.GetNodeList()

	for _, remoteNode := range nodeList {
		if (remoteNode.Id == id) ||
			(remoteNode.Addr == id) {
			return &remoteNode
		}
	}

	return nil
}

func main() {
	var port int
	var addr string
	var debug bool

	portHelpString := "The server port to bind to. Defaults to a random port."
	flag.IntVar(&port, "port", 0, portHelpString)
	flag.IntVar(&port, "p", 0, portHelpString)

	connectHelpString := "An existing node to connect to. If left blank, does not attempt to connect to another node."
	flag.StringVar(&addr, "connect", "", connectHelpString)
	flag.StringVar(&addr, "c", "", connectHelpString)

	debugHelpString := "Turn on debug message printing."
	flag.BoolVar(&debug, "debug", false, debugHelpString)
	flag.BoolVar(&debug, "d", false, debugHelpString)

	flag.Parse()

	raft.SetDebug(debug)

	// Initialize Raft with default config
	config := raft.DefaultConfig()

	// Parse address of remote Raft node
	var remoteNode *raft.RemoteNode
	if addr != "" {
		remoteNode = &raft.RemoteNode{Id: raft.AddrToId(addr, config.NodeIdSize), Addr: addr}
	}

	// Create Raft node
	fmt.Println("Starting a Raft node...")
	raftNode, err := raft.CreateNode(port, remoteNode, config)
	node := RaftNode{raftNode}

	if err != nil {
		fmt.Printf("Error starting Raft node: %v\n", err)
		return
	}

	fmt.Printf("Successfully created Raft node: %v\n", node)

	// Kick off shell
	shell := ishell.New()

	debugCommand := ishell.Cmd{
		Name: "debug",
		Help: "turn debug messages on or off, on by default",
		Func: func(c *ishell.Context) {
			shell.Println("Usage: debug <on|off>")
		},
	}

	debugCommand.AddCmd(&ishell.Cmd{
		Name: "on",
		Help: "turn debug messages on",
		Func: func(c *ishell.Context) {
			raft.SetDebug(true)
		},
	})

	debugCommand.AddCmd(&ishell.Cmd{
		Name: "off",
		Help: "turn debug messages off",
		Func: func(c *ishell.Context) {
			raft.SetDebug(false)
		},
	})

	shell.AddCmd(&debugCommand)

	shell.AddCmd(&ishell.Cmd{
		Name: "state",
		Help: "print out the current local and cluster state",
		Func: func(c *ishell.Context) {
			shell.Println(node.FormatState())
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "log",
		Help: "print out the local log cache",
		Func: func(c *ishell.Context) {
			shell.Println(node.FormatLogCache())
		},
	})

	enableCommand := ishell.Cmd{
		Name: "enable",
		Help: "enable communications with one or all nodes in the cluster",
		Func: func(c *ishell.Context) {
			shell.Println("Usage: enable all | enable <send|recv> <addr>")
		},
	}

	enableCommand.AddCmd(&ishell.Cmd{
		Name: "all",
		Help: "enable all communications with the cluster",
		Func: func(c *ishell.Context) {
			node.NetworkPolicy.PauseWorld(false)
		},
	})

	enableCommand.AddCmd(&ishell.Cmd{
		Name: "send",
		Help: "enable sending requests to a specific remote node",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Printf("Usage: enable send <addr>")
				return
			}

			remoteNode := node.findNode(c.Args[0])
			if remoteNode == nil {
				shell.Printf("Error: could not find node with given ID or address: %v\n", c.Args[0])
			}

			node.NetworkPolicy.RegisterPolicy(*node.GetRemoteSelf(), *remoteNode, true)
		},
	})

	enableCommand.AddCmd(&ishell.Cmd{
		Name: "recv",
		Help: "enable receiving requests from a specific remote node",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Printf("Usage: enable recv <addr>")
				return
			}

			remoteNode := node.findNode(c.Args[0])
			if remoteNode == nil {
				shell.Printf("Error: could not find node with given ID or address: %v\n", c.Args[0])
			}

			node.NetworkPolicy.RegisterPolicy(*remoteNode, *node.GetRemoteSelf(), true)
		},
	})

	shell.AddCmd(&enableCommand)

	disableCommand := ishell.Cmd{
		Name: "disable",
		Help: "disable communications with one or all nodes in the cluster",
		Func: func(c *ishell.Context) {
			shell.Println("Usage: disable all | disable <send|recv> <addr>")
		},
	}

	disableCommand.AddCmd(&ishell.Cmd{
		Name: "all",
		Help: "disable all communications with the cluster",
		Func: func(c *ishell.Context) {
			node.NetworkPolicy.PauseWorld(true)
		},
	})

	disableCommand.AddCmd(&ishell.Cmd{
		Name: "send",
		Help: "disable sending requests to a specific remote node",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Printf("Usage: disable send <addr>")
				return
			}

			remoteNode := node.findNode(c.Args[0])
			if remoteNode == nil {
				shell.Printf("Error: could not find node with given ID or address: %v\n", c.Args[0])
			}

			node.NetworkPolicy.RegisterPolicy(*node.GetRemoteSelf(), *remoteNode, false)
		},
	})

	disableCommand.AddCmd(&ishell.Cmd{
		Name: "recv",
		Help: "disable receiving requests from a specific remote node",
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				shell.Printf("Usage: disable recv <addr>")
				return
			}

			remoteNode := node.findNode(c.Args[0])
			if remoteNode == nil {
				shell.Printf("Error: could not find node with given ID or address: %v\n", c.Args[0])
			}

			node.NetworkPolicy.RegisterPolicy(*remoteNode, *node.GetRemoteSelf(), false)
		},
	})

	shell.AddCmd(&disableCommand)

	shell.AddCmd(&ishell.Cmd{
		Name: "leave",
		Help: "gracefully leave the cluster",
		Func: func(c *ishell.Context) {
			raft.Out.Println("Gracefully exiting local raft node...")
			node.GracefulExit()
			raft.Out.Println("Bye!")
		},
	})

	shell.Println(shell.HelpText())
	shell.Run()
}
