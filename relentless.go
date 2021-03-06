package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/signup", signupH)
	http.Handle("/", http.FileServer(http.Dir("app")))

	http.ListenAndServe(":80", nil)
}

type signup struct {
	Name    string        `json:"name"`
	Email   string        `json:"email"`
	Options signupOptions `json:"options"`
}

type signupOptions struct {
	Gap bool `json:"gap"`
}

func signupH(w http.ResponseWriter, r *http.Request) {
	var s signup
	d := json.NewDecoder(r.Body)
	d.Decode(&s)
	fmt.Println(s)

	notifyWithEmail("New recruit", s.Name+": "+s.Email, "jacob@relentlessinvest.com")
}

func notifyWithEmail(subject, body string, emails ...string) {
	hostname, _ := os.Hostname()
	if strings.HasSuffix(hostname, "local") {
		fmt.Println(subject)
		fmt.Println(body)
		return
	}
	emailAddresses := strings.Join(emails, ",")

	cmd := exec.Command("mail", "-s", subject, "-a", "Content-Type: text/html", "-t", emailAddresses)
	cmd.Stdin = strings.NewReader(body)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	if string(out) != "" {
		fmt.Printf("email response: %q\n", string(out))
	}
	return
}
