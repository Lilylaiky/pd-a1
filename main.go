package main

import "fmt"

// If you want to run this code, just press 'Ctrl'+'`' to open the terminal, and then input 'go run main.go' and press 'Enter'.

// I. statement

// only available in this file
func fun1(num int) int {
	// equal:
	// var ret int = num * num
	ret := num * num
	return ret
}

/* C/C++ - fun1:
int fun1(int num) {
	int ret = num * num;
	return ret;
}
*/

/* Python - fun1:
def fun1(num):
  ret = num * num
	return ret
*/

func Fun2(num int) {
	fmt.Println(num * num)
}

/* C/C++ - Func2:
void Fun2(int num) {
	::printf("%d\n", num * num);
}
*/

/* Python - Fun2:
def Fun2(num):
  print("{}", num * num)
*/

// II. branching

// use 'if-else'
func fun3(num int) int {
	if num < 4 {
		return num * num * num * num
	} else if num == 4 {
		return num * num * num
	} else {
		if num > 100 {
			return num
		}
		return num * num
	}
}

/* C/C++ - func3:
int fun3(int num) {
	if (num < 4) {
			return num * num * num * num;
	} else if (4 == num) {
		return num * num * num;
	} else {
		if (num > 100) {
			return num;
		}
		return num * num;
	}
}
*/

/* Python - func3:
def fun3(num):
  if num < 4:
		return num * num * num * num
	elif num == 4:
		return num * num * num
	else:
	  if num > 100:
		  return num
		return num * num
*/

// use 'switch-case'
func fun4(num int) {
	switch num {
	case 1:
		fmt.Println("The number is one!")
	case 2:
		fmt.Println("The number is two!")
	default:
		fmt.Printf("The number is %v!!!\n", num)
	}
}

/* C/C++ - fun4:
void fun4(int num) {
	swith (num) {
	case 1:
		::printf("The number is one!\n");
		break;
	case 2:
		::printf("The number is two!\n");
		break;
	default:
		::printf("The number is %v!\n", num);
	}
}
*/

/* Python - fun4:
There is no 'switch-case' in Python!
*/

// III. iteration

// use for loop - 1
func fun5() {
	num := 0
	for {
		if num > 5 {
			break
		}
		fmt.Print("*")
		num++
	}
	fmt.Println()
}

/* C/C++ - fun5:
void fun5() {
	int num = 0;
	while (1) {
		if (num > 5) {
			break;
		}
		::printf("*");
		++num;
	}
	::printf("\n");
}
*/

/* Rust - func5
fn fun5() {
	let mut num = 0;
	loop {
		if num > 5 {
			break;
		}
		print!("*");
		num += 1;
	}
	println!("");
}
*/

// equal
func fun6() {
	for i := 0; i < 6; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

/* C/C++ - fun6:
void fun6() {
	for (int i = 0; i < 6; ++i) {
		::printf("*");
	}
	::printf("\n");
}
*/

/* Rust - func5
fn fun5() {
	let mut num = 0;
	while num < 6 {
		print!("*");
		num += 1;
	}
	println!("");
}
*/

// IV. array iteration

func fun7() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for num := range nums {
		fmt.Printf("%v", num)
	}
	fmt.Printf("\n%v\n", nums[4])
}

/* C/C++ - fun7:
void fun7() {
	int nums[] {1, 2, 3, 4, 5, 6, 7};
	for (const auto& num : nums) {
    ::printf("%d", num);
  }
	::printf("\n%d\n", nums[4]);
}
*/

/* Rust - fun7:
fn fun7() {
	let nums = [1, 2, 3, 4, 5, 6, 7];
	for num in nums.iter() {
		print!("{}", num);
	}
	println!("\n{}", nums[4]);
}
*/

func main() {
	fmt.Println("I. statement")
	// I. statement
	fmt.Printf("%v\n", fun1(5))
	Fun2(6)

	fmt.Println("II. branching")
	// II. branching
	fmt.Printf("%v\n", fun3(2))
	fmt.Printf("%v\n", fun3(4))
	fmt.Printf("%v\n", fun3(6))
	fmt.Printf("%v\n", fun3(200))
	fun4(1)
	fun4(5)
	fun4(2)

	fmt.Println("III. iteration")
	// III. iteration
	fun5()
	fun6()

	fmt.Println("IV. array iteration")
	// IV. array iteration
	fun7()
}
