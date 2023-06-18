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
void reverse(node* &head){
    if (head==NULL || head->next==NULL){
        return;
    }
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
void printList(node* head){
    while(head!=NULL){
        cout<<head->data<<"->";
        head=head->next;
    }
    cout<<"NULL"<<endl;
}
int main(){
    node* list=NULL;
    insert(list,1);
    insert(list,2);
    insert(list,3);
    insert(list,4);
    insert(list,5);
    cout<<"Before : ";
    printList(list);
    reverse(list);
    cout<<"After  : ";
    printList(list);
    return 0;
}