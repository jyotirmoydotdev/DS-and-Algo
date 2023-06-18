#include<iostream>
#include<climits>
using namespace std;

int main(){
    int n=5;
    int arr[n];

    for (int i=0;i<n;i++){
        cin>>arr[i];
    }
    long long int ans[5];
    long long int max=LONG_MIN, min=LONG_MAX;
    //-------------------------------
    for (int i=0;i<n;i++){
        long long int temp=0;
        for (int j=0;j<n;j++){
            if (i!=j){
                temp=temp+arr[j];
            }
        }
        ans[i]=temp;
    }
    //------------------------------
    for(int i=0;i<5;i++){
        if (ans[i]>max){
		max=ans[i];
		}
		if (ans[i]<min){
		min=ans[i];
		}
    }
    cout<<min<<" "<<max<<endl;
    return 0;
}