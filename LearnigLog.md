june-12- 
Learnt 
UI:
1. ng new <project name>  - will create basic project structure for angular
2. window.location.origin - will give the url . 
Backend:
1.git switch -c <branch name> - will create branch and switch to the created branch


july 2:
select{} // waits  untils the case satisfies
sub.Drain() // waits for subscibed unprocessed msg to complete and unsubscibe 
sync.WaitGroup (and sync.Mutex) should be declared as values, not pointers


july 4:
//after setting header only atlast you should write resp 
w.WriteHeader(http.StatusAccepted)
	w.Write(resp) 

//  WaitGroup.wait()  should be called atlast 
//http  should be called as async and before wait so that port is listening, async  beacuse it will be struck on listenandserve line itself if called sync