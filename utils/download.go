package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	proxy = "http://localhost:7890"
	//proxy = ""
)

func DownloadFile(path string, file string, proxy string) error {
	// set http proxy
	if proxy != "" {
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			fmt.Println("Cannot parse proxy path", err)
			return err
		}
		http.DefaultTransport.(*http.Transport).Proxy = http.ProxyURL(proxyUrl)
	} else {
		http.DefaultTransport.(*http.Transport).Proxy = nil
	}
	client := http.Client{Transport: http.DefaultTransport}
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println("Cannot create request", err)
		return err
	}
	request.Header.Set("User-Agent", "tModLoader Server Manager")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Cannot get response", err)
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	size, err := strconv.Atoi(response.Header.Get("Content-Length"))
	if err != nil {
		return err
	}
	start := time.Now()
	var written int64
	// according file size creat buffer
	fmt.Println("Downloading", file, "size:", size/1024, "KB")
	buffer := make([]byte, 8*1024)
	var n int
	for {
		n, err = response.Body.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Cannot read response", err)
			return err
		}
		if n == 0 {
			break
		}
		written += int64(n)
		// print download progress
		out.Write(buffer[:n])
		elapsed := time.Since(start).Seconds()
		speed := float64(written) / elapsed / 1024 / 1024
		percent := float64(written) / float64(size) * 100
		fmt.Printf("\r%.2f%% %.2fMB/%.2fMB %.2fMB/s", percent, float64(written)/1024/1024, float64(size)/1024/1024, speed)
	}
	fmt.Println()
	fmt.Println("Download complete")
	return nil
}

func DownloadMod(name string) error {
	// name e.g. CalamityMod
	mirror := "https://mirror.sgkoi.dev"
	url := fmt.Sprintf("%s/tModLoader/download.php?Down=mod/%s.tmod", mirror, name)
	return DownloadFile(url, name+".tmod", proxy)
}
func DownloadTModLoader(name string) error {
	// name e.g. v2022.09.47.33
	release := "https://github.com/tModLoader/tModLoader/releases"
	url := fmt.Sprintf("%s/download/%s/tModLoader.zip", release, name)
	return DownloadFile(url, "./core/tModLoader.zip", proxy)
}
