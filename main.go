package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/2-coffee/Text-Search-Engine/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "filename.txt.", "sample text")
	flag.StringVar(&query, "q", "Some text", "search query")
	flag.Parse()
	log.Println("Full text search is in progress")
	start := time.Now()                        // keep track of how long it takes
	docs, err := utils.LoadDocuments(dumpPath) // loading

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs) // join sets
	log.Printf("Indexed %dd documents in %v", len(docs), time.Since(start))

	start = time.Now()

	matchedIds := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIds), time.Since(start))

	for _, id := range matchedIds {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

}
