PROJECT(SERVER)
SET(CMAKE_C_COMPILER g++)
SET(CMAK_CXX_COMPILER g++)
SET(CMAKE_BUILD_TYPE "Debug")
SET(CMAKE_CXX_FLAGS_DEBUG "$ENV{CXXFLAGS} -O0 -Wall -g3 -ggdb")
SET(CMAKE_CXX_FLAGS_RELEASE "$ENV{CXXFLAGS} -O3 -Wall")

add_compile_options(-std=c++0x)
SET(SRC_LIST main.cpp ErrorNoDefine.h Event.h EventCenter.h EventMutiIOAPI.h EventMutiIOEpoll.h TCPListener.h TCPServer.h Event.cpp EventCenter.cpp EventMutiIOAPI.cpp EventMutiIOEpoll.cpp TCPListener.cpp TCPServer.cpp)
MESSAGE(STATUS "This is BINARY dir " ${HELLO_BINARY_DIR})
MESSAGE(STATUS "This is SOURCE dir "${HELLO_SOURCE_DIR})
ADD_EXECUTABLE(server  ${SRC_LIST})
