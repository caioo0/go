package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing", source)

	p.progress = 0
	for p.progress < 100 {
		time.Sleep(time.Millisecond * 100)
		p.progress += 10
		fmt.Printf("Playing %s %d%%\n", source, p.progress)
	}

	fmt.Println("\nfinished Playing", source, "done")
}
