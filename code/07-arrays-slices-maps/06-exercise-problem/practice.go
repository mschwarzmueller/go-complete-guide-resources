package main
type product struct{
  title string
  id string
  price int
  }

func main() {
   hobbies:=[3]string{"badminton","dancing","scrolling"}
   fmt.Println(hobbies)

  fmt.Println(hobbies[0])
  fmt.Println(hobbies[1:3])

  mainhobbies1:=[] string { hobbies[0],hobbies[1] }
  fmt.Println(mainhobbies1)

  mainhobbies2:= make( []string , hobbies[0])
  mainhobbies2 = append(mainhobbies,hobbies[1])
  fmt.Println(mainhobbies2)

  mainhobbies2= hobbies[1:3]
  fmt.Println(mainhobbies2)

  dynamic_Arrays:= []string{"tolearngo","tomakeaproject"}
  fmt.Pritnln(dynamic_Arrays)


  dynamic_Arrays[1]="tomakelotsofproject"
  dynmain_Arrays= append(dynamic_Arrays,"tomastergo"}
                         
  product_list1:=[]product{ product{title:"vanila",id:"67",price:99},product{title:"chocolate",id:"89",price:65}
  fmt.Println(product_list1)

  
  
  
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
