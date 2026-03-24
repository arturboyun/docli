package cmd

import (
	"fmt"
	"log"
	"strings"

	client "github.com/arturboyun/docli/internal/docker"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List of running containers",
	RunE: func(cmd *cobra.Command, args []string) error {

		cli, err := client.New()

		if err != nil {
			log.Fatalf("failed to create docker client: %v", err)
		}
		defer cli.Close()

		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return err
		}

		containers, err := cli.ContainerList(client.ContainerListOptions{
			All: all,
		})
		if err != nil {
			return err
		}

		spacesFpattern := "%-14s %-26s %-30s %-24s %s\n"
		fmt.Printf(spacesFpattern, "CONTAINER ID", "NAMES", "IMAGE", "STATUS", "PORTS")
		for _, c := range containers.Items {
			portsText := make([]string, len(c.Ports))
			for i, p := range c.Ports {
				if p.PublicPort != 0 {
					portsText[i] = fmt.Sprintf("%d->%d/%s", p.PublicPort, p.PrivatePort, p.Type)
				} else {
					portsText[i] = fmt.Sprintf("%d/%s", p.PrivatePort, p.Type)
				}
			}
			fmt.Printf(spacesFpattern, c.ID[:12], strings.Trim(c.Names[0], "/"), c.Image, c.Status, strings.Join(portsText, ", "))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	psCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
}
