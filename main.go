package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"video/handlers"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	store := NewMockStorage{}
	server := NewServerApi(&store, ":3000")
	server.Run()
	
	runDefaultServer()
}

func runDefaultServer() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=backend_new sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select id, name from tasks")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type count struct {
		count int64
	}
	var cc int;
	// c := count{}
	for rows.Next() {
		err := rows.Scan(&cc)
		if err != nil {
			panic(err)
		}
		fmt.Println(cc)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(666, err)
	}


	// fmt.Println(db)

	server := http.NewServeMux()
	
	// server.Handle("/", http.FileServer(http.Dir(".")))

	// server.HandleFunc("/video", handleVideo)
	// server.HandleFunc("/post", handlePost)
	// server.HandleFunc("/upload", handleUpload)
	server.HandleFunc("/create-user", handlers.CreateUser())

	http.ListenAndServe(":3000", server)
}

func readCommandLine() {
	buf := make([]byte, 3)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		str := string(buf[:n])

		fmt.Println(str)
		fmt.Println(str == "help\n")

		if str == "q" {
			break
		}

		if str == "help" {
			fmt.Println("help 28/11")
		}
	}
	fmt.Println("after")
}

type body struct {
	Some string
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	b :=body{}
	d := json.NewDecoder(r.Body)
	err := d.Decode(&b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", b)
	w.Write([]byte("psi"))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	name := r.Header["File-Name"][0]
	if name == "" {
		panic("no name header")
	}
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
			panic(err)
	}

	fmt.Println(len(body))

	f.Write(body)
}

func handleVideo(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("file.mov")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			panic(err)
		}

		ran := r.Header.Get("range")
		if ran == ""{
			panic("no range header")
		}

		size := stat.Size()
		reg := regexp.MustCompile(`\D`)

		start, err := strconv.Atoi(reg.ReplaceAllString(ran, ""))
		if err != nil {
			panic(err)
		}

		step := 1024 * 100
		end := int(math.Min(float64(start + step), float64(size - 1)))

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start,end, size))
		w.Header().Add("Accept-Ranges", "bytes")
		w.Header().Add("Content-Type", "video/mp4")
		w.Header().Add("Content-Length", fmt.Sprintf("%d", end - start + 1))
		w.WriteHeader(206)

		_, seekErr := f.Seek(int64(start), io.SeekStart)
		if seekErr != nil {
			panic(seekErr)
		}

		_, copyErr := io.CopyN(w, f, int64(end - start + 1))
		if copyErr != nil {
			panic(copyErr)
		}
}
