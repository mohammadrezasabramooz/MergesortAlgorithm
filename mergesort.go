package main
import (
	"os"
	"log"
	"bufio"
	"strconv"
	"time"
	"fmt"
	"math"

)


func main() {
	sizeofarray:=10000000
	var array_database [] int
	array_database=make([]int,sizeofarray)
	//initialize
	fmt.Println("initializing...")
	start := time.Now()
	initializeArray(array_database)
	end := time.Now()
	fmt.Printf("initialize time: %v\n",end.Sub(start))
	//sorting process
	fmt.Println("Sorting...")
	start =	time.Now()
	array_database=sort(array_database)
	end = time.Now()
	fmt.Printf("sorting time: %v\n",end.Sub(start))
	fmt.Println("Saving new array in text file...")
	write(array_database,sizeofarray)
}

func initializeArray (array [] int){
	file, err := os.Open("array.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for i:=0;scanner.Scan();i++  {
		array[i], _ =(strconv.Atoi(scanner.Text()))

	}
}
func sort(m []int) []int {
	if len(m) <= 1 {
		return m
	}
	mid := int(math.Ceil(float64(len(m)) / 2.0))
	left := m[:mid]
	right := m[mid:]

	left = sort(left)
	right = sort(right)

	arr := merge(left, right)
	return arr
}



func merge(left, right []int) []int {

	var result []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}
var (
	newFile *os.File
	err     error
)

func write(array [] int,size int)  {
	newFile, err = os.Create("sorted_array.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	for i := 0; i < size; i++ {
		_, err=newFile.WriteString(strconv.Itoa(array[i])+"\n")
	}
	newFile.Close()
	fmt.Println("==> done creating file")


}
