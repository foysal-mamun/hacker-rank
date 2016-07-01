package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
)

func main() {

	N := 0
	fmt.Scan(&N)

	P := make([][]int, N)

	if N%2 != 0 {
		fmt.Println("0")
	} else {

		for i := 0; i < N; i++ {
			x, y := 0, 0
			fmt.Scan(&x)
			fmt.Scan(&y)
			P[i] = []int{x, y}
		}

		PP := [][]int{}
		pp := make([]int, N)
		for i, _ := range P {
			pp[i] = i
		}

		p, err := NewPerm(pp, nil) //generate a Permutator
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, err := p.Next(); err == nil; i, err = p.Next() {
			// fmt.Printf("%3d permutation: %v left %d\n",p.Index()-1,i.([]int),p.Left())
			PP = append(PP, i.([]int))
		}

		newPP := removeDuplicates(PP)
		total := 0
		for _, v := range newPP {

			l1 := N - 1
			l2 := (N / 2)
			a, b, c := v[0], v[l2-1], v[l2]
			d, e, f := v[l2-1], v[l2], v[l1]

			if findAngle(P[a], P[b], P[c]) && findAngle(P[d], P[e], P[f]) && getPostion(P[a], P[b], P[c]) < 0 && getPostion(P[d], P[e], P[f]) > 0 {
				total += 1
			}

			//fmt.Println(P[d], P[e], P[f])
			// fmt.Println(angl1, angl2)

		}

		// fmt.Println(len(newPP))
		// fmt.Println(len(PP))
		fmt.Println(total)

	}

}

func getPostion(p00, p11, p22 []int) int {

	x0, y0 := float64(p00[0]), float64(p00[1])
	x1, y1 := float64(p11[0]), float64(p11[1])
	x2, y2 := float64(p22[0]), float64(p22[1])

	// + left - right
	// a = ((b.x - a.x)*(c.y - a.y) - (b.y - a.y)*(c.x - a.x))

	return int(((x1-x0)*(y2-y0) - (y1-y0)*(x2-x0)))
}

func findAngle(p00, p11, p22 []int) bool {
	x0, y0 := float64(p00[0]), float64(p00[1])
	x1, y1 := float64(p11[0]), float64(p11[1])
	x2, y2 := float64(p22[0]), float64(p22[1])

	a := math.Pow(x1-x0, 2) + math.Pow(y1-y0, 2)
	b := math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)
	c := math.Pow(x2-x0, 2) + math.Pow(y2-y0, 2)

	d := math.Acos((a+b-c)/math.Sqrt(4*a*b)) * 180 / math.Pi

	return d <= 90
}

func checkDuplicate(a []int, b []int) bool {

	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}

	if reflect.DeepEqual(a, b) {
		return true
	}

	return false
}

func removeDuplicates(elements [][]int) [][]int {
	result := [][]int{}

	for i := 0; i < len(elements); i++ {

		exists := false
		for v := 0; v < i; v++ {
			if checkDuplicate(elements[v], elements[i]) {

				exists = true
				break
			}
		}

		if !exists {
			result = append(result, elements[i])
		}
	}
	return result
}

type Less func(i, j interface{}) bool

type sortable struct {
	value reflect.Value
	less  Less
}

func (s sortable) Len() int {
	return s.value.Len()
}
func (s sortable) Less(i, j int) bool {
	return s.less(s.value.Index(i).Interface(), s.value.Index(j).Interface())
}
func (s sortable) Swap(i, j int) {
	temp := reflect.ValueOf(s.value.Index(i).Interface())
	s.value.Index(i).Set(s.value.Index(j))
	s.value.Index(j).Set(temp)
}

type Permutator struct {
	idle   chan bool
	value  reflect.Value
	less   Less
	length int
	index  int
	amount int
}

//Reset the Permutator, next time invoke p.Next() will return the first permutation in lexicalorder
func (p *Permutator) Reset() {
	<-p.idle
	sort.Sort(sortable{p.value, p.less})
	p.index = 1
	p.idle <- true
}

//return the next n permuations, if n>p.Left(),return all the left permuations
//if all permutaions generated or n is illegal(n<=0),return a empty slice
func (p *Permutator) NextN(n int) interface{} {
	<-p.idle
	//if n<=0 or we generate all pemutations,just return a empty slice
	if n <= 0 || p.left() == 0 {
		p.idle <- true
		return reflect.MakeSlice(reflect.SliceOf(p.value.Type()), 0, 0).Interface()
	}

	var i, j int
	cap := p.left()
	if cap > n {
		cap = n
	}

	result := reflect.MakeSlice(reflect.SliceOf(p.value.Type()), cap, cap)

	if p.length == 1 {
		p.index++
		l := reflect.MakeSlice(p.value.Type(), p.length, p.length)
		reflect.Copy(l, p.value)
		p.idle <- true
		result.Index(0).Set(l)
		return result.Interface()
	}

	if p.index == 1 {
		p.index++
		l := reflect.MakeSlice(p.value.Type(), p.length, p.length)
		reflect.Copy(l, p.value)
		result.Index(0).Set(l)
	}

	for k := 1; k < cap; k++ {
		for i = p.length - 2; i >= 0; i-- {
			if p.less(p.value.Index(i).Interface(), p.value.Index(i+1).Interface()) {
				break
			}
		}
		for j = p.length - 1; j >= 0; j-- {
			if p.less(p.value.Index(i).Interface(), p.value.Index(j).Interface()) {
				break
			}
		}
		//swap
		temp := reflect.ValueOf(p.value.Index(i).Interface())
		p.value.Index(i).Set(p.value.Index(j))
		p.value.Index(j).Set(temp)
		//reverse
		reverse(p.value, i+1, p.length-1)

		//increase the counter
		p.index++
		l := reflect.MakeSlice(p.value.Type(), p.length, p.length)
		reflect.Copy(l, p.value)
		result.Index(k).Set(l)
	}
	p.idle <- true
	return result.Interface()
}

//Invoke Permutator.Index() to return the index of last permutation, which start from 1 to n! (n is the length of slice)
func (p Permutator) Index() int {
	<-p.idle

	j := p.index - 1
	p.idle <- true
	return j
}

//Generate a New Permuatator, the argument k must be a non-nil slice,and the less argument must be a Less function that implements compare functionality of k's element type
//if k's element is ordered,less argument can be nil
//for ordered in Golang, visit http://golang.org/ref/spec#Comparison_operators
//After generating a Permutator, the argument k can be modified and deleted,Permutator store a copy of k internel.Rght now, a Permutator can  be used concurrently

func NewPerm(k interface{}, less Less) (*Permutator, error) {
	v := reflect.ValueOf(k)
	//check to see if i is a slice
	if v.Kind() != reflect.Slice {
		return nil, errors.New("argument must be a slice")
	}
	if v.IsValid() != true {
		return nil, errors.New("argument must not be nil")
	}
	if v.Len() == 0 {
		return nil, errors.New("argument must not be empty")
	}

	l := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
	reflect.Copy(l, v)
	v = l

	length := v.Len()
	if less == nil {
		switch v.Type().Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			less = lessInt
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			less = lessUint
		case reflect.Float32, reflect.Float64:
			less = lessFloat
		case reflect.String:
			less = lessString
		default:
			return nil, errors.New("the element type of slice is not ordered,you must provide a function\n")
		}
	}
	//check to see if v is in increasing order,if not sort it
	i := 0
	for i = 0; i < length-1; i++ {
		if !less(v.Index(i).Interface(), v.Index(i+1).Interface()) {
			break
		}
	}
	if i != length-1 {
		sort.Sort(sortable{v, less})
	}
	s := &Permutator{value: v, less: less, length: length, index: 1, amount: factorial(length)}
	s.idle = make(chan bool, 1)
	s.idle <- true
	return s, nil
}

//generate the next permuation in lexcial order,if all permutations generated,return an error
func (p *Permutator) Next() (interface{}, error) {
	<-p.idle
	//check to see if all permutations generated
	if p.left() <= 0 {
		p.idle <- true
		return nil, errors.New("all Permutations generated")
	}

	var i, j int
	//the first permuation is just p.value
	if p.index == 1 {
		p.index++
		l := reflect.MakeSlice(p.value.Type(), p.length, p.length)
		reflect.Copy(l, p.value)
		p.idle <- true
		return l.Interface(), nil
	}

	//when we arrive here, there must be some permutations to generate
	for i = p.length - 2; i >= 0; i-- {
		if p.less(p.value.Index(i).Interface(), p.value.Index(i+1).Interface()) {
			break
		}
	}
	for j = p.length - 1; j >= 0; j-- {
		if p.less(p.value.Index(i).Interface(), p.value.Index(j).Interface()) {
			break
		}
	}
	//swap
	temp := reflect.ValueOf(p.value.Index(i).Interface())
	p.value.Index(i).Set(p.value.Index(j))
	p.value.Index(j).Set(temp)
	//reverse
	reverse(p.value, i+1, p.length-1)

	//increase the counter
	p.index++
	l := reflect.MakeSlice(p.value.Type(), p.length, p.length)
	reflect.Copy(l, p.value)
	p.idle <- true
	return l.Interface(), nil
}

//return the left permutation that can be generated
func (p Permutator) Left() int {
	<-p.idle
	j := p.left()
	p.idle <- true
	return j
}

//because we use left inside some methods,so we need a non-block version
func (p Permutator) left() int {
	return p.amount - p.index + 1
}

//reverse the slice v[i:j]
func reverse(v reflect.Value, i, j int) {
	length := j - i + 1

	if length < 2 {
		return
	}

	for length >= 2 {
		temp := reflect.ValueOf(v.Index(j).Interface())
		v.Index(j).Set(v.Index(i))
		v.Index(i).Set(temp)

		length -= 2
		i++
		j--
	}
}

//caculate n!,because this function can only be invoked by NewPerm,so we do not need the check if i>=0
func factorial(i int) int {
	result := 1
	for i > 0 {
		result *= i
		i--
	}
	return result
}

func lessUint(i, j interface{}) bool {
	return reflect.ValueOf(i).Uint() < reflect.ValueOf(j).Uint()
}
func lessInt(i, j interface{}) bool {
	return reflect.ValueOf(i).Int() < reflect.ValueOf(j).Int()
}
func lessFloat(i, j interface{}) bool {
	return reflect.ValueOf(i).Float() < reflect.ValueOf(j).Float()
}
func lessString(i, j interface{}) bool {
	return reflect.ValueOf(i).Interface().(string) < reflect.ValueOf(j).Interface().(string)
}

// function findAngle(p0,p1,p2) {
//   var a = Math.pow(p1.x-p0.x,2) + Math.pow(p1.y-p0.y,2),
//       b = Math.pow(p1.x-p2.x,2) + Math.pow(p1.y-p2.y,2),
//       c = Math.pow(p2.x-p0.x,2) + Math.pow(p2.y-p0.y,2);
//   return Math.acos( (a+b-c) / Math.sqrt(4*a*b) ) * 180/Math.PI;
// }

// a = {x:0, y:0};
// a = {x:0, y:2};
// b = {x:2, y:0};
// c = {x:2, y:2};

// findAngle( {x:0, y:0}, {x:0, y:2}, {x:2, y:0})
