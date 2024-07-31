package utils

import (
	"image"
	"math"
	"runtime"
	"sync"
)

func ParallelImageProcessing(size image.Point, callback func(x, y int)) {

	var ws sync.WaitGroup
	for x := 0; x < size.X; x++ {

		for y := 0; y < size.Y; y++ {
			ws.Add(1)
			go func(x, y int) {
				defer ws.Done()
				callback(x, y)
			}(x, y)

		}
	}
	ws.Wait()
}
func ParallelForEachPixel(size image.Point, f func(x int, y int)) {
	procs := runtime.GOMAXPROCS(0)
	var waitGroup sync.WaitGroup
	for i := 0; i < procs; i++ {
		startX := i * int(math.Floor(float64(size.X)/float64(procs)))
		var endX int
		if i < procs-1 {
			endX = (i + 1) * int(math.Floor(float64(size.X)/float64(procs)))
		} else {
			endX = size.X
		}
		for j := 0; j < procs; j++ {
			startY := j * int(math.Floor(float64(size.Y)/float64(procs)))
			var endY int
			if j < procs-1 {
				endY = (j + 1) * int(math.Floor(float64(size.Y)/float64(procs)))
			} else {
				endY = size.Y
			}
			waitGroup.Add(1)
			go func(sX int, eX int, sY int, eY int) {
				defer waitGroup.Done()
				for x := sX; x < eX; x++ {
					for y := sY; y < eY; y++ {
						f(x, y)
					}
				}
			}(startX, endX, startY, endY)
		}
	}
	waitGroup.Wait()
}
