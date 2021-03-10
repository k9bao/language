#include "demo1.hpp"
#include "test.h"

CASE("funadd test [tag1]") {
    EXPECT(add(1, 2) == 3);
}

CASE("classadd test [tag1]") {
    App app1;
    EXPECT(app1.add(1, 2) == 3);
    EXPECT_NOT(false);
}

CASE("expectThrowAs test [tag2]") {
    EXPECT_THROWS_AS(throwfun(), std::invalid_argument);
}

CASE("ignore test [.ignore tag]") {
    EXPECT(1 + 2 == 3);
}

CASE("hide test [hide]") {
    EXPECT(1 + 2 == 3);
}