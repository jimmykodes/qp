package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"os"
	"strconv"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		panic(err)
	}
}

func run(args []string) error {
	var (
		src  io.ReadCloser
		dest io.WriteCloser
		err  error
	)
	switch len(args) {
	case 0:
		src = os.Stdin
		dest = os.Stdout
	case 1:
		src, err = os.Open(args[0])
		if err != nil {
			return err
		}
		dest = os.Stdout
	case 2:
		src, err = os.Open(args[0])
		if err != nil {
			return err
		}
		dest, err = os.Create(args[1])
		if err != nil {
			return err
		}
	default:
		panic("invalid number of args")
	}
	defer src.Close()
	defer dest.Close()
	input, err := io.ReadAll(src)
	if err != nil {
		return err
	}
	input = bytes.TrimSpace(input)

	u := url.URL{RawQuery: string(input)}

	final  := make(map[string]interface{})

	for key, values := range u.Query() {
		if len(values) == 1 {
			final[key] = convert(values[0])
		} else {
			v := make([]interface{}, len(values))
			for i, value := range values {
				v[i] = convert(value)
			}
			final[key] = v
		}
	}
	return json.NewEncoder(dest).Encode(final)
}

func convert(v string) interface{} {
	if i, err := strconv.Atoi(v); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(v, 64); err == nil {
		return f
	}
	if b, err := strconv.ParseBool(v); err == nil {
		return b
	}
	return v
}
