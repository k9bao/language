include(FetchContent)

OPTION(ENABLE_GTEST "enable googletest" OFF)
OPTION(ENABLE_GMOCK "enable googlemock" OFF)
OPTION(ENABLE_GBENCH "enable google benchmark" ON)
OPTION(ENABLE_GFLAGS "enable gflags" OFF)

# option(BUILD_GMOCK "Builds the googlemock subproject" OFF)
# option(INSTALL_GTEST "Enable installation of googletest" OFF)
# option(BENCHMARK_ENABLE_INSTALL "Enable installation of benchmark" OFF)
option(BENCHMARK_ENABLE_TESTING "Enable installation of benchmark" OFF)
option(BENCHMARK_ENABLE_GTEST_TESTS "Enable installation of benchmark" OFF)

#############
# gtest
#############

if (ENABLE_GTEST)

FetchContent_Declare(
    googletest
    URL ${EP_MIRROR_DIR}/googletest-release-1.8.1.tar.gz 
    URL_MD5 2e6fbeb6a91310a16efe181886c59596
    DOWNLOAD_DIR ${EP_CACHE_DIR}
    DOWNLOAD_NO_PROGRESS ON
)
FetchContent_GetProperties(googletest)
if(NOT googletest_POPULATED)
    FetchContent_Populate(googletest)

    add_subdirectory(${googletest_SOURCE_DIR} EXCLUDE_FROM_ALL)
endif()

endif (ENABLE_GTEST)

#############
# benchmark
#############

if(ENABLE_GBENCH)

FetchContent_Declare(
    googlebenchmark
    URL https://codeload.github.com/google/benchmark/tar.gz/v1.5.0
    URL_MD5 eb1466370f3ae31e74557baa29729e9e
    DOWNLOAD_DIR ${EP_CACHE_DIR}
    DOWNLOAD_NO_PROGRESS ON
)
FetchContent_GetProperties(googlebenchmark)
if(NOT googlebenchmark_POPULATED)
    FetchContent_Populate(googlebenchmark)

    if(ANDROID)
        set(HAVE_STD_REGEX 1)
        set(HAVE_GNU_POSIX_REGEX 1)
        set(HAVE_POSIX_REGEX 1)
        set(HAVE_STEADY_CLOCK 1)
    endif()

    add_subdirectory(${googlebenchmark_SOURCE_DIR} EXCLUDE_FROM_ALL)
endif()

add_library(gbench ALIAS benchmark)

endif(ENABLE_GBENCH)
