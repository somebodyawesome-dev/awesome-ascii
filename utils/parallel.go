package utils

import (
	"image"
	"math"
	"runtime"
	"sync"
)

func ParallelImageProcess(size image.Point, f func(x int, y int)) {
	procs := runtime.GOMAXPROCS(0)
	xStep := int(math.Floor(float64(size.X) / float64(procs)))
	yStep := int(math.Floor(float64(size.Y) / float64(procs)))

	var waitGroup sync.WaitGroup
	for i := 0; i < procs; i++ {
		startX := i * xStep
		var endX int
		if i < procs-1 {
			endX = (i + 1) * xStep
		} else {
			endX = size.X
		}
		for j := 0; j < procs; j++ {
			startY := j * yStep
			var endY int
			if j < procs-1 {
				endY = (j + 1) * yStep
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
