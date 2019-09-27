package main

import (
	"cloud.google.com/go/storage"
	"crypto/md5"
	"fmt"
	c "github.com/sourcequench/league/common"
	"github.com/sourcequench/league/parser"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var templates = template.Must(template.ParseFiles("upload.html"))

func main() {
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/root", Root)
	appengine.Main()
	//	http.ListenAndServe(":9999", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	fileName := "uploaded.csv"
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		f, _, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		data, err := ioutil.ReadAll(f)

		bucketName, err := file.DefaultBucketName(ctx)
		if err != nil {
			log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
			return
		}

		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Errorf(ctx, "failed to create client: %v", err)
			return
		}
		defer client.Close()

		bucketHandle := client.Bucket(bucketName)

		wc := bucketHandle.Object(fileName).NewWriter(ctx)
		defer wc.Close()

		wc.ContentType = "text/plain"

		if _, err := wc.Write(data); err != nil {
			log.Errorf(ctx, "save: unable to open file from bucket %q, file %q: %v", bucketName, fileName, err)
			return
		}
		matches, errors := parser.Parse(fileName, ctx)
		w.Header().Set("Content-Type", "text/html")
		head := `
<html>
<head>
       <title>Uploaded</title>
</head>
<body>
<p>
`

		foot := `
</p>
</body>
</html>
`

		fmt.Fprintf(w, head)
		fmt.Fprintf(w, "wrote %d matches with %d errors<br/>", len(matches), len(errors))
		for _, e := range errors {
			fmt.Fprintf(w, "%v<br/>", e)
		}
		fmt.Fprintf(w, foot)
	}
}

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	page := `
<html>
<head>
       <title>Hello</title>
</head>
<body>
<p>
Hello World
</p>
</body>
</html>
`
	fmt.Fprintf(w, page)
}

func renderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
