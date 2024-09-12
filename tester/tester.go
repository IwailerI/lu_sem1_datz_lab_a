package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const CPP_PATH = "main.cpp"
const PY_PATH = "main.py"

var tests = []testGroup{
	{
		"Happy path, triangle exists, input is valid",
		[]testCase{
			{3, 4, 5, VALID, CRASH, CRASH},
			{30, 40, 50, VALID, CRASH, CRASH},
			{2.35, 1.23, 2.03, VALID, CRASH, CRASH},
		},
	},
	{
		"One of the sides is exactly the same length as sum of other 2",
		[]testCase{
			{1, 2, 3, INVALID, CRASH, CRASH},
			{0.1, 0.2, 0.3, INVALID, CRASH, CRASH},
			{37, 125, 37 + 125, INVALID, CRASH, CRASH},
		},
	},
	{
		"Some of the sides are invalid",
		[]testCase{
			{-4, 5, 6, INVALID, CRASH, CRASH},
			{4, 5, -6, INVALID, CRASH, CRASH},
			{4, -5, -6, INVALID, CRASH, CRASH},
			{4, 5, 0, INVALID, CRASH, CRASH},
			{4, -5, 0, INVALID, CRASH, CRASH},
		},
	},
}

type output byte

const (
	VALID output = iota
	INVALID
	CRASH
)

func (o output) String() string {
	switch o {
	case VALID:
		return "VALID"
	case INVALID:
		return "INVALID"
	case CRASH:
		return "CRASH"
	default:
		return fmt.Sprintf("<output %d>", o)
	}
}

type testCase struct {
	a, b, c float64
	result  output
	cppOut  output
	pyOut   output
}

func (tc testCase) rotate() testCase {
	return testCase{tc.b, tc.c, tc.a, tc.result, tc.cppOut, tc.pyOut}
}

type testGroup struct {
	note string
	tc   []testCase
}

func testCpp(t []testGroup) error {
	cmd := exec.Command("g++", CPP_PATH, "-o", "./a.out", "-D", "SIMPLE_IO")
	if err := cmd.Run(); err != nil {
		return err
	}
	defer os.Remove("a.out")

	log.Print("C++ compiled, testing...")

	for tgi, tg := range t {
		for i := range tg.tc {
			cmd := exec.Command("./a.out")
			input := new(bytes.Buffer)
			output := new(bytes.Buffer)
			cmd.Stdin = input
			cmd.Stdout = output

			fmt.Fprintf(input, "%f %f %f\n", tg.tc[i].a, tg.tc[i].b, tg.tc[i].c)

			if err := cmd.Run(); err != nil {
				tg.tc[i].cppOut = CRASH
				continue
			}

			var outputNum int
			fmt.Fscan(output, &outputNum)

			if outputNum == 1 {
				tg.tc[i].cppOut = VALID
			} else {
				tg.tc[i].cppOut = INVALID
			}
		}
		log.Printf("%d/%d groups.", tgi+1, len(t))
	}

	return nil
}

func testPy(t []testGroup) error {
	log.Print("Testing Python...")
	for tgi, tg := range t {
		for i := range tg.tc {
			cmd := exec.Command("python", PY_PATH, "--", "--stripped")
			input := new(bytes.Buffer)
			output := new(bytes.Buffer)
			cmd.Stdin = input
			cmd.Stdout = output

			fmt.Fprintf(input, "%v %v %v\n", tg.tc[i].a, tg.tc[i].b, tg.tc[i].c)

			if err := cmd.Run(); err != nil {
				tg.tc[i].pyOut = CRASH
				continue
			}

			var outputNum int
			fmt.Fscan(output, &outputNum)

			if outputNum == 1 {
				tg.tc[i].pyOut = VALID
			} else {
				tg.tc[i].pyOut = INVALID
			}
		}
		log.Printf("%d/%d groups.", tgi+1, len(t))
	}

	return nil
}

func buildFile(t []testGroup) error {
	f, err := os.Create("./test_cases_generated.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	sep := strings.Repeat("=", 60)

	for groupIdx, group := range t {
		fmt.Fprintf(w, "%s\nGroup %d:\n%s\n%s\n\n",
			sep, groupIdx+1, group.note, sep)

		for testIdx, t := range group.tc {
			fmt.Fprintf(w, "Test case: %d\n", (groupIdx+1)*100+testIdx+1)
			fmt.Fprintf(w, "Input: %v %v %v\n", t.a, t.b, t.c)
			fmt.Fprintf(w, "Expected: %v\n", t.result)
			fmt.Fprintf(w, "C++: %v\n", t.cppOut)
			fmt.Fprintf(w, "Python: %v\n\n", t.pyOut)

			if t.result != t.cppOut || t.result != t.pyOut {
				fmt.Fprintln(w, "TEST FAILED ---------------------------")
			}
		}
	}

	return nil
}

func main() {
	t := []testGroup{}
	for _, tg := range tests {
		newTg := testGroup{
			note: tg.note,
		}

		for _, test := range tg.tc {
			newTg.tc = append(newTg.tc,
				test, test.rotate(), test.rotate().rotate())
		}

		t = append(t, newTg)
	}
	log.Print("Tests generated")

	if err := testPy(t); err != nil {
		log.Panic(err)
	}
	log.Print("Python tested")

	if err := testCpp(t); err != nil {
		log.Panic(err)
	}
	log.Print("C++ tested")

	if err := buildFile(t); err != nil {
		log.Panic(err)
	}
	log.Print("Results saved")

}
