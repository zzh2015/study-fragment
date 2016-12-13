# cmake 版本要求
cmake_minimum_required (VERSION 2.8)
# 项目信息
project (sockets)
# compile flags
set(CMAKE_CXX_FLAGS "-std=c++11 -O2 -Wall")

include_directories(../libev/include)
set (LINK_DIR ../libev/lib)

link_directories (${LINK_DIR})
# 源文件
aux_source_directory(. DIR_SRCS)
# 编译
add_executable(socks5client ${DIR_SRCS})

target_link_libraries (socks5client ev pthread m)


