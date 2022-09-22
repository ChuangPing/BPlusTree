package bptree

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
)

//产生测试用例数量  50 100  300 500  1000 5000  10000
const TestCOUNT = 10000

//	初始化B+树
var bpTree = NewBPTree(5)

//	实验数据产生函数
func CreateTestData(sum int) []int {
	var testKeys []int
	//	产生1~50个有序数
	for i := 1; i <= sum; i++ {
		testKeys = append(testKeys, i)
	}
	return testKeys
}

type BPTreeTestDataStruct struct {
	Key      int64
	setValue interface{}
	getValue interface{}
}

func TestBPTree_Get(t *testing.T) {

	//	初始化B+树
	//bpTree := NewBPTree(4)
	//测试用例
	var testData []BPTreeTestDataStruct
	testKeys := CreateTestData(TestCOUNT)
	for _, key := range testKeys {
		testData = append(testData, BPTreeTestDataStruct{
			Key:      int64(key),
			setValue: int64(key),
			getValue: int64(key),
		})
	}
	//testData := []struct {
	//	key int64
	//	// 存入的值
	//	setValue interface{}
	//	// 正确获取的值
	//	getValue interface{}
	//}{

	//	10   11  12  16  20  21  22  40  50  51  52  55  56  57   60  69  77  78  79  80  85  86  87
	//{10, 10, 10},
	//{11, 11, 11},
	//{12, 12, 12},
	//{16, 16, 16},
	//{20, 20, 20},
	//{21, 21, 21},
	//{22, 22, 22},
	//{40, 40, 40},
	//{50, 50, 50},
	//{51, 51, 51},
	//{52, 52, 52},
	//{55, 55, 55},
	//{56, 56, 56},
	//{57, 57, 57},
	//{60, 60, 60},
	//{69, 69, 69},
	//{77, 77, 77},
	//{78, 78, 78},
	//{79, 79, 79},
	//{80, 80, 80},
	//{85, 85, 85},
	//{86, 86, 86},
	//{87, 87, 87},
	//}

	//	循环测试用例进行测试
	for index, data := range testData {
		//	将测试用例存入b+树
		bpTree.Set(data.Key, data.setValue)
		if index == len(testData)-1 {
			for _, data := range testData {
				res, count := getData(data.Key)
				//msg := fmt.Sprintf("BPlusTree search %v,count %v\n", res, count)
				msg := fmt.Sprintf(" %v,%v\n", res, count)
				writeLog(msg, "data/BPlusTree"+strconv.Itoa(TestCOUNT)+".txt")
				if res != data.getValue {
					t.Errorf("set key: %v,value %v input;expect %v; but got %v", data.Key, data.setValue, data.getValue, res)
				}
			}
		}

	}
}

func getData(key int64) (interface{}, int) {
	return bpTree.Get(key)
}

type BinaryTestDataStruct struct {
	searchKeys int
}

func TestBinarySearch(t *testing.T) {
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	keys := []int{10, 11, 12, 16, 20, 21, 22, 40, 50, 51, 52, 55, 56, 57, 60, 69, 77, 78, 79, 80, 85, 86, 87}
	res, count := BinarySearch(keys, 20)
	fmt.Printf("search %v,count %v\n", res, count)
	//return

	testKeys := CreateTestData(TestCOUNT)
	//var testData []struct{searchKey int}
	var testData []BinaryTestDataStruct
	for _, testkey := range testKeys {
		testData = append(testData, BinaryTestDataStruct{searchKeys: testkey})
	}

	for _, data := range testData {
		res, count := BinarySearch(testKeys, data.searchKeys)
		//msg := fmt.Sprintf("binarySearch: search %v,count %v\n", res, count)
		msg := fmt.Sprintf(" %v,%v\n", res, count)
		writeLog(msg, "data/binarySearch"+strconv.Itoa(TestCOUNT)+".txt")
	}

}

// 写入文件函数
func writeLog(msg string, fileName string) {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	fileHandle, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	defer fileHandle.Close()
	// NewWriter 默认缓冲区大小是 4096
	// 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
	buf := bufio.NewWriterSize(fileHandle, len(msg))

	buf.WriteString(msg)

	err = buf.Flush()
	if err != nil {
		log.Println("flush error :", err)
	}
}

//
//func TestBPTRand(t *testing.T) {
//	bpt := NewBPTree(3)
//
//	for i := 0; i < 12; i++ {
//		key := rand.Int()%20 + 1
//		t.Log(key)
//		bpt.Set(int64(key), key)
//	}
//
//	data, _ := json.MarshalIndent(bpt.GetData(), "", "    ")
//	t.Log(string(data))
//}
