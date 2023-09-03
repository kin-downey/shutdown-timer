package main

import (
	"fmt"
	"time"
	"os"
	"math"
	"os/exec"
)


func main() {
	// 現在時刻を取得する
	t := time.Now()

	// 引数が2つでなければ終了する（Arg[0]は勝手に入るので，条件式は3と比較している）
	if (len(os.Args) != 3) {
		fmt.Println("invalid argument")
		return
	}
	str_time := os.Args[1] + " " + os.Args[2]

	// 引数で取得した時刻をJSTに変換する
	arg_time, _ := time.ParseInLocation("2006-01-02 15:04:05", str_time, time.Local)

	// 現在時刻と引数で取得した時刻の差分（分）を取得する
	diff := arg_time.Sub(t)
	diff_min := math.Ceil(diff.Minutes())
	// 差分が0より小さければ終了する
	if (diff_min < 0) {
		fmt.Println("invalid argument\nargument time must be later than now")
		return
	}

	// shutdownコマンドを実行する
	out, err := exec.Command("sudo", "shutdown", "-h", "+" + fmt.Sprintf("%.0f", diff_min)).Output()

	// エラーが発生した場合は終了する
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("結果：%s\n", out)
}