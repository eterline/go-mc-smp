
# go-mc-smp

Go client for Minecraft JSON-RPC


## Installation

```bash
go get github.com/eterline/go-mc-smp
```

## Features

- Players management: list, kick
- Allowlist management: get, add, remove, clear
- Server control: status, save, stop, system messages
- Gamerules: list, update
- Fully asynchronous JSON-RPC over WebSocket
- Context support and configurable call timeout
- Another in process
## Usage

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/eterline/go-mc-smp"
)

func main() {
	smp, err := gomcsmp.NewClient("127.0.0.1", 9100, "YOUR_RPC_TOKEN", gomcsmp.WithCallTimeout(10*time.Second))
	if err != nil {
		panic(err)
	}
	defer smp.Close()

	msg := gomcsmp.SystemMessage{
		ReceivingPlayers: []gomcsmp.Player{gomcsmp.NewPlayer("EterLine")},
		Overlay:          false,
		Message:          gomcsmp.NewMessage("hello", "hello"),
	}

	sent, err := smp.ServerSystemMessage(context.Background(), msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent:", sent)
}
```
Sending a System Message
![Use Screen](./screen/system_message.png)

## License

[MIT](https://choosealicense.com/licenses/mit/)
