package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jimsnab/go-lane"
)

func main() {
	l := lane.NewLogLane(context.Background())
	l.Info("Starting the service...")

	diskLane, err := lane.NewDiskLane(context.Background(), "test.log")
	if err != nil {
		fmt.Println("Error creating DiskLane:", err)
		return
	}
	defer diskLane.Close()

	for i := 0; i < 6; i++ {
		diskLane.Debug("Debug log message using DiskLane: ", i)
		diskLane.Info("Info log message using DiskLane: ", i)
		time.Sleep(1 * time.Second)
	}
}
