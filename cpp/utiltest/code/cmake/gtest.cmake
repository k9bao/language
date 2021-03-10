include(ExternalProject)

OPTION(ENABLE_GTEST "enable googletest" ON)
OPTION(ENABLE_GMOCK "enable googlemock" OFF)
OPTION(ENABLE_GBENCH "enable google benchmark" ON)
OPTION(ENABLE_GFLAGS "enable gflags" OFF)

#############
# gtest
#############

if (ENABLE_GTEST)

get_library_name(gtest_NAME gtest DEBUG)
get_library_name(gtest_main_NAME gtest_main DEBUG)
get_library_name(gmock_NAME gmock DEBUG)
get_library_name(gmock_main_NAME gmock_main DEBUG)

ExternalProject_Add( 
    gtest_build

    URL ${EP_MIRROR_DIR}/googletest-release-1.8.1.tar.gz 
    URL_MD5 2e6fbeb6a91310a16efe181886c59596
    DOWNLOAD_DIR ${EP_CACHE_DIR}
    DOWNLOAD_NO_PROGRESS ON

    LOG_CONFIGURE ON

    CMAKE_GENERATOR ${CMAKE_GENERATOR}
    CMAKE_ARGS
        -DBUILD_GMOCK:BOOL=${ENABLE_GMOCK}
        -DINSTALL_GTEST:BOOL=ON
        -DCMAKE_INSTALL_PREFIX=<INSTALL_DIR>
        -DCMAKE_BUILD_TYPE=${CMAKE_BUILD_TYPE}
        ${CMAKE_RT_FLAGS}
    BUILD_BYPRODUCTS <INSTALL_DIR>/lib/${gtest_NAME}
)
ExternalProject_Get_Property(gtest_build INSTALL_DIR)

add_cc_interface(gtest 
    INCS ${INSTALL_DIR}/include
    LIBS ${INSTALL_DIR}/lib/${gtest_NAME}
    DEPS gtest_build
)
add_cc_interface(gtest_main 
    INCS ${INSTALL_DIR}/include
    LIBS ${INSTALL_DIR}/lib/${gtest_main_NAME}
    DEPS gtest_build
)

if(ENABLE_GMOCK)

add_cc_interface(gmock
    INCS ${INSTALL_DIR}/include
    LIBS ${INSTALL_DIR}/lib/${gmock_NAME}
    DEPS gtest_build
)
add_cc_interface(gmock_main 
    INCS ${INSTALL_DIR}/include
    LIBS ${INSTALL_DIR}/lib/${gmock_main_NAME}
    DEPS gtest_build
)

endif(ENABLE_GMOCK)


endif (ENABLE_GTEST)

#############
# benchmark
#############

if(ENABLE_GBENCH)

get_library_name(gbench_NAME benchmark DEBUG)
get_library_name(gbench_main_NAME benchmark_main DEBUG)

ExternalProject_Add(
    gbench_build

    URL ${EP_MIRROR_DIR}/benchmark-1.5.0.tar.gz  
    URL_MD5 eb1466370f3ae31e74557baa29729e9e
    DOWNLOAD_DIR ${EP_CACHE_DIR}
    DOWNLOAD_NO_PROGRESS ON

    LOG_CONFIGURE ON

    CMAKE_GENERATOR ${CMAKE_GENERATOR}
    
    CMAKE_ARGS
        -DBENCHMARK_ENABLE_GTEST_TESTS:BOOL=OFF
        -DBENCHMARK_ENABLE_TESTING:BOOL=OFF
        -DCMAKE_INSTALL_PREFIX=<INSTALL_DIR>
        -DCMAKE_BUILD_TYPE=${CMAKE_BUILD_TYPE}
        -DCMAKE_DEBUG_POSTFIX=d
        ${CMAKE_RT_FLAGS}
    BUILD_BYPRODUCTS <INSTALL_DIR>/lib/${gbench_NAME}
)
ExternalProject_Get_Property(gbench_build INSTALL_DIR)

add_cc_interface(gbench 
    INCS ${INSTALL_DIR}/include
    LIBS ${INSTALL_DIR}/lib/${gbench_NAME}
    DEPS gbench_build
)
add_cc_interface(gbench_main 
    INCS ${INSTALL_DIR}/include
    LIBS ${INSTALL_DIR}/lib/${gbench_main_NAME}
    DEPS gbench_build
)

if(WIN32)
    target_link_libraries(gbench INTERFACE Shlwapi.lib)
endif()

endif(ENABLE_GBENCH)

#############
# gflags
#############

if(ENABLE_GFLAGS)

get_library_name(gflags_NAME gflags_nothreads_static DEBUG_POSTFIX _debug)

ExternalProject_Add( 
    gflags_build

    URL ${EP_MIRROR_DIR}/gflags-2.2.2.tar.gz 
    URL_MD5 1a865b93bacfa963201af3f75b7bd64c
    DOWNLOAD_DIR ${EP_CACHE_DIR}
    DOWNLOAD_NO_PROGRESS ON
    CMAKE_GENERATOR ${CMAKE_GENERATOR}
    CMAKE_ARGS
        -DREGISTER_INSTALL_PREFIX:BOOL=OFF
        -DCMAKE_INSTALL_PREFIX=<INSTALL_DIR>
        -DCMAKE_BUILD_TYPE=${CMAKE_BUILD_TYPE}
        ${CMAKE_RT_FLAGS}
    BUILD_BYPRODUCTS <INSTALL_DIR>/lib/${gflags_NAME}
)
ExternalProject_Get_Property(gflags_build INSTALL_DIR)

add_cc_interface(gflags 
    INCS ${INSTALL_DIR}/include
    LIBS ${INSTALL_DIR}/lib/${gflags_NAME}
    DEPS gflags_build
)

endif(ENABLE_GFLAGS)
