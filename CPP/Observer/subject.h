#ifndef SUBJECT_H_
#define SUBJECT_H_

#include <stdint.h>
#include <vector>

class Observer;

class Subject {
    public:
        Subject();
        virtual ~Subject();

        virtual void Attach(Observer *);
        virtual void Detach(Observer *);
        virtual void Notify();
        virtual void SetState(const int32_t) = 0;
        virtual int32_t GetState() const = 0;
    private:
        std::vector<Observer *> observers_;
};

class ConcreteSubject : public Subject {
    public:
        ConcreteSubject();
        ~ConcreteSubject();
        void SetState(const int32_t);
        int32_t GetState() const;
    private:
        int32_t datas_;
};
#endif
