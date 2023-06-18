//Input Format

//A single integer,n .

//Constraints

//1<=n<=10**9

//Output Format

//If 1<=n<=9 , then print the lowercase English word corresponding to the number (e.g., one for 1 ,two for 1, etc.); otherwise, print Greater than 9.

#include<iostream>
using namespace std;

int main(){
	int n;
	cin>>n;

	string l[10]={"zero","one","two","three","four","five","six","seven","eight","nine"};

	if (0<=n and n<=9){
		cout<<l[n]<<endl;
	}
	else if (n>9){
		cout<<"Greater than 9"<<endl;
	}
	else {
		cout<<"Lesser than 0"<<endl;
	}

	return 0;

}
