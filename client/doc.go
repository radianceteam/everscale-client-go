// Package provides wrapper to EVER-SDK dll that must be compiled separately.
//
// For methods available see Client interface.
//
// # Example usage
//
//		import "github.com/radianceteam/everscale-client-go/Client"
//
//		func main() {
//	   client, err := client.NewClient(client.Config{
//		    Network: &client.NetworkConfig{ServerAddress: null.NewString("net.ton.dev", true)},
//		  }, client.WrapperConfig{MaxCGOConcurrentThreads: 10})
//
//		  if err != nil {
//		    return err
//		  }
//		  defer c.Close()
//
//		  version, err := c.ClientVersion()
//		}
package client
