package main

import "TSM-Server/utils"

func main() {
	//r := server.NewRouter()
	//r.Run(":9000")
	//utils.DownloadTModLoader("v2022.09.47.33")
	//utils.DownloadTModLoader("v2022.09.47.33")
	utils.Unzip("./core/tModLoader.zip", "./core/tModLoader")
}
