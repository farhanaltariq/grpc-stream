package main

import (
	pb "github.com/farhanaltariq/protomessager/go/protomessagergateway"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("Hello World")
	a := &pb.MapRequest{}
	logrus.Println(a)
}
