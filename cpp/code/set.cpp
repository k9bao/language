#include <iostream>
#include <set>
#include <unordered_set>
using namespace std;

//https://www.tutorialspoint.com/cpp_standard_library/set.htm
int main() {
    set<int> l1;                        //红黑树实现的非重复集合
    l1.insert(2);                       //插入元素
    l1.insert(3);                       //插入元素
    l1.insert(1);                       //插入元素
    assert(l1.size() == 3);             //取出大小
    assert(l1.max_size() >= l1.size()); //最大容量
    assert(l1.count(2) == 1);           //判断是否有2
    assert(l1.count(4) == 0);           //判断是否有4
    for_each(l1.begin(), l1.end(), [](int v) { cout << v << ","; });
    cout << endl;
    for_each(l1.cbegin(), l1.cend(), [](int v) { cout << v << ","; });
    cout << endl;
    l1.clear();
    assert(l1.empty() == true);
    l1.emplace(3); //就地构建，不执行拷贝和赋值，效率高一些
    assert(l1.size() == 1);
    l1.erase(l1.begin());
    assert(l1.empty());
    l1.emplace_hint(l1.begin(), 5); //具体含义待明确
    set<int> l2;
    l2.insert(4);
    l1.swap(l2);
    assert(l1.find(4) != l1.end());
    assert(l2.find(5) != l2.end());

    multiset<int> ll1;

    return 0;
}