package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type V struct {
	x, y, z    float64
	vx, vy, vz float64
}

var vectors = []V{}

func readInput(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		splits := strings.Split(line, " @ ")
		pos := strings.Split(splits[0], ", ")
		x, ok := strconv.Atoi(strings.Trim(pos[0], " "))
		if ok != nil {
			panic(ok)
		}
		y, ok := strconv.Atoi(strings.Trim(pos[1], " "))
		if ok != nil {
			panic(ok)
		}
		z, ok := strconv.Atoi(strings.Trim(pos[2], " "))
		if ok != nil {
			panic(ok)
		}
		vel := strings.Split(splits[1], ", ")
		vx, ok := strconv.Atoi(strings.Trim(vel[0], " "))
		if ok != nil {
			panic(ok)
		}
		vy, ok := strconv.Atoi(strings.Trim(vel[1], " "))
		if ok != nil {
			panic(ok)
		}
		vz, ok := strconv.Atoi(strings.Trim(vel[2], " "))
		if ok != nil {
			panic(ok)
		}
		if vx == 0 || vy == 0 || vz == 0 {
			panic("zero")
		}
		v := V{float64(x), float64(y), float64(z), float64(vx), float64(vy), float64(vz)}
		vectors = append(vectors, v)
	}
}

func findIntersection(v1, v2 V) (float64, float64, bool, bool) {
	x1, y1 := v1.x, v1.y
	x2, y2 := v1.x+v1.vx*1e9, v1.y+v1.vy*1e9
	x3, y3 := v2.x, v2.y
	x4, y4 := v2.x+v2.vx*1e9, v2.y+v2.vy*1e9
	denom := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if denom == 0 {
		return .0, .0, false, false
	}
	x := ((x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)) / denom
	y := ((x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)) / denom
	f1 := (x > x1) == (v1.vx > 0)
	f2 := (x > x3) == (v2.vx > 0)
	return x, y, f1, f2
}

// const L1, L2 = 7, 27

const (
	L1 = 200000000000000
	L2 = 400000000000000
)

func firstPart(f *os.File) {
	readInput(f)
	// fmt.Println(vectors)
	n := len(vectors)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v1, v2 := vectors[i], vectors[j]
			// fmt.Println(v1, v2)
			x, y, f1, f2 := findIntersection(v1, v2)
			if i == 14 && j == 141 {
				fmt.Printf("%f %f %f %f\n", v1.x, v1.y, v1.vx, v1.vy)
				fmt.Printf("%f %f %f %f\n", v2.x, v2.y, v2.vx, v2.vy)
				fmt.Printf("%v %v %f %f\n", f1, f2, x, y)
			}
			if f1 && f2 && L1 <= x && x <= L2 && L1 <= y && y <= L2 {
				fmt.Println(i, j)
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func findPossibleVelocities(vs map[int][]int) []int {
	vels := map[int]bool{}
	for i := -1000; i <= 1000; i++ {
		vels[i] = true
	}
	for v, ps := range vs {
		for i := 0; i < len(ps); i++ {
			for j := i + 1; j < len(ps); j++ {
				for k := range vels {
					if vels[k] && k-v != 0 && (ps[i]-ps[j])%(k-v) != 0 {
						vels[k] = false
					}
				}
			}
		}
	}
	validVs := []int{}
	for v, ok := range vels {
		if ok {
			validVs = append(validVs, v)
		}
	}
	return validVs
}

func findT(x1, vx1, y1, vy1, x2, vx2, y2, vy2, vx, vy int) (int, bool) {
	dvx1 := vx - vx1
	dvx2 := vx - vx2
	dvy1 := vy - vy1
	dvy2 := vy - vy2
	denom := dvy2*dvx1 - dvx2*dvy1
	if denom == 0 {
		return -1, false
	}
	num := (y2-y1)*dvx2 - (x2-x1)*dvy2
	if num%denom != 0 {
		return -1, true
	}
	t1 := num / denom
	// x := v1.x - float64(dvx1) * t1
	// t2 := (v2.x - x) / float64(dvx2)
	return t1, true
}

func secondPart(f *os.File) {
	readInput(f)
	// fmt.Println(vectors)
	n := len(vectors)
	// n := 10
	vxs := map[int][]int{}
	vys := map[int][]int{}
	vzs := map[int][]int{}
	for i := 0; i < n; i++ {
		v := vectors[i]
		x := int(v.x)
		vx := int(v.vx)
		if xs, ok := vxs[vx]; ok {
			vxs[vx] = append(xs, x)
		} else {
			vxs[vx] = []int{x}
		}
		y := int(v.y)
		vy := int(v.vy)
		if ys, ok := vys[vy]; ok {
			vys[vy] = append(ys, y)
		} else {
			vys[vy] = []int{y}
		}
		z := int(v.z)
		vz := int(v.vz)
		if zs, ok := vzs[vz]; ok {
			vzs[vz] = append(zs, z)
		} else {
			vzs[vz] = []int{z}
		}
	}
	fmt.Println(vxs)
	fmt.Println(vys)
	fmt.Println(vzs)
	pvxs := findPossibleVelocities(vxs)
	pvys := findPossibleVelocities(vys)
	pvzs := findPossibleVelocities(vzs)
	fmt.Println(pvxs)
	fmt.Println(pvys)
	fmt.Println(pvzs)
	for _, vx := range pvxs {
		for _, vy := range pvys {
			for _, vz := range pvzs {
				valid := true
			outer:
				for i := 0; i < n-1; i++ {
					t := -1
					first := true
					for j := i + 1; j < n; j++ {
						v1, v2 := vectors[i], vectors[j]
						v2 = vectors[j]
						t00, ok0 := findT(int(v1.x), int(v1.vx), int(v1.y), int(v1.vy), int(v2.x), int(v2.vx), int(v2.y), int(v2.vy), vx, vy)
						t11, ok1 := findT(int(v1.x), int(v1.vx), int(v1.z), int(v1.vz), int(v2.x), int(v2.vx), int(v2.z), int(v2.vz), vx, vz)
						t22, ok2 := findT(int(v1.y), int(v1.vy), int(v1.z), int(v1.vz), int(v2.y), int(v2.vy), int(v2.z), int(v2.vz), vy, vz)
						if !(ok0 && ok1 && ok2) {
							continue
						}
						if first {
							t = t00
							first = false
						}
						if t00 < 0 || t11 < 0 || t22 < 0 || t != t00 || t != t11 || t != t22 {
							valid = false
							break outer
						}
					}
				}
				if valid {
					v1, v2 := vectors[0], vectors[1]
					t, ok := findT(int(v1.x), int(v1.vx), int(v1.y), int(v1.vy), int(v2.x), int(v2.vx), int(v2.y), int(v2.vy), vx, vy)
					if !ok {
						panic("try another pair")
					}
					x := int(v1.x + (v1.vx-float64(vx))*float64(t))
					y := int(v1.y + (v1.vy-float64(vy))*float64(t))
					z := int(v1.z + (v1.vz-float64(vz))*float64(t))
					fmt.Println(x + y + z)
				}
			}
		}
	}
}

func main() {
	inputPath := os.Args[1]
	fmt.Println(inputPath)
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if len(os.Args) == 2 {
		fmt.Println("First part")
		firstPart(f)
	} else {
		fmt.Println("Second part")
		secondPart(f)
	}
}
