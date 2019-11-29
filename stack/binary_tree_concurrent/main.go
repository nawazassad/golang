package main

import("fmt"
      "time")
/*
import (
  "encoding/json"
  "fmt"
  "log"
)
*/

type Tree struct{
  Root *Node
}

type Node struct{
  Value interface{}
  Link [2]*Node
  Parent  *Node
  Count int
}

func insert_to_left(node *Node, value int){
  fmt.Println("---1---")
  if node.Link[0] == nil{
    node.Link[0] = &Node{Value: value}
    return
  }else{
    insert_to_left(node.Link[0], value)
    return
  }
  return
}

func check_empty_link(node *Node) bool{
  fmt.Println("---2---", node.Value)
  if node.Link[0] == nil || node.Link[1] == nil{
    return true
  }
  return false
}

func (node *Node)insert(value int){

  // insert when either of the left and right child is empty
  if node.Link[0] == nil{
    node.Link[0] = &Node{Value: value}
    return
  }else if node.Link[1] == nil{
    node.Link[1] = &Node{Value: value}
    return
  }

  // search tree for empty child on the left side of the Root
  if check_empty_link(node.Link[0]){
    node.Link[0].insert(value)
    return
  }

  // search tree for empty child on the right side of the Root
  if check_empty_link(node.Link[1]){
    node.Link[1].insert(value)
    return
  }

  // when there is no empty child, Blindly insert it to the left most child
  insert_to_left(node, value)
  return

}


func (tree *Tree)push(value int){
  if tree.Root ==nil{
    tree.Root = &Node{Value: value,}
    return
  }
  tree.Root.insert(value)
  return
}

func check_for_single_child(node *Node)bool{

  if node.Link[0] == nil{
  }else{
    check_for_single_child(node.Link[0])
  }

  // when there is no left child it will start parsing the right
  if node.Link[1] == nil{
    if node.Link[0] != nil{
      // when there left child but when there is a right child it will be removed
      node.Link[0] = nil
      return true
    }
  }else{
    check_for_single_child(node.Link[1])
  }
  return false
}

func depth_of_left_nodes(node *Node, left_depth int)(*Node, int){
  if node.Link[1] ==nil{
    return node, left_depth
  }else{
    last_node, count := depth_of_left_nodes(node.Link[1], left_depth+1)
    return last_node, count
  }
}

func depth_of_right_nodes(node *Node, right_depth int)(*Node, int){
  if node.Link[1] ==nil{
    return node, right_depth
  }else{
    last_node, count := depth_of_right_nodes(node.Link[1], right_depth+1)
    return last_node, count
  }
}

func remove_last_node(node *Node, last_node *Node){

  if node.Link[1] == last_node{
    fmt.Println("Popping element--->", node.Link[1].Value)
    node.Link[1] = nil
    return
  }else{
    remove_last_node(node.Link[1], last_node)
    return
  }
}



func (node *Node)remove(){

  // check for single child and remove it in the left side of the root
  if check_for_single_child(node.Link[0]){
    return
  }

  // check for single child and remove it in the right side of the root
  if check_for_single_child(node.Link[1]){
    return
  }

  // when there is no single child on either side of the Root
  // measuring the depth of both the side of the Root and getting the last node of the either side
  last_left_node, left_depth   := depth_of_left_nodes(node.Link[0], 0)
  last_right_node, right_depth := depth_of_right_nodes(node.Link[1], 0)

  // depending on the depth of the tree respect child is removed
  if left_depth > right_depth{
    remove_last_node(node.Link[0], last_left_node)
  }else{
    remove_last_node(node.Link[1], last_right_node)
  }
}

func (tree *Tree)pop(){
  if tree.Root == nil{
    return
  }else if tree.Root.Link[0] == nil && tree.Root.Link[1] == nil{
    tree.Root = nil
    return
  }
  tree.Root.remove()
}

func producers(tree *Tree, c1, c2 chan string, number int){
  for i:=1; i< number+1; i++{
    //fmt.Println("Pushgin-->", i)
    tree.push(i)
    c1 <- "pushed"
    <- c2
  }
  close(c1)
}


func consumers(tree *Tree){
  tree.pop()
}

func main(){
  tree := Tree{}
  var number int
  var c1 = make(chan string)
  var c2 = make(chan string)

  fmt.Println("Enter the number of producers you want: ")
  fmt.Scanf("%d", &number)

  fmt.Println("Now we Produce :")
  start := time.Now()

  go producers(&tree, c1, c2, number)
  for _ = range c1{
    consumers(&tree)
    //fmt.Println("Popping")
    c2 <- "consumed"
  }
  close(c2)

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}

/*
  tree.push(1)
  tree.push(2)
  tree.push(3)
  tree.push(4)
  tree.push(5)
  tree.push(6)
  tree.push(7)
  //tree.push(8)
  fmt.Println("***Pavilion***")
  data, err := json.MarshalIndent(tree, "", "  ")
  fmt.Println("***Pavilion***")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", data)

  tree.pop()


  data, err = json.MarshalIndent(tree, "", "  ")
  fmt.Println("***Pavilion***")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", data)
}

*/
