package main
 
import (
 //   "crypto/md5"
    "fmt"
    "html/template"
    "io"
    "net/http"
    "os"
//    "strconv"
 //   "time"
)
 
func upload(w http.ResponseWriter, r *http.Request) {
 
    if r.Method == "GET" {
        // GET
        t, _ := template.ParseFiles("upload.gtpl")
 
        t.Execute(w, nil)
 
    } else if r.Method == "POST" {
        // Post
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
 
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./assets/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
 
        io.Copy(f, file)
 
    } else {
        fmt.Println("Unknown HTTP " + r.Method + "  Method")
    }
}

func index_handler(w http.ResponseWriter, r *http.Request) {
    // MAIN SECTION HTML CODE
    fmt.Fprintf(w, "<a href=http://127.0.0.1/upload>go to uplaod</a>")
    fmt.Fprintf(w, "<a href=http://127.0.0.1/about>go to about</a>")
    fmt.Fprintf(w, "<title>Go</title>")
    fmt.Fprintf(w, "<img src='assets/beach.jpg' alt='gopher' style='width:200px;height:200px;'>")
    fmt.Fprintf(w, "<img src='assets/Capture.JPG' alt='gopher' style='width:200px;height:200px;'>")
    fmt.Fprintf(w, "<img src='assets/nin_nyc.jpg' alt='gopher' style='width:200px;height:200px;'>")
    fmt.Fprintf(w, "<img src='assets/balatero_fam.jpg' alt='gopher' style='width:200px;height:200px;'>")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
    // ABOUT SECTION HTML CODE
    fmt.Fprintf(w, "<title>Go/about/</title>")
    fmt.Fprintf(w, "Expert web design by JT Skrivanek")
}

 
func main() {
    http.HandleFunc("/", index_handler)
    http.HandleFunc("/about/", about_handler)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
    http.HandleFunc("/upload", upload)
    http.ListenAndServe(":80", nil) // setting listening port
}