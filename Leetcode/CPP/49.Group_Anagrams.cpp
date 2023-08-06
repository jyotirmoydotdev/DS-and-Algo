#include<iostream>
#include<vector>
#include<map>
using namespace std;

vector<vector<string>> groupAnagrams(vector<string>& strs) {
    vector<vector<string>> ans;
      map<string, vector<string>> mp;
      for (string s : strs) {
        string sorted_s = s;
        sort(sorted_s.begin(), sorted_s.end());
        mp[sorted_s].push_back(s);
      }
      for (auto it : mp) {
        ans.push_back(it.second);
      }
    return ans;
}

int main(){
    vector<string> vec={"eat","tea","tan","ate","nat","bat"};
    vector<vector<string>> ans=groupAnagrams(vec);
    for (const auto& group : ans ){
      cout<<"[";
      for (const auto& word: group ){
        cout<<word<<", ";
      }
      cout<<"]"<<endl;
    }
    return 0;
}