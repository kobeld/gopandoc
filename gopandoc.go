package gopandoc

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const (
	op_html        = `html`
	op_markdown    = `markdown-auto_identifiers` // Don't auto generate id for header
	op_atx_headers = `--atx-headers`             // User atx header with "#"
)

func Check() error {
	path, err := exec.LookPath("pandoc")
	log.Println(path)
	return err
}

func ToHtml(mdStr string) (htmlStr string, err error) {

	htmlStr, err = bash(fmt.Sprintf("pandoc -f %s -t %s %s",
		op_markdown, op_html, op_atx_headers), mdStr)
	if err != nil {
		return
	}

	return
}

func ToMarkdown(htmlStr string) (mdStr string, err error) {

	mdStr, err = bash(fmt.Sprintf("pandoc -f %s -t %s %s", op_html,
		op_markdown, op_atx_headers), htmlStr)
	if err != nil {
		return
	}

	return
}

func bash(bash, content string) (out string, err error) {
	cmd := exec.Command("/bin/sh", "-c", bash)
	cmd.Stdin = strings.NewReader(content)
	var buf bytes.Buffer
	cmd.Stderr = &buf
	cmd.Stdout = &buf
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		return
	}

	out = buf.String()
	return
}
