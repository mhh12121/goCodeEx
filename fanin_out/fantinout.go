package main

import "sync"

//Convert a list of input (*Image) to a channel that emits the data in the list
func GenReq(images ...*Image) <-chan *Image {
	out := make(chan *Image)
	go func() {
		for _, n := range images {
			out <- n
		}
		close(out)
	}()
	return out
}

//FunOut: receives the input from a channel and return a channel that emits the result.
func FunOut(in <-chan *Image) <-chan *Image {
	out := make(chan *Image)
	go func() {
		for n := range in {
			out <- DoImageProcessing(n)
		}
		close(out)
	}()
	return out
}

//FanIn:  converts a list of channels to a single channel
func FanIn(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan *Image)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan *Image) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := GenReq(image1, image2, image3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := FanOut(in)
	c2 := FanOut(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		//return output here.
	}
}
