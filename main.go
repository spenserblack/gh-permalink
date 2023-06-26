package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/cli/safeexec"
	"golang.org/x/exp/slices"
)

var helpOptions = []string{"-h", "-help", "--help"}
var linePattern = regexp.MustCompile(`^(?P<start>\d+)(?:-(?P<end>\d+))?$`)

func main() {
	// TODO Make dynamic?
	protocol := "https"

	repo, err := repository.Current()
	if err != nil {
		onError(err)
	}
	rev, err := head()
	if err != nil {
		onError(err)
	}

	args := os.Args[1:]
	argLen := len(args)
	if argLen < 1 || argLen > 2 {
		printUsage(os.Stderr, 1)
	}
	if slices.Contains(helpOptions, args[0]) {
		printUsage(os.Stdout, 0)
	}

	filename := args[0]
	// NOTE In GitHub, lines start at 1.
	var (
		lineStart int
		lineEnd   int
	)
	if argLen == 2 {
		lineMatch := linePattern.FindStringSubmatch(args[1])
		if lineMatch == nil {
			printUsage(os.Stderr, 1)
		}
		lineStart, _ = strconv.Atoi(lineMatch[1])
		if lineStart < 1 {
			onError(fmt.Errorf("line numbers start at 1, got %d", lineStart))
		}
		if lineMatch[2] != "" {
			lineEnd, _ = strconv.Atoi(lineMatch[2])
			if lineEnd < lineStart {
				onError(fmt.Errorf("line end must be greater than line start, got %d", lineEnd))
			}
		}
	}


	url := fmt.Sprintf(
		"%s://%s/%s/%s/blob/%s/%s",
		protocol,
		repo.Host, repo.Owner, repo.Name,
		rev,
		filename,
	)
	if lineStart > 0 {
		url = fmt.Sprintf("%s#L%d", url, lineStart)
	}
	if lineEnd > 0 {
		url = fmt.Sprintf("%s-L%d", url, lineEnd)
	}
	fmt.Println(url)
}

func onError(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

// Head gets the rev of the HEAD commit.
func head() (string, error) {
	git, err := safeexec.LookPath("git")
	if err != nil {
		return "", err
	}
	cmd := exec.Command(git, "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	// NOTE Command output ends in a newline
	return strings.TrimRight(string(out), "\n"), nil
}

func printUsage(out *os.File, code int) {
	fmt.Fprintf(out, `usage: gh permalink FILENAME [LINE | START_LINE-END_LINE]`)
	os.Exit(code)
}
