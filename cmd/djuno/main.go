package main

import (
	"cosmossdk.io/simapp/params"
	junocmd "github.com/forbole/juno/v5/cmd"
	initcmd "github.com/forbole/juno/v5/cmd/init"
	migratecmd "github.com/forbole/juno/v5/cmd/migrate"
	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	startcmd "github.com/forbole/juno/v5/cmd/start"

	desmosapp "github.com/desmos-labs/desmos/v5/app"

	parsecmd "github.com/desmos-labs/djuno/v2/cmd/parse"
	desmosdb "github.com/desmos-labs/djuno/v2/database"
	"github.com/desmos-labs/djuno/v2/x"
)

func main() {
	// Setup the config
	parseCfg := parsecmdtypes.NewConfig().
		WithRegistrar(x.NewModulesRegistrar()).
		WithEncodingConfigBuilder(makeEncodingConfig).
		WithDBBuilder(desmosdb.Builder)

	cfg := junocmd.NewConfig("djuno").
		WithParseConfig(parseCfg)

	// Run the command
	rootCmd := junocmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		junocmd.VersionCmd(),
		initcmd.NewInitCmd(cfg.GetInitConfig()),
		startcmd.NewStartCmd(cfg.GetParseConfig()),
		parsecmd.NewParseCmd(cfg.GetParseConfig()),
		migratecmd.NewMigrateCmd(cfg.GetName(), cfg.GetParseConfig()),
	)

	executor := junocmd.PrepareRootCmd(cfg.GetName(), rootCmd)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func makeEncodingConfig() params.EncodingConfig {
	config := desmosapp.MakeEncodingConfig()
	return params.EncodingConfig{
		InterfaceRegistry: config.InterfaceRegistry,
		Codec:             config.Marshaler,
		TxConfig:          config.TxConfig,
		Amino:             config.Amino,
	}
}
