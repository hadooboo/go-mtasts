package mtasts_test

import (
	"context"
	"fmt"

	"github.com/hadooboo/go-mtasts/v2"
)

func ExampleCache_Get() {
	c := mtasts.NewRAMCache()
	policy, err := c.Get(context.Background(), "gmail.com")
	if err != nil {
		fmt.Println("Oh noes!", err)
		return
	}

	fmt.Println("Allowed MXs:", policy.MX)
}
