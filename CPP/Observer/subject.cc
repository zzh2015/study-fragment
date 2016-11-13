#include "subject.h"

#include "observer.h"

Subject::Subject() {}
Subject::~Subject() {}

void Subject::Attach(Observer *obj)
{
    observers_.push_back(obj);
}

void Subject::Detach(Observer *obj)
{
    for (auto it = observers_.cbegin(); it != observers_.cend(); ++it) {
        if (*it == obj) {
            observers_.erase(it);
            break;
        }
    }
}

void Subject::Notify()
{
    for (auto it : observers_) {
        it->Update();
    }
}

ConcreteSubject::ConcreteSubject() : Subject(), datas_(0) {}
ConcreteSubject::~ConcreteSubject() {}


void ConcreteSubject::SetState(const int32_t data)
{
    datas_ = data;
}

int32_t ConcreteSubject::GetState() const
{
    return datas_;
}
