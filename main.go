package main
import (
"io/ioutil"
"net/http"
"regexp"
"strconv"
)
func main(){
	number:=1
	output :=make([]LianJia,0)
	url:=""
	for i:=8;i<=37;i++{
		end:=1
		for number=1;number<=end;number++{
			//https://sz.lianjia.com/ditiezufang/li110460692s100021208/rt200600000001erp5000/#contentList
			//https://sz.lianjia.com/ditiezufang/li110460692s100021208/pg2rt200600000001erp5000/#contentList
			//8-37-----------------``----
			url = "https://sz.lianjia.com/ditiezufang/li110460692s1000212"+GETi(i)+"/pg" +strconv.Itoa(number)+"rt200600000001erp5000/#contentList"
			println(url)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Add("cache-control", "no-cache")
			res, err := http.DefaultClient.Do(req)
			if err!=nil{
				println("GET error")
				return
			}
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)

			ioutil.WriteFile("response.html",body,0644)
			//fmt.Println(res)
			//fmt.Println(string(body))
			/*	x:=MatchLine(string(body))
				for _,v:=range x{
					println(v)
				}
			*/

			reg0 := regexp.MustCompile(`content__title--hl">[^<]*`)
			result0:=reg0.FindString(string(body))
			println(result0)
			end=GetNumber(result0)
			reg := regexp.MustCompile(`(content__list--item--main[^元]*元)`)
			result:=reg.FindAllString(string(body),-1)
			//println(number,result)
			for _,v:=range result{
				vv,err:=myCut(v,i) // clean and get job detail
				if err !=nil{
					println(err.Error(),vv.title)
				}
				output=append(output,vv)
				//myCut(v)
				//myShow(vv)
			}
		}
	}

	println(len(output))
	Output(output)

}
func GETi(i int )string{
	out:=strconv.Itoa(i)
	if len(out)==1{
		k:="0"+out
		return k
	}
	return out
}
func GetNumber(text string)int{
	b:=string(text[20:len(text)])

	out,err:=strconv.Atoi(b)

	if err==nil{
		w:=(out-1)/30
		return w+1
	}else{
		return -1
	}
}
func Output(result []LianJia)error{
	out:=""
	for _,value:=range result{
		out=out+value.price+","+value.square+","+strconv.Itoa(value.Station)+"\n"
	}

	ioutil.WriteFile("result33.csv",[]byte(out),0644)
	return nil;
}

//https://www.zhipin.com/job_detail/?query=golang&city=101280600&industry=&position=
//https://www.zhipin.com/c101280600/?query=golang&page=2&ka=page-2
type LianJia struct{
	title string
	price string
	address string
	square string
	Station int

}
func myCut(text string ,i int)(LianJia,error){
	a:=LianJia{}
	reg,err := regexp.Compile(`[\d]*㎡`)
	if err!=nil{
		return a,err
	}
	result:=reg.FindAll([]byte(text),-1)
	for _,v:=range result{
		b:=v[:len(v)-3]
		//println(string(b))
		a.square=string(b)
	}
	reg2,err := regexp.Compile(`[\d]*</em> 元`)
	if err!=nil{
		return a,err
	}
	result2:=reg2.FindAll([]byte(text),-1)
	for _,v:=range result2{
		b:=v[:len(v)-9]
		//println(string(b))
		a.price=string(b)
	}
	a.Station=i
	/*reg3,err := regexp.Compile(`[\d]*㎡`)
	if err!=nil{
		return a,err
	}
	result3:=reg3.FindAll([]byte(text),-1)
	for _,v:=range result3{
		println(string(v))
	}*/


	return a,nil
}