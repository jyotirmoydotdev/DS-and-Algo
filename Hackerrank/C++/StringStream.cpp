#include <sstream>
#include <vector>
#include <iostream>
using namespace std;

vector<int> parseInts(string str) {
	stringstream s(str);
    vector<int> vec;
    char ch;
    do{
        int temp;
        s>>temp;
        vec.push_back(temp);
    }while(s>>ch);
    return vec;
}

int main() {
    string str;
    cin >> str;
    vector<int> integers = parseInts(str);
    for(int i = 0; i < integers.size(); i++) {
        cout << integers[i] << "\n";
    }
    
    return 0;
}