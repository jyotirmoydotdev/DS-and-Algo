#include<iostream>
using namespace std;

void a(int n1,int n2){
	int plus,m_plus;
	plus=n1+n2;
	m_plus=0;
	if (plus==0){
		m_plus=n2;
	}
	else{
		m_plus=plus;
	}
	cout<<m_plus<<endl;
	return ;
}

void b(int n1,int n2){
	int sub,m_sub;
	sub=n1-n2;
	m_sub=0;

	if (sub<0){
		m_sub=sub*-1;
	}
	else {
		m_sub=sub;
	}
	cout<<m_sub<<endl;
	return ;
}

int main(){
	int n1,n2;
	cin>>n1>>n2;

	a(n1,n2);
	b(n1,n2);

	return 0;
}
