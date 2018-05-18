// ipc.h
#ifndef WRAP_CPP_H
#define WRAP_CPP_H

#pragma once
#include <stdio.h>
struct data {
    char *modelPath;
    char *imgPath;
    char *app;
    int projectType;
};

#ifdef __cplusplus
extern "C"{
#endif
extern void CBGateway(struct data Data);
extern int CardIdentify(struct data Data);
#ifdef __cplusplus
}
#endif

#endif // WRAP_CPP_H
