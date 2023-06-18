#include<iostream>
using namespace std;

int main(){
    int n;
    cin>>n;
    int arr[n];

    for (int i=0;i<n;i++){
        cin>>arr[i];
    }

    float neg,pos,zero;
    for (int i=0;i<n;i++){
        if (arr[i]==0){
            zero=zero+1;
        }
        else if (arr[i]>0){
            pos=pos+1;
        }
        else{
            neg=neg+1;
        }
    }

    pos=pos/n;
    neg=neg/n;
    zero=zero/n;

    cout<<pos<<endl<<neg<<endl<<zero<<endl;
    return 0;
}