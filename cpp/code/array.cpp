#include <array>
#include <iostream>
using namespace std;

//https://www.tutorialspoint.com/cpp_standard_library/set.htm
int main() {
    const int count = 5;
    array<int, count> l1; //红黑树实现的非重复集合
    l1.fill(2);           //插入元素 2,2,2,2,2
    for_each(l1.begin(), l1.end(), [](int v) { cout << v << ","; });
    cout << endl;
    for (size_t i = 0; i < count; i++) {
        l1[i] = i + 1;
    }
    for_each(l1.begin(), l1.end(), [](int v) { cout << v << ","; }); //1,2,3,4,5
    cout << endl;

    assert(l1.at(2) == 3);                                    //取出大小
    assert(l1.size() == count && l1.size() == l1.max_size()); //最大容量

    return 0;
}