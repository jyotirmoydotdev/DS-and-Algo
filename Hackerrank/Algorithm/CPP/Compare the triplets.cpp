#include<iostream>
using namespace std;

int main(){
    int arra[3];
    int arrb[3];
    for (int i=0;i<3;i++){
        cin>>arra[i];
    }
    for (int i=0;i<3;i++){
        cin>>arrb[i];
    }
    int a=0,b=0;
    for (int i=0;i<3;i++){
        if (arra[i]>arrb[i]){
            a=a+1;
        }
        else if(arra[i]<arrb[i]){
            b=b+1;
        }
    }
    cout<<a<<" "<<b<<endl;
    return 0;
}