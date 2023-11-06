package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

var (
	DebugOn = false
)

func main() {
	flag.BoolVar(&DebugOn, "debug", false, "enable debugging statements")
	flag.Parse()

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("can't read input: %v", err)
	}

	value := strings.TrimSuffix(string(bytes), "\n")

	parsed := ParseIfJSON(value)
	if s, ok := parsed.(string); ok {
		os.Stdout.WriteString(s)
		os.Stdout.WriteString("\n")
		return
	}

	output, err := json.MarshalIndent(parsed, "", "    ")
	if err != nil {
		log.Fatalf("can't unmarshal data: %s", err)
	}

	os.Stdout.Write(output)
	os.Stdout.WriteString("\n")
}

func ParseIfJSON(value any) any {

	debug("parsing %#v", value)

	var bytes []byte
	switch val := value.(type) {
	case []byte:
		bytes = val
	case string:
		bytes = []byte(val)
	case []any:
		for i, v := range val {
			val[i] = ParseIfJSON(v)
		}
		return val
	case map[string]any:
		for k, v := range val {
			val[k] = ParseIfJSON(v)
		}
		return val
	default:
		debug("it's a %T, not a string/[]byte/composite", value)
		return value
	}

	var data any
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		debug("parsing %#v failed: %v", value, err)
		return value
	}

	debug("parsed as %T", data)
	return ParseIfJSON(data)
}

func debug(f string, args ...any) {
	if DebugOn {
		log.Printf(f, args...)
	}
}
