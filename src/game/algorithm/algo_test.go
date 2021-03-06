package algorithm

import (
	"testing"
	"fmt"
	"lib/utils"
)

// 测试
func TestExist(t *testing.T) {
	//cards := []byte{0x07, 0x08, 0x09, 0x18, 0x18, 0x41,0x43,0x44}
	//value := existHu(cards,0)
	//t.Log(cards, value, strconv.FormatInt(int64(value), 2))
	//cs4 := []byte{0x08, 0x08, 0x18, 0x18, 0x18, 0x41, 0x41, 0x41, 0x41}
	//var ok bool = existFour2(cs4)
	//t.Log("ok -> ", ok)
	//cs3 := []byte{0x01, 0x05, 0x09, 0x11, 0x15, 0x19, 0x21, 0x25, 0x29,0x41,0x43,0x51,0x52,0x53}
	//cs3 := []byte{0x01, 0x05, 0x09, 0x11, 0x19, 0x21, 0x29,0x41,0x42,0x43,0x44,0x51,0x52,0x53}
	cs3 := []byte{65, 9, 4, 1 ,17, 20 ,25 ,35, 39, 66, 68, 82, 83 ,83}
	var thirteen = ExistHu(cs3,[]uint32{},[]uint32{},[]uint32{},65,0)
	t.Log("thirteen -> ", fmt.Sprintf("%b", thirteen),thirteen&HU_SINGLE)
	//cs7 := []byte{0x01, 0x01, 0x02, 0x02, 0x11, 0x11, 0x19,0x19,0x21,0x21,0x44,0x44,0x52,0x52}
	//cs7 := []byte{0x01, 0x01, 0x01, 0x01, 0x11, 0x11, 0x19, 0x19, 0x21, 0x21, 0x44, 0x44, 0x52, 0x52}
	//var seven uint32 = exist7pair(cs7)
	//t.Log("seven -> ", seven)
	//cs7kong := []byte{0x01, 0x01, 0x01, 0x01, 0x11, 0x11, 0x21, 0x21, 0x21, 0x21, 0x44, 0x44, 0x44, 0x44}
	//var haveKong int = haveKong(cs7kong)
	//t.Log("haveKong -> ", haveKong)
}

// 大七对, 4刻子+2, 不可以吃/清一色/字一色
// 有序slices cs 包含吃，碰，杠
func TestHuTypeDetect(t *testing.T) {
	//大七对
	//cs := []byte{0x01, 0x01, 0x01, 0x11, 0x11, 0x11, 0x21, 0x21, 0x21,0x44,0x44, 0x51, 0x51, 0x51}
	//清一色
	//cs := []byte{0x01, 0x02, 0x03, 0x04, 0x04, 0x04,0x05,0x05,0x05,0x06,0x07,0x08,0x09,0x09}
	//字一色
	//cs := []byte{0x41, 0x41, 0x42, 0x42, 0x43, 0x43, 0x44, 0x44, 0x51,0x51,0x51,0x52,0x52,0x52}
	//cs := []byte{0x11,0x11,0x11,0x11,0x24,0x24,0x24,0x44,0x44,0x44,0x07,0x07,0x14,0x14,0x14}
	//cs := []byte{0x01,0x02,0x03,0x04, 0x14,0x15,0x16, 0x27,0x27,0x27,0x41,0x41,0x41}
	//cs := []byte{0x16,0x17,0x18,0x24,0x24,0x26,0x27,0x28,0x42,0x43,0x44,0x51,0x52,0x53}
	//cs := []byte{0x01,0x01,0x04,0x04,0x04,0x14,0x15,0x16,0x25,0x25,0x25,0x17,0x17,0x17}
	//cs := []byte{0x11,0x11,0x12,0x13,0x14,0x15,0x16,0x17,0x07,0x08,0x09,0x41,0x43,0x44}
	//cs := []byte{0x03,0x03,0x03,0x03,0x05,0x05,0x22,0x22,0x24,0x24,0x15,0x15,0x53,0x53}
	//cs := []byte{0x51, 0x52, 0x53, 0x51, 0x52, 0x53, 0x51, 0x52, 0x53, 0x51, 0x52, 0x53, 0x54, 0x54}
	//var val uint32 = HuTypeDetect(false, false, true, cs, 0)
	//t.Log("val -> ", val)
}

// 测试
func TestNextSeat(t *testing.T) {
	t.Log("NextSeat -> ", NextSeat(1))
	t.Log("NextSeat -> ", NextSeat(2))
	t.Log("NextSeat -> ", NextSeat(3))
	t.Log("NextSeat -> ", NextSeat(4))
}

// 测试
func TestRemove(t *testing.T) {
	cs2chow := []byte{0x01, 0x01, 0x01, 0x01, 0x11, 0x11, 0x21, 0x21, 0x21, 0x21, 0x44, 0x44, 0x44, 0x44}
	cs2chow = Remove(0x21, cs2chow)
	t.Log("cs2chow -> ", cs2chow, " len -> ", len(cs2chow))
	cs2chow = Remove(0x11, cs2chow)
	t.Log("cs2chow -> ", cs2chow, " len -> ", len(cs2chow))
	cs2chow = Remove(0x44, cs2chow)
	t.Log("cs2chow -> ", cs2chow, " len -> ", len(cs2chow))
}

// 测试
func TestCode(t *testing.T) {
	//c2 := byte(5460561 >> 8 & 0xFF)
	//t.Log("c2 -> ", c2)
	//card := EncodeChow(0x23,0x24,0x25)
	//t.Logf("card -> %d ", card)
	//var card uint32 = 1381651
	//var card uint32 = 21073
	//var card uint32 = 5396
	var card uint32 = 1382680
	c1, c2, c3 := DecodeChow(card)
	t.Log("c1, c2, c3 -> ", c1, c2, c3)
	//c1, c2 := DecodeChow2(card)
	//t.Log("c1, c2 -> ", c1, c2)
	//var card2 uint32 = 1447444
	//var card2 uint32 = (22 << 8) | 20
	var card2 uint32 = 5652
	t.Logf("card2 -> %d ", card2)
	c21, c22 := DecodeChow2(card2)
	t.Log("c1, c2 -> ", c21, c22)
}

// 测试胡牌
func Test_hu(t *testing.T) {
	//cards := []byte{0x02,0x03,0x04,0x26,0x26,0x41,0x42,0x42,0x43,0x44,0x26}
	//cards := []byte{0x07, 0x08, 0x09, 0x18, 0x18, 0x41,0x43,0x44}
	//cards := []byte{0x26,0x26,0x41,0x42,0x43,0x43,0x44,0x45}
	//value := existHu(cards,0)
	//t.Log(cards, value, strconv.FormatInt(int64(value), 2))
	//value = existHu3n(cards, len(cards))
	//t.Logf("%x, %d, %v\n", cards, value, strconv.FormatInt(int64(value), 2))
	//ok := Exist(0x18, cards, 3)
	//t.Log("ok -> ", ok)
	//cs, st := RemoveE(0x18, cards, 3)
	//t.Log("cs -> ", cs, " st -> ", st)
	//s, c := DecodePeng(0x413)
	//t.Logf("s %d, c %x", s, c)
	//cards2 := []byte{0x41,0x42,0x42,0x43,0x44,0x44}
	//v2 := existHu3n2_zi(cards2, len(cards2))
	//t.Log(v2)
}

// 测试胡牌
func Test_zuhe(t *testing.T) {
	//b := []int{1,2,3,4,5}
	//t.Log(zuheResult(len(b), 3))
	//ms := map[int]int{0:1,1:2,2:1,3:2}
	ms := []int{2, 2, 2, 2, 1}

	t.Log(ms)
	//ms := []int{1,2,1,2}
	//ms[1] -= 1
	//t.Log(ms)
	mss := map[int]int{0: 1, 1: 2, 2: 1, 3: 2}
	cs := mss
	delete(cs, 0)
	t.Log(cs, mss)
	list1 := []int{1, 2, 1, 1, 1}
	list2 := make([]int, len(list1)-1)
	copy(list2, append(list1[:1], list1[2:]...))
	list2[1] = 100
	t.Log(list1)
	t.Log(list2)

	cards := []byte{0x01, 0x02, 0x03, 0x23, 0x24, 0x25, 0x26, 0x26, 0x41, 0x42, 0x42, 0x43, 0x44, 0x44} //数做将,字顺子
	//cards := []byte{0x01,0x02,0x03,0x23,0x24,0x25,0x26,0x26,0x42,0x42,0x42,0x44,0x44,0x44} //数做将,字刻子
	//cards := []byte{0x01,0x02,0x03,0x23,0x24,0x25,0x41,0x41,0x41,0x43,0x44} //字做将,刻子折做将
	//cards := []byte{0x01, 0x02, 0x03, 0x11, 0x02, 0x03,0x01, 0x02, 0x03, 0x41, 0x41, 0x41, 0x42, 0x42} //字做将,杠折做将
	//cards := []byte{1, 1, 1, WILDCARD, WILDCARD ,WILDCARD ,41 ,41 ,65, 65 ,66 ,66 ,67 ,67}
	//cards := []byte{1, 1, 1, WILDCARD, WILDCARD ,WILDCARD ,41 ,41 ,65, 65 ,66 ,66 ,67 ,67}
	//cards := []byte{1, 1 ,2, WILDCARD ,3 ,WILDCARD ,4 ,4 ,5 ,5, 6, WILDCARD, 7 ,7}
	//cards := []byte{0x01,0x53,0x03,  WILDCARD, WILDCARD, 0x15, 0x15, 0x21, 0x22, 0x23, 0x42, 0x42, 0x42, 0x15}
	//Sort(cards,0,len(cards)-1)
	//v := ExistNCaiNKe(cards,[]uint32{},[]uint32{},[]uint32{},0x02)
	//v := existluanfeng(cards)
	//t.Log("是否胡牌：", v, fmt.Sprintf("%b", v), cards)

	//cards := []byte{0x07,0x13, 0x13, 0x14, 0x14, 0x15, 0x15, 0x16, 0x17, 0x17, 0x18, 0x19, WILDCARD, WILDCARD} //数做将,字顺子
	//cards := []byte{0x03,0x03,0x03,0x23,0x23,0x23,0x26,0x26,0x26,WILDCARD,WILDCARD,0x44,0x44,0x44} //数做将,字顺子

	cards = []byte{ 0x01, 0x01, 0x02, 0x03, 0x04, 0x05,0x06, 0x07, 0x08, 0x09, 0x09,0x09,0x09}
	//11 12 13  53 14 15 13 23 24 27 27
	//v := ExistHu(cards, []uint32{}, []uint32{EncodePeng(1,0x01)}, []uint32{}, 0, 0x19)
	v := ExistHu(cards, []uint32{}, []uint32{}, []uint32{}, 0x01, 0x53)
	t.Logf("%v  %v %v ", v&HU_ONE_SUIT > 0, v&HU > 0, v&HU_GUI_WEI2 > 0)

	t.Log(utils.Md5("XG0e2Ye/KAUJRXaMNnJ5UH1haBvh2FXOoAggE6f2Utw6022912345678900987654321"))
	utils.Md5("XG0e2Ye/KAUJRXaMNnJ5UH1haBvh2FXOoAggE6f2Utw" + "60229" + "1234567890" + "0987654321")
	//"081c319e6d948ec8109b6b32abfb4fff"
	// 081c319e6d948ec8109b6b32abfb4fff

	//cards := []byte{0x01,0x01, 0x02,WILDCARD, 0x04, 0x05, 0x06, 0x07}

	//cards = replaceWildcard(cards, 0x43, false)
	//胡
	//
	//ti:= time.Now().Unix()
	//var v uint32
	////for i:=0;i<1000000;i++{
	//	Sort(cards, 0, len(cards)-1)
	//	v=existHu(cards, []uint32{},[]uint32{},[]uint32{},0,0)
	////}

	//t.Logf("-==-%d %b %b",time.Now().Unix()-ti ,v,HU_SEVEN_PAIR)

	//<-time.After(time.Hour)

	//t.Logf("1个财神%v 2个财神%v 3个财神%v",v&HU_GUI_WEI1 > 0,v&HU_GUI_WEI2> 0,v&HU_GUI_WEI3> 0)
	//t.Logf("1财1刻%v 2财1刻%v 3财1刻%v",v&HU_CAI_1 > 0,v&HU_CAI_2> 0,v&HU_CAI_3> 0)
	//t.Logf("是胡牌：%+b %+b %+b",v&HU_CAI_1,v&HU_CAI_2,v&HU_CAI_3)
}
