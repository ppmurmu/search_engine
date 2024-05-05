package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	utils "github.com/ppmurmu/search_engine/utils"
)

func main() {
	fmt.Println("hello woeld")
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml", "wiki abstract dump path")
	// if .xml file used then no decoder is needed in the utils.LoadDocuments function.
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	log.Println("Running Full Text Search")
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since((start)))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()

	matchedId := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedId), time.Since(start))
	for _, id := range matchedId {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

}
