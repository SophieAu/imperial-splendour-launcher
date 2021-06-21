package backend

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// Handler is an abstraction of all direct system interaction
type Handler interface {
	MkdirAll(filePath string) error
	WriteFile(filePath string, data []byte) error
	ReadFile(filePath string) ([]byte, error)
	MoveFile(source, destination string) error
	RunCommand(name string, arg ...string) error
	StartCommand(name string, arg ...string) error
	DoesFileExist(path string) (bool, error)
	GetDirContentByName(dirname string) ([]string, error)
	DownloadFile(url string, targetFilePath string) error
	ZipFiles(filename string, files []string) error
	Exit(exitCode int)
}

type SystemHandler struct {
}

func rename(source, destination string) error {
	return os.Rename(source, destination)
}

func copyPasteDeleteFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return errors.New("Couldn't open source file: " + err.Error())
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return errors.New("Couldn't open dest file: " + err.Error())
	}

	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return errors.New("Writing to output file failed: " + err.Error())
	}

	// The copy was successful, so now delete the original file
	if os.Remove(sourcePath); err != nil {
		return errors.New("Failed removing original file: " + err.Error())
	}

	return nil
}

func (w *SystemHandler) MkdirAll(filePath string) error {
	return os.MkdirAll(filePath, os.FileMode(0777))
}

func (w *SystemHandler) WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, os.FileMode(0644))
}

func (w *SystemHandler) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (w *SystemHandler) StartCommand(name string, arg ...string) error {
	return exec.Command(name, arg...).Start()
}
func (w *SystemHandler) RunCommand(name string, arg ...string) error {
	return exec.Command(name, arg...).Run()
}

func (w *SystemHandler) MoveFile(source, destination string) error {
	if err := rename(source, destination); err != nil {
		return copyPasteDeleteFile(source, destination)
	}
	return nil
}

func (w *SystemHandler) DoesFileExist(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil

	} else if os.IsNotExist(err) {
		return false, nil

	} else {
		return false, err
	}
}

// zipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func (w *SystemHandler) ZipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

func (w *SystemHandler) GetDirContentByName(dirname string) ([]string, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	fileList := []string{}
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}
	return fileList, nil
}

func (w *SystemHandler) DownloadFile(url string, targetFilePath string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (w *SystemHandler) Exit(exitCode int) {
	os.Exit(exitCode)
}
