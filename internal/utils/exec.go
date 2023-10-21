package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func ExecuteFFmpegCommand(inputFile string, preArgs []string, args []string) error {
	var newArgs []string
	newArgs = append([]string{"-hide_banner"})
	newArgs = append(newArgs, preArgs...)
	newArgs = append(newArgs, "-i", inputFile)
	newArgs = append(newArgs, args...)

	fmt.Println(newArgs)

	cmd := exec.Command("ffmpeg", newArgs...)

	// Capture both standard error and standard output
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	// Execute the command
	if err := cmd.Run(); err != nil {
		// Log the standard error
		log.Error().Err(err).Msgf("ffmpeg error: %v\n%s", err, stderr.String())
		return fmt.Errorf("ffmpeg failed: %v\n%s", err, stderr.String())
	}

	// If you also want to handle or log the standard output, you can use stdout.String() here

	return nil
}

func ExecuteFFprobeCommand(args []string, file string) ([]byte, error) {
	args = append(args, file)
	cmd := exec.Command("ffprobe", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func GenerateStoryboardImage(args []string, inputFile string, outputFile string) error {
	// set env vars
	os.Setenv("MAGICK_HEIGHT_LIMIT", "1000MP")
	os.Setenv("MAGICK_WIDTH_LIMIT", "1000MP")
	os.Setenv("MAGICK_DISK_LIMIT", "5GB")

	args = append(args, inputFile, outputFile)

	cmd := exec.Command("montage", args...)
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
