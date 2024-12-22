package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/Max-Gabriel-Susman/argus-tf-go-poc/internal/object"
	"github.com/Max-Gabriel-Susman/argus-tf-go-poc/internal/stream"
	"github.com/Max-Gabriel-Susman/argus-tf-go-poc/internal/tensorflow"
)

func worker(
	workerId int,
	ingress stream.Source,
	detection *object.Detection,
	// openCV *opencv.OpenCVImageProcessor,
) {
	fmt.Printf("WORKER ID [%v] - SOURCE [%v]\n", workerId, ingress.Label)

	// openCV.CaptureStreamVideo(camera.Input, runDetectionUseCase.Execute)
}

// List of ingresses to run object detection
func getingresses() (ingresses []stream.Source) {
	ingresses = []stream.Source{
		{
			ID:    "1",
			Label: "Camera 1",
			Input: "rtsp://host:port",
		},
	}

	return ingresses
}

func main() {
	wg := sync.WaitGroup{}

	tf := tensorflow.NewTensorflowMachineLearning()
	// openCV := opencv.NewOpenCVImageProcessor()

	modelPath, err := filepath.Abs("data/models/ssd_mobilenet_v1_coco_2018_01_28/saved_model")
	// _, err := filepath.Abs("data/models/ssd_mobilenet_v1_coco_2018_01_28/saved_model")
	if err != nil {
		log.Fatal(err)
	}

	model, tfSession, err := tf.LoadSavedModel(modelPath)
	if err != nil {
		log.Fatal(err)
	}
	defer model.Session.Close()
	defer tfSession.Close()

	detection := object.NewDetection()

	log.Println("Argus Stream Engine Service online.")

	for index, camera := range getingresses() {
		wg.Add(1)
		defer wg.Done()

		// go worker(index+1, camera, openCV, detection)
		go worker(index+1, camera, detection)
	}

	// wg.Wait() // all goroutines are asleep - deadlock!
}
