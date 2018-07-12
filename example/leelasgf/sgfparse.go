package main

import (
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
	"path/filepath"
	"strings"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"gitee.com/larry_dev/goban"
	"flag"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type LeelaTemp struct {
	Id          int64
	Url         string
	BoardSize   int
	HandsCount  int
	GameDate    string
	BlackPlayer string
	WhitePlayer string
	Result      string
	Stones      int
	Komi        float64
	Difficulty  int
}

const (
	endpoint   string = "oss-cn-hangzhou.aliyuncs.com"
	accessID   string = "fLiqoGFZxpl0Iled"
	accessKey  string = "xL9ZLWwA59jOFPW9NR5ZgHI5aFHcDl"
	bucketName string = "yikeweiqi"
)

var dirName string

func main() {
	flag.StringVar(&dirName, "d", "", "文件")
	flag.Parse()
	if dirName == "" {
		panic(errors.New("请输入正确的文件夹"))
	}
	log.Println("读取文件")
	client, err := oss.New(endpoint, accessID, accessKey)
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
	orm, err := xorm.NewEngine("mysql", "dannyvan:Pa$$w0rd@tcp(101.231.109.3)/tgame?charset=utf8")
	if err != nil {
		log.Printf("filepath.Walk() returned %v\n", err)
		return
	}
	err = filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if (info == nil) {
			return err
		}
		if info.IsDir() {
			return nil
		}
		println(info.Name())
		if strings.Contains(info.Name(), ".sgf") {
			v, _ := ioutil.ReadFile(path)
			kifu := goban.ParseSgf(string(v))
			ossName := "kifu/" + info.Name()
			var (
				date string
				pw   string
				pb   string
				re   string
			)
			if temp, has := kifu.Root.Info["DT"]; has {
				date = temp[0]
			}
			if temp, has := kifu.Root.Info["PB"]; has {
				pb = temp[0]
			}
			if temp, has := kifu.Root.Info["PW"]; has {
				pw = temp[0]
			}
			if temp, has := kifu.Root.Info["RE"]; has {
				re = temp[0]
			}
			l := &LeelaTemp{
				Url:         "res.yikeweiqi.com/" + ossName,
				BoardSize:   kifu.Size,
				HandsCount:  kifu.NodeCount,
				GameDate:    date,
				BlackPlayer: pb,
				WhitePlayer: pw,
				Stones:      kifu.Handicap,
				Komi:        kifu.Komi,
				Difficulty:  5,
				Result:      re,
			}
			_, err := orm.Insert(l)
			if err == nil {
				os.Remove(path)
				options := []oss.Option{
					oss.ContentType("text/plain"),
				}
				bucket.PutObject("kifu/"+info.Name(), strings.NewReader(kifu.ToSgf()), options...)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("filepath.Walk() returned %v\n", err)
	}
}
