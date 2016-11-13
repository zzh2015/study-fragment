#include "observer.h"

#include <iostream>
#include "subject.h"

ConcreteObserverOne::ConcreteObserverOne(Subject *obj) : Observer(obj), datas_(0) {
    subjects_->Attach(this);
}

ConcreteObserverOne::~ConcreteObserverOne() {
    subjects_->Detach(this);
}

void ConcreteObserverOne::Update()
{
    datas_ = subjects_->GetState();
    Draw();
}

void ConcreteObserverOne::Draw()
{
    std::cout << "Draw Line: " << datas_ << std::endl;
}

ConcreteObserverTwo::ConcreteObserverTwo(Subject *obj) : Observer(obj), datas_(0) {
    subjects_->Attach(this);
}

ConcreteObserverTwo::~ConcreteObserverTwo() {
    subjects_->Detach(this);
}
void ConcreteObserverTwo::Update()
{
    datas_ = subjects_->GetState();
    Draw();
}

void ConcreteObserverTwo::Draw()
{
    std::cout << "Draw Circle: " << datas_  << std::endl;
}


