package cmd

import "fmt"

// Stub functions to simulate nix commands
func ListFlakes() []string {
	return []string{"flake1", "flake2", "flake3"}
}

func ListProfilePkgs() []string {
	return []string{"htop", "ripgrep", "fzf"}
}

func InstallPkg(pkg string) {
	fmt.Println("Installing package:", pkg)
}
