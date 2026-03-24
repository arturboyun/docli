package client

import (
	"context"

	mobyclient "github.com/moby/moby/client"
)

type Client struct {
	ctx context.Context
	cli *mobyclient.Client
}

func New() (*Client, error) {
	ctx := context.Background()
	cli, err := mobyclient.New(mobyclient.FromEnv)
	if err != nil {
		return nil, err
	}
	return &Client{ctx: ctx, cli: cli}, nil
}

func (c *Client) Close() error {
	return c.cli.Close()
}

type ContainerListOptions struct {
	All bool
}

func (c *Client) ContainerList(options ContainerListOptions) (mobyclient.ContainerListResult, error) {
	containers, err := c.cli.ContainerList(c.ctx, mobyclient.ContainerListOptions{
		All: options.All,
	})
	if err != nil {
		return mobyclient.ContainerListResult{}, err
	}
	return containers, nil
}

type ContainerLogsOptions struct {
	ShowStdout bool
	ShowStderr bool
	Follow     bool
	Timestamps bool
	Tail       string
}

func (c *Client) ContainerLogs(containerID string, options ContainerLogsOptions) (mobyclient.ContainerLogsResult, error) {
	logs, err := c.cli.ContainerLogs(c.ctx, containerID, mobyclient.ContainerLogsOptions{
		ShowStdout: options.ShowStdout,
		ShowStderr: options.ShowStderr,
		Follow:     options.Follow,
		Timestamps: options.Timestamps,
		Tail:       options.Tail,
	})
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (c *Client) ContainerCopyFrom(containerID, srcPath string) (mobyclient.CopyFromContainerResult, error) {
	res, err := c.cli.CopyFromContainer(c.ctx, containerID, mobyclient.CopyFromContainerOptions{
		SourcePath: srcPath,
	})
	if err != nil {
		return mobyclient.CopyFromContainerResult{}, err
	}
	return res, nil
}

type ConsoleSize struct {
	Height, Width uint
}

type ExecAttachOptions struct {
	TTY         bool
	ConsoleSize ConsoleSize
}

func (c *Client) ExecAttach(containerID string, cmd []string, options ExecAttachOptions) (mobyclient.ExecAttachResult, error) {
	res, err := c.cli.ExecAttach(
		c.ctx,
		containerID,
		mobyclient.ExecAttachOptions{
			TTY: options.TTY,
			ConsoleSize: mobyclient.ConsoleSize{
				Height: options.ConsoleSize.Height,
				Width:  options.ConsoleSize.Width,
			},
		},
	)
	if err != nil {
		return mobyclient.ExecAttachResult{}, err
	}
	return res, nil
}
