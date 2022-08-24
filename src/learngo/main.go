package main // 모든 Go 코드는 패키지에 포함돼야 하고 패키지 선언으로 시작함.

import "fmt"

func main()  { // Go 프로그램의 시작점. 
	
	
	//다양한 for 문 사용방법
		sum, i := 0 ,0

		for i := 0; i < 10; i++ {
			sum += i;
		}
		
		for i < 10{
			sum += i;
			i++;
		}
		
		// for문의 조건식 생략
		for {
			if i >= 10 {
				break
			}
			sum += i;
			i++;
		}
		fmt.Println(sum)

	c := 'a'

	switch {

	// case에 조건식 사용
	case '0' <= c && c <='9':
		fmt.Println("%c는 숫자", c)
	case 'a' <= c && c <='z':
		fmt.Println("%c는 소문자", c)
	case 'A' <= c && c <='Z':
		fmt.Println("%c는 대문자", c)
	}
}