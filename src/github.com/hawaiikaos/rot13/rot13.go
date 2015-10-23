package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
	//"bytes"
	//"strconv"
	//"encoding/binary"
	//"math"
)

type rot13Reader struct {
	r io.Reader
}

/*func (r13 *rot13Reader) Read(b []byte) (int, error) {
    n, err := r13.r.Read(b)
    for i := 0; i <= n; i++ {
        b[i] = rot13(b[i])
    }
    return n, err
}*/

//for x := range b {
		/*letter := int(b[x])
		b[x] = b[x] + 13
		break*/
		//a := "A"
		//b := []byte(a)

func (r *rot13Reader) Read(b []byte) (int, error) {
	/*x, err := r.r.Read(b)
	fmt.Println(x)
	fmt.Println(err)*/
	/*for i := 0; i <= x; i++ {
		buf := bytes.NewBuffer(b)
		d, e := binary.ReadVarint(buf)
		f := float64(d)
		if (math.Abs(f) > 32 && math.Abs(f) < 45) {
			if (d > 38) || (d == -39) {
				if d < 0 {
					d = int64(math.Abs(float64(d)))
					d = d - 6
				} else {
					d = d -6
					d = int64(math.Abs(float64(d)))
				}
			} else {
				if d < 0 {
					d = d - 6
					d = int64(math.Abs(float64(d)))
				} else {
					d = d + 7
					d = 0 - d
				}
			}
		} else if (math.Abs(f) > 48 && math.Abs(f) < 61) {
			if (d > 54) || (d == -55) {
				if d < 0 {
					d = int64(math.Abs(float64(d)))
					d = d - 6
				} else {
					d = d -6
					d = int64(math.Abs(float64(d)))
				}
			} else {
				if d < 0 {
					d = d - 6
					d = int64(math.Abs(float64(d)))
				} else {
					d = d + 7
					d = 0 - d
				}
			}
		}
	
		binary.PutVarint(b, d)
		fmt.Println(b)
		//fmt.Println(x)
		fmt.Println(e)
	}*/
	//return len(b), nil
	return 0, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	
	//result := bytes.Compare(b,[]byte("@"))
	//c := []byte("B")
	//b++
	//fmt.Println(b)
	
	//i, err := strconv.Atoi("-42")
	//s := strconv.Itoa(-42)
	
	//buf := bytes.NewBuffer(b) // b is []byte
	//myfirstint, err := binary.ReadVarint(buf)
	
	
	
	//fmt.Println(d)
	//fmt.Println(e)
	/*switch {
	case result < 0:
		fmt.Println("lower")
	case result > 0:
		fmt.Println("higher")
		result = bytes.Compare(b,[]byte("Z"))
		switch {
		case result > 0:
	case result == 0:
		fmt.Println("equal to")
	}*/
}
