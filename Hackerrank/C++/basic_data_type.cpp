//INPUT FORMAT

//Input consists of the following space-separated values: int, long, char, float, and double, respectively.

//OUTPUT FORMAT

//Print each element on a new line in the same order it was received as input. Note that the floating point value should be correct up to 3 decimal places and the double to 9 decimal places.


#include<iostream>
using namespace std;

int main(){

	int a;
	long  b;
	char c;
	float d;
	double e;


	cin>>a>>b>>c>>d>>e;
	cout<<a<<endl<<b<<endl<<c<<endl;
	printf("%.3f",d);
	cout<<endl;
	printf("%.9lf",e);
	cout<<endl;

	return 0;
}
