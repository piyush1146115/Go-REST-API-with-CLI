package api

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var MyRouter = mux.NewRouter().StrictSlash(true)
var wait time.Duration
var server *http.Server

type Article struct {
	Id      string `json: "id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

var users = map[string]string{
	"test":  "secret",
	"user1": "password1",
	"user2": "password2",
}

func isAuthorised(username, password string) bool {
	pass, ok := users[username]
	if !ok {
		return false
	}
	return password == pass
}

func CreateDB() {
	Articles = []Article{
		Article{Id: "1", Title: "Test title", Desc: "Test Description", Content: "Hello World"},
		Article{Id: "2", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "3", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func CreateServer(port string) {
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "The duration for which our server gracefully wait for existing connections to finish")
	flag.Parse()
	//fmt.Println(port)
	port = ":" + port

	server = &http.Server{
		Addr:         port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      MyRouter,
	}
}

func StartServer() {
	//fmt.Println("From startserver " + server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	GracefulShutDown()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println(w, "Homepage Endpoint Hit")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var notFound bool
	notFound = true
	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			err := json.NewEncoder(w).Encode(article)
			notFound = false
			return
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

	if notFound == true {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Return single article Endpoint hit!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content-Type", "application/json")
	//username, password, ok := r.BasicAuth()
	//
	//if !ok {
	//	w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte(`{"message": "No basic auth present"}`))
	//	return
	//}
	//
	//if !isAuthorised(username, password) {
	//	w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
	//	w.WriteHeader(http.StatusUnauthorized)
	//	w.Write([]byte(`{"message": "Invalid username or password"}`))
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Endpoint Hit: returnAllArticles")
	err := json.NewEncoder(w).Encode(Articles)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Get all articles endpoint hit!")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	err := json.NewEncoder(w).Encode(article)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Create Endpoint Hit")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	var notFound bool
	notFound = true

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			notFound = false
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

	if notFound == true {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(w, "DELETE Endpoint Hit")
}

func updateArticles(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var notFound bool
	notFound = true

	reqbody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqbody, &article)

	for index, art := range Articles {
		// if our id path parameter matches one of our articles

		if art.Id == id {
			notFound = false
			// updates our Articles array to remove the article
			Articles[index] = article
		}
	}

	if notFound == true {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(article)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(w, "UPDATE Endpoint Hit")
}

func handleUnknownRequests(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func handleRequests() {
	MyRouter.HandleFunc("/", homePage)
	MyRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("Get")
	MyRouter.HandleFunc("/articles", returnAllArticles).Methods("Get")
	MyRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	MyRouter.HandleFunc("/article/{id}", updateArticles).Methods("PUT")
	MyRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	MyRouter.PathPrefix("/").HandlerFunc(handleUnknownRequests)
}

func GracefulShutDown() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)
	log.Println("Shutting down!!")
	os.Exit(0)
}

func init() {
	CreateDB()
	handleRequests()
}
