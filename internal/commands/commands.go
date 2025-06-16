package commands

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"
)

func StreamCommandOutput(ctx context.Context, command string, out chan<- string, arg ...string) error {
	cmd := exec.CommandContext(ctx, command, arg...)
	log.Printf("%v", cmd)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout for command '%s': %w", command, err)
	}

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to run command '%s': %w", command, err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			cmd.Process.Kill()
			close(out)
			return ctx.Err()
		case out <- scanner.Text():
		}
	}

	if err = scanner.Err(); err != nil {
		cmd.Wait()
		close(out)
		return fmt.Errorf("scanner error for command '%s': %w", command, err)
	}

	err = cmd.Wait()
	log.Println(err)

	close(out)

	return err
}
