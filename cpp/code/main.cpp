#include <iostream>
using namespace std;

int TestPoint(float* pOut){
    float in = 4.3;
    //*pOut = in;
    ::memcpy(pOut, &in, sizeof(float));
    return 0;
}

int TestPPoint(float** pOut){
    float in = 4.4;
    (*(*pOut)) = in;
    //::memcpy(*pOut, &in, sizeof(float));
    return 0;
}

int main(int, char**) {
    std::cout << "Hello, world!\n";
    float in = 15;
    std::cout<<in<<endl;

    TestPoint(&in);
    std::cout<<in<<endl;

    float* pIn = &in;
    TestPPoint(&pIn);
    std::cout<<in<<endl;
}