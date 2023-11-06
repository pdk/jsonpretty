# jsonpretty

Read (probably) JSON input, and output it with indentation.

Attempts to parse input as JSON, and "pretty print" it (i.e. print with
indentation). If the input is not JSON, then the input is echoed to stdout
without any changes.

Also, JSON parsing will be attempted on any string values found within a
successfully parsed result.

Example:

```
echo '{"a": true, "b": 4,"c": { "x": "y", "y": "{\"a\": 23}"}, "j": "[1, \"treu\", \"true\", false]"}' | go run jsonpretty.go 
{
    "a": true,
    "b": 4,
    "c": {
        "x": "y",
        "y": {
            "a": 23
        }
    },
    "j": [
        1,
        "treu",
        true,
        false
    ]
}
```

Debug logging can be enabled with a `-debug` flag.
