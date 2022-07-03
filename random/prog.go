package main

// types, the implementation is using value semantics for both.

type IP []byte     // nasa net package sa standard lib
type IPMask []byte // based on existing data structure === walang iba "slice of bytes"
// pinakaimportanteng data struct sa GO
// mask is using a value receive and returning a value of type IP.
// this method is using value semantics for type IP.

func (ip IP) Mask(mask IPMask) IP { // nakakalito naman ung func na to bagosakin
	if len(mask) == IPv6len && len(ip) == IPv4len && allFF(mask[:12]) {
		mask = mask[12:]
	}
	if len(mask) == IPv4len && len(ip) == IPv6len && bytesEqual(ip[:12], v4InV6Prefix) {
		ip = ip[12:]
	}
	n := len(ip)
	if n != len(mask) {
		return nil
	}
	out := make(IP, n)
	for i := 0; i < n; i++ {
		out[i] = ip[i] & mask[i]
	}
	return out
}

// ipEmptryString accepts a value of type IP and returns a value of type
// the Function is using value semantics for type IP.

func ipEmptryString(ip IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}

// should time use value or pointer semantics? if you need to modify
// a value should you mutate the value or create a new one?
// cguro: have your own copy? tas mutate it in ISOLATION. para walang madamay

type Time struct {
	sec  int64
	nsec int32
	loc  *location // using pointer here & *
}

// factory functions dictate the semantics that will be used. The Now
// returns a value of type Time. This means we should be using value
// semantics and copy Time values.

func Now() Time {
	sec, nsec := now()
	return Time{sec + unixToInternal, nsec, Local}
}

func (t Time) Add (d Duration) Time {
	t.sec += int64(d / 1e9)
	nsec := int32(t.nsec) + int32(d%1e9)
	if nsec >= 1e9{
		t.sec++
		nsec -= 1e9
	} else if nsec < 0 {
		t.sec--
		nsec += ie9
	}
	t.nsec = nsec
	return t
}

// div accepts a value of type Time and returns values of built-in type
// the function is using value semantics for type Time.

func div(t Time, d Duration) (qmod2 int, r Duration) {

	// code here
}

// the only use pointer semantics for the 'Time' api are these
// unmarshal related functions

func (t *Time) UnmarshalBinary(data []byte) error {}
func (t *Time) GobDecode(data []byte) error {}
func (t *Time) UnmarshalJSON(data []byte) error {}
func (t *Time) UnmarshalText(data []byte) error {}

// factory functions dictate the semantics that will be used. the open 
// returns a pointer of type File. This means we should be using pointer
// semantics and share File values.

func Open(name string) (file *File, err error) {
	return OpenFile(name, O_RDONLY, 0)
}

// chdir is using a pointer receive. this method is using pointer
// semantics for File.

func (f *File) Chdir() error {
	if f == nil {
		return ErrInvalid
	}
	if e := syscall.Fchdir(f.fd); e != nil {
		return &PathError{"chdir", f.name, e}
	}
	return nil
}

// epipecheck accepts a pointer of type File.
// the function is using pointer semantics for type File.

func epipecheck(file *File, e error) {
	if e == syscall.EPIPE {
		if atomic.AddInt32(&file.nepipe, 1) >= 10 {
			sigpipe()
		}
	} else {
		atomic.StoreInt32((&file.nepipe, 0)
	}
}