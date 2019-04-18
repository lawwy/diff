package diff

import (
	"reflect"
	"testing"
)

type CompareTestData struct {
	A    string
	B    string
	want []Part
}

var datas []CompareTestData = []CompareTestData{
	CompareTestData{
		A: "abcde",
		B: "aobdi",
		want: []Part{
			Part{"a", false, false},
			Part{"o", false, true},
			Part{"b", false, false},
			Part{"c", true, false},
			Part{"d", false, false},
			Part{"e", true, false},
			Part{"i", false, true},
		},
	},
	// CompareTestData{
	// 	A: "一二三四五",
	// 	B: "一六二四七",
	// 	want: []Part{
	// 		Part{"一", false, false},
	// 		Part{"六", false, true},
	// 		Part{"二", false, false},
	// 		Part{"三", true, false},
	// 		Part{"四", false, false},
	// 		Part{"五", true, false},
	// 		Part{"七", false, true},
	// 	},
	// },
	// CompareTestData{
	// 	A: "百度发现有些人竟然把字符串转成数组再根据下标取。",
	// 	B: "百度有些人居然把字符串转成字符数组再根据下标取，简直了",
	// 	want: []Part{
	// 		Part{"百度", false, false},
	// 		Part{"发现", true, false},
	// 		Part{"有些人", false, false},
	// 		Part{"竟", true, false},
	// 		Part{"居", false, true},
	// 		Part{"然把字符串转成", false, false},
	// 		Part{"字符", false, true},
	// 		Part{"数组再根据下标取", false, false},
	// 		Part{"。", true, false},
	// 		Part{"，简直了", false, true},
	// 	},
	// },
}

func TestStringDiff2(t *testing.T) {
	for _, data := range datas {
		r := DiffStrings(data.A, data.B)
		if !reflect.DeepEqual(r, data.want) {
			t.Fatalf("compare fail")
		}
	}
}

type HtmlTokenizeTestData struct {
	data string
	want []string
}

var htmlTokenizeDatas []HtmlTokenizeTestData = []HtmlTokenizeTestData{
	HtmlTokenizeTestData{
		data: "hh<p>hi<span style=\"font-weight: bold;\">ab</span>nba</p>\n<p></p>\n",
		want: []string{"h", "h", "<p>", "h", "i", "<span style=\"font-weight: bold;\">", "a", "b", "</span>", "n", "b", "a", "</p>", "\n", "<p>", "</p>", "\n"},
	},
	HtmlTokenizeTestData{
		data: "<dd>你好世界。</hhh>世界，你好<kkk>",
		want: []string{"<dd>", "你", "好", "世", "界", "。", "</hhh>", "世", "界", "，", "你", "好", "<kkk>"},
	},
}

func TestHtmlTokenize(t *testing.T) {
	for _, data := range htmlTokenizeDatas {
		r := HtmlTokenize(data.data)
		if !reflect.DeepEqual(r, data.want) {
			t.Fatalf("html tokenize fail")
		}
	}
}

var htmlsData []CompareTestData = []CompareTestData{
	CompareTestData{
		A: "<p>世界你好。</p> <span>在线编纂系统真好用！</span><b>黑人问号</b>",
		B: "<h1>世界你不好。</h1> <span style=\"right: 200px;\">智能编纂系统好用吗？</span><b>黑人问号</b>",
		want: []Part{
			Part{"<p>", true, false},
			Part{"<h1>", false, true},
			Part{"世界你", false, false},
			Part{"不", false, true},
			Part{"好。", false, false},
			Part{"</p>", true, false},
			Part{"</h1>", false, true},
			Part{" ", false, false},
			Part{"<span>在线", true, false},
			Part{"<span style=\"right: 200px;\">智能", false, true},
			Part{"编纂系统", false, false},
			Part{"真", true, false},
			Part{"好用", false, false},
			Part{"！", true, false},
			Part{"吗？", false, true},
			Part{"</span><b>黑人问号</b>", false, false},
		},
	},
}

func TestHtmlDiff(t *testing.T) {
	for _, data := range htmlsData {
		r := DiffHtmls(data.A, data.B)
		if !reflect.DeepEqual(r, data.want) {
			t.Fatalf("html diff fail")
		}
	}
}
