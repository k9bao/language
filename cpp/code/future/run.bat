rm -rf build
cmake -B build -G Ninja
cmake --build build

::在build目录生成可执行程序
cd build
future-package-task
future-promise
cd ..