package main
import("fmt")

type Avl struct{
  data int
  l_child *Avl
  r_child *Avl
  parent *Avl
  balance int
}

func right_left_rotate(avl *Avl){
  p1 := avl.parent
  c1 := avl.r_child
  c2 := c1.l_child

  avl.r_child = c2
  c2.r_child  = c1
  c1.l_child  = nil

  if p1.r_child == avl{
    p1.r_child = c2
  }else{
    p1.l_child = c2
  }
  c2.l_child = avl
  avl.r_child = nil
}

func left_right_rotate(avl *Avl){
  p1 := avl.parent
  c1 := avl.l_child
  c2 := c1.r_child

  avl.l_child = c2
  c2.l_child  = c1
  c1.r_child = nil

  if p1.r_child == avl{
    p1.r_child = c2
  }else{
    p1.l_child = c2
  }

  c2.r_child = avl
  avl.l_child = nil
}


func right_rotate(avl *Avl){
  p1 := avl.parent
  c1 := avl.r_child
  c2 := c1.r_child

  p1.r_child = c1
  c1.r_child = c2
  c1.l_child = avl

  avl.r_child = nil
  c2.r_child  = nil
}


func left_rotate(avl *Avl){
  fmt.Println("Parent data: ", avl.parent.data)
  fmt.Println("My data: ", avl.data)

  p1 := avl.parent
  c1 := avl.l_child
  c2 := c1.l_child

  fmt.Println("Last Element: ", c2.data)

  p1.r_child = c1
  c1.r_child = avl
  c1.l_child = c2

  c1.r_child.l_child = nil
  c1.l_child.l_child = nil


  c1.r_child.r_child = nil
  c1.l_child.r_child = nil

  fmt.Println("p1-->", p1.data, p1.r_child.data)
  fmt.Println("c1-->", c1.r_child.data, c1.l_child.data)

}

func double_rotation(avl *Avl){
  fmt.Println("Calling double rotation")
}

func insert(avl *Avl, data int) (int){
  if data < avl.data{
    if avl.l_child == nil{
      child := Avl{data: data, balance: 1, parent: avl}
      avl.l_child = &child
      if avl.r_child == nil{
        return avl.balance + 2
      }
      return avl.balance
    }else{
      if avl.r_child ==nil{
        avl.l_child.balance = avl.balance + 2
      }
      balance := insert(avl.l_child, data)
      fmt.Println("Balance: ",balance, data)
      if balance == 4{
        left_rotate(avl)
      }else if balance == 5{
        fmt.Println("Calling from Here***1***")
        left_right_rotate(avl)
      }

    }
  }else{
    //fmt.Println("***Inside Right***")
    if avl.r_child == nil{
      fmt.Println("***nil_for_first_and_last***", data, avl.balance)
      child := Avl{data: data, parent: avl}
      avl.r_child = &child
      if avl.l_child == nil{
        return avl.balance + 3
      }
      return avl.balance
    }else{
      //fmt.Println("***other_Data***", data)
      if avl.l_child == nil{
        avl.r_child.balance = avl.balance + 3
      }
      balance := insert(avl.r_child, data)
      fmt.Println("Balance: ",balance, data)
      if balance == 6{
        fmt.Println("*****Right-Rotating*****")
        right_rotate(avl)
      }else if balance == 5{
        fmt.Println("Calling from Here***2***")
        right_left_rotate(avl)
      }

    }
  }
  return avl.balance
}

func print_tree(t *Avl){


  if t.l_child == nil{

    if t.r_child ==nil{
      fmt.Println(t.data)
      return
    }
  }else{
    print_tree(t.l_child)
  }


  if t.r_child == nil{

    if t.r_child == nil{
      fmt.Println(t.data)
      return
    }
  }else{
    print_tree(t.r_child)
  }
  fmt.Println(t.data)
  return
}

//***removing***
func check_for_single_child(node *Avl)bool{
  fmt.Println("--->NODE-->1", node.data)

  if node.l_child == nil{
  }else{
    check_for_single_child(node.l_child)
    fmt.Println("--->NODE-->2", (*node).data)
  }

  if node.r_child == nil{
    if node.l_child != nil{
      fmt.Println("found a single value-->", node.data)
      fmt.Println("found a single value-->", node)
      node.l_child = nil
      fmt.Println("found a single value-->", node)
      return true
    }
  }else{
    check_for_single_child(node.r_child)
  }
  return false
}

func depth_of_left_nodes(node *Avl, left_depth int)(*Avl, int){
  if node.r_child ==nil{
    return node, left_depth
  }else{
    last_node, count := depth_of_left_nodes(node.r_child, left_depth+1)
    return last_node, count
  }
}

func depth_of_right_nodes(node *Avl, right_depth int)(*Avl, int){
  if node.r_child ==nil{
    return node, right_depth
  }else{
    last_node, count := depth_of_right_nodes(node.r_child, right_depth+1)
    return last_node, count
  }
}

func remove_last_node(node *Avl, last_node *Avl){

  if node.r_child == last_node{
    fmt.Println("Popping element--->", node.r_child.data)
    node.r_child = nil
    return
  }else{
    remove_last_node(node.r_child, last_node)
    return
  }
}

func remove(node *Avl){
  if check_for_single_child(node.l_child){
    return
  }

  if check_for_single_child(node.r_child){
    return
  }

  last_left_node, left_depth   := depth_of_left_nodes(node.l_child, 0)
  last_right_node, right_depth := depth_of_right_nodes(node.r_child, 0)

  fmt.Println("Last_LeftNOde-->", last_left_node.data)
  fmt.Println("Last_RightNOde-->", last_right_node.data)


  if left_depth > right_depth{
    remove_last_node(node.l_child, last_left_node)
  }else{
    remove_last_node(node.r_child, last_right_node)
  }

}
///* left rotate example
func main(){
  var a = Avl{data: 41}
  insert(&a, 20)
  insert(&a, 55)
  insert(&a, 11)
  insert(&a, 50)
  insert(&a, 65)
  insert(&a, 29)
  insert(&a, 26)
  insert(&a, 23)
  print_tree(&a)
  remove(&a)
  print_tree(&a)
}
//*/

/* right rotate example
func main(){
  var a = Avl{data: 41}
  insert(&a, 20)
  a.balance = 0
  insert(&a, 55)
  a.balance = 0
  insert(&a, 65)
  a.balance = 0
  insert(&a, 70)
  a.balance = 0
  print_tree(&a)
}
*/


/*left right rotate example
func main(){
  var a = Avl{data: 41}
  insert(&a, 20)
  a.balance = 0
  insert(&a, 65)
  a.balance = 0
  insert(&a, 50)
  a.balance = 0
  insert(&a, 55)
  a.balance = 0
  print_tree(&a)
}
*/


/* right left roatate example
func main(){
  var a = Avl{data: 41}
  insert(&a, 20)
  a.balance = 0
  insert(&a, 65)
  a.balance = 0
  insert(&a, 11)
  a.balance = 0
  insert(&a, 26)
  a.balance = 0
  insert(&a, 27)
  a.balance = 0
  insert(&a, 21)
  a.balance = 0
  insert(&a, 70)
  a.balance = 0
  print_tree(&a)
}
*/
