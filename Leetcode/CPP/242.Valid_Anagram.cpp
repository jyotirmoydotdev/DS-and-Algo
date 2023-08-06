#include<iostream>
#include<string>
#include <vector>
#include<unordered_map>
#include<map>
using namespace std;

bool isAnagram(string s, string t){
    if (s.size()!=t.size()){
        return false;
    }
    vector<char> s_vec(s.begin(), s.end());
    vector<char> t_vec(t.begin(), t.end());

    map<char,int> m;
    map<char,int> n;
    
    for (int i=0;i<s.size();i++){
        m[s_vec[i]]++;
    }
    for (int i=0;i<t.size();i++){
        n[t_vec[i]]++;
    }
    if (m==n){
        return true;
    }
    return false;
}

int main(){
    string s="rat";
    string t="artt";
    bool ans=isAnagram(s,t);
    if (ans){cout<<"isAnagram: Yes"<<endl;}
    else{cout<<"isAnagram: No"<<endl;}
    return 0;
}
