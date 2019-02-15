# unicode-mask
A library to mask Unicode strings 

Creating a simple funciton to mask certain charecters of a uncode strings was surprisingly ticky in GO.


```
import "masker"
import "fmt"

func main() {
  str := "こんにちは世界"
  masked := masker.Mask(str, 3, '*')
  reverseMasked := masker.MaskReverse(str,3,'*')
  fmt.Println(str)
  fmt.Println(masked)
  fmt.Println(reverseMasked)
}
```

```
こんにちは世界
こんに****
****は世界
```
