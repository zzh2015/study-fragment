/*************************************************************************
 > File Name: product.h
 > Author:
 > Mail:
 > Created Time: 2016年11月17日 星期四 11时10分15秒
 ************************************************************************/

#ifndef PRODUCT_H_
#define PRODUCT_H_

#include <string>

class Product {
    public:
        Product(const std::string &);
        virtual ~Product();

        virtual void ShowName() = 0;
    protected:
        std::string name_;
};

class ToyProduct : public Product {
    public:
        ToyProduct();
        ~ToyProduct();

        void ShowName();
};

class GameProduct : public Product {
    public:
        GameProduct();
        ~GameProduct();

        void ShowName();
};

#endif
