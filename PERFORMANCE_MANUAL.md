# General
Provided method of data extraction is basically an `awk` on steroids:
* The most common operation is looking for char/string in a string - this is what string splitting utilities like awk/gawk do when looking for the next field separator.
* There's nothing like in depth lookup like in regex matching. We just find first character/string in the rest of line. It is enough in a huge majority of cases (and it is rather something wrong with the input if this is not enough).

So, the conclusion is it would be basically enough to check performance on regular column-formatted files as it reflects common usage. The file can be generated with [utility](https://github.com/sirkon/ldetool/blob/master/columngen.7z). Don't forget to get a package needed:
```bash
go get github.com/sirkon/message
```
And here is the [benchmark code](https://github.com/sirkon/ldetool/blob/master/benchmarker.7z). Following packages are needed:
```bash
go get github.com/sirkon/message
go get github.com/youtube/vitess/go/cgzip
```
Compile both and make preparations:
```bash
go install columngen
go install main
time columngen 100000000 > data
wc -l data
```
Now, `data` file is cached and we can test performance. We will extract 1st and 4th columns and output them separated again with | in the stdout.

#### LDE with short lookup optimization (faster for short lookups of 1-6 characters)
```
$ time ./main < data | wc -l

real	0m6.565s
user	0m7.020s
sys	0m0.832s
```
#### LDE with regular lookup (faster for longer lookups)
```
$ time ./main < data | wc -l
100000000

real	0m7.454s
user	0m7.916s
sys	0m0.740s
```
#### Gawk
```
$ time  gawk -F '|' '{ print $1 "|" $4 }' data | wc -l
100000000

real	1m2.773s
user	1m3.208s
sys	0m1.520s
```
#### Mawk, should be quite fast
```
$ time  mawk -F '|' '{ print $1 "|" $4 }' data | wc -l
100000000

real	0m20.785s
user	0m20.992s
sys	0m1.008s
```
It is indeed, only three times slower. Still, it doesn't signal errors, it only can work upon columned files and, finally,
it just doesn't have an infrastructure.
#### sed
```
$ time sed -E 's/^(.*?)\|.*?\|.*?\|.*?\|(.*?)\|.*$/\1|\2/g' data | wc -l
100000000

real	7m44.861s
user	7m47.444s
sys	0m5.600s
```
OMG, that was SLOW
#### Go with regular expression with group capture 
Program
```go
package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"

	"github.com/sirkon/message"
	"github.com/youtube/vitess/go/cgzip"
)

func main() {
	var zreader io.Reader
	if len(os.Args) != 2 {
		zreader = os.Stdin
	} else {
		rawreader, err := os.Open(os.Args[1])
		if err != nil {
			message.Critical(err)
		}
		defer rawreader.Close()
		zreader, err = cgzip.NewReaderBuffer(rawreader, 512*1024)
		if err != nil {
			message.Critical(err)
		}
	}
	reader := bufio.NewReaderSize(zreader, 128*1024)
	scanner := bufio.NewScanner(reader)

	dest := bufio.NewWriter(os.Stdout)
	defer dest.Flush()
	buf := &bytes.Buffer{}
	r := regexp.MustCompile(`^(.*?)\|.*?\|.*?\|.*?\|(.*?)\|.*$`)
	for scanner.Scan() {
		data := r.FindSubmatch(scanner.Bytes())
		if len(data) == 3 {
			buf.Reset()
			buf.Write(data[1])
			buf.WriteByte('|')
			buf.Write(data[2])
			buf.WriteByte('\n')
			dest.Write(buf.Bytes())
		}
	}
	if scanner.Err() != nil {
		message.Critical(scanner.Err())
	}
}
```
Launching
```bash
$ time ./goregex < data | wc -l

100000000

real	2m4.064s
user	2m6.928s
sys	0m5.616s
```
Harder to use than LDE generated code, slower, harder to reason when something goes wrong. It was a bit easier to write LDE 
rule than a regexp and significantly easier to access extracted data
```perl
parser =
    Name(string) '|'     # Take a text until | into Name as a string ([]byte, actually), then pass |
    _ '|'                # We are at the start of column 2, find | and pass it
    _ '|'                # find | and pass it to go at column 4
    _ '|'                # find | and pass again
    Count(string) '|';   # take the content of 5th column right to the | and exit
``` 
vs
```
^(.*?)\|.*?\|.*?\|.*?\|(^.*?)\|.*$
```
#### Ragel with

```ragel
package main

import (
        "bufio"
        "io"
        "os"
                "github.com/sirkon/message"
                        "github.com/youtube/vitess/go/cgzip"

        )

// Ragel based parsing
type Ragel struct {
    Name  []byte
    Count []byte
}

%% machine fields;
%% write data;

// Extract extracts field from
func (r *Ragel) Extract(data []byte) (ok bool, error error) {
    cs, p, pe := 0, 0, len(data)
    var pos = 0
    %%{
        action shot       { pos = p + 1             }
        action take_name  { r.Name = data[pos:p+1]  }
        action take_count { r.Count = data[pos:p+1] }

        field = (any -- ( "|" ))* ;
        main :=
             field@take_name "|" field "|" field "|" field "|"@shot field@take_count ;
        write init;
        write exec;
    }%%
    return true, nil
}

func main() {
        var zreader io.Reader
        if len(os.Args) != 2 {
                zreader = os.Stdin
        } else {
                rawreader, err := os.Open(os.Args[1])
                if err != nil {
                        message.Critical(err)
                }
                defer rawreader.Close()
                zreader, err = cgzip.NewReaderBuffer(rawreader, 512*1024)
                if err != nil {
                        message.Critical(err)
                }
        }
        reader := bufio.NewReaderSize(zreader, 128*1024)
        scanner := bufio.NewScanner(reader)

    r := &Ragel{}
                dest := bufio.NewWriter(os.Stdout)
        defer dest.Flush()
        var ok bool
        var err error
        for scanner.Scan() {
            if ok, err = r.Extract(scanner.Bytes()); !ok {
                    if err == nil {
                            panic("Unknown error")
                        }
                        panic(err)
                }
                                        dest.Write(r.Name)
                        dest.WriteByte('|')
                        dest.Write(r.Count)
                        dest.WriteByte('\n')
        }
                if scanner.Err() != nil {
                message.Critical(scanner.Err())
        }
}
```
compiled with `ragel -Z -G2 ragel.template` (turned to be the fastest)
```
$ time ./ragel < data | wc -l
100000000

real	0m7.216s
user	0m7.700s
sys	0m0.816s
```
Quite fast, you see. Still, the dataset is nearly ideal for Ragel, but it wasn't enough to beat LDE with short lookup optimization and it is just barely faster than generic LDE which would destroy it on longer fields.
