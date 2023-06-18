#include<iostream>
using namespace std;
class Solution{
    public:
    void printNos(int n){
        for (int i=1;i<=n;i++){
            cout<<i<<" ";
        }
    }
};

int main(){
    int n;
    cin>>n;
    Solution ob;
    ob.printNos(n);
    cout<<endl;
}