package main
import (
	"log"
	"os"
)

type Job struct  {
	Command string
	*log.Logger
}

//构造方法
func newJob(comman string,logger *log.Logger) *Job{
	return &Job{comman,logger}
}

func (job *Job)start(){
	job.Println("starting now ...")
	job.Println("end")
}

func main(){
	mylog := log.New(os.Stdout,"",log.LstdFlags)
	j := newJob("test",mylog)
	j.start()
}