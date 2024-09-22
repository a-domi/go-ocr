FROM golang:latest

RUN apt-get update -qq

RUN apt-get install -y -qq libtesseract-dev libleptonica-dev

ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata/

RUN apt-get install -y -qq \
    libtesseract-dev \
    libleptonica-dev \
    tesseract-ocr-eng \
    tesseract-ocr-deu \
    tesseract-ocr-jpn \
    build-essential \
    cmake \
    pkg-config \
    libjpeg-dev \
    libpng-dev \
    libtiff-dev \
    libavcodec-dev \
    libavformat-dev \
    libswscale-dev \
    libv4l-dev \
    libxvidcore-dev \
    libx264-dev \
    libgtk-3-dev \
    libatlas-base-dev \
    gfortran \
    wget \
    unzip \
    curl

# OpenCVのインストール
# RUN wget -O opencv.zip https://github.com/opencv/opencv/archive/4.5.5.zip && \
#     unzip opencv.zip && \
#     mv opencv-4.5.5 /opencv && \
#     cd /opencv && \
#     mkdir build && \
#     cd build && \
#     cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local .. && \
#     make -j$(nproc) && \
#     make install && \
#     ldconfig

# PKG_CONFIG_PATH環境変数の設定
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig

WORKDIR ${GOPATH}/src/github.com/ghost/cool-project
RUN go mod init

RUN go get -t github.com/otiai10/gosseract/v2

# gocvのインストール
RUN go get -u -d gocv.io/x/gocv

