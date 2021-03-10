#include <fstream>
#include <iostream>

using namespace std;

template <class T>
struct is_numberic : std::is_arithmetic<typename std::remove_cv<typename std::remove_reference<T>::type>::type> {};

struct OutputStream {
    std::ostream &os;
    OutputStream(std::ostream &os) : os(os) {}
    OutputStream(OutputStream &os) : os(os.os) {}

    template <class T>
    OutputStream &write(const T &t) {
        if (os) {
            os.write((const char *)&t, sizeof(T));
        }
        return *this;
    }

    template <class T>
    OutputStream &write(const T *t, size_t n) {
        if (os) {
            os.write((const char *)t, sizeof(T) * n);
        }
        return *this;
    }
    operator bool() const { return !os.fail(); }
};

// int,float,double
template <class T>
static typename std::enable_if<is_numberic<T>::value, OutputStream>::type &
operator<<(OutputStream &os, const T &t) {
    return os.write(t);
}

int main() {
    cout << "streamDemo" << endl;
    remove("test.bin");
    ofstream output;
    output.open("test.bin", std::ios::binary);
    int i = 1;
    output << i;
    OutputStream o(output);
    o << i;
    output.close();
    return 0;
}