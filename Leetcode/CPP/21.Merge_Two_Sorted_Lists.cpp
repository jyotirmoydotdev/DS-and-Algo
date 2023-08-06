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
    }else{
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
        cout<<head->data<<" ";
        head=head->next;
    }
    cout<<"NULL"<<endl;
}
node* merge(node* head1,node* head2){
    node* ans=new node();
    node* curr=ans;
    while (head1 && head2)
    {   
        if (head1->data<=head2->data){
            curr->next=head1;
            head1=head1->next;
        }
        else{
            curr->next=head2;
            head2=head2->next;
        }
        curr = curr->next;
    }
    curr->next = head1 ? head1 : head2;
    return ans->next;
}
int main(){
    node* head1=NULL;
    insert(head1,1);
    insert(head1,2);
    insert(head1,4);
    node* head2=NULL;
    insert(head2,1);
    insert(head2,3);
    insert(head2,4);
    merge(head1,head2);
    print(head1);
    return 0;
}