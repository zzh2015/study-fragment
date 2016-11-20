#include <algorithm>
#include <mutex>
#include <stdint.h>
#include <stdio.h>
#include <thread>
#include <vector>

std::vector<uint8_t> kVector;
std::mutex kMutex;

// std::lock lock(m1, m2); lock could lock more than one mutext
// std::lock_guard lg(m1, std::adopt_lock);
// std::lock_guard lg(m2, std::adopt_lock);
void add_to_vector(const uint8_t element) {
    std::lock_guard<std::mutex> guard(kMutex);
    for (uint8_t i=0; i<element; ++i) {
        kVector.push_back(element);
    }
}

void access_from_vector() {
    std::lock_guard<std::mutex> guard(kMutex);
    std::for_each(kVector.begin(), kVector.end(), [](uint8_t e){
                printf("element: %u\n", e);
            });
}

int main() {
    uint8_t value = 16;
    std::thread addThread(add_to_vector, value);
    std::thread accessThread(access_from_vector);

    addThread.join();
    accessThread.join();
}
