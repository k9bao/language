#include <fstream>
#include <iostream>
#include <stdlib.h>
#include <time.h> 

using namespace std;
int main(){
    ofstream output;    
    output.open("test.txt", std::ios::binary);
    srand((unsigned)time(NULL)); 
    for(int i = 0; i < 10;i++ ) 
            output << rand() << endl;
    output << endl; 
    return 0;
}