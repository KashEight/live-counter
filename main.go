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
  <title>生きてて偉い！</title>
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
      <h1>あなたはこのサイトができてから <span id="day">{{ .Day }}</span> 日 <span id="hour">{{ .Hour }}</span> 時 <span id="minute">{{ .Minute }}</span> 分 <span id="second">{{ .Second }}</span> 秒間生きました とても偉い！</h1>
      <a href="https://twitter.com/intent/tweet?text=あなたはこのサイトができてから%20{{ .Day }}%20日%20{{ .Hour }}%20時%20{{ .Minute }}%20分%20{{ .Second }}%20秒間生きました%20とても偉い！" class="twitter-share-button" data-show-count="false" data-size="large" data-lang="ja"></a>
      <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
    </div>
  </div>

  <script>
    const secEl = document.getElementById("second")
    const minEl = document.getElementById("minute")
    const hourEl = document.getElementById("hour")
    const dayEl = document.getElementById("day")

    setInterval(() => {
      let sec = parseInt(secEl.innerText)
      let min = parseInt(minEl.innerText)
      let hour = parseInt(hourEl.innerText)
      let day = parseInt(dayEl.innerText)

      sec++

      if (sec === 60) {
        min++
        sec = 0
      }

      if (min === 60) {
        hour++
        min = 0
      }

      if (hour === 24) {
        day++
        hour = 0
      }

      secEl.innerText = sec
      minEl.innerText = min
      hourEl.innerText = hour
      dayEl.innerText = day
    }, 1000)
  </script>
</body>
</html>
`

type Time struct {
	Day    int
	Hour   int
	Minute int
	Second int
}

func main() {
	baseDay, _ := time.Parse(time.RFC3339, "2021-03-19T09:00:00+09:00") // このコードを書いた日が 2021/3/19 なので

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentDay := time.Now()
		hourSub := currentDay.Hour() - baseDay.Hour()

		if hourSub < 0 {
			hourSub = 24 - hourSub
		}

		t := Time{
			Day:    currentDay.Day() - baseDay.Day(),
			Hour:   hourSub,
			Minute: currentDay.Minute() - baseDay.Minute(),
			Second: currentDay.Second() - baseDay.Second(),
		}

		tmpl := template.Must(template.New("").Parse(HtmlTemplate))

		err := tmpl.Execute(w, t)

		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
