//Input Format

//Input will contain four integers - a,b,c,d , one per line.

//Output Format

//Return the greatest of the four integers.
//PS: I/O will be automatically handled.

//There are two way with funtion and without funtion

//with funtion

#include<iostream>
using namespace std;

int largest(int a, int b, int c, int d){
	int largest;
	largest=max(max(a,b),max(c,d));
	return largest;
}

int main(){

	int a,b,c,d;
	cin>>a>>b>>c>>d;

	cout<<largest(a,b,c,d)<<endl;

	return 0;
}

//----whitout funtion-----
//
//if (a>b and a>c and a>d){
//	cout<<a<<endl;
//}
//else if (b>a and b>c and b>d){
//	cout<<b<<endl;
//}
//else if (c>a and c>b and c>d){
//	cout<<c<<endl;
//else {
//	cout<<d<<endl;
//}
