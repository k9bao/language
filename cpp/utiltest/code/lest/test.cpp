#include "test.h"

struct launch : lest::confirm {
    using confirm::confirm;
    launch &operator()(lest::test testing) {
        try {
            ++selected;
            testing.behaviour(output(testing.name));
            os << "[passed] " << testing.name << std::endl;
        } catch (lest::message const &e) {
            ++failures;
            report(os, e, output.context());
            os << "[failed] " << testing.name << std::endl;
        }
        return *this;
    }
    ~launch() {
        os << "== Run " << selected << " cases, "
           << "passed: " << (selected - failures) << " "
           << "failed: " << failures << "\n";
    }
};

lest::tests &specification() {
    static lest::tests tests;
    return tests;
}

int main(int argc, char **argv) {
    return lest::run<launch>(specification(), lest::texts(argv + 1, argv + argc));
}
