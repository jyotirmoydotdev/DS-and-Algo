#include<iostream>
using namespace std;

int main(){
	int num;
	cin>>num;

	if (num%2==1 and num>0){
		cout<<"Weird"<<endl;
	}
	else if(num%2==0 and 2<=num and num<=5 ){
		cout<<"Not Weird"<<endl;
	}
	else if(num%2==0 and 6<=num and num<=20){
		cout<<"Weird"<<endl;
	}
	else if(num%2==0 and num>20 and num>0){
		cout<<"Not Weird"<<endl;
	}
	else{
		cout<<"enter a vaild number"<<endl;
	}
	return 0;
}


