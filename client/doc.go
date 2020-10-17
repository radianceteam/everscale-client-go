// Package provides wrapper to TON-SDK dll that must be compiled separately.
//
// For methods available see Client interface.
//
// Example usage
//  import "github.com/radianceteam/ton-client-go/client"
//      ...
//		c, err := client.NewClient(client.Config{
//			Network: client.ConfigNetwork{ServerAddress: "net.ton.dev"},
//		})
//		if err != nil {
//			return err
//		}
//		defer c.Close()
//
//		version, err := c.ClientVersion()
package client
