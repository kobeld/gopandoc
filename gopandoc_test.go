package gopandoc

import "testing"

func TestCheck(t *testing.T) {
	err := Check()
	if err != nil {
		t.Error(err)
	}
}

func TestToHtml(t *testing.T) {
	return
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

var sampleHTML = `<h1 id="an-exhibit-of-markdown">An exhibit of Markdown</h1>
<p>This note demonstrates some of what <a href="http://daringfireball.net/projects/markdown/">Markdown</a> is capable of doing.</p>
<p><em>Note: Feel free to play with this page. Unlike regular notes, this doesn't automatically save itself.</em></p>
<h2 id="basic-formatting">Basic formatting</h2>
<p>Paragraphs can be written like so. A paragraph is the basic block of Markdown. A paragraph is what text will turn into when there is no reason it should become anything else.</p>
<p>Paragraphs must be separated by a blank line. Basic formatting of <em>italics</em> and <strong>bold</strong> is supported. This <em>can be <strong>nested</strong> like</em> so.</p>
<h2 id="lists">Lists</h2>
<h3 id="ordered-list">Ordered list</h3>
<ol style="list-style-type: decimal">
<li>Item 1</li>
<li>A second item</li>
<li>Number 3</li>
<li>Ⅳ</li>
</ol>
<p><em>Note: the fourth item uses the Unicode character for <a href="http://www.fileformat.info/info/unicode/char/2163/index.htm">Roman numeral four</a>.</em></p>
<h3 id="unordered-list">Unordered list</h3>
<ul>
<li>An item</li>
<li>Another item</li>
<li>Yet another item</li>
<li>And there's more...</li>
</ul>
<h2 id="paragraph-modifiers">Paragraph modifiers</h2>
<h3 id="code-block">Code block</h3>
<pre><code>Code blocks are very useful for developers and other people who look at code or other things that are written in plain text. As you can see, it uses a fixed-width font.</code></pre>
<p>You can also make <code>inline code</code> to add code into other things.</p>
<h3 id="quote">Quote</h3>
<blockquote>
<p>Here is a quote. What this is should be self explanatory. Quotes are automatically indented when they are used.</p>
</blockquote>
<h2 id="headings">Headings</h2>
<p>There are six levels of headings. They correspond with the six levels of HTML headings. You've probably noticed them already in the page. Each level down uses one more hash character.</p>
<h3 id="headings-can-also-contain-formatting">Headings <em>can</em> also contain <strong>formatting</strong></h3>
<h3 id="they-can-even-contain-inline-code">They can even contain <code>inline code</code></h3>
<p>Of course, demonstrating what headings look like messes up the structure of the page.</p>
<p>I don't recommend using more than three or four levels of headings here, because, when you're smallest heading isn't too small, and you're largest heading isn't too big, and you want each size up to look noticeably larger and more important, there there are only so many sizes that you can use.</p>
<h2 id="urls">URLs</h2>
<p>URLs can be made in a handful of ways:</p>
<ul>
<li>A named link to <a href="http://www.markitdown.net/">MarkItDown</a>. The easiest way to do these is to select what you want to make a link and hit <code>Ctrl+L</code>.</li>
<li>Another named link to <a href="http://www.markitdown.net/">MarkItDown</a></li>
<li>Sometimes you just want a URL like <a href="http://www.markitdown.net/">http://www.markitdown.net/</a>.</li>
</ul>
<h2 id="horizontal-rule">Horizontal rule</h2>
<p>A horizontal rule is a line that goes across the middle of the page.</p>
<hr />
<p>It's sometimes handy for breaking things up.</p>
<h2 id="images">Images</h2>
<p>Markdown can also contain images. I'll need to add something here sometime.</p>
<h2 id="finally">Finally</h2>
<p>There's actually a lot more to Markdown than this. See the official <a href="http://daringfireball.net/projects/markdown/basics">introduction</a> and <a href="http://daringfireball.net/projects/markdown/syntax">syntax</a> for more information. However, be aware that this is not using the official implementation, and this might work subtly differently in some of the little things.</p> 
`

func BenchmarkToMarkdownWithCommandLineCalling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToMarkdown(sampleHTML)
	}
	return
}

func BenchmarkToMarkdownWithZMQCalling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZToMarkdown(sampleHTML)
	}
	return
}
