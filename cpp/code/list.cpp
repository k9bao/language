#include <iostream>
#include <list>

using namespace std;

int main() {
    list<int> l1;                                                     //支持双向插入和删除
    l1.push_back(1);                                                  //向后插入
    l1.insert(l1.begin(), 3);                                         //在begin()后边插入元素
    l1.push_front(2);                                                 //向前插入
    for_each(l1.begin(), l1.end(), [](int &v) { cout << v << ","; }); //2,3,1,
    cout << endl;
    assert(l1.front() == 2); //取出前边元素
    assert(l1.back() == 1);  //取出后边元素
    assert(l1.size() == 3);  //获取list大小
    l1.pop_back();           //弹出末尾元素
    l1.pop_front();          //弹出头部元素
    assert(l1.front() == l1.back() && l1.front() == 3 && l1.size() == 1);
    l1.clear(); //清空元素
    assert(l1.empty() == true);
    l1.resize(3, 2); //2,2,2，重置为3个2
    assert(l1.size() == 3);
    l1.erase(l1.begin()); //删除首元素
    assert(l1.size() == 2);

    l1 = {5, 2, 1};
    list<int> l2(2, 4);           //4,4
    l1.merge(l2, greater<int>()); //合并降序list，原来list必须都是降序
    assert(l2.empty() == true);
    for_each(l1.begin(), l1.end(), [](int &v) { cout << v << ","; }); //5,4,4,2,1,
    cout << endl;

    l1 = {1, 2, 5};
    l2.assign(2, 4);                                      //4,4
    l1.merge(l2, [](int v1, int v2) { return v1 < v2; }); //合并升序list，原来list必须都是升序
    assert(l2.empty() == true);
    for_each(l1.begin(), l1.end(), [](int &v) { cout << v << ","; }); //4,4,1,2,5,
    cout << endl;

    return 0;
}