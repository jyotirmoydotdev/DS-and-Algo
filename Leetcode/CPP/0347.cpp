#include<iostream>
#include<map>
#include<vector>

using namespace std;

vector<int> topKFrequent(vector<int>& nums, int k) {
        vector<int> ans;
        
        return ans;
}

int main(){
    vector<int> nums={1,1,1,2,2,3};
    int k=2;
    vector<int> ans= topKFrequent(nums,k);
    for (int i=0;i<ans.size();i++){
        cout<<ans[i]<<", ";
    }
    cout<<endl;
    return 0;
}