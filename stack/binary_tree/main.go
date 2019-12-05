package main

import("fmt"
        "time"
        "encoding/json"
        "log"
)

type Tree struct{
  Root *Node
}

type Node struct{
  Value interface{}
  Link [2]*Node
}

func insert_to_left(node *Node, value int){
  //fmt.Println("---1---")
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
  //fmt.Println("---2---", node.Value)
  if node.Link[0] == nil || node.Link[1] == nil{
    //fmt.Println("Returning True")
    return true
  }
  return false
}

func (node *Node)insert(value int){
  //fmt.Println("----Inserting----", node.Value)

  if node.Link[0] == nil{
    node.Link[0] = &Node{Value: value}
    //fmt.Println("----Returning----")
    return
  }else if node.Link[1] == nil{
    node.Link[1] = &Node{Value: value}
    return
  }

  //fmt.Println("Before check_empty_link--->1", value, node.Link[0].Value)
  if check_empty_link(node.Link[0]){
    node.Link[0].insert(value)
    return
  }

  //fmt.Println("Before check_empty_link--->2", value)
  if check_empty_link(node.Link[1]){
    node.Link[1].insert(value)
    return
  }

  insert_to_left(node, value)
  return

}


func (tree *Tree)Push(value int){
  if tree.Root ==nil{
    tree.Root = &Node{Value: value,}
    return
  }
  tree.Root.insert(value)
  //fmt.Println(tree.Root.Link[0].Value)
  return
}

func check_for_single_child(node *Node)bool{

  if node.Link[0] == nil{
  }else{
    check_for_single_child(node.Link[0])
  }

  if node.Link[1] == nil{
    if node.Link[0] != nil{
      //fmt.Println("Found Single element--->", node.Link[0].Value)
      //fmt.Println("Found Single element--->", node)
      //fmt.Println("Popping element-------->", node.Link[0].Value)
      node.Link[0] = nil
      //fmt.Println("Found Single element--->", node)
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
    //fmt.Println("Popping element--->", node.Link[1].Value)
    //fmt.Println("Popping element--->", node)
    node.Link[1] = nil
    //fmt.Println("Popping element--->", node)
    return
  }else{
    remove_last_node(node.Link[1], last_node)
    return
  }
}



func (node *Node)remove(){

  //fmt.Println("checking for signle child--->", node.Link[0].Value)
  //fmt.Println("*******************1**********************", node.Value)
  if check_for_single_child(node.Link[0]){
    return
  }

  //fmt.Println("*******************2**********************", node.Value)
  //fmt.Println("*******************2**********************", node)
  if check_for_single_child(node.Link[1]){
    return
  }
  //fmt.Println("value--->", node.Value)
  last_left_node, left_depth   := depth_of_left_nodes(node.Link[0], 0)
  last_right_node, right_depth := depth_of_right_nodes(node.Link[1], 0)

  //fmt.Println("Last_LeftNOde-->", last_left_node.Value)
  //fmt.Println("Last_RightNOde-->", last_right_node.Value)


  if left_depth > right_depth{
    remove_last_node(node.Link[0], last_left_node)
  }else{
    remove_last_node(node.Link[1], last_right_node)
  }
}

func (tree *Tree)Pop(){
  //fmt.Println("---1---")
  if tree.Root == nil{
    //fmt.Println("---2---")
    return
  } else if tree.Root.Link[0] == nil{ 
    //fmt.Println("---3---")
    tree.Root = nil
    return
  } else if tree.Root.Link[1] == nil && tree.Root.Link[0] != nil{
    tree.Root.Link[0] = nil
    return
  }
  _, left_depth := depth_of_left_nodes(tree.Root.Link[0], 0)


  //fmt.Println("*****************************************************>", left_depth, right_depth)
  //fmt.Println("*****************************************************>", tree.Root.Value, tree.Root.Link[0].Value)

  if left_depth == 0{
    //fmt.Println("When depth is zeroooo")
    if tree.Root.Link[1] != nil{
      tree.Root.Link[1] = nil
      return
    }else{
      tree.Root.Link[0] = nil
      return
    }
  }
  //fmt.Println("---4---")
  tree.Root.remove()
}

func producers(tree *Tree, number int){
  for i:=1; i<number+1; i++{
    tree.Push(i)
  }
}

func consumers(tree *Tree, number int){
  for i:=1; i<number+1; i++{
    tree.Pop()
  }
}

func main(){
  tree := Tree{}

/*
  var number int
  fmt.Println("Enter the number of producers you want: ")
  fmt.Scanf("%d", &number)

  start := time.Now()

  producers(&tree, number)
  data, err := json.MarshalIndent(tree, "", "  ")
  //fmt.Println("***Pavilion***")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", data)
  consumers(&tree, number)
  data, err = json.MarshalIndent(tree, "", "  ")
  //fmt.Println("***Pavilion***")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", data)

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)

}
*/


  start := time.Now()
  tree.Push(1)
  tree.Push(2)
  tree.Push(3)
  tree.Push(4)
  tree.Push(5)
  tree.Push(6)
  //tree.Push(7)
  //tree.push(8)
  fmt.Println("***Pavilion***")
  data, err := json.MarshalIndent(tree, "", "  ")
  fmt.Println("***Pavilion***")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", data)
/*
  tree.Pop()
  tree.Pop()
  tree.Pop()


  data, err = json.MarshalIndent(tree, "", "  ")
  fmt.Println("***Pavilion***")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", data)
*/

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}



