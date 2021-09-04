package main

import (
	"fmt"

	cond "packages/condition"
	learn_func "packages/function"
	"packages/hello"
)

func main() {
	fmt.Println("\t---- Package in go ----")

	fmt.Println("Step1: go mod init <project name>.")
	fmt.Println("Step2: Make directory, package name should be same directory name.")
	fmt.Println("Step3: Make public function with Uppercase such as Myfucntion.")
	fmt.Println("Step4: Import package/function with <root project name>/<package path>")
	fmt.Println("\tin this case use \"packages/hello\".")
	fmt.Println("Step5: Use package with <package name>.<public function name>")
	fmt.Println("\tin this case use hello.Hello(), Now you can use function from other package!\n")
	fmt.Println()

	pkg_hello := hello.HelloWorld()
	fmt.Println(pkg_hello)
	fmt.Println()

	cond.Conditions()
	fmt.Println()

	learn_func.Function()
	fmt.Println()
	learn_func.ValueFunction()
	fmt.Println()

	fmt.Println()
	fmt.Println("Tips: Re new package name you can do like this...")
	fmt.Println("\tnew_name <root project>/<package path>")
}
