#include <atomic>
#include <iostream>

using namespace std;

class demoClass {
public:
    demoClass() : m1(1), m3{3}, m4(4) {}

    atomic<int> m1;
    atomic<int> m2{2};
    atomic<int> m3;
    atomic<int> m4;
    // atomic<int> m5(1);//expected identifier before numeric constant

    static atomic<int> mc1;
    static atomic<int> mc2;
    static atomic<int> mc3;
    static atomic<int> mc4;
    // const atomic<int> mc5(1); //expected identifier before numeric constant
};
atomic<int> demoClass::mc1(1);
atomic<int> demoClass::mc2{2};
atomic<int> demoClass::mc3{3};
atomic<int> demoClass::mc4{4};

int main() {
    cout << "hello world" << endl;
    atomic<int> i1(5);
    atomic<int> i2{6};
    cout << i1.load() << "," << i2 << endl;

    demoClass d1;
    cout << d1.m1 << "," << d1.m2 << "," << d1.m3 << "," << d1.m4 << endl;
    cout << d1.mc1 << "," << d1.mc2 << "," << d1.mc3 << "," << d1.mc4 << endl;
    return 0;
}