package gopandoc

import (
	"bytes"
	"os/exec"
	"strings"
)

const (
	CMD            = `pandoc`
	OP_HTML        = `html`
	OP_MARKDOWN    = `markdown-auto_identifiers` // Don't auto generate id for header
	OP_ATX_HEADERS = `--atx-headers`             // User atx header with "#"
)

func Check() error {
	_, err := exec.LookPath(CMD)
	return err
}

func ToHtml(mdStr string) (htmlStr string, err error) {

	cmd := exec.Command(CMD, "-f", OP_MARKDOWN, "-t", OP_HTML)

	cmd.Stdin = strings.NewReader(mdStr)
	b, err := cmd.Output()
	if err != nil {
		return
	}

	htmlStr = (string)(b)

	return
}

func ToMarkdown(htmlStr string) (mdStr string, err error) {

	cmd := exec.Command(CMD, "-f", OP_HTML, "-t", OP_MARKDOWN, OP_ATX_HEADERS)

	cmd.Stdin = strings.NewReader(htmlStr)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	mdStr = out.String()

	return
}
