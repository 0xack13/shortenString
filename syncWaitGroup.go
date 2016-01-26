package main

import(
  "net/http"
  "fmt"
  "os"
  "io/ioutil"
  "sync"
)

func main() {

    var wg sync.WaitGroup
    var urls = []string{
            "http://www.golang.org/",
            "http://www.google.com/",
            "http://www.bing.com/",
    }
    for _, url := range urls {
            // Increment the WaitGroup counter.
            wg.Add(1)
            // Launch a goroutine to fetch the URL.
            go func(url string) {
                    // Decrement the counter when the goroutine completes.
                    defer wg.Done()
                    // Fetch the URL.
                    response, err := http.Get(url)
		if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
	    fmt.Printf("*********************************")
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
	
	    fmt.Printf("*********************************")
    }
            }(url)
    }
    // Wait for all HTTP fetches to complete.
    wg.Wait()

}
