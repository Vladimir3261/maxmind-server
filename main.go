package main

import (
	"os"
	"gopkg.in/ini.v1"
	"net/http"
	"io"
	"compress/gzip"
	"path/filepath"
	"path"
	"log"
	"github.com/oschwald/geoip2-golang"
	"net"
	"encoding/json"
)

var config *ini.Section

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	config = cfg.Section("")

	if len(os.Args) < 3 {
		log.Println("list or count subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		var dbFile string
		if os.Args[2] != "-" {
			dbFile = os.Args[2]
		} else {
			dbFile = update()
		}
		runServer(dbFile)
	case "update":
		update()
	default:
		log.Println("Invalid action")
		os.Exit(1)
	}
}

func update() string {
	dbUrl := config.Key("database_url").String()
	if dbUrl == "" {
		panic("No database url provided")
		os.Exit(1)
	}

	dbPath := config.Key("database_dir").String()
	if dbPath == "" {
		panic("No database_dir path in configuration")
		os.Exit(1)
	}

	downLoadFilename := path.Base(dbUrl)
	downLoadedFilePath := dbPath + "/" + downLoadFilename

	log.Printf("Downloading file from %s to %s \r\n", dbUrl, downLoadedFilePath)
	err := DownloadFile(downLoadedFilePath, dbUrl)
	if err != nil {
		panic(err)
	}

	log.Printf("File success downloaded and saved at %s \r\n", downLoadedFilePath)

	log.Printf("Unpacking downloaded file %s file to %s \r\n", downLoadedFilePath, dbPath)

	err, datFile := ungzip(downLoadedFilePath, dbPath)

	if err != nil {
		panic(err)
		os.Exit(1)
	}

	log.Printf("File %s success unpacked \r\n", datFile)

	err = os.Remove(downLoadedFilePath)

	if err != nil {
		log.Println("Warning! Cant remove downloaded filr")
	} else {
		log.Printf("Downloaded file %s has been removed \r\n", downLoadedFilePath)
	}

	log.Println("Trying to run server...")

	return datFile
}

func DownloadFile(path string, url string) error {
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func ungzip(source, target string) (error, string) {
	reader, err := os.Open(source)
	if err != nil {
		return err, ""
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err, ""
	}
	defer archive.Close()

	target = filepath.Join(target, archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		return err, ""
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err, target
}

func runServer(dbFile string) {
	httpServerAddress := config.Key("http_server").String()
	if httpServerAddress == "" {
		panic("No http_server config provided")
	}

	db, err := geoip2.Open(dbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	type Response struct {
		Status int
		Data   *geoip2.Country
		Error  string
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		ip := net.ParseIP(request.URL.Query().Get("ip"))
		encoder := json.NewEncoder(writer)

		record, err := db.Country(ip)

		r := Response{Status: http.StatusOK, Error: ""}

		if err != nil {
			r.Error = err.Error()
		} else {
			r.Data = record
		}

		encoder.Encode(r)
	})

	err = http.ListenAndServe(httpServerAddress, nil)
	if err != nil {
		log.Fatal("Failed on start http server: ", err)
		os.Exit(1)
	}
}
