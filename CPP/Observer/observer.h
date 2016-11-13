#ifndef OBSERVER_H_
#define OBSERVER_H_

#include <stdint.h>

class Subject;

class Observer {
    public:
        Observer(Subject *obj) : subjects_(obj) {};
        virtual ~Observer() {};
        virtual void Update() = 0;
    protected:
        virtual void Draw() = 0;

        Subject *subjects_;
};

class ConcreteObserverOne : public Observer {
    public:
        ConcreteObserverOne(Subject *);
        ~ConcreteObserverOne();

        void Update();
    protected:
        void Draw();
    private:
        int32_t datas_;
};

class ConcreteObserverTwo : public Observer {
    public:
        ConcreteObserverTwo(Subject *);
        ~ConcreteObserverTwo();

        void Update();
    protected:
        void Draw();
    private:
        int32_t datas_;
};
#endif
