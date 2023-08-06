#include<iostream>
#include<string>
using namespace std;

int romanToInt(string s) {
        int n=s.length();
        char str[n+1];

        strcpy(str,s.c_str());

        long long int num=0;
        for (int i=0;i<n+1;i++){
            if(s[i] == 'I' && s[i+1] == 'V'){
                num += 4;
                i++;
            }
            else if(str[i] == 'I' && s[i+1] == 'X'){
                num += 9;
                i++;
            }
            else if(str[i] == 'X' && s[i+1] == 'L'){
                num += 40;
                i++;
            }
            else if(str[i] == 'X' && s[i+1] == 'C'){
                num += 90;
                i++;
            }
            else if(str[i] == 'C' && s[i+1] == 'D'){
                num += 400;
                i++;
            }
            else if(str[i] == 'C' && s[i+1] == 'M'){
                num += 900;
                i++;
            }
            else if (str[i]=='I'){
                num=num+1;
            }
            else if(str[i]=='V'){
                num=num+5;
            }
            else if(str[i]=='X'){
                num=num+10;
            }
            else if(str[i]=='L'){
                num=num+50;
            }
            else if(str[i]=='C'){
                num=num+100;
            }
            else if(str[i]=='D'){
                num=num+500;
            }
            else if(str[i]=='M'){
                num=num+1000;
            }
            else{
                break;
            }
        }
        return num;
    }

int main(){
    string rn="III";
    int ans=romanToInt(rn);
    cout<<ans<<endl;
    return 0;
}