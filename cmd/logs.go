package cmd

import (
	"log"
	"os"
	"strconv"

	client "github.com/arturboyun/docli/internal/docker"
	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [OPTIONS] CONTAINER",
	Short: "Fetch the logs of a container",
	RunE: func(cmd *cobra.Command, args []string) error {

		cli, err := client.New()

		if err != nil {
			log.Fatalf("failed to create docker client: %v", err)
		}
		defer cli.Close()

		if len(args) != 1 {
			return cmd.Help()
		}

		containerID := args[0]

		follow, err := cmd.Flags().GetBool("follow")
		if err != nil {
			return err
		}

		timestamps, err := cmd.Flags().GetBool("timestamps")
		if err != nil {
			return err
		}

		tail, err := cmd.Flags().GetInt64("tail")
		if err != nil {
			return err
		}

		tailStr := "all"
		if tail >= 0 {
			tailStr = strconv.FormatInt(tail, 10)
		}

		opts := client.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     follow,
			Timestamps: timestamps,
			Tail:       tailStr,
		}

		logs, err := cli.ContainerLogs(containerID, opts)
		if err != nil {
			return err
		}
		defer logs.Close()

		_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, logs)
		if err != nil {
			log.Fatalf("failed to copy logs output: %v", err)
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().BoolP("follow", "f", false, "Follow log output")
	logsCmd.Flags().BoolP("timestamps", "t", false, "Show timestamps")
	logsCmd.Flags().Int64P("tail", "", -1, "Number of lines to show from the end of the logs")
}
