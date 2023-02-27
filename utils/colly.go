package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func GetPage(page int) string {
	url := fmt.Sprintf("https://mirror.sgkoi.dev/?page=%d", page)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Accept-Language
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	req.Header.Set("User-Agent", "tModLoader Server Manager")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(body))

	return string(body)
}

func GetModList(page int) map[string][]string {
	body := GetPage(page)

	// 获取<tbody>标签中的内容
	re := regexp.MustCompile(`(?s)<tbody>(.*)</tbody>`)
	tbody := re.FindStringSubmatch(body)[1]

	// 获取<tr>标签中的内容
	re = regexp.MustCompile(`(?s)<tr>(.*?)</tr>`)
	trs := re.FindAllStringSubmatch(tbody, -1)

	var titles []string
	var images []string
	var sizes []string
	var authors []string
	var origins []string

	re = regexp.MustCompile(`(?s)<td>(.*?)</td>`)

	for _, tr := range trs {
		tds := re.FindAllStringSubmatch(tr[1], -1)

		title := regexp.MustCompile(`>([^<]+)<`).FindStringSubmatch(tds[0][1])[1]
		titles = append(titles, title)

		image := regexp.MustCompile(`<img[^>]*alt="([^"]*)"[^>]*>`).FindStringSubmatch(tds[0][1])

		if len(image) > 0 {
			images = append(images, image[1])
		} else {
			images = append(images, "")
		}

		size := regexp.MustCompile(`>([^<]+)<`).FindStringSubmatch(tds[1][1])[1]
		sizes = append(sizes, size)

		author := regexp.MustCompile(`title="([^"]*)"`).FindStringSubmatch(tds[2][1])[1]
		authors = append(authors, author)

		origin := regexp.MustCompile(`href="([^"]+)"`).FindStringSubmatch(tds[3][1])[1]
		origins = append(origins, origin)
	}

	var mods = make(map[string][]string)
	mods["titles"] = titles
	mods["images"] = images
	mods["sizes"] = sizes
	mods["authors"] = authors
	mods["origins"] = origins

	return mods
}
