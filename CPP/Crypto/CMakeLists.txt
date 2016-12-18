# cmake 版本要求
cmake_minimum_required (VERSION 2.8)
# 项目信息
project (crypto C)
# compile flags
set(CMAKE_C_FLAGS "-O2 -Wall")


# 源文件
aux_source_directory(. DIR_SRCS)

# 编译
add_executable(chacha20.out ${DIR_SRCS})

target_link_libraries (chacha20.out sodium)


