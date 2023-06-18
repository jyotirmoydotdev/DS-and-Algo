#include<iostream>
#include<vector>
using namespace std;

int main(){
	int n;
	cin>>n;
	std::vector<int>arr;


	for (int i=1;i<=n;i++){
		int input;
		cin>>input;
		arr.push_back(input);
	}
	cout<<arr<<endl;
	return 0;
}
