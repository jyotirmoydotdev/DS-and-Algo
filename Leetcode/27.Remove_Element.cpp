#include<iostream>
#include<vector>
using namespace std;

// Solution <--- Start --->
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        /*
        int i=0;
        if((nums.size())==0){return -1;}
        for(int j=0;j<nums.size();j++){
            if(nums[j]!=val){
                nums[i]=nums[j];
                i++;
            }
        }
        return i;
        */
        int i=0;
        while(i<nums.size()){
            if(nums[i] == val) {
                nums.erase(nums.begin() + i);
            } else {
                i++;
            }
        }
        return nums.size();
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
        vector<int> vec={0,1,2,2,3,0,4,2};
        vector<int> OldVec=vec;
        int val=2;
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
        int val;
        cin>>val;
        */
        Solution s;
        int size=s.removeElement(vec,val);
        cout<<"Input : ";printVec(OldVec);
        cout<<"Output: ";printVec(vec,size);
        cout<<"Size  : "<<size<<endl;
    }
    return 0;
}