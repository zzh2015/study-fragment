/*************************************************************************
 > File Name: parallel_accumulate.cc
 > Author:
 > Mail:
 > Created Time: 2016年11月18日 星期五 13时37分37秒
 ************************************************************************/
#include <algorithm>
#include <iostream>
#include <iterator>
#include <numeric>
#include <thread>
#include <vector>

template<typename It, typename T>
struct AccumulateBlock {
    void operator() (It first, It last, T &result) {
        result = std::accumulate(first, last, result);
    }
};

template<typename It, typename T>
T ParallelAccumulate(It first, It end, T init) {
    auto length = std::distance(first, end);

    if (!length) {
        return init;
    }

    auto minPerThread = 25;
    auto maxThreads = (length + minPerThread - 1) / minPerThread;
    auto hardwareThread = std::thread::hardware_concurrency();

    auto numThreads = std::min(static_cast<unsigned long>(hardwareThread != 0 ? hardwareThread : 2), static_cast<unsigned long>(maxThreads));
    auto blockSize = length / numThreads;

    std::vector<T> results(numThreads);
    std::vector<std::thread> threads(numThreads - 1);

    auto start = first;
    for (unsigned long i=0; i<numThreads-1; ++i) {
        auto end = start;
        std::advance(end, blockSize);
        threads[i] = std::thread(AccumulateBlock<It, T>(), start, end, std::ref(results[i]));
        start = end;
    }

    AccumulateBlock<It, T>()(start, end, results[numThreads - 1]);
    std::for_each(threads.begin(), threads.end(), [](std::thread &th){
                th.join();
            });

    return std::accumulate(results.begin(), results.end(), init);
}

int main() {
    std::vector<unsigned long> v;
    for (auto i=10000; i<110000; ++i) {
        v.push_back(i);
    }

    std::cout << "thread id: " << std::this_thread::get_id() << std::endl;
    unsigned long init = 0;
    std::cout << "accumulate: " << std::accumulate(v.begin(), v.end(), init) << std::endl;
    init = 0;
    std::cout << "parallel accumulate: " << ParallelAccumulate(v.begin(), v.end(), init) << std::endl;

    return 0;
}
