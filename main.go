package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/gdamore/tcell"
)

const (
	pic = "assets/30c621a657fb4a0bf4234e1f20f7ce91333fd712.png"
)

type pomo struct {
	duration time.Duration
	msg      string
}

func (p pomo) String() string {
	return fmt.Sprintf("(%s) %s", p.duration.String(), p.msg)
}

var (
	re         = regexp.MustCompile("(?i)^y")
	quitKey    = tcell.NewEventKey(tcell.Key, 'q', tcell.ModNone)
	restartKey = tcell.NewEventKey(tcell.Key, 'r', tcell.ModNone)
)

func drawToScreen(timer <-chan pomo) {
	for p := range timer {
		fmt.Println(p)
	}
}

func main() {
	var err error

	workDuration := flag.Duration("work", 20*time.Second, "default work duration")
	breakDuration := flag.Duration("break", 5*time.Second, "default break duration")
	longBreakDuration := flag.Duration("longbreak", 10*time.Second, "default long break duration")
	flag.Parse()

	fullPomodoro := []pomo{
		{
			duration: *workDuration,
			msg:      "time to start working",
		},
		{
			duration: *breakDuration,
			msg:      "time for a break",
		},
		{
			duration: *workDuration,
			msg:      "back to work!",
		},
		{
			duration: *breakDuration,
			msg:      "another break",
		},
		{
			duration: *workDuration,
			msg:      "last work session!",
		},
		{
			duration: *longBreakDuration,
			msg:      "big break time :)",
		},
	}

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	if s.Init() != nil {
		log.Fatal(err)
	}
	s.DisableMouse()
	s.SetContent(0, 0, 'h', nil, tcell.StyleDefault)
	s.Show()
	for {
		s.PollEvent()

	}
	s.Fini()

	// timer := make(chan pomo)
	// go drawToScreen(timer)

	// for {
	// 	for _, p := range fullPomodoro {
	// 		// timer <- p
	// 		// err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	// 		err = beeep.Notify("GoModoro", p.msg, pic)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		time.Sleep(p.duration)
	// 	}
	// 	fmt.Print("Do you want to restart? [y/N] ")
	// 	var input string
	// 	fmt.Scanln(&input)
	// 	if !re.MatchString(input) {
	// 		os.Exit(0)
	// 	}
	// }
}
