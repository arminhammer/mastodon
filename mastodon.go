package main

import (
    "fmt"
    "runtime"
    "net/http"
    "flag"
    "os"
    //"time"
    //"./command"
)

/*
func handler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
*/

// Job ...
type Job struct {
    url string
    count int
}

func worker(id int, jobs <-chan Job, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        //time.Sleep(time.Second)
        results <- j * 2
    }
}

func makeRequest(url string) (int) {
    //client := &http.Client{
    //       CheckRedirect: redirectPolicyFunc,
    //}
    resp, _ := http.Get(url)
    return resp.StatusCode
}

func main() {

    role := flag.String("role","command","Select whether to run as a server, worker, or command")
    url := flag.String("url","false","url to load-test [command mode only]")

    num := flag.Int("n",1,"number of requests [command mode only]")
    conc := flag.Int("c",runtime.NumCPU(),"number of requests [command mode only]")

    flag.Parse()

    switch {
    case len(flag.Args()) > 0:
        fmt.Println("Unable to process command line arguments because of: ", flag.Args())
        os.Exit(1)
    case *role == "server":
        fmt.Println("Starting in server mode")
    case *role == "worker":
        fmt.Println("Starting in worker mode")
    case *role == "command" && *url != "false":
        fmt.Println("Starting in command mode")
        fmt.Println("Load-testing", *url, "", *num,"times with concurrency", *conc)


        jobs := make(chan Job)
        results := make(chan int, *num)

        for w := 0; w <= *conc; w++ {
            go worker(w, jobs, results)
        }

        for j := 0; j <= *num; j++ {
            jobs <- j
        }
        close(jobs)

        for a := 0; a <= *num; a++ {
            result := <-results
            fmt.Println("result: ", result)
        }

        /*for i := 0; i < *num; i++ {
            results := make(chan string)
            go makeRequest(*url)
            //status := makeRequest(*url)
            //fmt.Println(status)
        }*/
    default:
        fmt.Println("Improper arguments.")
        fmt.Println(*role)
        fmt.Println(*url)
    }

    /*if len(flag.Args()) > 0 {
    } else if()

    fmt.Printf("%s\n", *role)
    if(*role=="command") {
    fmt.Printf("%s\n", *url)
}

fmt.Println("tail:", flag.Args())
*/
//workerCount :=  runtime.NumCPU()

//fmt.Printf("%d\n", workerCount)

//http.HandleFunc("/", handler)
//http.ListenAndServe(":3699", nil)

}
