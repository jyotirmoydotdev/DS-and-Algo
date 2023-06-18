#include<iostream>
#include<vector> 
using namespace std;
int linearSearch(int arr[],int size,int key){
  for (int i=0;i<size;i++){
    if (arr[i]==key){
      return i;
    }
  }
  return -1;
}

int search(vector<int> vec,int key){
  int start=0;
  int end=vec.size() - 1;
  while(start<=end){
    int mid=(start+end)/2;
    if (key==vec[mid]){
      return mid;
    }
    else if(vec[mid]>key){
      end=mid-1;
    }
    else{
      start=mid+1;
    }
  }
  return start;
}

int main(){
  vector<int> vec={1,3,5,6};
  cout<<search(vec,7)<<endl;
  return 0;
}