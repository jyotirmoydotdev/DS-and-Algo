#include<iostream>
#include<vector>
using namespace std;
vector<int> plusOne(vector<int> digits){
    int size=digits.size()-1;
    for (int i=size;i>=0;i--){
        if (digits[i]==9){
            digits[i]=0;
        }
        else{
            digits[i]++;
            return digits;
        }
    }
    digits.insert(digits.begin(), 1);
    return digits;
}

int main(){
    vector<int> digits={7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6};
    for (int i=0;i<digits.size();i++){
        cout<<digits[i]<<" ";
    }cout<<endl;
    vector<int> vec=plusOne(digits);
    for (int i=0;i<vec.size();i++){
        cout<<vec[i]<<" ";
    }
    cout<<endl;
    return 0;
}