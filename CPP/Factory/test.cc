/*************************************************************************
 > File Name: test.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月17日 星期四 11时48分16秒
 ************************************************************************/
#include "factory.h"
#include "product.h"

int main()
{
    std::unique_ptr<Factory> fac(new ProductFactory());

    std::unique_ptr<Product> toy = fac->ProductCreate(ProductOptions::ToyOption);
    toy->ShowName();

    std::unique_ptr<Product> game = fac->ProductCreate(ProductOptions::GameOption);
    game->ShowName();

    return 0;
}
