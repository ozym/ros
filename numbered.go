package ros

import (
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

func scanFlags(s *scanner.Scanner) map[string]string {
	ws := s.Whitespace
	mo := s.Mode
	id := s.IsIdentRune
	defer func() {
		s.Whitespace = ws
		s.Mode = mo
		s.IsIdentRune = id
	}()
	s.Whitespace = 1<<'\t' | 1<<'-' | 1<<'\r' | 1<<' ' | 1<<','
	s.IsIdentRune = func(ch rune, i int) bool {
		return unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}
	s.Mode = scanner.ScanIdents

	flags := make(map[string]string)

	var tok rune
	for {
		var key string
		if tok = s.Scan(); tok == scanner.EOF || tok == '\n' {
			break
		}
		key = s.TokenText()
		if tok = s.Scan(); tok == scanner.EOF || tok == '\n' {
			break
		}
		flags[key] = s.TokenText()
	}

	return flags
}

func skipComment(s *scanner.Scanner) {
	ws := s.Whitespace
	defer func() {
		s.Whitespace = ws
	}()
	s.Whitespace = 1 << '\r'

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		if tok == '\n' {
			break
		}
	}
}

func scanComment(s *scanner.Scanner) string {
	ws := s.Whitespace
	defer func() {
		s.Whitespace = ws
	}()
	s.Whitespace = 1 << '\r'

	comment := ""

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		if tok == '\n' {
			break
		}
		comment = comment + s.TokenText()
	}
	return comment
}

func scanNumberedItem(s *scanner.Scanner, f map[string]string) (map[string]string, error) {

	var key string

	res := make(map[string]string)

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		if tok == '\n' && s.Peek() == '\n' {
			break
		}

		switch s.TokenText() {
		case "#":
			skipComment(s)
		case ";;;":
			res["comment"] = scanComment(s)
		default:
			if s.Peek() == '=' {
				key = s.TokenText()
			} else if key == "" {
				for _, v := range s.TokenText() {
					if _, ok := f[string(v)]; ok {
						res[f[string(v)]] = "yes"
					}
				}
			} else if s.TokenText() != "\n" && s.TokenText() != "=" {
				u, err := strconv.Unquote(s.TokenText())
				if err != nil {
					u = s.TokenText()
				}
				if _, ok := res[key]; ok {
					res[key] = res[key] + " " + u
				} else {
					res[key] = u
				}
				if _, ok := res["comment"]; !ok {
					res["comment"] = ""
				}
			}
		}
	}

	return res, nil
}

func ScanNumberedItemList(results string) ([]map[string]string, error) {
	var list []map[string]string

	var s scanner.Scanner
	s.Init(strings.NewReader(results))

	s.Mode = scanner.ScanIdents | scanner.ScanStrings
	s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == ':' || ch == ';' || ch == '/' || ch == '-' || ch == ',' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}

	var flags map[string]string

	var tok rune
	for tok != scanner.EOF {
		if tok = s.Scan(); tok == '\n' || tok == scanner.EOF {
			continue
		}
		if s.TokenText() != "Flags:" {
			no := s.TokenText()

			item, err := scanNumberedItem(&s, flags)
			if err != nil {
				return nil, err
			}

			ans := make(map[string]string)
			ans["number"] = no
			for k, v := range item {
				ans[k] = strings.TrimSpace(v)
			}
			for _, v := range flags {
				if _, ok := ans[v]; !ok {
					ans[v] = "no"
				}
			}

			list = append(list, ans)
		} else {
			flags = scanFlags(&s)
		}
	}

	return list, nil
}

func ScanFirstNumberedItemList(results string) (map[string]string, error) {
	list, err := ScanNumberedItemList(results)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}
