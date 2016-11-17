/*************************************************************************
  > File Name: factory.h
  > Author:
  > Mail:
  > Created Time: 2016年11月17日 星期四 11时22分28秒
 ************************************************************************/

#ifndef FACTORY_H_
#define FACTORY_H_

#include <memory>

class Product;
enum class ProductOptions {
    ToyOption,
    GameOption
};

class Factory {
    public:
        Factory() {}
        virtual ~Factory() {}

        virtual std::unique_ptr<Product> ProductCreate(ProductOptions) = 0; //?
};

class ProductFactory : public Factory {
    public:
        ProductFactory() : Factory() {}
        ~ProductFactory() {}

        std::unique_ptr<Product> ProductCreate(ProductOptions);
};

#endif
