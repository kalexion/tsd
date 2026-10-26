package executor

import (
	"fmt"
	"os"
	"path/filepath"
)

var initFileContent = []byte(`session "example-session" {
    window "example-window-1" {
        run = [
            "echo hello",
        ]
    }

    window "example-window-2" {
        run = [
            "echo hello again",
        ]
    }

    focus = "example-window-1"
}
`)

func Init(currentDirectory string, configFileName string) (retErr error) {
	filePath := filepath.Join(currentDirectory, configFileName)

	file, err := os.OpenFile(
		filePath,
		os.O_CREATE|os.O_EXCL|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return fmt.Errorf("create config file: %w", err)
	}

	// Remove the incomplete file if initialization fails.
	defer func() {
		if err := file.Close(); err != nil && retErr == nil {
			retErr = fmt.Errorf("close config file: %w", err)
		}

		if retErr != nil {
			if err := os.Remove(filePath); err != nil {
				retErr = fmt.Errorf(
					"%w; remove incomplete config file: %v",
					retErr,
					err,
				)
			}
		}
	}()

	n, err := file.Write(initFileContent)
	if err != nil {
		return fmt.Errorf("write config file: %w", err)
	}
	if n != len(initFileContent) {
		return fmt.Errorf(
			"write config file: %w",
			fmt.Errorf("short write: wrote %d of %d bytes", n, len(initFileContent)),
		)
	}

	return nil
}
