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
    node(){};
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
    }cout<<"NULL"<<endl;
}
node* remove(node* head){
    node* curr=head;
    while(curr!=NULL && curr->next!=NULL){
        if (curr->data == curr->next->data){
            curr->next=curr->next->next;
        }
        else{
            curr=curr->next;
        }
    }
    return head;
}
int main(){
    node* list=NULL;
    insert(list,1);
    insert(list,1);
    insert(list,2);
    insert(list,4);
    insert(list,5);
    remove(list);
    print(list);
    return 0;
}