package godictbot

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Dict struct {
	Field    string
	Headword string
	Means    string
}

func (self *Dict) String() string {
	data := []string{
		fmt.Sprintf("[%v]", strings.TrimSpace(self.Field)),
		fmt.Sprintf("* %v", strings.TrimSpace(self.Headword)),
		fmt.Sprintf("  %v", strings.TrimSpace(self.Means)),
		"",
	}

	return strings.Join(data, "\n")
}

func DictList(c appengine.Context, query string) []*Dict {
	query = fmt.Sprintf("http://dic.daum.net/search.do?q=%v", query)
	client := urlfetch.Client(c)
	res, err := client.Get(query)
	if err != nil {
		return nil
	}

	html, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil
	}

	input := string(html)
	regex := `(?s)<[^>]+class\s*=\s*"\s*?tit_searchfd\s*">([^>]+?)</h4>.+?class="\s*link_txt\s*">([^>]+?)</a>.+?class="\s*txt_means_[^"]+?\s*">([^>]+?)</div>`

	dictList := make([]*Dict, 0, 4)
	re, err := regexp.Compile(regex)
	if nil == err {
		res := re.FindAllStringSubmatch(input, -1)
		if len(res) <= 0 {
			return nil
		}

		length := len(res)
		for i := 0; i < length; i++ {
			dict := &Dict{
				Field:    res[i][1],
				Headword: res[i][2],
				Means:    res[i][3],
			}
			dictList = append(dictList, dict)
		}
	}

	return dictList
}
