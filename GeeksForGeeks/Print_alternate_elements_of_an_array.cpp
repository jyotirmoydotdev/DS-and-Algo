#include<iostream>
using namespace std;

void al(int ar[],int n){
    for (int i=0;i<n;i++){
        if (i%2==0 or i==0){
           cout<<ar[i]<<" ";
        }
    }
}

int main(){
    int n;
    cin>>n;
    int ar[n];
    for (int i=0;i<n;i++){
        cin>>ar[i];
    }
    al(ar, n);
    return 0;
}