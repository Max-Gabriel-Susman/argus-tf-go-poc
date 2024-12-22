package main 

func main() {
	wg := sync.WaitGroup{}

	tf := tensorflow.NewTensorflowMachineLearning()
	openCV := opencv.NewOpenCVImageProcessor()

	modelPath, err := filepath.Abs("data/models/ssd_mobilenet_v1_coco_2018_01_28/saved_model")
	if err != nil {
		log.Fatal(err)
	}

	model, tfSession, err := tf.LoadSavedModel(modelPath)
	if err != nil {
		log.Fatal(err)
	}
	defer model.Session.Close()
	defer tfSession.Close()

	runDetectionUseCase := usecase.NewRunDetectionUseCase(tf, openCV)
	 
	log.Println("Argus Stream Engine Service online.")

	for index, camera := range getCameras() {
		wg.Add(1)
		defer wg.Done()

		go worker(index+1, camera, openCV, runDetectionUseCase)
	}

	wg.Wait()
}