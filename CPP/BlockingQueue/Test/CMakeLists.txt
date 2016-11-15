aux_source_directory(. DIR_SRCS)

add_executable(test_mutex ${DIR_SRCS} ../condition.cc)

target_link_libraries(test_mutex pthread)
