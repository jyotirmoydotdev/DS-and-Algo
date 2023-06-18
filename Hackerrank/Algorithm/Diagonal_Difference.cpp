#include<iostream>
#include<math.h>
using namespace std;


int main(){
    int n;
    cin>>n;

    int arr[n];

    int l=0,r=0;
    for (int i=0;i<n;i++){
        for (int j=0;j<n;j++){
            int temp;
            cin>>temp;
            if (i==j){
                l = l + temp;
            }
            if (i+j==(n-1)){
                r=r+temp;
            }
        }
    }
    int ans=l-r;
    ans =abs(ans);
    cout<<ans<<endl;
    return 0;
} 
