#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;

int main(){
    long long int n;
    cin>>n;
    vector<long long int> v;
    for(int i=0;i<n;i++){
        long long int temp;
        cin>>temp;
        v.push_back(temp);
    }

    long long int q;
    cin>>q;
    vector<long long int> y;
    for(int i=0;i<q;i++){
        long long int temp;
        cin>>temp;
        y.push_back(temp);
    }

    for (int i=0;i<q;i++){
        bool loop;
        for(int j=0;j<n;j++){
            if(v[j]==y[i]){
                cout<<"Yes "<<j+1<<endl;
                loop=true;
                break;
            }
            else{
                loop=false;
            }
        }
        if (loop==false){
           for(int x=0;x<v.size();x++){
                if (y[i]<v[x]){
                    cout<<"No "<<x+1<<endl;
                    break;
                }
            }
        }
    }
    return 0;
}