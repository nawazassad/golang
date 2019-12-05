package main
import("fmt")

type Tree struct{
  Root *Avl
}

type Avl struct{
  Data int
  L_child *Avl
  R_child *Avl
  Parent *Avl
  Balance int
}

func right_left_rotate(avl *Avl){
  p1 := avl.Parent
  c1 := avl.R_child
  c2 := c1.L_child

  avl.R_child = c2
  c2.R_child  = c1
  c1.L_child  = nil

  if p1.R_child == avl{
    p1.R_child = c2
  }else{
    p1.L_child = c2
  }
  c2.L_child = avl
  avl.R_child = nil
}

func left_right_rotate(avl *Avl){
  p1 := avl.Parent
  c1 := avl.L_child
  c2 := c1.R_child

  avl.L_child = c2
  c2.L_child  = c1
  c1.R_child = nil

  if p1.R_child == avl{
    p1.R_child = c2
  }else{
    p1.L_child = c2
  }

  c2.R_child = avl
  avl.L_child = nil
}


func right_rotate(avl *Avl){
  fmt.Println("right_rotate----called----1", avl.Data)
  p1 := avl.Parent
  c1 := avl.R_child
  c2 := c1.R_child
  fmt.Println("right_rotate----called----2", c1)

  p1.R_child = c1
  c1.R_child = c2
  c1.L_child = avl
  fmt.Println("right_rotate----called----3")

  avl.R_child = nil
  c2.R_child  = nil

  c1.Parent = p1
  avl.Parent = c1
  c2.Parent  = c1
  fmt.Println("right_rotate----called----4")
}


func left_rotate(avl *Avl){
  fmt.Println("Parent Data: ", avl.Parent.Data)
  fmt.Println("My Data: ", avl.Data)

  p1 := avl.Parent
  c1 := avl.L_child
  c2 := c1.L_child

  fmt.Println("Last Element: ", c2.Data)

  p1.R_child = c1
  c1.R_child = avl
  c1.L_child = c2

  c1.R_child.L_child = nil
  c1.L_child.L_child = nil


  c1.R_child.R_child = nil
  c1.L_child.R_child = nil

  fmt.Println("p1-->", p1.Data, p1.R_child.Data)
  fmt.Println("c1-->", c1.R_child.Data, c1.L_child.Data)

}

func double_rotation(avl *Avl){
  fmt.Println("Calling double rotation")
}

func change_dynamics(avl *Avl)(*Avl){
  fmt.Println("change in dynamics---", avl.Data)
  avl = avl.R_child
  lc := avl.L_child
  p1 := avl.Parent

  fmt.Println("---1---")
  lc.R_child = p1
  lc.L_child = p1.L_child
  p1.Parent = lc
  fmt.Println("---2---")
  fmt.Println("---3---")
  p1.L_child = nil
  p1.R_child = nil
  fmt.Println("---3---")


  root := avl
  return root

}


func (avl *Avl)insert(Data, ld, rd int) (int, int, int, *Avl){
  fmt.Println("--->INSERTCALLED<---", avl.Data)
  if Data < avl.Data{
    if avl.L_child == nil{
      child := Avl{Data: Data, Balance: 1, Parent: avl}
      avl.L_child = &child
      if avl.R_child == nil{
        return avl.Balance + 2,ld, rd, avl
      }
      return avl.Balance, ld+1,rd, avl
    }else{
      if avl.R_child ==nil{
        avl.L_child.Balance = avl.Balance + 2
      }
      Balance,_, _, _ := avl.L_child.insert(Data, ld, rd)
      //fmt.Println("Balance: ",Balance, Data)
      if Balance == 4{
        left_rotate(avl)
      }else if Balance == 5{
        //fmt.Println("Calling from Here***1***")
        left_right_rotate(avl)
      }

    }
  }else{
    //fmt.Println("***Inside Right***")
    if avl.R_child == nil{
      //fmt.Println("***nil_for_first_and_last***", Data, avl.Balance)
      fmt.Println("***Inserting-Data***", Data)
      avl.R_child = &Avl{Data: Data, Parent: avl}
      //fmt.Println("Inserted to right")
      if avl.L_child == nil{
        fmt.Println("Returning with incrementing--->", avl.Balance + 3)
        return avl.Balance + 3, ld, rd+1, avl
      }

      fmt.Println("Returning wihtout incrementing--->", avl.Balance)
      return avl.Balance, ld+1, rd, avl
    }else{
      //fmt.Println("***other_Data***", Data)
      if avl.L_child == nil{
        avl.R_child.Balance = avl.Balance + 3
      }
      fmt.Println("CAlling Insert with AVL:", avl.Data, "With R_child:", avl.R_child)
      Balance, ld, rd, _ := avl.R_child.insert(Data, ld, rd)
      //fmt.Println("Balance: ",Balance, "Data:", Data, "AVL.DATA", avl.Data)
      _, depth := depth_of_right_nodes(avl, 0)
      //fmt.Println("depth------>", depth)
      if Balance == 6{
        fmt.Println("*****Right-Rotating*****", avl.Data)
        if rd -ld == 2{
          if avl.L_child ==nil{
            //fmt.Println("************INSIDE*******calling right_rotate")
            right_rotate(avl)
          }else{
            root := change_dynamics(avl)
            return 0, ld, rd, root
          }
        }
      }else if Balance == 5{
        //fmt.Println("Calling from Here***2***")
        right_left_rotate(avl)
      }else if Balance == 3 && depth == 3{
        fmt.Println("Heloo am hererrrrrrrrrrrrrrrr", avl.Data)
        root := &Avl{}
        if avl.L_child == nil{
            right_rotate(avl.R_child)
        }else{
          root =  change_dynamics(avl)
        }
        fmt.Println("Changed Root value--->", root.Data)
        return 0, ld, rd, root

      }

    }
  }
  return avl.Balance, ld, rd, avl
}

func print_tree(t *Avl){


  if t.L_child == nil{

    if t.R_child ==nil{
      fmt.Print(t.Data)
      return
    }
  }else{
    print_tree(t.L_child)
  }


  if t.R_child == nil{

    if t.R_child == nil{
      fmt.Print(t.Data)
      return
    }
  }else{
    print_tree(t.R_child)
  }
  fmt.Print(t.Data)
  return
}

//***removing***
func check_for_single_child(node *Avl)bool{
  fmt.Println("--->NODE-->1", node.Data)

  if node.L_child == nil{
  }else{
    check_for_single_child(node.L_child)
    fmt.Println("--->NODE-->2", (*node).Data)
  }

  if node.R_child == nil{
    if node.L_child != nil{
      fmt.Println("found a single value-->", node.Data)
      fmt.Println("found a single value-->", node)
      node.L_child = nil
      fmt.Println("found a single value-->", node)
      return true
    }
  }else{
    check_for_single_child(node.R_child)
  }
  return false
}

func depth_of_left_nodes(node *Avl, left_depth int)(*Avl, int){
  if node.R_child ==nil{
    return node, left_depth
  }else{
    last_node, count := depth_of_left_nodes(node.R_child, left_depth+1)
    return last_node, count
  }
}

func depth_of_right_nodes(node *Avl, right_depth int)(*Avl, int){
  fmt.Println("depth_of_right_nodes called ---------------->", node.Data, right_depth)
  if node.R_child ==nil{
    return node, right_depth
  }else{
    last_node, count := depth_of_right_nodes(node.R_child, right_depth+1)
    return last_node, count
  }
}

func remove_last_node(node *Avl, last_node *Avl){

  if node.R_child == last_node{
    fmt.Println("Popping element--->", node.R_child.Data)
    node.R_child = nil
    return
  }else{
    remove_last_node(node.R_child, last_node)
    return
  }
}

func remove(node *Avl){
  if check_for_single_child(node.L_child){
    return
  }

  if check_for_single_child(node.R_child){
    return
  }

  last_left_node, left_depth   := depth_of_left_nodes(node.L_child, 0)
  last_right_node, right_depth := depth_of_right_nodes(node.R_child, 0)

  fmt.Println("Last_LeftNOde-->", last_left_node.Data)
  fmt.Println("Last_RightNOde-->", last_right_node.Data)


  if left_depth > right_depth{
    remove_last_node(node.L_child, last_left_node)
  }else{
    remove_last_node(node.R_child, last_right_node)
  }

}

func (tree *Tree)Push(value int){
  if tree.Root == nil{
    tree.Root = &Avl{Data: value}
    return
  }

  left_depth :=0
  right_depth :=0
  if tree.Root.L_child == nil{
    if tree.Root.R_child != nil{
      if tree.Root.R_child.Data < value{
        fmt.Println("Here************************************")
        root := tree.Root
        tree.Root = root.R_child //&Avl{Data: value}
        tree.Root.R_child = &Avl{Data: value}
        tree.Root.L_child = root
        tree.Root.L_child.R_child = nil
        tree.Root.L_child.Parent = tree.Root
        tree.Root.R_child.Parent = tree.Root
        return
      }
    }
  }
  if tree.Root.L_child != nil && tree.Root.R_child != nil{
    _, left_depth = depth_of_left_nodes(tree.Root.L_child, 0)
    _, right_depth = depth_of_right_nodes(tree.Root.R_child, 0)
  }
  fmt.Println("Left depth:", left_depth, "Right DEpth:", right_depth, "*****************************************")
  //if left_depth == 0 && right_depth ==2:


  //fmt.Println("coming here")
  _, _,_, tree.Root = tree.Root.insert(value, left_depth, right_depth)
  fmt.Println("//////////////************",tree.Root.Data, "***********/////////////////////")
  //fmt.Println("returning here")
  //fmt.Println(tree.Root.R_child.Data)
}

func producer(tree *Tree, number int){
  for i:=1; i<number+1; i++{
    fmt.Println("--------------------------------------------------",i,"--------------------------------------------------")
    tree.Push(i)
  }
}

func main(){
  number := 9
  tree := Tree{}
  producer(&tree, number)
  print_tree(tree.Root)
  fmt.Println()
}


