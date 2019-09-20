package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":9999", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	page := `
<html>
<head>
	<title>Lines</title>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.8.0/Chart.min.js"></script>
	<style>
	canvas{
		-moz-user-select: none;
		-webkit-user-select: none;
		-ms-user-select: none;
	}
	</style>
</head>

<body>
<canvas id="line-chart" width="800" height="450"></canvas>
<script>
new Chart(document.getElementById("line-chart"), {
  type: 'line',
  data: {
	labels: ["2018-01-09","2018-01-16","2018-01-23","2018-01-30","2018-02-06","2018-02-13","2018-02-20","2018-02-27","2018-03-06","2018-03-13","2018-03-20","2018-03-27","2018-04-03","2018-04-10","2018-04-17","2018-04-24","2018-05-08","2018-05-1"],
    	datasets: [{ 
		data: [45, null, 49, 45, 41, null, null, 41, 44, null, null, 44, null, null, 41, 44, 48, 42 ],
		label: "KurtR",
		borderColor: '#E49988',
		fill:false,
                spanGaps:true
	}, {
		data: [102, 104, 102, 104, 106, 104, 102, 104, 106, 104, 102, null, 100, null, null, 102, 98, 100 ],
		label: "SoyC",
		borderColor: '#E4C788',
		fill:false,
                spanGaps:true
	}, {
		data: [null, 106, 104, 102, 100, 102, 100, 98, 100, 102, null, 100, 102, 104, 106, 108, 106, null ],
		label: "BryanH",
		borderColor: '#D3E488',
		fill:false,
                spanGaps:true
	}, {
		data: [86, 84, 82, 84, 86, 88, 86, null, 88, 86, 88, 90, 92, 90, 88, null, null, 86 ],
		label: "BillA",
		borderColor: '#A5E488',
		fill:false,
                spanGaps:true
	}, {
		data: [114, 112, null, 114, 112, 110, 112, 110, 108, 110, 112, 110, 108, 110, 112, 110, null, null ],
		label: "KenB",
		borderColor: '#88E499',
		fill:false,
                spanGaps:true
	}, {
		data: [50, 52, 50, 48, 46, 44, null, 46, 44, 42, 40, 38, 36, 34, null, null, 30, null ],
		label: "DonP",
		borderColor: '#88E4C7',
		fill:false,
                spanGaps:true
	}, {
		data: [46, 48, 46, 42, 40, null, null, 40, 42, null, null, 46, 44, 44, 41, 46, null, null ],
		label: "DaveR",
		borderColor: '#88AFE4',
		fill:false,
                spanGaps:true
	}, {
		data: [59, 61, 63, 65, 67, 69, 67, 69, 67, 69, 71, 69, 71, 73, 71, null, null, null ],
		label: "BobB",
		borderColor: '#8F88E4',
		fill:false,
                spanGaps:true
	}, {
		data: [97, 95, 97, 95, 97, 99, 97, 99, 97, 99, 101, null, 103, 101, 99, 97, null, null ],
		label: "RyanS",
		borderColor: '#BD88E4',
		fill:false,
                spanGaps:true
	}, {
		data: [null, null, 92, 94, 92, 94, 92, null, 94, null, 94, 94, 96, 94, 96, 98, 96, null ],
		label: "JKB",
		borderColor: '#E488DD',
		fill:false,
                spanGaps:true
	}, {
		data: [75, 73, null, 75, 77, 75, 73, 71, 73, null, 69, 71, 69, 71, 73, 73, null, null ],
		label: "LouisC",
		borderColor: '#E488AF',
		fill:false,
                spanGaps:true
	}, {
		data: [109, 107, null, 109, 111, 109, 111, 109, 111, null, 109, 111, 113, 117, 119, 115, 113, null ],
		label: "JohnL",
		borderColor: '#E48F88',
		fill:false,
                spanGaps:true
	}, {
		data: [90, 92, 90, 92, 90, 92, 94, null, 96, 94, 92, 90, null, 92, 94, null, 94, null ],
		label: "JeffD",
		borderColor: '#E48888',
		fill:false,
                spanGaps:true
	}, {
		data: [24, 22, 24, null, 22, 24, 26, null, 24, null, null, 24, 26, 28, 26, null, 28, 30 ],
		label: "KevinH",
		borderColor: '#88E488',
		fill:false,
                spanGaps:true
	}, {
		data: [96, 98, 100, 102, 104, 102, null, 104, 102, 104, 102, 104, 102, 100, 102, 100, null, null ],
		label: "WinstonW",
		borderColor: '#E4E488',
		fill:false,
                spanGaps:true
	}, {
		data: [84, 82, null, 84, 86, 84, 86, 84, 82, 84, 86, 84, 82, 80, 78, 76, null, null ],
		label: "JimO",
		borderColor: '#B6E488',
		fill:false,
                spanGaps:true
      }
    ]
  },
  options: {
    title: {
      display: true,
      text: 'Spring 2018',
    }
  }
});
</script>
</body>

</html>
	`
	fmt.Fprintf(w, page)
}
