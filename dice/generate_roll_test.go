package dice

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func generateRoll2(prev [][]byte, d Die) [][]byte {
	var buff bytes.Buffer
	del := byte(Delim[0])
	for _, roll := range prev {
		for _, s := range d.Sides {
			buff.Write(roll)
			if len(roll) != 0 {
				buff.WriteByte(del)
			}
			buff.WriteString(s)
			buff.WriteByte('\n')
		}
	}
	return bytes.Split(buff.Bytes(), []byte{'\n'})
}

func generateRoll5(prev [][]byte, d Die) [][]byte {
	var buff bytes.Buffer
	var nl byte = 0x0a
	for _, roll := range prev {
		for _, s := range d.Sides {
			buff.Write(roll)
			if len(roll) != 0 {
				buff.WriteString(Delim)
			}
			buff.WriteString(s)
			buff.WriteString("\n")
		}
	}
	return bytes.Split(buff.Bytes(), []byte{nl})
}

func generateRoll3(prev [][]byte, d Die) [][]byte {
	var buff bytes.Buffer
	var nl byte = '\n'
	var byt []byte
	del := byte(Delim[0])

	for _, roll := range prev {
		for _, s := range d.Sides {
			buff.Write(roll)
			if len(roll) != 0 {
				byt = []byte(s)
			} else {
				byt = append([]byte{del}, []byte(s)...)
			}
			buff.Write(append(byt, nl))
		}
	}
	return bytes.Split(buff.Bytes(), []byte{nl})
}

func generateRoll4(prev *[][]byte, d Die) {
	del := byte(Delim[0])
	for _, roll := range *prev {
		for _, s := range d.Sides {
			if len(roll) != 0 {
				roll = append(roll, del)
			}
			roll = append(roll, []byte(s)...)
			roll = append(roll, '\n')
		}
	}
}

func generateRoll6(prev bytes.Buffer, d Die) bytes.Buffer {
	var err error
	var line []byte
	var out bytes.Buffer
	var wg sync.WaitGroup

	del := []byte(Delim)
	length := prev.Len()

	for err != io.EOF {
		// Read a line.
		line, err = prev.ReadBytes('\n')
		if err == io.EOF && length > 0 {
			break
		}
		wg.Add(len(d.Sides))
		for _, s := range d.Sides {
			go func(data []byte, side string) {
				defer wg.Done()
				// Only add the delimiter if
				// there is previous data to append.
				if len(data) > 1 {
					out.Write(data[:len(data)-1])
					out.Write(del)
				}
				out.WriteString(side)
				out.WriteByte('\n')
			}(line, s)
		}
		wg.Wait()
	}
	return out
}

func BenchmarkGenRoll(b *testing.B) {
	d6 := D6()
	f, err := os.Open(filepath.Join(os.Getenv("GOPATH"), "data", "8d6.csv"))
	if err != nil {
		b.Fatal(err)
	}
	raw, err := ioutil.ReadAll(f)
	if err != nil {
		b.Fatal(err)
	}
	data := bytes.Split(raw, []byte("\n"))
	b.ResetTimer()
	b.Run("1", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var buff bytes.Buffer
			for _, elem := range data {
				buff.Write(elem)
			}
			out := generateRoll(buff, d6)
			_ = bytes.Split(out.Bytes(), []byte{'\n'})
		}
	})
	b.ResetTimer()
	b.Run("2", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generateRoll2(data, d6)
		}
	})
	b.ResetTimer()
	b.Run("5", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generateRoll5(data, d6)
		}
	})
	b.ResetTimer()
	b.Run("3", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = generateRoll3(data, d6)
		}
	})
	b.ResetTimer()
	b.Run("6", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var buff bytes.Buffer
			for _, elem := range data {
				buff.Write(elem)
			}
			out := generateRoll6(buff, d6)
			_ = bytes.Split(out.Bytes(), []byte{'\n'})
		}
	})
	b.ResetTimer()
	b.Run("4", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			generateRoll4(&data, d6)
		}
	})
}
