#include<iostream>
using namespace std;
class node{
    public:
    int data;
    node* next;
    node(int val){
        data=val;
        next=NULL;
    }
};
void insert(node* &head,int val){
    if (head==NULL){
        node* temp=new node(val);
        temp->next=head;
        head=temp;
    }
    else{
        node* temp=new node(val);
        node* last=head;
        while(last->next!=NULL){
            last=last->next;
        }
        last->next=temp;
        temp->next=NULL;
    }
}
void print(node* head){
    while(head!=NULL){
        cout<<head->data<<"->";
        head=head->next;
    }
    cout<<"NULL"<<endl;
}
void reverse(node* &head){
    node* previous=NULL;
    node* current=head;
    node* next=NULL;
    while(current!=NULL){
        next=current->next;
        current->next=previous;
        previous=current;
        current=next;
    }
    head=previous;
}
bool palindrome(node* head){
    node* slow=head;
    node* fast=head;
    while(fast->next && fast->next->next){
        slow=slow->next;
        fast=fast->next->next;
    }
    slow=slow->next;
    reverse(slow);
    while(slow!=NULL){
        if (head->data!=slow->data){
            return false;
        }
        head=head->next;
        slow=slow->next;
    }
    return true;
}
int main(){
    node* list=NULL;
    insert(list,1);
    insert(list,1);
    insert(list,2);
    insert(list,1);
    insert(list,1);
    if (palindrome(list)){
        cout<<"Palindrome"<<endl;
    }
    else{
        cout<<"Not Palindrome"<<endl;
    }
    return 0;
}