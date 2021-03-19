package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

const HtmlTemplate = `
<html>
<head>
  <meta charset="UTF-8">
  <title>You lived {{ . }} days!</title>
  <style>
    * {
      margin: 0;
      padding: 0;
    }

    body {
      height: 100vh;
    }

    .container {
      display: flex;
      align-items: center;
      justify-content: center;
      text-align: center;
      height: 100%;
    }

    .column-box {
      flex-direction: column;
    }

    h1 {
      margin-bottom: 12px;
    }
  </style>
</head>

<body>
  <div class="container">
    <div class="column-box">
      <h1>あなたはこのサイトができてから {{ . }} 日間生きました とても偉い！</h1>
      <a href="https://twitter.com/intent/tweet?text=あなたはこのサイトができてから%20{{ . }}%20日間生きました%20とても偉い！" class="twitter-share-button" data-show-count="false" data-size="large"></a>
      <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
    </div>
  </div>
</body>
</html>
`

func main() {
	baseDay, _ := time.Parse("2006-01-02", "2021-03-19") // このコードを書いた日が 2021/3/19 なので

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentDay := time.Now()
		daySub := currentDay.Day() - baseDay.Day() + 1

		t := template.Must(template.New("").Parse(HtmlTemplate))

		err := t.Execute(w, daySub)

		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
