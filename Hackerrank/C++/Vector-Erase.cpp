#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;

int main(){
    long long int n;
    cin>>n;
    vector<int> v;
    for(int i=0;i<n;i++){
        long long int temp;
        cin>>temp;
        v.push_back(temp);
    }
    long long int q1;
    cin>>q1;
    v.erase(v.begin()+q1-1);
    long long int r1,r2;
    cin>>r1>>r2;
    v.erase(v.begin()+r1-1,v.begin()+r2-1);
    
    cout<<v.size()<<endl;
    for(int i=0;i<v.size();i++){
        cout<<v[i]<<" ";
    }
    return 0;
}