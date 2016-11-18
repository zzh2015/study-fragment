# cmake 版本要求
cmake_minimum_required (VERSION 2.8)
# 项目信息
project (Thread)
# compile flags
set(CMAKE_CXX_FLAGS "-O2 -Wall -std=c++11")
# 源文件
# aux_source_directory(. DIR_SRCS)
# 编译
add_executable(tg thread_guard_test.cc)
target_link_libraries(tg pthread)

add_executable(tv thread_product_test.cc)
target_link_libraries(tv pthread)

add_executable(ta parallel_accumulate.cc)
target_link_libraries(ta pthread)
