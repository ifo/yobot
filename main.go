package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	gzb "github.com/ifo/gozulipbot"
)

func main() {
	var (
		port      = flag.Int("port", 7070, "Server port")
		emailAddr = os.Getenv("ZULIP_EMAIL")
		zAPIKey   = os.Getenv("ZULIP_KEY")
		yoAPIKey  = os.Getenv("YO_KEY")
	)
	flag.Parse()

	bot, err := gzb.MakeBot(emailAddr, apiKey, []string{"bot-test"})
	if err != nil {
		log.Fatal(err)
	}

	c := Context{
		YoKey: yoAPIKey,
		ZBot:  &bot,
	}

	http.HandleFunc("/", injectContext(index, c))

	log.Printf("Starting server on port %d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type Context struct {
	YoKey string
	ZBot  *gzb.Bot
}

func index(w http.ResponseWriter, r *http.Request, c Context) {
	user := r.FormValue("username")
	link := r.FormValue("link")
	fmt.Println(user, link)

}

func injectContext(fn func(w http.ResponseWriter, r *http.Request, c Context), c Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, c)
	}
}
