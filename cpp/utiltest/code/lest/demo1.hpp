#include <stdio.h>
#include <cstring>
#include <stdexcept>

int add(int a, int b) {
    return a + b;
}

int throwfun(){
    throw std::invalid_argument("check arg fail");
}

class App {
public:
    int add(int a, int b) {
        return a + b;
    }
};