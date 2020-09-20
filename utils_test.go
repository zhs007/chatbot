package chatbot

import (
	"testing"
)

func Test_SplitCommandString(t *testing.T) {
	type strarr struct {
		src  string
		dest []string
	}

	arrok := []strarr{
		strarr{src: "getdtdata -m gamedatareport -s 2019-04-17 -e 2019-04-17", dest: []string{"getdtdata", "-m", "gamedatareport", "-s", "2019-04-17", "-e", "2019-04-17"}},
		strarr{src: "getdtdata", dest: []string{"getdtdata"}},
		strarr{src: "   getdtdata", dest: []string{"getdtdata"}},
		strarr{src: "   getdtdata   ", dest: []string{"getdtdata"}},
		strarr{src: "getdtdata   ", dest: []string{"getdtdata"}},
		strarr{src: " getdtdata  -m  gamedatareport  -s    2019-04-17   -e 2019-04-17   ", dest: []string{"getdtdata", "-m", "gamedatareport", "-s", "2019-04-17", "-e", "2019-04-17"}},
		strarr{src: " getdtdata  -m  gamedatareport  -s    \"2019-04-17\"   -e 2019-04-17   ", dest: []string{"getdtdata", "-m", "gamedatareport", "-s", "2019-04-17", "-e", "2019-04-17"}},
		strarr{src: " getdtdata  -m  gamedatareport  -s    \"2019-04-17   \"   -e 2019-04-17   ", dest: []string{"getdtdata", "-m", "gamedatareport", "-s", "2019-04-17   ", "-e", "2019-04-17"}},
		strarr{src: "duckling -l zh_CN -t \"这个星期四要去看复联\"", dest: []string{"duckling", "-l", "zh_CN", "-t", "这个星期四要去看复联"}},
	}

	for _, v := range arrok {
		cr := SplitCommandString(v.src)
		if len(cr) != len(v.dest) {
			t.Fatalf("Test_SplitCommandString Err src:%v dest:%v ret:%v", v.src, v.dest, cr)
		}

		for i, sv := range cr {
			if sv != v.dest[i] {
				t.Fatalf("Test_SplitCommandString Err src:%v dest:%v ret:%v", v.src, v.dest, cr)
			}
		}
	}

	t.Logf("Test_SplitCommandString OK")
}

func Test_SplitMultiCommandString(t *testing.T) {
	type strarr struct {
		src  string
		dest []string
	}

	arrok := []strarr{
		strarr{
			src:  "reqtask techinasia -m jobs\n\r\n\n\r\n",
			dest: []string{"reqtask techinasia -m jobs"},
		},
		strarr{
			src:  "reqtask techinasia -m jobs\nreqtask techinasia -m jobs\r\n\n\r\nreqtask techinasia -m jobs",
			dest: []string{"reqtask techinasia -m jobs", "reqtask techinasia -m jobs", "reqtask techinasia -m jobs"},
		},
		strarr{
			src:  "reqtask techinasia -m \"jobs\"\nreqtask techinasia -m jobs\r\n\n\r\nreqtask techinasia -m jobs",
			dest: []string{"reqtask techinasia -m \"jobs\"", "reqtask techinasia -m jobs", "reqtask techinasia -m jobs"},
		},
		strarr{
			src:  "reqtask techinasia -m \"jobs\nreqtask techinasia -m jobs\r\n\"\n\r\nreqtask techinasia -m jobs",
			dest: []string{"reqtask techinasia -m \"jobs\nreqtask techinasia -m jobs\r\n\"", "reqtask techinasia -m jobs"},
		},
	}

	for _, v := range arrok {
		cr := SplitMultiCommandString(v.src)
		if len(cr) != len(v.dest) {
			t.Fatalf("Test_SplitMultiCommandString Err src:%v dest:%v ret:%v", v.src, v.dest, cr)
		}

		for i, sv := range cr {
			if sv != v.dest[i] {
				t.Fatalf("Test_SplitMultiCommandString Err src:%v dest:%v ret:%v", v.src, v.dest, cr)
			}
		}
	}

	t.Logf("Test_SplitMultiCommandString OK")
}
