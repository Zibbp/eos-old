package utils

import (
	"io"
	"os"
	"os/exec"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ExecuteFFmpegCommand(inputFile string, args []string) error {
	args = append([]string{"-hide_banner", "-loglevel", "error", "-i", inputFile}, args...)
	cmd := exec.Command("ffmpeg", args...)
	// Initialize your zerolog logger.
	logger := log.Logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel) // Or whatever log level you desire.

	// Create a pipe for the standard error.
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// Create a pipe for the standard output.
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// Start the command.
	if err := cmd.Start(); err != nil {
		return err
	}

	// Create a multiwriter that writes to both standard error and the logger.
	stderrMulti := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, logger)

	// Create a multiwriter that writes to both standard output and the logger.
	stdoutMulti := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, logger)

	// Copy the command's output to the logger.
	go io.Copy(stderrMulti, stderr)
	go io.Copy(stdoutMulti, stdout)

	// Wait for the command to finish.
	return cmd.Wait()
}

func GenerateStoryboardImage(inputFile string, outputFile string) error {
	args := []string{"-geometry", "+0+0", "-tile", "1x", inputFile, outputFile}
	cmd := exec.Command("montage", args...)
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
