set(base_SRCS
    post.c
  )

add_executable(curl-post ${base_SRCS})
target_link_libraries(curl-post curl)

add_executable(curl-get get.c)
target_link_libraries(curl-get curl)

add_executable(curl-https-get get_https.c)
target_link_libraries(curl-https-get curl)
