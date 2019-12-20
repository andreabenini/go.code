package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path"
	"strings"
)

// Message types
const clientCLOSE = 0
const clientOPEN = 1
const clientMESSAGE = 2
const serverSHUTDOWN = 100

// Chat message struct
type clientConnection struct {
	socket  net.Conn
	event   uint8
	message string
}

// Function used on client connection
func handleClientConnection(client net.Conn, channel chan clientConnection) {
	fmt.Printf("client: %s\n", client.RemoteAddr().String())
	var message string
	Loop := true
	// client init message to router
	channel <- clientConnection{event: clientOPEN, socket: client}
	for Loop {
		// Get some input from user
		netData, err := bufio.NewReader(client).ReadString('\n')
		if err != nil {
			channel <- clientConnection{event: clientCLOSE, socket: client}
			fmt.Println(err)
			return
		}
		message = strings.ToLower(strings.TrimSpace(string(netData))) // TODO: sanitize input when needed
		// now decide what to do with it
		switch message {
		// Close client connection
		case "exit":
			Loop = false
		// Request server shutdown
		case "shutdown":
			channel <- clientConnection{event: serverSHUTDOWN}
		// Generic message to other clients
		default:
			channel <- clientConnection{event: clientMESSAGE, message: message}
		}
	}
	// Inform router about client connection closing
	channel <- clientConnection{event: clientCLOSE, socket: client}
	client.Close()
} /**/

// Server message router
func router(channel chan clientConnection, serverConnection net.Listener) {
	serverLoop := true
	clients := make([]net.Conn, 0) // slice with connected clients
	for serverLoop {
		operation := <-channel
		switch operation.event {

		// Append new client to connection list
		case clientOPEN:
			clients = append(clients, operation.socket)

		// Swap last element with closing socket
		case clientCLOSE:
			for i := 0; i < len(clients); i++ {
				if clients[i] == operation.socket {
					fmt.Printf("Closing client %d/%d\n", i+1, len(clients))
					// swap last element with this one (the deleted one)
					if len(clients) > 1 {
						clients[i] = clients[len(clients)-1]
					}
					clients = clients[:len(clients)-1] // remove last element (now the deleted one)
					i = len(clients) + 1
				}
			}

		// Shutting down entire tcp server
		case serverSHUTDOWN:
			serverConnection.Close()            // Close server Accept()
			for i := 0; i < len(clients); i++ { // Close client connections
				clients[i].Close()
			}
			serverLoop = false // and close the message router as well

		// forwarding message to all clients
		case clientMESSAGE:
			for i := 0; i < len(clients); i++ {
				clients[i].Write([]byte(operation.message + "\n"))
			}

		// TODO: Deal with it carefully or add new commands here
		default:
			fmt.Printf("Unknown operation: %d", operation.event)
		}
	}
	close(channel) // closing communication channel
} /**/

/**
 * main()
 * @param Args[1] (int) TCP port number
 */
func main() {
	// Get TCP Port from Args[1], if any
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		fmt.Printf("   %s [interface:]port\n", path.Base(arguments[0]))
		return
	}
	PORT := ":" + arguments[1]
	// Open TCP (ipv4) socket on specific port
	connection, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()
	fmt.Println("'exit'      - close client")     // some fancy output here
	fmt.Println("'shutdown'  - server shutdown")  //
	controlChannel := make(chan clientConnection) // Create control channel between go threads
	go router(controlChannel, connection)         // Run chat control router (with [controlChannel])
	for {
		client, err := connection.Accept() // Accept() an incoming client connection
		if err != nil {                    // Server shutdown due to Accept() failure
			fmt.Println("Closing server")
			return
		}
		go handleClientConnection(client, controlChannel) // fork it to a different process
	}
} /**/
