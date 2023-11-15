package cli

import (
	"fmt"

	"github.com/web-of-things-open-source/tm-catalog-cli/internal/commands"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/remotes"
)

func FetchThingModelByName(name, remoteName string) error {
	remote, err := remotes.Get(remoteName)
	if err != nil {
		Stderrf("Could not ìnitialize a remote instance for %s: %v\ncheck config", remoteName, err)
		return err
	}

	fn := &commands.FetchName{}
	err = fn.Parse(name)
	if err != nil {
		Stderrf("Could not parse name %s: %v", name, err)
		return err
	}

	thing, err := commands.FetchThingByName(fn, remote)
	if err != nil {
		Stderrf("Could not fetch from remote: %v", err)
		return err
	}
	fmt.Println(string(thing))
	return nil
}
