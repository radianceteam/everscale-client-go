package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TylerBrock/colorjson"
	"github.com/spf13/cobra"
	"github.com/volatiletech/null"

	"github.com/radianceteam/ton-client-go/client"
)

func newClient() (*client.Client, error) {
	return client.NewClient(client.Config{
		Network: &client.NetworkConfig{ServerAddress: null.NewString("net.ton.dev", true)},
	}, client.WrapperConfig{MaxCGOConcurrentThreads: 10})
}

func printResult(data interface{}) error {
	f := colorjson.NewFormatter()
	f.Indent = 4
	s, err := f.Marshal(data)
	fmt.Println(string(s))

	return err
}

func main() { // nolint funlen
	c, err := newClient()
	if err != nil {
		log.Fatal(err)
	}
	getAccount := &cobra.Command{
		Use:   "get_account [account id]",
		Short: "Find account with id specified and prints it",
		Args:  cobra.ExactArgs(1),
		// 0:5223ad38763f4985141d54d49d509144dc2beecead7386d2b909855e7198e8a5
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			res, err := c.NetQueryCollection(&client.ParamsOfQueryCollection{
				Collection: "accounts",
				Filter:     json.RawMessage(`{"id":{"eq":"` + id + `"}}`),
				Limit:      null.Uint32From(20),
				Result:     "id acc_type balance",
			})
			if err != nil {
				return err
			}
			if len(res.Result) != 0 {
				return printResult(res.Result[0])
			}
			fmt.Printf("account with id %q not found\n", id)

			return nil
		},
	}

	genKeys := &cobra.Command{
		Use:   "gen_keys",
		Short: "Generates keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := c.CryptoGenerateRandomSignKeys()
			if err != nil {
				return err
			}
			fmt.Printf("public: %s\nsecret: %s\n", res.Public, res.Secret)

			return nil
		},
	}

	getGenesis := &cobra.Command{
		Use:   "get_config",
		Short: "Gets network configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := c.NetQueryCollection(&client.ParamsOfQueryCollection{
				Collection: "blocks",
				Filter:     json.RawMessage(`{"workchain_id":{"eq":-1}, "key_block":{"eq": true}}`),
				Limit:      null.Uint32From(1),
				Order: []client.OrderBy{{
					Path:      "seq_no",
					Direction: client.DescSortDirection,
				}},
				Result: `
id
master {
  config {
	p15 {
	  validators_elected_for
	  elections_start_before
	  elections_end_before
	  stake_held_for
	}
	p16 {
	  max_validators
	  max_main_validators
	  min_validators
	}
	p17 {
	  min_stake
	  max_stake
	  min_total_stake
	  max_stake_factor
	}
	p34 {
	  utime_since
	  utime_until
	  total
	  total_weight
	}
  }
}
`,
			})
			if err != nil {
				return err
			}
			if len(res.Result) != 0 {
				return printResult(res.Result[0])
			}
			fmt.Printf("genesis block not found\n")

			return nil
		},
	}
	rootCmd := &cobra.Command{Use: "cli"}
	rootCmd.AddCommand(getAccount, genKeys, getGenesis)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
