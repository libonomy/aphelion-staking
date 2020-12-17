package main

import (
	"fmt"
	"os"

	"github.com/libonomy/aphelion-staking/libs/log"
	"github.com/libonomy/aphelion-staking/privval"
)

var (
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Expected three args: <old path> <new key path> <new state path>")
		fmt.Println(
			"Eg. ~/.aphelion/config/priv_validator.json" +
				" ~/.aphelion/config/priv_validator_key.json" +
				" ~/.aphelion/data/priv_validator_state.json",
		)
		os.Exit(1)
	}
	err := loadAndUpgrade(args[0], args[1], args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadAndUpgrade(oldPVPath, newPVKeyPath, newPVStatePath string) error {
	oldPV, err := privval.LoadOldFilePV(oldPVPath)
	if err != nil {
		return fmt.Errorf("Error reading OldPrivValidator from %v: %v\n", oldPVPath, err)
	}
	logger.Info("Upgrading PrivValidator file",
		"old", oldPVPath,
		"newKey", newPVKeyPath,
		"newState", newPVStatePath,
	)
	oldPV.Upgrade(newPVKeyPath, newPVStatePath)
	return nil
}
