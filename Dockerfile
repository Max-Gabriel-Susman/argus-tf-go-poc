
RUN apt-get -y install --no-install-recommends \
  libprotobuf-dev \
  protobuf-compiler

RUN wget -q --no-check-certificate https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.11.0.tar.gz && \
  tar -C /usr/local -xzf libtensorflow-cpu-linux-x86_64-2.11.0.tar.gz && \
  rm -rf libtensorflow-cpu-linux-x86_64-2.11.0.tar.gz && \
  ldconfig /usr/local/lib
