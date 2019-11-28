package main

//import("fmt")
import (
  "encoding/json"
  "fmt"
  "log"
)

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
    fmt.Println("Returning True")
    return true
  }
  return false
}

func (node *Node)insert(value int){
  fmt.Println("----Inserting----", node.Value)

  if node.Link[0] == nil{
    //node.Link[0] = &Node{Value: value, Parent: node}
    node.Link[0] = &Node{Value: value}
    fmt.Println("----Returning----")
    return
  }else if node.Link[1] == nil{
    //node.Link[1] = &Node{Value: value, Parent: node}
    node.Link[1] = &Node{Value: value}
    return
  }

  fmt.Println("Before check_empty_link--->1", value, node.Link[0].Value)
  if check_empty_link(node.Link[0]){
    node.Link[0].insert(value)
    return
  }

  fmt.Println("Before check_empty_link--->2", value)
  if check_empty_link(node.Link[1]){
    node.Link[1].insert(value)
    return
  }

  insert_to_left(node, value)
  return

}

/*
func (node *Node)insert1(value int){

  if node.L_child == nil{
    node.L_child = &Node{Value: value, }
    return
  }else if node.R_child == nil{
    node.R_child = &Node{Value: value}
    return
  }

  if node.count == 0 {
    node.L_child.insert(value)
  }else{
    node.R_child.insert(value)
  }
  return
}
*/

func (tree *Tree)push(value int){
  if tree.Root ==nil{
    tree.Root = &Node{Value: value,}
    return
  }
  tree.Root.insert(value)
  fmt.Println("**1**")
  //fmt.Println(tree.Root.Link[0].Value)
  fmt.Println("**2**")
  return
}

func main(){
  tree := Tree{}
  tree.push(1)
  tree.push(2)
  tree.push(3)
  tree.push(4)
  tree.push(5)
  tree.push(6)
  tree.push(7)
  tree.push(8)
  fmt.Println("***Pavilion***")
  data, err := json.MarshalIndent(tree, "", "  ")
  fmt.Println("***Pavilion***")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", data)
}
