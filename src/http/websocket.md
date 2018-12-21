# Go WebSocket test
Server side
```go
func WSServer(ws *websocket.Conn) {
	var err error
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(WSServer))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServeTLS(":8000", "srv.cert", "srv.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

---
client side
```javascript
<script type="text/javascript">
    function WebSocketTest() {
        if ("WebSocket" in window) {
            console.log("WebSocket is supported by your Browser!");

            var ws = new WebSocket("wss://localhost:8000/ws");

            ws.onopen = function() {
                ws.send("Message from client");
                console.log("Message is sent...");
            };

            ws.onmessage = function(evt) {
                console.log("Message is received...", evt.data);
            };

            ws.onclose = function() {
                console.log("Connection is closed...");
            };
        } else {
            console.log("WebSocket NOT supported by your Browser!");
        }
    }
</script>
```
