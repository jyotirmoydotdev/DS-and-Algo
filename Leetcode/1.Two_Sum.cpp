#include<iostream>
#include<vector>
using namespace std;

// Solution <--- Start --->
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> ans;
        for(int i=0;i<nums.size();i++){
            for(int j=i+1;j<nums.size();j++){
                if((nums[i]+nums[j])==target){
                    ans.push_back(i);
                    ans.push_back(j);
                    return ans;
                }
            }
        }
    return ans; 
    }
};
// Solution <--- End --->

void printVec(vector<int> vec){
    for(int i=0;i<vec.size();i++){
       cout<<vec[i]<<" ";
    }cout<<endl;
}

int main(){
    int testcase=1; 
    // we will use only one test case;
    //cin>>t;
    for(int i=1;i<=testcase;i++){
        vector<int> vec={1,2,3,4,5,6,7,8};
        int target=7;
        /*
        vector<int> vec;
        int target;
        cout<<"Enter size : ";
        int n;
        cin>>n;
        for(int i=0;i<n;i++){
            int temp;
            cin>>temp;
            vec.push_back(temp);
        }
        cout<<"Enter the target : ";
        cin>>target;
        */
        Solution s;
        vector<int> ans=s.twoSum(vec,target);
        cout<<"Input : ";printVec(vec);
        cout<<"Target: "<<target<<endl;
        cout<<"Output: ";printVec(ans);
    }
    return 0;
}