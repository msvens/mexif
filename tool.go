package mexif

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

const StayOpenArg = "-stay_open"
const ExecuteArg = "-execute"
const Cmd = "exiftool"

const JsonArg = "-j"
const GroupArg = "-g2"


var initArgs = []string{StayOpenArg, "True", "-@", "-", "-common_args"}


type MExifTool struct {
	mutex    sync.Mutex
	stdin   io.WriteCloser
	stdout  io.ReadCloser
	scanout *bufio.Scanner
	closed bool
}

func NewMExifTool(flags ...string) (*MExifTool, error) {
	flags = append(initArgs, flags...)

	tool := MExifTool{closed:true}

	cmd := exec.Command(Cmd, flags...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	tool.stdout = stdout
	tool.stdin = stdin

	tool.scanout = bufio.NewScanner(stdout)
	tool.scanout.Split(splitReadyToken)

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	tool.closed = false

	return &tool, nil
}



func (tool *MExifTool) Close() error {
	tool.mutex.Lock()
	defer tool.mutex.Unlock()

	//Execute stay open false
	fmt.Fprintln(tool.stdin, StayOpenArg)
	fmt.Fprintln(tool.stdin, "False")
	fmt.Fprintln(tool.stdin, ExecuteArg)

	var errs []error
	if err := tool.stdout.Close(); err != nil {
		errs = append(errs, fmt.Errorf("error while closing stdout: %w", err))
	}

	if err := tool.stdin.Close(); err != nil {
		errs = append(errs, fmt.Errorf("error while closing stdin: %w", err))
	}
	tool.closed = true

	if len(errs) > 0 {
		return fmt.Errorf("error while closing exiftool: %w", errs)
	}

	return nil
}

func (tool *MExifTool) ExifCompact(path string) (*ExifCompact, error) {
	if d, err := tool.ExifData(path); err == nil {
		return NewExifCompact(d), nil
	} else {
		return nil, err
	}
}

func (tool *MExifTool) ExifData(path string) (*ExifData, error) {
	root, err := tool.Unmarshal(path)
	if err != nil {
		return nil, err
	}
	return NewExifData(root), nil
}

func (tool *MExifTool) Unmarshal(path string) (map[string]interface{},error) {
	bytes, err := tool.Read(path)
	if err != nil {
		return nil, err
	}
	var f []interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		return nil, err;
	}
	if len(f) < 1 {
		return nil, fmt.Errorf("no data")
	}
	return f[0].(map[string]interface{}), nil
}

func (tool *MExifTool) Read(path string) ([]byte, error) {
	return tool.ReadWithFlags(path)
}

func (tool *MExifTool) ReadWithFlags(path string, flags ...string) ([]byte, error) {
	tool.mutex.Lock()
	defer tool.mutex.Unlock()

	if tool.closed {
		return nil, fmt.Errorf("MExifTool is closed")
	}
	for _, f := range flags {
		fmt.Fprintln(tool.stdin, f)
	}
	fmt.Fprintln(tool.stdin, JsonArg)
	fmt.Fprintln(tool.stdin, GroupArg)
	fmt.Println(path)
	fmt.Fprintln(tool.stdin, path)
	fmt.Fprintln(tool.stdin, ExecuteArg)

	//read output
	if !tool.scanout.Scan() {
		return nil, fmt.Errorf("Failed to read output")
	} else {
		results := tool.scanout.Bytes()
		fmt.Println("len of results: ",len(results))
		sendResults := make([]byte, len(results), len(results))
		copy(sendResults, results)
		return sendResults, nil
	}
}


func splitReadyToken(data []byte, atEOF bool) (int, []byte, error) {
	delimPos := bytes.Index(data, []byte("{ready}\n"))
	delimSize := 8

	//windows
	if delimPos == -1 {
		delimPos = bytes.Index(data, []byte("{ready}\r\n"))
		delimSize = 9
	}

	if delimPos == -1 { // still no token found
		if atEOF {
			return 0, data, io.EOF
		} else {
			return 0, nil, nil
		}
	} else {
		if atEOF && len(data) == (delimPos+delimSize) { // nothing left to scan
			return delimPos + delimSize, data[:delimPos], bufio.ErrFinalToken
		} else {
			return delimPos + delimSize, data[:delimPos], nil
		}
	}
}

