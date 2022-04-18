package main

import "fmt"

func main() {
	fmt.Println(caesarCipher("Always4Look5on-the-Bright-Side-of-Life", 27, 99))
}

// a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z
// 97-122
// A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z
// 65-90
//func caesarCipher(s string, k int32) string {
//	var t []string
//	var ans string
//	k = k % 26
//	fmt.Println(k)
//	for _, v := range s {
//		if v >= 97 && v <= 122 {
//			if v+k > 122 {
//				k = k - 26
//			}
//			t = append(t, string(v+k))
//		} else if v >= 65 && v <= 90 {
//			if v+k > 90 {
//				k = k - 26
//			}
//			t = append(t, string(v+k))
//		} else {
//			t = append(t, string(v))
//		}
//	}
//	for _, v := range t {
//		ans = ans + string(v)
//	}
//	return ans
//}

// n = 1,100
// k = 0,100
func caesarCipher(s string, k int32, n int32) string {
	var t []string
	var ans string
	k = k % 26
	for _, v := range s {
		if v >= 97 && v+k <= 122 {
			t = append(t, string(v+k))
		} else if v >= 65 && v+k <= 90 {
			t = append(t, string(v+k))
		} else if (v >= 65 && v <= 90 && v+k > 90) || (v >= 97 && v <= 122 && v+k > 122) {
			t = append(t, string(v+k-26))
		} else {
			t = append(t, string(v))
		}
	}
	for _, v := range t {
		ans = ans + string(v)
	}

	// k = k % 26

	return ans
}
