package main

import (
	"cloud.google.com/go/logging"
	"cloud.google.com/go/storage"
	"crypto/md5"
	"fmt"
	"github.com/gonum/stat"
	c "github.com/sourcequench/league/common"
	"github.com/sourcequench/league/npl"
	"github.com/sourcequench/league/parser"
	"google.golang.org/appengine/file"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var templates = template.Must(template.ParseFiles("report.html", "upload.html", "material.html"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/report", Report)
	http.HandleFunc("/material", Material)
	log.Printf("Listening on port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// Create plumbing for logging and context.
	ctx := r.Context()
	client, err := logging.NewClient(ctx, "league-253800")
	if err != nil {
		return
	}
	defer client.Close()
	log := client.Logger("server-log")

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
			log.Log(logging.Entry{Payload: fmt.Sprintf("failed to get default GCS bucket name: %v", err)})
			return
		}

		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Log(logging.Entry{Payload: fmt.Sprintf("failed to create client: %v", err)})
			return
		}
		defer client.Close()

		bucketHandle := client.Bucket(bucketName)

		wc := bucketHandle.Object(fileName).NewWriter(ctx)
		defer wc.Close()

		wc.ContentType = "text/plain"

		if _, err := wc.Write(data); err != nil {
			log.Log(logging.Entry{Payload: fmt.Sprintf("save: unable to open file from bucket %q, file %q: %v", bucketName, fileName, err)})
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
	// Initial skills for unchanged matches
	Iskills     map[string]float64
	IPlayerDist map[string][]float64
	// Overall mean, stddev for this strategy
	IDist []float64
	// Final skills for unchanged matches
	Fskills map[string]float64
	// 3/2/1 skills map
	Threeskills     map[string]float64
	ThreeDist       []float64
	ThreePlayerDist map[string][]float64
	ThreeWins       map[string]float64
	// 2/1/0 skills map
	Twoskills     map[string]float64
	TwoDist       []float64
	TwoPlayerDist map[string][]float64
	TwoWins       map[string]float64
	// +-1 skills map
	Oneskills     map[string]float64
	OneDist       []float64
	OnePlayerDist map[string][]float64
	OneWins       map[string]float64
	// No change skills map
	Noskills     map[string]float64
	NoDist       []float64
	NoPlayerDist map[string][]float64
	NoWins       map[string]float64
	// +-2 skills map
	Skills     map[string]float64
	Dist       []float64
	PlayerDist map[string][]float64
	Wins       map[string]float64
}

func Material(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client, err := logging.NewClient(ctx, "league-253800")
	if err != nil {
		return
	}
	defer client.Close()
	log := client.Logger("server-log")

	fileName := "uploaded.csv"
	matches, e := parser.Parse(fileName, ctx)
	if len(matches) == 0 {
		fmt.Println("shit - didn't work")
	}
	if len(e) != 0 {
		log.Log(logging.Entry{Payload: fmt.Sprintf("failed to parse: %v", err)})
		fmt.Printf("shit - errors: %v", e)
	}
	w.Header().Set("Content-Type", "text/html")
	iskills := c.InitialSkill(matches)
	fskills := c.FinalSkill(matches)
	// 3/2/1
	threematches := c.UpdateMatches(matches, npl.ThreeTwoOne{})
	threeskills := c.FinalSkill(threematches)
	threediffs := c.PercentDiff(threematches)
	threemu, threesigma := stat.MeanStdDev(threediffs, nil)
	threeperuser := c.PerUserPercentDiff(threematches)
	threedist := c.PlayerNormal(threeperuser)
	threewins := c.WinRecord(threematches)

	// 2/1/0
	twomatches := c.UpdateMatches(matches, npl.TwoOneZero{})
	twoskills := c.FinalSkill(twomatches)
	twodiffs := c.PercentDiff(twomatches)
	twomu, twosigma := stat.MeanStdDev(twodiffs, nil)
	twoperuser := c.PerUserPercentDiff(twomatches)
	twodist := c.PlayerNormal(twoperuser)
	twowins := c.WinRecord(twomatches)

	// +-2
	zmatches := c.UpdateMatches(matches, npl.Two{})
	skills := c.FinalSkill(zmatches)
	zdiffs := c.PercentDiff(zmatches)
	zmu, zsigma := stat.MeanStdDev(zdiffs, nil)
	zperuser := c.PerUserPercentDiff(zmatches)
	zdist := c.PlayerNormal(zperuser)
	zwins := c.WinRecord(zmatches)

	// +-1
	onematches := c.UpdateMatches(matches, npl.One{})
	oneskills := c.FinalSkill(onematches)
	onediffs := c.PercentDiff(onematches)
	onemu, onesigma := stat.MeanStdDev(onediffs, nil)
	oneperuser := c.PerUserPercentDiff(onematches)
	onedist := c.PlayerNormal(oneperuser)
	onewins := c.WinRecord(onematches)

	// No change
	nomatches := c.UpdateMatches(matches, npl.NoChange{})
	noskills := c.FinalSkill(nomatches)
	nodiffs := c.PercentDiff(nomatches)
	nomu, nosigma := stat.MeanStdDev(nodiffs, nil)
	noperuser := c.PerUserPercentDiff(nomatches)
	nodist := c.PlayerNormal(noperuser)
	nowins := c.WinRecord(nomatches)

	d := Dump{
		Iskills:         iskills,
		Fskills:         fskills,
		Threeskills:     threeskills,
		ThreeDist:       []float64{threemu, threesigma},
		ThreePlayerDist: threedist,
		ThreeWins:       threewins,
		Twoskills:       twoskills,
		TwoWins:         twowins,
		TwoDist:         []float64{twomu, twosigma},
		TwoPlayerDist:   twodist,
		Oneskills:       oneskills,
		OneDist:         []float64{onemu, onesigma},
		OnePlayerDist:   onedist,
		OneWins:         onewins,
		Noskills:        noskills,
		NoDist:          []float64{nomu, nosigma},
		NoPlayerDist:    nodist,
		NoWins:          nowins,
		Skills:          skills,
		Dist:            []float64{zmu, zsigma},
		PlayerDist:      zdist,
		Wins:            zwins,
	}
	renderTemplate(w, "material", d)
}
func Report(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client, err := logging.NewClient(ctx, "league-253800")
	if err != nil {
		return
	}
	defer client.Close()
	log := client.Logger("server-log")

	fileName := "uploaded.csv"
	matches, e := parser.Parse(fileName, ctx)
	if len(matches) == 0 {
		fmt.Println("shit - didn't work")
	}
	if len(e) != 0 {
		log.Log(logging.Entry{Payload: fmt.Sprintf("failed to parse: %v", err)})
		fmt.Printf("shit - errors: %v", e)
	}
	w.Header().Set("Content-Type", "text/html")
	iskills := c.InitialSkill(matches)
	fskills := c.FinalSkill(matches)
	// 3/2/1
	threematches := c.UpdateMatches(matches, npl.ThreeTwoOne{})
	threeskills := c.FinalSkill(threematches)
	threediffs := c.PercentDiff(threematches)
	threemu, threesigma := stat.MeanStdDev(threediffs, nil)
	threeperuser := c.PerUserPercentDiff(threematches)
	threedist := c.PlayerNormal(threeperuser)
	threewins := c.WinRecord(threematches)

	// 2/1/0
	twomatches := c.UpdateMatches(matches, npl.TwoOneZero{})
	twoskills := c.FinalSkill(twomatches)
	twodiffs := c.PercentDiff(twomatches)
	twomu, twosigma := stat.MeanStdDev(twodiffs, nil)
	twoperuser := c.PerUserPercentDiff(twomatches)
	twodist := c.PlayerNormal(twoperuser)
	twowins := c.WinRecord(twomatches)

	// +-2
	zmatches := c.UpdateMatches(matches, npl.Two{})
	skills := c.FinalSkill(zmatches)
	zdiffs := c.PercentDiff(zmatches)
	zmu, zsigma := stat.MeanStdDev(zdiffs, nil)
	zperuser := c.PerUserPercentDiff(zmatches)
	zdist := c.PlayerNormal(zperuser)
	zwins := c.WinRecord(zmatches)

	// +-1
	onematches := c.UpdateMatches(matches, npl.One{})
	oneskills := c.FinalSkill(onematches)
	onediffs := c.PercentDiff(onematches)
	onemu, onesigma := stat.MeanStdDev(onediffs, nil)
	oneperuser := c.PerUserPercentDiff(onematches)
	onedist := c.PlayerNormal(oneperuser)
	onewins := c.WinRecord(onematches)

	// No change
	nomatches := c.UpdateMatches(matches, npl.NoChange{})
	noskills := c.FinalSkill(nomatches)
	nodiffs := c.PercentDiff(nomatches)
	nomu, nosigma := stat.MeanStdDev(nodiffs, nil)
	noperuser := c.PerUserPercentDiff(nomatches)
	nodist := c.PlayerNormal(noperuser)
	nowins := c.WinRecord(nomatches)

	d := Dump{
		Iskills:         iskills,
		Fskills:         fskills,
		Threeskills:     threeskills,
		ThreeDist:       []float64{threemu, threesigma},
		ThreePlayerDist: threedist,
		ThreeWins:       threewins,
		Twoskills:       twoskills,
		TwoWins:         twowins,
		TwoDist:         []float64{twomu, twosigma},
		TwoPlayerDist:   twodist,
		Oneskills:       oneskills,
		OneDist:         []float64{onemu, onesigma},
		OnePlayerDist:   onedist,
		OneWins:         onewins,
		Noskills:        noskills,
		NoDist:          []float64{nomu, nosigma},
		NoPlayerDist:    nodist,
		NoWins:          nowins,
		Skills:          skills,
		Dist:            []float64{zmu, zsigma},
		PlayerDist:      zdist,
		Wins:            zwins,
	}
	renderTemplate(w, "report", d)
}

func renderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
