package main

import (
	"net/http"
	"encoding/json"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

/*
keyId		STRING(필수)	인증키		OepnAPI 에서 발급된 인증키
serviceId	STRING(필수)	서비스명	요청대상인 해당 서비스명
dataType	STRING(필수)	요청파일 타입	xml : xml파일 , json : json파일
startIdx	STRING(필수)	요청시작위치	정수입력
endIdx		STRING(필수)	요청종료위치	정수입력
url = "http://openapi.foodsafetykorea.go.kr/api/keyId/serviceId/dataType/startIdx/endIdx"
 */
const(
	keyId = "kid"
	serviceId = "I0950" //식품첨가물공전
	dataType = "json"

	url = "http://openapi.foodsafetykorea.go.kr/api/"
)

const (
	DB_HOST = "tcp(192.168.0.202:3306)"
	DB_NAME = "easysafetest"
	DB_USER = /*"root"*/ "con"
	DB_PASS = /*""*/ "concon"
)
/*
PRDLST_CD	품목코드
PC_KOR_NM	품목한글명
TESTITM_CD	시험항목코드
T_KOR_NM	시험항목 한글명
FNPRT_ITM_NM	세부항목명
SPEC_VAL	기준규격값
SPEC_VAL_SUMUP	기준규격값 요약
VALD_BEGN_DT	유효개시일자
VALD_END_DT	유효종료일자
SORC		출처
MXMM_VAL	최대값
MIMM_VAL	최소값
INJRY_YN	위해여부
UNIT_NM		단위명
total_count		총 출력 대상줄수 (한번에 1000개 씩만 조회가 가능하니 회수 돌릴때 사용한다.)
 */
type Apirep struct {

	Sid struct {
		    Result struct {
				   MSG  string        `json:"MSG"`
				   CODE  string        `json:"CODE"`
			   }	`json:"RESULT"`
		    Total_count  string        `json:"total_count"`
		    Row []struct {
			    UPC  string        `json:"PRDLST_CD"`
			    Name  string        `json:"PC_KOR_NM"`
			    Company string       `json:"CMPNY_NM"`
		    }        `json:"row"`
	    }	`json:"I2570"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"

	db, err := sql.Open("mysql", dsn)
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	stmtIns, err := db.Prepare("Insert into tbl_fds_foodupc(upc, name, company) values( ?, ?, ?)")
	check(err)
	defer stmtIns.Close()


	//url = "http://openapi.foodsafetykorea.go.kr/api/keyId/serviceId/dataType/startIdx/endIdx"
	url := url + "/" + keyId +"/" + serviceId +"/"+ dataType +"/"

	res, err := http.Get(url+"1/1000")
	check(err)
	defer res.Body.Close()

	var repsruct Apirep
	json.NewDecoder(res.Body).Decode(&repsruct)

	endIdx, err := strconv.Atoi(repsruct.Sid.Total_count)
	check(err)



	for i := 0; i <= endIdx/1000; i++{
		//
		t := (i+1)*1000
		res, err := http.Get(url+ strconv.Itoa(t-999) +"/" + strconv.Itoa(t))
		if err != nil{
			panic(err)
		}
		defer res.Body.Close()

		json.NewDecoder(res.Body).Decode(&repsruct)

		fmt.Print(strconv.Itoa(t-999) +" ~ " + strconv.Itoa(t))
		fmt.Println("\tMSG : "+ repsruct.Sid.Result.MSG + "\tCODE : "+repsruct.Sid.Result.CODE)

		for j := range repsruct.Sid.Row{
			//Insert into tbl_fds_foodupc(upc, name, company) values( ?, ?, ?)
			stmtIns.Exec(repsruct.Sid.Row[j].UPC, repsruct.Sid.Row[j].Name, repsruct.Sid.Row[j].Company)
		}
	}
}
