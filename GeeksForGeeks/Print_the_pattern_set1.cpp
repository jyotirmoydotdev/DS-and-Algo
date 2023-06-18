#include<iostream>
using namespace std;

// 1 code tried but failed
/*
void printpat(int n){
    for (int i=n;i>0;i--){
        for(int j=i;j>0;j--){
            for(int h=i;h>0;h--){
                cout<<j<<" ";
            }
        }
        cout<<"$";
    }
}
*/
void printpat(int n){
    for (int i=n;i>0;i--){
        for (int j=n;j>0;j--){
            for (int h=i;h>0;h--){
                cout<<j<<" ";
            }
        }
        cout<<"$";
    }
}

int main(){
    /*int t;
    cin>>t;
    while(t--){*/
        int n;
        cin>>n;
        printpat(n);
        cout<<endl;
    //}
}