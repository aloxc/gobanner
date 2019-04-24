package banner

import (
	"io"
	"log"
	"os"
	"runtime"
	"text/template"
	"time"
)

var logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

// SetLog permits the user change the default logger instance
func SetLog(l *log.Logger) {
	if l == nil {
		return
	}
	logger = l
}

type vars struct {
	GoVersion      string
	GOOS           string
	GOARCH         string
	NumCPU         int
	GOPATH         string
	GOROOT         string
	Compiler       string
	AnsiColor      ansiColor
	AnsiBackground ansiBackground
}

func (v vars) Env(env string) string {
	return os.Getenv(env)
}

// See https://github.com/golang/go/blob/f06795d9b742cf3292a0f254646c23603fc6419b/src/time/format.go#L9-L41
func (v vars) Now(layout string) string {
	return time.Now().Format(layout)
}

// Init load the banner and prints it to output
// All errors are ignored, the application will not print the banner in case of error.
func Init(out io.Writer, isEnabled, isColorEnabled bool) {
	if !isEnabled {
		logger.Println("The banner is not enabled.")
		return
	}
	show(out, isColorEnabled, GO_BANNER)
}

func show(out io.Writer, isColorEnabled bool, content string) {
	t, err := template.New("banner").Parse(content)

	if err != nil {
		logger.Printf("Error trying to parse the banner file, err: %v", err)
		return
	}

	t.Execute(out, vars{
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		os.Getenv("GOPATH"),
		runtime.GOROOT(),
		runtime.Compiler,
		ansiColor{isColorEnabled},
		ansiBackground{isColorEnabled},
	})
}
