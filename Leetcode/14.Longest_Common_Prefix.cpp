#include<iostream>
#include<vector>
using namespace std;

string prefix(vector<string> strs){
    int small=INT_MAX;
    for (int i=0;i<strs.size();i++){
        if (small>strs[i].length()){
            small=i;
        }
    }
    string ans="";
    for (int i=0;i<strs[small].length();i++){
        char Letter=strs[small][i];
        for (int j=0;j<strs.size();j++){
            if ( Letter !=strs[j][i]){
                return ans;
            }
        }
        ans=ans+Letter;
    }
    return ans;
}

int main(){
    vector<string> strs={"flower","flow","floght"};
    string ans=prefix(strs);
    cout<<"Prefix: "<<ans<<endl;
    return 0;
}