#include<iostream>
using namespace std;

int main(){
    int a[3],b[3];
    
    for(int i=0;i<3;i++){
        cin>>a[i];
    }
    for(int i=0;i<3;i++){
        cin>>b[i];
    }
    
    int ar=0,br=0;
    for (int i=0;i<3;i++){
       if (a[i]>b[i]){
           ar=ar+1;
       }
       else if(a[i]<b[i]){
           br=br+1;
       }
       else{
           ar=ar+0;
           br=br+0;
       }
    }
    cout<<ar<<" "<<br<<endl;
    return 0;
}
