#include<iostream>
#include<vector>
using namespace std;

// Solution <--- Start --->
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        if((nums.size())==0){cout<<"Empty"<<endl;return false;}
        for(int i=0;i<nums.size();i++){
            for(int j=i+1;j<nums.size();j++){
                if(nums[i]==nums[j]){
                    return true;
                }
            }
        }
        return false;
    }
};
// Solution <--- End --->

void printVec(vector<int> vec){
    for(int i=0;i<vec.size();i++){
       cout<<vec[i]<<" ";
    }cout<<endl;
}
void printVec(vector<int> vec,int size){
    for(int i=0;i<size;i++){
       cout<<vec[i]<<" ";
    }
    for(int i=0;i<vec.size()-size;i++){
        cout<<"_ ";
    }
    cout<<endl;
}

int main(){
    int testcase=1; 
    // we will use only one test case;
    //cin>>t;
    for(int i=1;i<=testcase;i++){
        vector<int> vec={1,2,3,1};
        /*
        vector<int> vec;
        cout<<"Enter size : ";
        int n;
        cin>>n;
        for(int i=0;i<n;i++){
            int temp;
            cin>>temp;
            vec.push_back(temp);
        }
        */
        Solution s;
        bool ans=s.containsDuplicate(vec);
        cout<<"Input : ";printVec(vec);
        cout<<"Output: ";
        if(ans){cout<<"true"<<endl;}else{cout<<"False"<<endl;}
    }
    return 0;
}