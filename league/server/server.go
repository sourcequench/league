package main

import (
	"cloud.google.com/go/storage"
	"crypto/md5"
	"fmt"
	"github.com/gonum/stat"
	c "github.com/sourcequench/league/common"
	"github.com/sourcequench/league/npl"
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

var templates = template.Must(template.ParseFiles("report.html", "upload.html"))

func main() {
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/", Root)
	http.HandleFunc("/report", Report)
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

type Dump struct {
	// Initial skills map
	Iskills map[string]float64
	// 3/2/1 skills map
	Threeskills map[string]float64
	// 2/1/0 skills map
	Twoskills map[string]float64
	// +-2 skills map
	Skills map[string]float64
	// Overall mean, stddev for this strategy
	IDist []float64
	// Per player man, stddev for this strategy
	IPlayerDist     map[string][]float64
	ThreeDist       []float64
	ThreePlayerDist map[string][]float64
	TwoDist         []float64
	TwoPlayerDist   map[string][]float64
	Dist            []float64
	PlayerDist      map[string][]float64
}

func Report(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	fileName := "uploaded.csv"
	matches, _ := parser.Parse(fileName, ctx)
	w.Header().Set("Content-Type", "text/html")
	iskills := c.InitialSkill(matches)
	// 3/2/1
	threematches := c.UpdateMatches(matches, npl.ThreeTwoOne{})
	threeskills := c.FinalSkill(threematches)
	threediffs := c.PercentDiff(threematches)
	threemu, threesigma := stat.MeanStdDev(threediffs, nil)
	threeperuser := c.PerUserPercentDiff(threematches)
	threedist := c.PlayerNormal(threeperuser)

	// 2/1/0
	twomatches := c.UpdateMatches(matches, npl.TwoOneZero{})
	twoskills := c.FinalSkill(twomatches)
	twodiffs := c.PercentDiff(twomatches)
	twomu, twosigma := stat.MeanStdDev(twodiffs, nil)
	twoperuser := c.PerUserPercentDiff(twomatches)
	twodist := c.PlayerNormal(twoperuser)

	// +-2
	zmatches := c.UpdateMatches(matches, npl.Two{})
	skills := c.FinalSkill(zmatches)
	zdiffs := c.PercentDiff(zmatches)
	zmu, zsigma := stat.MeanStdDev(zdiffs, nil)
	zperuser := c.PerUserPercentDiff(zmatches)
	zdist := c.PlayerNormal(zperuser)

	d := Dump{
		Iskills:         iskills,
		Threeskills:     threeskills,
		Twoskills:       twoskills,
		Skills:          skills,
		ThreeDist:       []float64{threemu, threesigma},
		ThreePlayerDist: threedist,
		TwoDist:         []float64{twomu, twosigma},
		TwoPlayerDist:   twodist,
		Dist:            []float64{zmu, zsigma},
		PlayerDist:      zdist,
	}
	renderTemplate(w, "report", d)
}

func renderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
