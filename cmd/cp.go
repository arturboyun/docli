package cmd

import (
	"log"
	"strings"

	client "github.com/arturboyun/docli/internal/docker"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp [OPTIONS] CONTAINER:PATH DEST",
	Short: "Copy files/folders between a container and the local filesystem",
	RunE: func(cmd *cobra.Command, args []string) error {
		cli, err := client.New()

		if err != nil {
			log.Fatalf("failed to create docker client: %v", err)
		}

		defer cli.Close()

		if len(args) != 2 {
			return cmd.Help()
		}

		src := args[0]
		dest := args[1]

		log.Printf("Copying from %s to %s", src, dest)

		parts := strings.SplitN(src, ":", 2)
		if len(parts) != 2 {
			return cmd.Help()
		}
		containerID := parts[0]
		srcPath := parts[1]

		_, err = cli.ContainerCopyFrom(containerID, srcPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)
}
