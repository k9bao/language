#include <iostream>
#include <utility>
void reference(int& v) {
    std::cout << "lvalue reference" << std::endl;
}
void reference(int&& v) {
    std::cout << "rvalue reference" << std::endl;
}
template <typename T>
void pass(T&& v) {//v is a reference, it is also an lvalue
    reference(v);
}

int main() {
    std::cout << "rvalue pass:";
    pass(1);//output: lvalue reference

    std::cout << "lvalue pass:";
    int v1 = 1;
    pass(v1);//output: lvalue reference

    std::cout << "------------" << std::endl;
    std::cout << "rvalue pass:";
    reference(1);//output: rvalue reference

    std::cout << "lvalue pass:";
    int v2 = 1;
    reference(v2);//output: lvalue reference

    return 0;
}