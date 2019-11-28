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

func check_for_single_child(node *Node)bool{
  fmt.Println("--->NODE-->1", node.Value)

  if node.Link[0] == nil{
  }else{
    check_for_single_child(node.Link[0])
    fmt.Println("--->NODE-->2", (*node).Value)
  }

  if node.Link[1] == nil{
    if node.Link[0] != nil{
      fmt.Println("found a single value-->", node.Value)
      fmt.Println("found a single value-->", node)
      node.Link[0] = nil
      fmt.Println("found a single value-->", node)
      return true
    }
  }else{
    check_for_single_child(node.Link[1])
  }
  return false
}

func check_for_single_child1(node *Node)bool{
  fmt.Println("Value: ", node.Value, "Left: ") //, node.Link[0].Value, node.Link[1].Value)
  fmt.Println("Value:-->", node.Link[0])
  if node.Link[0] == nil{
    fmt.Println("Returning false")
    return false
  }else if node.Link[1] == nil{
    return true
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

  fmt.Println("checking for signle child--->", node.Link[0].Value)
  fmt.Println("*******************1**********************", node.Value)
  if check_for_single_child(node.Link[0]){
    return
  }

  fmt.Println("*******************2**********************", node.Value)
  fmt.Println("*******************2**********************", node)
  if check_for_single_child(node.Link[1]){
    return
  }
  fmt.Println("value--->", node.Value)
  last_left_node, left_depth   := depth_of_left_nodes(node.Link[0], 0)
  last_right_node, right_depth := depth_of_right_nodes(node.Link[1], 0)

  fmt.Println("Last_LeftNOde-->", last_left_node.Value)
  fmt.Println("Last_RightNOde-->", last_right_node.Value)


  if left_depth > right_depth{
    remove_last_node(node.Link[0], last_left_node)
  }else{
    remove_last_node(node.Link[1], last_right_node)
  }
}

func (tree *Tree)pop(){
  if tree.Root == nil{
    return
  }
  tree.Root.remove()
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
