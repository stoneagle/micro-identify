// ipc.cpp
#include "ipc.h"
#include "clsAPI.h"
#include <iostream>
using namespace std;
using namespace cv;
#include <cv.hpp>

struct thread_param{
	BookReading* objBR;
	cv::Mat* ptr_img;
	int modelid;
	int coverid;
	int threadID;
};

int CardIdentify(struct data card) {
    std::string model_file(card.modelPath, strlen(card.modelPath));
    std::string file(card.imgPath, strlen(card.imgPath));

    if(model_file[model_file.size() - 1] == '/') {
        model_file = model_file.substr(0, model_file.size()-1);
    }

    cv::Mat img = cv::imread(file, -1);
    if (img.empty()) {
        CBGateway(card);
        return 0;
    }

    int modelid = 1;
    thread_param ti;
    ti.objBR = new BookReading(model_file);
    ti.ptr_img = &img;
    ti.modelid = modelid;
    ti.threadID = 1;

    stru_predrst prediction;
    prediction = ti.objBR->readcover_fixedsize(*(ti.ptr_img));
    delete ti.objBR;
    return prediction.bookid;
}
