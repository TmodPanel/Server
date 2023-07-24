package utils

import (
	"TSM-Server/cmd/setting"
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func setProxy(proxy string) {
	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)
		http.DefaultTransport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	} else {
		http.DefaultTransport.(*http.Transport).Proxy = nil
	}
}

func downloadFile(proxy, path, file string) error {
	setProxy(proxy)

	resp, err := http.Get(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	var written int64

	fmt.Println("Downloading", file, "size:", size/1024, "KB")

	buffer := make([]byte, 8*1024)
	start := time.Now()
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		written += int64(n)
		out.Write(buffer[:n])

		elapsed := time.Since(start).Seconds()
		speed := float64(written) / elapsed / 1024 / 1024
		percent := float64(written) / float64(size) * 100
		fmt.Printf("\r%.2f%% %.2fMB/%.2fMB %.2fMB/s", percent, float64(written)/1024/1024, float64(size)/1024/1024, speed)
	}

	fmt.Println("\nDownload complete")
	return nil
}

func DownloadMod(name string) error {
	// name e.g. CalamityMod
	mirror := "https://mirror.sgkoi.dev"
	url := fmt.Sprintf("%s/tModLoader/download.php?Down=mod/%s.tmod", mirror, name)
	return downloadFile(setting.Proxy, url, filepath.Join(setting.ModPath, name+".tmod"))
}
func DownloadTModLoader(name string) error {
	// name e.g. v2022.09.47.33
	release := "https://github.com/tModLoader/tModLoader/releases"
	url := fmt.Sprintf("%s/download/%s/tModLoader.zip", release, name)
	return downloadFile(setting.Proxy, url, filepath.Join(setting.ModPath, "tModLoader.zip"))
}

func CopyFile(source, dest string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func WriteToFile(filename, data string, perm os.FileMode) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func read(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var lines []string
	for {
		line, _ := r.ReadString('\n')
		if line == "" {
			break
		}
		lines = append(lines, strings.Trim(line, "\r\n"))
	}
	return lines, nil
}
func write(file string, lines []string, filter string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, line := range lines {
		if !strings.Contains(line, filter) {
			fmt.Fprintf(f, "%s\n", line)
		}
	}
	return nil
}

func ReadUserList() ([]string, error) {
	lines, err := read("./config/userList.txt")
	if err != nil {
		return nil, err
	}
	return lines, err
}
func Unzip(zipFilePath string, destDirectory string) error {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		fmt.Println("open zip file failed:", err)
		return err
	}
	defer r.Close()

	// 遍历zip文件中的所有文件
	for _, f := range r.File {
		// 打开zip文件中的文件
		rc, err := f.Open()
		if err != nil {
			fmt.Println("open file in zip failed:", err)
			return err
		}
		// 创建目标文件
		path := filepath.Join(destDirectory, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
			continue
		} else {
			os.MkdirAll(filepath.Dir(path), os.ModePerm)
		}
		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			fmt.Println("create file failed:", err)
			return err
		}

		// 将zip文件中的内容写入目标文件
		_, err = io.Copy(targetFile, rc)
		if err != nil {
			fmt.Println("write file failed:", err)
			return err
		}
	}

	fmt.Println("Zip文件解压缩完成！")
	return nil
}
func RemoveFile(filepath string) error {
	err := os.Remove(filepath)
	return err
}
func RemoveDir(dir string) error {
	err := os.RemoveAll(dir)
	return err
}
