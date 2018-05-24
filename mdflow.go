package mdflow

import (
	"io/ioutil"

	"fmt"

	"bytes"

	"gopkg.in/russross/blackfriday.v2"
)

func ParseFile(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	md := blackfriday.New()
	n := md.Parse(data)
	buf := bytes.NewBuffer([]byte{})
	n.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		return walkFunc(node, entering, buf)
	})

	fmt.Println(buf.String())
	return nil
}

/*
const (
	Document NodeType = iota
	BlockQuote
	List
	Item
	Paragraph
	Heading
	HorizontalRule
	Emph
	Strong
	Del
	Link
	Image
	Text
	HTMLBlock
	CodeBlock
	Softbreak
	Hardbreak
	Code
	HTMLSpan
	Table
	TableCell
	TableHead
	TableBody
	TableRow
)
*/

func walkFunc(node *blackfriday.Node, entering bool, buf *bytes.Buffer) blackfriday.WalkStatus {

	switch node.Type {
	case blackfriday.Document:
	case blackfriday.BlockQuote:
	case blackfriday.List:

	case blackfriday.Item:
	case blackfriday.Paragraph:
	case blackfriday.Heading:
		if entering {
			buf.WriteString(fmt.Sprintf("h%d. ", node.HeadingData.Level))
		}
	case blackfriday.HorizontalRule:
	case blackfriday.Emph:
	case blackfriday.Strong:
	case blackfriday.Del:
	case blackfriday.Link:
	case blackfriday.Image:
	case blackfriday.Text:
		if entering {
			buf.Write(node.Literal)
		}
	case blackfriday.HTMLBlock:
	case blackfriday.CodeBlock:
	case blackfriday.Softbreak:
	case blackfriday.Hardbreak:
	case blackfriday.Code:
	case blackfriday.HTMLSpan:
	case blackfriday.Table:
	case blackfriday.TableCell:
	case blackfriday.TableHead:
	case blackfriday.TableBody:
	case blackfriday.TableRow:
	default:
	}
	if entering {
		fmt.Println("entering ", node.Type)
		fmt.Println(string(node.Literal))
	} else {
		fmt.Println("exiting ", node.Type)
	}

	return blackfriday.GoToNext
}
