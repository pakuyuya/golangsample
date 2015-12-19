package varconst

type T struct {
  mem1 int
}

func varconst() {
  var str1 string = "1"
  var strPointer *string = &str1
  
  const (
    multivar1 = 1
    multivar2 = 2
  )
  
  var aryInt [10]int = [10]int{1, 2, 3, 4, 5, }
  var aryInt2 = [...]int{1, 2, 3, 4, 5}         // 型推論
  
  var t1 *T = new(T)
  var t = new(T)
  
  var m map[string]int = make(map[string]int)
}