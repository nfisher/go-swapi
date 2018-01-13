package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/nfisher/swapi"
	"github.com/nfisher/swapi/search"
)

type RuntimeConfig struct {
	JsonPath string
}

func main() {
	var config RuntimeConfig

	flag.StringVar(&config.JsonPath, "filename", "swapi.json", "Star Wars JSON API file path.")
	flag.Parse()

	u, err := swapi.LoadUniverse(config.JsonPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := search.NewCorpus(u)
	index := search.NewIndex(true, 14)
	err = index.Train(p.Corpus)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Enter your query:")
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if line == "" {
			break
		}

		qi, err := index.Query(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		v, err := p.Retrieve(qi)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%#v\n", v)
	}
}
