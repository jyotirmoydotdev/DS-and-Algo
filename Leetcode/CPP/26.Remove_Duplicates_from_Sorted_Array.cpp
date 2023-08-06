#include<iostream>
#include<vector>
using namespace std;

// Solution <--- Start --->
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int i=0;
        if (nums.size()==0){return -1;}
        for(int j=1;j<nums.size();j++){
            if(nums[i]!=nums[j]){
                i++;
                nums[i]=nums[j];
            }
        }
        return i+1;
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
        vector<int> vec={0,0,1,1,1,2,2,3,4,4,5};
        vector<int> OldVec=vec;
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
        int k=s.removeDuplicates(vec);
        cout<<"Input : ";printVec(OldVec);
        cout<<"Output: ";printVec(vec,k);
        cout<<"k : "<<k<<endl;
    }
    return 0;
}