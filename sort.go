package main
import(
        "fmt"
)
func main(){
        var arr [10]int = [10]int{15,1,23,14,11,6,8,4,10,30}
        var arr0 [10]int = arr
        var tmp int     
        for i :=0; i < len(arr); i ++ {
                for j :=1; j < len(arr)-i; j++{
                        if(arr[j]>arr[j-1]){
                                tmp = arr[j]
                                arr[j] = arr[j-1]
                                arr[j-1] = tmp
                        }
                }       
        }
        fmt.Println("排序之前",arr0,"\n排序之后",arr)


}