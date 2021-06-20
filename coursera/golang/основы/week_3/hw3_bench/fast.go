package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type User struct {
	Browsers []string
	Email    string
	Name     string
}

var linesPool = sync.Pool{New: func() interface{} {
	lines := make([]byte, 500*1024)
	return lines
}}
var seenBrowsers = make(map[string]byte)
var r = regexp.MustCompile("@")

func FastSearch(out io.Writer) {
	foundUsers := ""
	i := -1

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	user := &User{}
	for sc.Scan() {
		data := sc.Bytes()
		//data := []byte(line)
		user.UnmarshalJSON(data)

		i++

		isAndroid := false
		isMSIE := false
		for _, browser := range user.Browsers {

			okAndroid := strings.Contains(browser, "Android")
			okMSIE := strings.Contains(browser, "MSIE")
			isAndroid = isAndroid || okAndroid
			isMSIE = isMSIE || okMSIE

			if okAndroid || okMSIE {
				if _, ok := seenBrowsers[browser]; !ok {
					seenBrowsers[browser] = 0
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}
		email := r.ReplaceAllString(user.Email, " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email)

	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

// easyjson codegen

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9f2eff5fDecodeGithubComAndrewalfCourseraGoMailruHw3BenchUser(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9f2eff5fEncodeGithubComAndrewalfCourseraGoMailruHw3BenchUser(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9f2eff5fEncodeGithubComAndrewalfCourseraGoMailruHw3BenchUser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9f2eff5fEncodeGithubComAndrewalfCourseraGoMailruHw3BenchUser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9f2eff5fDecodeGithubComAndrewalfCourseraGoMailruHw3BenchUser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9f2eff5fDecodeGithubComAndrewalfCourseraGoMailruHw3BenchUser(l, v)
}
