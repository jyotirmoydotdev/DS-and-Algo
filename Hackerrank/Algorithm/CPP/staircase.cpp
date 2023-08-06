#include<iostream>
using namespace std;

int main()
{
   int n;
   cin>>n;
   for (int i=0,z=n-1;i<=n-1;i++,z--){
                for (int space=z;space>=1;space--){
                        cout<<" ";
                }
                for (int j=0;j<=i;j++){
                        cout<<"#";
                }
                cout<<endl;
        }
        return 0;
}
