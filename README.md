
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
- Server notification event listening


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

## Usage with notification events

```go
package main

import (
	"context"
	"fmt"

	gomcsmp "github.com/eterline/go-mc-smp"
)

func main() {
	smp, err := gomcsmp.NewClient("127.0.0.1", 9100, "YOUR_RPC_TOKEN")
	if err != nil {
		panic(err)
	}
	defer smp.Close()

	ctx := context.TODO()

	// USE WITH NEW CHANNEL CREATE!
	// IF YOU WILL USE THAT IN for{}!
	// THAT CAN CAUSE GORUTINES LEAK!
	gamerulesCh := smp.NotifyGamerulesUpdates(ctx)
	playersJoinedCh := smp.NotifyPlayersJoined(ctx)
	playersLeftCh := smp.NotifyPlayersLeft(ctx)
	serverSavedCh := smp.NotifyServerSaved(ctx)

	for {
		select {
		case u, ok := <-gamerulesCh:
			if !ok {
				return
			}
			fmt.Println("gamerule update", u.Key, u.Value)

		case p, ok := <-playersJoinedCh:
			if !ok {
				return
			}
			fmt.Println("player joined", p.Name, p.ID)

		case p, ok := <-playersLeftCh:
			if !ok {
				return
			}
			fmt.Println("player left", p.Name, p.ID)

		case <-serverSavedCh:
			fmt.Println("server saved world")
		}
	}
}
```

![Use Screen](./screen/server_events.png)
![Use Screen](./screen/image.png)

## License

[MIT](https://choosealicense.com/licenses/mit/)
