#include<iostream>
using namespace std;

int main(){
    int n;
    cin>>n;
    int arr[n];

    for (int i=0;i<n;i++){
        cin>>arr[i];
    }
    int arra[n];
    for (int i=0;i<n;i++){
        arra[i]=arr[i];
    }
    for (int i=0;i<n;i++){
        if (arr[i]>=38 and (arr[i]%5)!=0){
            int r=arr[i]%5;
            int sub=5-r;
            arr[i]=sub+arr[i];
        }
    }
    int ans[n];
    for (int i=0;i<n;i++){
        if ((arr[i]-arra[i])<3){
            ans[i]=arr[i];
        }
        else if((arr[i]-arra[i])==3){
            ans[i]=arra[i];
        }
        else if((arr[i]-arra[i])>3){
            ans[i]=arra[i];
        }
        else{
            ans[i]=arra[i];
        }
    }
    for (int i=0;i<n;i++){
        cout<<ans[i]<<endl;
    }
    return 0;
}