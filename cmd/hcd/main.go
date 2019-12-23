package main

import (
	app "github.com/cosmos/sdk-tutorials/hellochain"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
	"github.com/tendermint/tendermint/libs/cli"
)

func main() {

	params := starter.NewServerCommandParams(
		"hdc",                  // name of the command
		"hellochain AppDaemon", // description
		starter.NewAppCreator(app.NewHelloChainApp), // method for constructing an app
		starter.NewAppExporter(app.NewHelloChainApp), // method for exporting chain state
	)


	serverCmd := starter.NewServerCommand(params)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(serverCmd, "HC", starter.DefaultNodeHome)
	err := executor.Execute()

	if err != nil {
		panic(err)
	}
}
