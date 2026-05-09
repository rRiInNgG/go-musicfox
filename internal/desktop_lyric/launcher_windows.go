//go:build windows

package desktop_lyric

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/go-musicfox/go-musicfox/internal/configs"
)

var defaultCommands = []string{
	"Lyricify Lite.exe",
	"LyricifyLite.exe",
	"Lyricify.exe",
}

func autoStart() {
	cfg := configs.AppConfig.Main.Lyric.Desktop
	if !cfg.Enable || !cfg.AutoStart {
		return
	}

	command := strings.TrimSpace(cfg.Command)
	if command == "" {
		command = findDefaultCommand()
	}
	if command == "" {
		slog.Warn("desktop lyric helper is enabled but no launch command was found")
		return
	}

	cmd := exec.Command(command, cfg.Args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Start(); err != nil {
		slog.Warn("failed to start desktop lyric helper", "command", command, "error", err)
		return
	}
	slog.Info("started desktop lyric helper", "command", command)
}

func findDefaultCommand() string {
	for _, candidate := range defaultCommands {
		if path, err := exec.LookPath(candidate); err == nil {
			return path
		}
	}

	for _, dir := range defaultSearchDirs() {
		for _, name := range defaultCommands {
			path := filepath.Join(dir, name)
			if _, err := os.Stat(path); err == nil {
				return path
			}
		}
	}
	return ""
}

func defaultSearchDirs() []string {
	localAppData := os.Getenv("LOCALAPPDATA")
	programFiles := os.Getenv("ProgramFiles")
	programFilesX86 := os.Getenv("ProgramFiles(x86)")

	var dirs []string
	for _, base := range []string{localAppData, programFiles, programFilesX86} {
		if base == "" {
			continue
		}
		dirs = append(dirs,
			filepath.Join(base, "Programs", "Lyricify Lite"),
			filepath.Join(base, "Lyricify Lite"),
			filepath.Join(base, "Lyricify"),
		)
	}
	return dirs
}
