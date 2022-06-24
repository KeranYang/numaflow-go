# NumaFlow Golang SDK

This SDK provides the interfaces to implement [Numaflow](https://github.com/numaproj/numaflow) User Defined Functions or Sinks in Golang.

## Implement User Defined Functions

```golang
package main

import (
	"context"

	funcsdk "github.com/numaproj/numaflow-go/function"
)

func handle(ctx context.Context, key, msg []byte) (funcsdk.Messages, error) {
	return funcsdk.MessagesBuilder().Append(funcsdk.MessageToAll(msg)), nil
}

func main() {
	funcsdk.Start(context.Background(), handle)
}
```

## Implement User Defined Sinks

```golang
package main

import (
	"context"
	"fmt"

	sinksdk "github.com/numaproj/numaflow-go/sink"
)

func handle(ctx context.Context, msgs []sinksdk.Message) (sinksdk.Responses, error) {
	result := sinksdk.ResponsesBuilder()
	for _, m := range msgs {
		fmt.Println(string(m.Payload))
		result = result.Append(sinksdk.ResponseOK(m.ID))
	}
	return result, nil
}

func main() {
	sinksdk.Start(context.Background(), handle)
}

```
