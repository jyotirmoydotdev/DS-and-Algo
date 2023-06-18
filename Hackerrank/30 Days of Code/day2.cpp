#include<iostream>
#include<cmath>
using namespace std;

int main(){
	double meal_cost;
	double tip_p,tax_p;

	cin>>meal_cost;
	cin>>tip_p;
	cin>>tax_p;
	
	double tip,tax;
	tip=(tip_p*meal_cost)/100;
	tax=(tax_p*meal_cost)/100;
	
	double total_meal_cost=meal_cost+tip+tax;

	total_meal_cost=round(total_meal_cost);
	cout<<total_meal_cost<<endl;

	return 0;
}

