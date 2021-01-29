package main

import (
	"github.com/ethereum/hive/hivesim"
)

func main() {
	suite := hivesim.Suite{
		Name: "trace",
		Description: `
The tracing test suite uses a running node to test transaction tracing. It makes sure
the traces are correct and adhere to the EIP-3155 standard.`,
	}
	suite.Add(&hivesim.ClientTestSpec{
		Name:        "client launch",
		Description: `This is a meta-test that launches the client and runs the rest of the tests.`,
		Parameters:  hivesim.Params{},
		Files:       map[string]string{},
		Run:         runAllTests,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func runAllTests(t *hivesim.T, c *hivesim.Client) {
	t.Fatal("Not implemented")
}
