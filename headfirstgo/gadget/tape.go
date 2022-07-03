package gadget

import "fmt"

// TapePlayer struct
type TapePlayer struct {
	Batteries string
}

// Play method
func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

// Stop method
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

// TapeRecorder struct
type TapeRecorder struct {
	Microphones int
}

// Play method
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

// Record method
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

// Stop method
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
