/*************************************************************************
 > File Name: product.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月17日 星期四 11时10分21秒
 ************************************************************************/

#include "product.h"

#include <iostream>

Product::Product(const std::string &name) : name_(name) {}
Product::~Product() {}

ToyProduct::ToyProduct() : Product("toy product") {}
ToyProduct::~ToyProduct() {}
void ToyProduct::ShowName() {
    std::cout << "name_: " << name_ << std::endl;
}

GameProduct::GameProduct() : Product("game product") {}
GameProduct::~GameProduct() {}

void GameProduct::ShowName() {
    std::cout << "name_: " << name_ << std::endl;
}

