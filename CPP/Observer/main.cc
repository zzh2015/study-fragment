#include "observer.h"
#include "subject.h"

#include <memory>

int main(int argc, char **argv)
{
    std::shared_ptr<Subject> sub(new ConcreteSubject());

    std::shared_ptr<Observer> one(new ConcreteObserverOne(sub.get()));
    {
        std::shared_ptr<Observer> two(new ConcreteObserverTwo(sub.get()));
        sub.get()->SetState(1024);
        sub.get()->Notify();
    }
    sub.get()->SetState(2048);
    sub.get()->Notify();

    return 0;
}
