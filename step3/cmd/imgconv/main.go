package main

import (
	"bufio"
	"fmt"
	"os"
)

func run() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("引数が足りません。")
	}
	
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])

	from, to := os.Args[1], os.Args[2];

	inFile, err := os.Open(from);
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)		
	}

	// 関数終了時にファイルを閉じる
	defer inFile.Close()	

	outFile, err2 := os.Create(to);
	
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err2.Error())
		os.Exit(1)		
	}
	
	// 関数終了時にファイルを閉じる
	defer outFile.Close()
	
	scanner := bufio.NewScanner(inFile)
	
	for i := 1; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
		fmt.Fprintf(outFile, "%d:%s\n", i, scanner.Text())
	}

	return scanner.Err()
}


func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	
}
