# cmake 版本要求
cmake_minimum_required (VERSION 2.8)
# 项目信息
project (BlockingQueue)
# compile flags
set(CMAKE_CXX_FLAGS "-O2 -Wall -std=c++11")
# 子目录
add_subdirectory(Test)
# 源文件
# aux_source_directory(. DIR_SRCS)
# 编译
# add_executable(observer ${DIR_SRCS})

