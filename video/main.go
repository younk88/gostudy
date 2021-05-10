package main

import (
	"os"
	"fmt"
	"github.com/abema/go-mp4"
	"github.com/sunfish-shogi/bufseekio"
)

func main()  {
	f := os.Args[1]
	fs, _ := os.Open(f)
	defer fs.Close()
	sinfo, _ := fs.Stat()
	fmt.Printf("%v\n", sinfo)
	info, _ := mp4.Probe(bufseekio.NewReadSeeker(fs, 1024, 4))             
	fmt.Printf("%v\n", info)                              
	fmt.Printf("track %v  duration:%d\n", info.Tracks, info.Duration)
	for _, track := range info.Tracks {
		if track != nil && track.AVC != nil {
			fmt.Printf("track %v\n", *track.AVC)
		}
	}
}