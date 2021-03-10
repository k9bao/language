cmake_minimum_required(VERSION 3.0.0)

project(utiltest)

include(CTest)
enable_testing()

if(MSVC)
    add_compile_options("/utf-8")
    add_compile_options("/std:c++latest")
endif()

add_subdirectory(lest)

include(cmake/gtest2.cmake)

add_subdirectory(gtest)
