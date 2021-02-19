package concurrency

// WebsiteChecker is a function that accepts a string and returns a boolean
type WebsiteChecker func(string) bool

// result private struct
type result struct {
	string
	bool
}

// CheckWebsites checks a list of urls with a given function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)

	// creating a channel of type result
	resultChannel := make(chan result)

	for _, url := range urls {
		// creating an anonumous function to start a go routine
		// and calling it right after declaring it
		go func(u string) {
			// sending a result struct for each call to wc to the resultChannel
			// <- is the send statement (channel <- value)
			resultChannel <- result{u, wc(u)} // this is happening in parallel
		}(url) // binding the url
	}

	for i := 0; i < len(urls); i++ {
		// receive expression which assigns a value received from
		// a channel to a variable (variable := <- channel)
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
