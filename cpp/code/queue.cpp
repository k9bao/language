#include <iostream>
#include <queue>

using namespace std;
//queue 先进先出，可以取到二端的内容
int main() {
    queue<int> l1;           //支持双向插入和删除
    l1.push(2);              //从后边插入元素
    l1.push(3);              //从后边插入元素
    l1.push(1);              //从后边插入元素
    assert(l1.front() == 2); //取出前边元素
    assert(l1.back() == 1);  //取出后边元素
    assert(l1.size() == 3);  //获取list大小
    l1.pop();                //只能弹出头部元素
    l1.pop();                //只能弹出头部元素
    assert(l1.front() == l1.back() && l1.front() == 1 && l1.size() == 1);
    while (l1.empty() == false) //不支持clear
        l1.pop();
    assert(l1.empty() == true);
    l1.emplace(3); //就地构建，不执行拷贝和赋值，效率高一些
    assert(l1.size() == 1 && l1.front() == 3);

    queue<int> l2;
    l2.emplace(5);
    l1.swap(l2);
    assert(l1.front() == 5 && l2.front() == 3);

    return 0;
}