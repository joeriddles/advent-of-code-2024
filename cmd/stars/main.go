package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"

	"github.com/anaskhan96/soup"
	"github.com/joeriddles/advent-of-code-2024/pkg/util"
)

var AOC_URL = &url.URL{
	Scheme: "https",
	Host:   "adventofcode.com",
}

var STARS_PATTERN = regexp.MustCompile(`Day (\d+), (one|two) stars?`)

func main() {
	// Get stars from AoC
	jar, err := cookiejar.New(nil)
	if err != nil {
		util.LogErr(err)
		os.Exit(1)
	}

	session, exists := os.LookupEnv("AOC_SESSION")
	if !exists {
		util.LogErrf("set AOC_SESSION environment variable\n")
	}

	jar.SetCookies(AOC_URL, []*http.Cookie{
		{
			Name:  "SESSION",
			Value: session,
		},
	})

	c := http.Client{Jar: jar}
	resp, err := c.Get(fmt.Sprintf("%v://%v", AOC_URL.Scheme, AOC_URL.Host))
	if err != nil {
		util.LogErr(err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		util.LogErrf("error status code: %v\n", resp.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		util.LogErrf("failed to read body")
		os.Exit(1)
	}

	doc := soup.HTMLParse(string(body))
	labels := []string{}
	for _, link := range doc.FindAll("a") {
		dayEl := link.Find("span", "class", "calendar-day")
		if dayEl.Error != nil {
			continue
		}
		day := dayEl.Text()

		part1Complete := link.Find("span", "class", "calendar-mark-complete").Error == nil
		part2Complete := link.Find("span", "class", "calendar-mark-verycomplete").Error == nil

		stars := ""
		if part2Complete {
			stars = "⭐️⭐️"
		} else if part1Complete {
			stars = "⭐️"
		}

		label := fmt.Sprintf("Day %v: %v  \n", day, stars)
		util.LogSuccessf(label)
		labels = append(labels, label)
	}

	// Write star labels to README
	readme, err := os.ReadFile("README.md")
	if err != nil {
		util.LogErrf("failed to read README: %v\n", err.Error())
		os.Exit(1)
	}

	i := regexp.MustCompile("---\n").FindIndex(readme)
	if i == nil {
		util.LogErrf("failed to find `---` in README")
		os.Exit(1)
	}

	text := readme[:i[1]]
	for _, label := range labels {
		text = append(text, []byte(label)...)
	}

	err = os.WriteFile("README.md", text, 0644)
	if err != nil {
		util.LogErrf("failed to save README: %v\n", err.Error())
		os.Exit(1)
	}
}