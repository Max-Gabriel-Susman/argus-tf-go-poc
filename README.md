# Argus TensorFlow Go Proof of concept

This repository is a proof of concept for doing object detection in the go runtime using tensor flow as an ml solution  

go get "github.com/wamuir/graft/tensorflow" "github.com/wamuir/graft/tensorflow/op" "gocv.io/x/gocv" "golang.org/x/image/bmp"

tensorflow installation:
```
sudo wget -q --no-check-certificate https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.11.0.tar.gz 

sudo tar -C /usr/local -xzf libtensorflow-cpu-linux-x86_64-2.11.0.tar.gz 

sudo rm -rf libtensorflow-cpu-linux-x86_64-2.11.0.tar.gz

sudo ldconfig /usr/local/lib
```

opencv installation: 
```
git clone https://github.com/hybridgroup/gocv.git 

cd gocv 

make install

cd ..
  rm -rf gocv
```