package stdlib

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url`
}

type user struct {
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	Languages []string `json:"languages"`
}

type FormInput struct {
	Name string `json:"name"`
	CompanyName string `json: "company_name, omitempty"`
}

type Bottle struct {
	Name string `json:"name"`
	Price int `json:"price,omitempty"`
	KCal *int `json:"kcal,omitempty"`
}

type Marshal interface {
	MarshalJson() ([]byte, error)
}

type UnMarshal interface {
	UnmarshalJson() ([]byte, error)
}

type Record struct {
	ProcessID string `json:"process_id"`
	DeletedAt JSTime `json:"deleted_at"`
}

type JSTime time.Time

func (t JSTime) MarshalJson() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	v := strconv.Itoa(int(tt.UnixMilli()))
	return []byte(v), nil
}

type Country struct {
	Name string `csv:"国名"`
	ISOCode string `csv:"ISOコード"`
	Population int `csv:"人口"`
}

type record struct {
	Number int `csv:"number"`
	Message string `csv:message"`
}



func main() {
	// f, err := os.Open("ip.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// var resp ip
	// if err := json.NewDecoder(f).Decode(&resp); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", resp)

	// jsonのエンコード方法
	// var b bytes.Buffer
	// u := user{
	// 	UserID: "001",
	// 	UserName: "gopher",
	// }
	// _ = json.NewEncoder(&b).Encode(u)
	// fmt.Printf("%v\n", b.String())

	// u := user{
	// 	UserID: "001",
	// 	UserName: "gopher",
	// }
	// b, _ := json.Marshal(u)
	// fmt.Println(string(b))

	// in := FormInput{Name: "山田太郎"}

	// b, _ := json.Marshal(in)
	// fmt.Println(string(b))

	// b := Bottle {
	// 	Name: "ミネラルウォーター",
	// 	Price: 0,
	// 	KCal: Int(0),
	// }

	// out, _ := json.Marshal(b)
	// fmt.Println(string(out))

	// f, err := os.Open("country.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// r := csv.NewReader(f)
	// // r.Comma = '\t'
	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(record)
	// }

	// records := [][]string {
	// 	{"書籍名", "出版年", "ページ数"},
	// 	{"Go言語によるWebアプリケーション開発", "2016", "280"},
	// 	{"Go言語による並列処理", "2018", "256"},
	// 	{"Go言語で作るインタプリタ", "2018", "316"},
	// }
	// f, err := os.OpenFile("oreilly.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// w := csv.NewWriter(f)
	// defer w.Flush() //バッファにあるデータを書き込む

	// for _, record := range records {
	// 	if err := w.Write(record); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// if err := w.Error(); err != nil {
	// 	log.Fatal(err)
	// }

	// gocsv
	// lines := []Country {
	// 	{Name: "アメリカ合衆国", ISOCode: "US/USA", Population: 310232863},
	// 	{Name: "日本", ISOCode: "JP/JPN", Population: 127288000},
	// 	{Name: "中国", ISOCode: "CN/CHN", Population: 1330044000},
	// }

	// f, err := os.Create("country.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// if err := gocsv.MarshalFile(&lines, f); err != nil {
	// 	log.Fatal(err)
	// }

	// f, err := os.Open("country.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// var lines []Country
	// if err := gocsv.UnmarshalFile(f, &lines); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, v := range lines {
	// 	fmt.Printf("%+v\n", v)
	// }

	// 大規模なcsv
	// c := make(chan interface{})
	// go func() {
	// 	defer close(c)
	// 	for i := 0; i < 1000*1000; i++ {
	// 		c <- record{
	// 			Message: "Hello",
	// 			Number: i + 1,
	// 		}
	// 	}
	// 	return 
	// }()
	// if err := gocsv.MarshalChan(c, gocsv.DefaultCSVWriter(os.Stdout)); err != nil {
	// 	log.Fatal(err)
	// }
	out := excelize.NewFile()
	out.SetCellValue("Sheet1", "A1", "Hello Excel")
	if err := out.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	//ファイルの読み込み
	in, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	cell, err := in.GetCellValue("Sheet1", "A1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cell)
}

func Int(v int) *int {
	return &v
}