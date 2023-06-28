package utils

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
)

func GetPage(page int) (string, error) {
	url := fmt.Sprintf("https://mirror.sgkoi.dev/?page=%d", page)
	req, _ := http.NewRequest("GET", url, nil)

	// Accept-Language
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	req.Header.Set("User-Agent", "tModLoader Server Manager")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)

	if _, err := io.Copy(buf, resp.Body); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetModList(page int) (map[string][]string, error) {
	body, err := GetPage(page)
	if err != nil {
		return nil, err
	}

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

	return mods, nil
}

func IpAddress() string {
	ip := "localhost"
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		addrs, _ := net.InterfaceAddrs()
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
				}
			}
		}
	} else {
		body, _ := io.ReadAll(resp.Body)
		ip = string(body)
	}
	return ip
}
