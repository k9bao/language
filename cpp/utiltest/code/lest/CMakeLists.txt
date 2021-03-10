file(GLOB test_main_SRCS *_test.cpp)
add_executable(lest test.cpp ${test_main_SRCS})
target_include_directories(lest INTERFACE .)
