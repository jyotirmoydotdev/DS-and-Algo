// 21. Merge Two Sorted Lists
#include<iostream>
#include<vector>
using namespace std;
void merge(vector<int> &nums1,int m,vector<int> nums2,int n){
    for(int i=0;i<n;i++){
        nums1[m++]=nums2[i];
    }
    sort(nums1.begin(),nums1.end());
}
int main(){
    vector<int> vec1={1,2,3,0,0,0};
    int vec1Size=3;
    vector<int> vec2={2,5,6};
    int vec2Size=vec2.size();
    merge(vec1,vec1Size,vec2,vec2Size);
    for(int i=0;i<vec1.size();i++){
        cout<<vec1[i]<<" ";
    }cout<<endl;
    return 0;
}