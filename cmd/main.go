package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"b1_projet-JS_Forum/internal/db"
	"b1_projet-JS_Forum/internal/handlers"
)

func main() {

	db.InitDB()
	db.CreateTables()
	db.CreateAdminIfNotExists()

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/delete-user", handlers.DeleteUser)
	http.HandleFunc("/create-post", handlers.CreatePost)
	http.HandleFunc("/profile", handlers.Profile)
	http.HandleFunc("/send_message", handlers.SendMessage)
	http.HandleFunc("/photos", handlers.Photos)
	http.HandleFunc("/soirees", handlers.Soirees)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/templates/static"))))
	http.Handle("/web/fonts/", http.StripPrefix("/web/fonts/", http.FileServer(http.Dir("web/fonts"))))
	http.Handle("/web/assets/", http.StripPrefix("/web/assets/", http.FileServer(http.Dir("web/assets"))))
	http.HandleFunc("/post", handlers.ViewPost)
	http.HandleFunc("/create-comment", handlers.CreateComment)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/all-post", handlers.Home)
	http.HandleFunc("/send-message", handlers.SendMessageAdmin)
	http.HandleFunc("/admin", handlers.AdminPage)
	http.HandleFunc("/admin/promote", handlers.PromoteUser)

	// Ouvre automatiquement le navigateur
	go func() {
		time.Sleep(500 * time.Millisecond)
		openBrowser("http://localhost:8080")
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	}

	if cmd != nil {
		cmd.Start()
	}
}
