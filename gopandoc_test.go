package gopandoc

import (
	"testing"
)

func TestCheck(t *testing.T) {
	err := Check()
	if err != nil {
		t.Error(err)
	}
}

func TestToHtml(t *testing.T) {
	mdStr := "**Hello World!**"
	expected_htmlStr := "<p><strong>Hello World!</strong></p>\n"

	htmlStr, err := ToHtml(mdStr)
	if err != nil {
		t.Error(err)
	}
	if htmlStr != expected_htmlStr {
		t.Error("ToHtml doesn't work as expected!")
	}

	mdStr = "# This is the H1 title\n## This is the H2 title"
	expected_htmlStr = "<h1>This is the H1 title</h1>\n<h2>This is the H2 title</h2>\n"

	htmlStr, err = ToHtml(mdStr)
	if err != nil {
		t.Error(err)
	}
	if htmlStr != expected_htmlStr {
		t.Error("ToHtml doesn't work as expected!")
		t.Errorf("ToMarkdown doesn't work as expeted! \n -- Expected: %+v \n -- Actually: %+v",
			expected_htmlStr, htmlStr)
	}

}

func TestToMarkdown(t *testing.T) {
	htmlStr := "<p><strong>Hello World!</strong></p>"
	expected_mdStr := "**Hello World!**\n"

	mdStr, err := ToMarkdown(htmlStr)
	if err != nil {
		t.Error(err)
	}

	if mdStr != expected_mdStr {
		t.Errorf("ToMarkdown doesn't work as expeted! \n -- Expected: %+v \n -- Actually: %+v",
			expected_mdStr, mdStr)
	}

	htmlStr = "<h1>This is the H1 title</h1>"
	expected_mdStr = "# This is the H1 title\n"

	mdStr, err = ToMarkdown(htmlStr)
	if err != nil {
		t.Error(err)
	}

	if mdStr != expected_mdStr {
		t.Errorf("ToMarkdown doesn't work as expeted! \n -- Expected: %+v \n -- Actually: %+v",
			expected_mdStr, mdStr)
	}
}
