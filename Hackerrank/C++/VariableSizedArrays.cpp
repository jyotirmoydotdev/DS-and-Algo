#include<iostream>
#include<vector>
using namespace std;

int main(){
    int n,q; // n refer to the vec[n] and q is refer to quries
    cin>>n>>q;
    // making 2d vector and limiting the vec[n] with n size
    vector<vector<int> > vec(n); 

    // Entering the element
    for(int i=0;i<n;i++){
        int m;
        cin>>m;
        vec[i]=vector<int>(m); // making a new vector in vec[n]
        for (int j=0;j<m;j++){
            int x;
            cin>>x;
            vec[i][j]=x;
        }
    }

    // Queries
    for (int i=0;i<q;i++){
        int x,y;
        cin>>x>>y;
        cout<<vec[x][y]<<endl;
    }
    return 0;
}