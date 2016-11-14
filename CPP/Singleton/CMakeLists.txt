# cmake 版本要求
cmake_minimum_required (VERSION 2.8)
# 项目信息
project (Singleton C CXX)
# compile flags
set(CMAKE_CXX_FLAGS "-O2 -Wall -std=c++11")
# 源文件
aux_source_directory(. DIR_SRCS)
# 编译
add_executable(singleton ${DIR_SRCS})
# 依赖库
target_link_libraries(singleton pthread)
