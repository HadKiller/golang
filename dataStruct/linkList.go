package main


//默认使用头结点
//头指针
type linkList struct {
	linkHead *node
}
//节点
type node struct {
	next *node
	data interface{}
}

func (n *node)InitLink(){
	n.next=nil
	//头结点data存储长度
	n.data=1
}
func (this*node)insertTail(n *node){

}

func (this*node)Add(n *node) *node{
	this.GetLinkLength()
	return n
}

func(i*node)reverse(){
	if i==nil||i.next==nil{
		return
	}
	head:=i
	first:=i.next
	for first.next!=nil{
		temp:=first.next
		first.next=temp.next
		temp.next=head.next
		head.next=temp
	}

}
//获得长度,头结点存储尾节点
func(this *node)GetLinkLength()int{
	if this==nil{
		return -1
	}
	length:=1
	head:=this
	for head!=nil{
		length++
		head=head.next
	}
	this.data=head
	return length
}



