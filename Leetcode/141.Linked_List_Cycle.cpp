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
void insertCycle(node* &head,int pos){
    node* last=head;
    node* posPointer=NULL;
    int counter=1;
    while(last->next!=NULL){
        if (last->data==pos){
            posPointer=last;
        }
        last=last->next;
    }
    last->next=posPointer;
}
bool hasCycle(node* head){
    if (head==NULL || head->next==NULL){
        return true;
    }else{
        node* slow=head;
        node* fast=head;
        while(fast!=NULL && fast->next!=NULL){
            slow=slow->next;
            fast=fast->next->next;
            if (slow==fast){
                return true;
            }
        }
        return false;
    }
}
int main(){
    node* list=NULL;
    insert(list,1);
    insert(list,2);
    insert(list,3);
    insert(list,4);
    insert(list,5);
    cout<<"Befor :"<<hasCycle(list)<<endl;
    insertCycle(list,3);
    cout<<"After :"<<hasCycle(list)<<endl;
    return 0;
}