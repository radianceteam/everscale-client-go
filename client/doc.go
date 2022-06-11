// Package provides wrapper to EVER-SDK dll that must be compiled separately.
//
// For methods available see Client interface.
//
// Example usage
//  import "github.com/radianceteam/everscale-client-go/Client"
//      ...
//		c, err := client.NewClient(client.Config{
//			Network: Client.ConfigNetwork{ServerAddress: "net.ton.dev"},
//		})
//		if err != nil {
//			return err
//		}
//		defer c.Close()
//
//		version, err := c.ClientVersion()
package client
