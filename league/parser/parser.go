package parser

import (
	"bufio"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	c "github.com/sourcequench/league/common"
	pb "github.com/sourcequench/league/proto"
	//	"google.golang.org/appengine/file"
	"io"
	"os"
	"strconv"
	"strings"
)

func CloudOpen(f string, ctx context.Context) (io.Reader, error) {
	// Create a cloud storage client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating storage client: %v", err)
	}
	defer client.Close()

	// Get the cloud storage bucket name.
	//	bucketName, err := file.DefaultBucketName(ctx)
	bucketName := "league-253800.appspot.com"
	// Set the bucket of the client to be the default bucket.
	bucket := client.Bucket(bucketName)

	// Create a Reader for the file.
	rc, err := bucket.Object(f).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to open file from bucket %q, file %q: %v", bucketName, f, err)
	}
	return rc, nil
}

func Parse(f string, ctx context.Context) ([]c.Match, []error) {
	var matches []c.Match
	var errors []error

	var rdr io.Reader
	// Open from a local file in string if we didn't get context.
	if ctx == nil {
		fh, err := os.Open(f)
		if err != nil {
			return nil, []error{fmt.Errorf("local file open error %v", err)}
		}
		rdr = fh
	} else {
		fh, err := CloudOpen(f, ctx)
		if err != nil {
			return nil, []error{fmt.Errorf("unable to open cloud file %q: %v", f, err)}
		}
		rdr = fh
	}

	scanner := bufio.NewScanner(rdr)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		p1skill, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			errors = append(errors, fmt.Errorf("error getting p1skill: %v\n", fields))
			continue
		}
		p1needs, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			errors = append(errors, fmt.Errorf("error getting p1needs: %v\n", fields))
			continue
		}
		p1got, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			errors = append(errors, fmt.Errorf("error getting p1got: %v\n", fields))
			continue
		}
		p2skill, err := strconv.ParseFloat(fields[6], 64)
		if err != nil {
			errors = append(errors, fmt.Errorf("error getting p2skill: %v\n", fields))
			continue
		}
		p2needs, err := strconv.ParseFloat(fields[7], 64)
		if err != nil {
			errors = append(errors, fmt.Errorf("error getting p2needs: %v\n", fields))
			continue
		}
		p2got, err := strconv.ParseFloat(fields[8], 64)
		if err != nil {
			errors = append(errors, fmt.Errorf("error getting p2got: %v\n", fields))
			continue
		}

		p1name := fields[0]
		p2name := fields[5]

		m := c.Match{
			P1name:  p1name,
			P2name:  p2name,
			Date:    fields[10],
			P1needs: p1needs,
			P1got:   p1got,
			P2needs: p2needs,
			P2got:   p2got,
			P1skill: p1skill,
			P2skill: p2skill,
		}
		matches = append(matches, m)
	}
	if err := scanner.Err(); err != nil {
		errors = append(errors, fmt.Errorf("scanner error: %v\n", err))
	}

	return matches, errors
}

func ProtoOut(matches []c.Match) error {
	season := &pb.Season{}
	for _, match := range matches {
		m := pb.Match{
			P1Name:  match.P1name,
			P2Name:  match.P2name,
			P1Needs: match.P1needs,
			P2Needs: match.P2needs,
			P1Got:   match.P1got,
			P2Got:   match.P2got,
			P1Skill: match.P1skill,
			P2Skill: match.P2skill,
			Date:    match.Date,
		}
		season.Matches = append(season.Matches, &m)
	}

	marshaler := &jsonpb.Marshaler{Indent: "  ", EmitDefaults: false, EnumsAsInts: false, OrigName: true}

	f, err := os.Create("matches.json")
	if err != nil {
		return fmt.Errorf("Failed to create file:", err)
	}
	w := bufio.NewWriter(f)

	err = marshaler.Marshal(w, season)
	if err != nil {
		return fmt.Errorf("Failed to write season:", err)
	}
	return nil
}
