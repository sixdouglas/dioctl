package gpio

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	. "github.com/stianeikeland/go-rpio"
	"math"
	"syscall"
	"time"
)

const DurationBetweenTwoLocks = 275
const FirstLockDuration = 9900
const SecondLockDuration = 2675

const OnHighDuration = 310
const OnLowDuration = 1340
const OffHighDuration = 310
const OffLowDuration = 310

const NanoToMicro = 1000
const MilliToMicro = 1000

func PulseIn(rpio IRpio, pin rpio.Pin, state rpio.State, timeout int64) (int64, error) {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Input()

	var start time.Time
	var end time.Time
	var current time.Time
	start = time.Now()

	for rpio.ReadPin(pin) != state {
		current = time.Now()
		if current.Sub(start).Nanoseconds() /NanoToMicro >= timeout {
			return 0, syscall.ETIMEDOUT
		} else {
			start = current
		}
	}

	for rpio.ReadPin(pin) != state {
		current = time.Now()
		if current.Sub(start).Nanoseconds() /NanoToMicro >= timeout {
			return 0, syscall.ETIMEDOUT
		} else {
			end = current
		}
	}

	return end.Sub(start).Nanoseconds() / NanoToMicro, nil
}

func ReadCode(rpio IRpio, pin rpio.Pin) (uint64, error) {

	i := 0
	var t int64 = 0
	//previous received bit
	var prevBit uint64 = 0
	//current bit
	var bit uint64 = 0

	//reset of the remote id
	var sender uint64 = 0
	//reset of the group id
	group := false
	//reset on/off state
	on := false
	//reset button row id
	var recipient uint64 = 0

	t, err := PulseIn(rpio, pin, Low, 1000000)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	///lock 1
	for t < 2700 || t > 2800 {
		t, err = PulseIn(rpio, pin, Low, 1000000)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	}

	// Data
	for i < 64 {
		t, err = PulseIn(rpio, pin, Low, 1000000)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}

		//Bit definition (0 or 1)
		if t > 180 && t < 420 {
			bit = 0
		} else if t > 1280 && t < 1480 {
			bit = 1
		} else {
			i = 0
			break
		}

		if i%2 == 1 {
			if (prevBit ^ bit) == 0 {
				// must be either 01 or 10 but must not be 00 or 11 otherwise we stop de detection, this is a bad read
				i = 0
				break
			}

			if i < 53 {
				// the 26 first (0-25) bits are remote for the id
				sender <<= 1
				sender |= prevBit
			} else if i == 53 {
				// the 26th bit is for the group
				group = prevBit == 0
			} else if i == 55 {
				// the 27th bit is the state (on/off)
				on = prevBit == 0
			} else {
				// the last 4 bits (28-32) are for button row id
				recipient <<= 1
				recipient |= prevBit
			}
		}

		prevBit = bit
		i++
	}

	//If data are correctly detected
	if i > 0 {

		fmt.Println("------------------------------")
		fmt.Println("Detected data:")
		fmt.Printf("  sender %d\n", sender)

		if group {
			fmt.Println(" group command")
		} else {
			fmt.Println(" command not grouped")
		}

		if on {
			fmt.Println(" on")
		} else {
			fmt.Println(" off")
		}

		fmt.Printf(" recipient %d\n", recipient)
	} else {
		fmt.Println("NO DATA...")
	}

	return sender, nil
}

func delayMicroseconds(delay int64) {
	var start = time.Now()
	var current time.Time

	for ; ; {
		current = time.Now()
		if current.Sub(start).Nanoseconds() / NanoToMicro >= delay {
			break
		}
	}
}

// Send basic heartbeat (from one state to another)
//    1 = 310µs high then 1340µs low
//    0 = 310µs high then 310µs low
func sendBit(rpio IRpio, pin rpio.Pin, b bool) {
	if b {
		rpio.WritePin(pin, High)
		delayMicroseconds(OnHighDuration) //275 originally, but tweaked.
		rpio.WritePin(pin, Low)
		delayMicroseconds(OnLowDuration) //1225 originally, but tweaked.
	} else {
		rpio.WritePin(pin, High)
		delayMicroseconds(OffHighDuration) //275 originally, but tweaked.
		rpio.WritePin(pin, Low)
		delayMicroseconds(OffLowDuration) //275 originally, but tweaked.
	}
}

// Compute 2^power, used for the conversion of decimal to binary
func power2(power uint64) uint64 {
	return uint64(math.Pow(2, float64(power)))
}

// Convert a number to binary for the sender
func itob(integer uint64, length uint64) []bool {
	var bit = make([]bool, length)
	var i uint64
	const one uint64 = 1
	for i = 0; i < length; i++ {
		if (integer / power2(length - one - i)) == one {
			integer -= power2(length - one - i);
			bit[i] = true
		} else {
			bit[i] = false
		}
	}
	return bit
}

// Send a pulse defined as: 0 =01 et 1 =10
// This is Manchester Coding, to avoid errors
func sendPair(rpio IRpio, pin rpio.Pin, b bool) {
	sendBit(rpio, pin, b)
	sendBit(rpio, pin, !b)
}

// Signal sending function
func transmit(rpio IRpio, pin rpio.Pin, blnOn bool, sender []bool, interrupter []bool) {
	var i int

	// Lock Sequence to wake the receiver
	rpio.WritePin(pin, High)
	delayMicroseconds(DurationBetweenTwoLocks) // just pure noise before starting to reset receiver delays
	rpio.WritePin(pin, Low)
	delayMicroseconds(FirstLockDuration)       // first lock of 9900µs
	rpio.WritePin(pin, High)                   // high again
	delayMicroseconds(DurationBetweenTwoLocks) // wait 275µs between two locks
	rpio.WritePin(pin, Low)                    // second lock of 2675µs
	delayMicroseconds(SecondLockDuration)
	rpio.WritePin(pin, High) // Back to High position to end the lock

	// Send sender code (for example 272946 = 1000010101000110010 in binary)
	for i = 0; i < len(sender); i++ {
		sendPair(rpio, pin, sender[i])
	}

	// Send bit for a grouped command or not in our case (26th bit)
	sendPair(rpio, pin, false)

	// The actual command bit to tell if it has to be on or off (27th bit)
	sendPair(rpio, pin, blnOn)

	// Sending the last 4 bits, representing the switch code, here 0 (encode sur 4 bit donc 0000)
	// nb: for official Chacon remote, switches are named like from 0 to X
	// switch 1 = 0 (so 0000),
	// switch 2 = 1 (so 1000),
	// switch 3 = 2 (si 0100) etc...
	for i = 0; i < 4; i++ {
		sendPair(rpio, pin, interrupter[i])
	}

	rpio.WritePin(pin, High)                   // coupure données, verrou
	delayMicroseconds(DurationBetweenTwoLocks) // attendre 275µs
	rpio.WritePin(pin, Low)                    // verrou 2 de 2675µs pour signaler la fermeture du signal
	delayMicroseconds(SecondLockDuration)      // attendre 275µs
	rpio.WritePin(pin, High)                   // coupure données, verrou
}

func SendCommand(rpio IRpio, pin rpio.Pin, senderId uint64, interrupterId uint64, on bool) {
	sender := itob(senderId, 26)
	interrupter := itob(interrupterId, 4)

	rpio.PinMode(pin, Output)

	if on {
		// send it 5 times to be sure
		for i := 0; i < 5; i++ {
			transmit(rpio, pin, true, sender, interrupter) // send ON
			delayMicroseconds(10 * MilliToMicro) // wait 10 ms (otherwise socket ignores us)
		}

	} else {
		for i := 0; i < 5; i++ {
			transmit(rpio, pin, false, sender, interrupter) // send OFF
			delayMicroseconds(10 * MilliToMicro) // wait 10 ms (otherwise socket ignores us)
		}
	}
}


func Analyse(irpio IRpio, pin rpio.Pin, delay int64) {
	var startTime = time.Now()
	var previousTime = startTime
	var currentTime = previousTime
	var times = make([]int64, 1000)
	var initialState = rpio.ReadPin(pin)
	var previousState = initialState
	var currentState = initialState

	var timeDiff int64
	for i := 0; i < len(times) ; i++ {
		for ; ; {
			currentState = rpio.ReadPin(pin)
			if currentState != previousState {
				currentTime = time.Now()
				timeDiff = currentTime.Sub(previousTime).Nanoseconds()
				times[i] = timeDiff
				previousState = currentState
				previousTime = currentTime
				break
			}
		}
		if timeDiff / NanoToMicro >= delay {
			break
		}
	}

	fmt.Println("Here is a list of all elements currently defined and their status")
	fmt.Println("")
	fmt.Println("  +----------+-------+------------+")
	fmt.Println("  |    i     | STATE |    TIME    |")
	fmt.Println("  +----------+-------+------------+")
	currentState = initialState
	for i := 0 ; i < len(times); i++ {

		fmt.Printf("  | %4d | %8d | %10d |\n", i, currentState, times[i] / NanoToMicro)
		if currentState == rpio.Low {
			currentState = rpio.High
		} else {
			currentState = rpio.Low
		}
	}
	fmt.Println("  +----------+-------+---------------------+")

}