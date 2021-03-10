#pragma once

#define lest_FEATURE_AUTO_REGISTER 1
#include "lest.hpp"

#define CASE(name) lest_CASE(specification(), name)
#define SCENARIO(name) lest_SCENARIO(specification(), name)

extern lest::tests &specification();

// bye
