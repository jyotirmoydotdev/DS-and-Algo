/*

Writing this because i know i may forget this
Problem : find how many tall candels are there

*/

#include<iostream>
#include<climits>
using namespace std;

int main(){
    int n;
    cin>>n;
    int arr[n];

    //-------Inserting All candels Height---------
    for (int i=0;i<n;i++){
        cin>>arr[i];
    }

    //------Finding the tallest candel-----------
    int mx=INT_MIN;
    for (int i=0;i<n;i++){
		if (arr[i]>mx){
			mx=arr[i];
		}
    }

    //-----Finding how many tall candels are there-
    int ans=0;
    for (int i=0;i<n;i++){
        if (mx==arr[i]){
            ans++;
        }
    }

    cout<<ans<<endl;
    return 0;
}