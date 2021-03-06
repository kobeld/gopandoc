package gopandoc

import (
	"bytes"
	"fmt"
	gopandoc "github.com/theplant/pandora/clients/go"
	"log"
	"os/exec"
	"runtime/debug"
	"strings"
)

const (
	op_html        = `html`
	op_markdown    = `markdown-escaped_line_breaks` // Don't auto generate id for header
	op_atx_headers = `--atx-headers`                // User atx header with "#"
	no_wrap        = "--no-wrap"
)

var (
	EnablePandocZMQMode = false
)

func Check() error {
	path, err := exec.LookPath("pandoc")
	log.Println(path)
	return err
}

func ToHtml(mdStr string) (htmlStr string, err error) {
	if EnablePandocZMQMode {
		return gopandoc.ToHTML(mdStr)
	}
	htmlStr, err = bash(fmt.Sprintf("pandoc -f %s -t %s %s",
		op_markdown, op_html, op_atx_headers), mdStr)
	if err != nil {
		return
	}

	return
}

func ToMarkdown(htmlStr string) (mdStr string, err error) {
	if EnablePandocZMQMode {
		return gopandoc.ToMarkdown(htmlStr)
	}
	mdStr, err = bash(fmt.Sprintf("pandoc -f %s -t %s %s %s", op_html,
		op_markdown, op_atx_headers, no_wrap), htmlStr)
	if err != nil {
		return
	}

	return
}

func bash(bash, content string) (out string, err error) {
	var buf bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", bash)
	cmd.Stdin = strings.NewReader(content)
	cmd.Stderr = &buf
	cmd.Stdout = &buf

	err = cmd.Run()
	if err != nil {
		printStackAndError(err)
		cmd.Process.Release()
		buf.Reset()
		return
	}

	out = buf.String()

	// Clean up resource
	cmd.Process.Kill()
	buf.Reset()
	debug.FreeOSMemory()

	return
}

func printStackAndError(err error) {
	log.Printf("********** Debug Error message: %+v ***********\n", err)
	debug.PrintStack()
}
