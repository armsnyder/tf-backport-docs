package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func render(rootDir, providerName string, specs []*docFileSpec) error {
	if err := writeImportScripts(rootDir, providerName, specs); err != nil {
		return fmt.Errorf("could not write import scripts: %w", err)
	}

	if err := writeTerraformExamples(rootDir, providerName, specs); err != nil {
		return fmt.Errorf("could not write Terraform examples: %w", err)
	}

	if err := processGoFiles(rootDir, providerName, specs); err != nil {
		return fmt.Errorf("could not process Golang files: %w", err)
	}

	return nil
}

func writeImportScripts(rootDir, providerName string, specs []*docFileSpec) error {
	for _, spec := range specs {
		if spec.ImportComment == "" || spec.ImportCommand == "" {
			continue
		}

		var buf bytes.Buffer
		scanner := bufio.NewScanner(strings.NewReader(spec.ImportComment))
		for scanner.Scan() {
			buf.WriteByte('#')
			if len(scanner.Bytes()) > 0 {
				buf.WriteByte(' ')
				buf.Write(bytes.TrimSpace(scanner.Bytes()))
			}
			buf.WriteByte('\n')
		}
		buf.WriteString(spec.ImportCommand)
		buf.WriteByte('\n')

		exampleDir := filepath.Join(rootDir, "examples", spec.parent, fmt.Sprintf("%s_%s", providerName, spec.name))
		if err := os.MkdirAll(exampleDir, 0755); err != nil {
			return err
		}

		err := func() error {
			f, err := os.Create(filepath.Join(exampleDir, "import.sh"))
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err := buf.WriteTo(f); err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func processGoFiles(rootDir, providerName string, specs []*docFileSpec) error {
	goFilenameToSpec := make(map[string]*docFileSpec)
	for _, spec := range specs {
		switch spec.parent {
		case "resources":
			goFilenameToSpec[fmt.Sprintf("resource_%s_%s.go", providerName, spec.name)] = spec
		case "data-sources":
			goFilenameToSpec[fmt.Sprintf("data_source_%s_%s.go", providerName, spec.name)] = spec
		}
	}

	return filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		spec := goFilenameToSpec[filepath.Base(path)]
		if spec == nil {
			return nil
		}

		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		b, err := rewriteGoContents(src, spec)
		if err != nil {
			return fmt.Errorf("could not rewrite Golang file contents for file %q: %w", path, err)
		}

		if err := os.Remove(path); err != nil {
			return err
		}

		return os.WriteFile(path, b, 0644)
	})
}

func rewriteGoContents(src []byte, spec *docFileSpec) ([]byte, error) {
	attributeKeyPattern := regexp.MustCompile(`"([a-z_]+)": {`)
	typePattern := regexp.MustCompile(`Type:\s+schema.[A-Za-z]+,`)
	resourcePattern := regexp.MustCompile(`return &schema\.Resource{`)
	scanner := bufio.NewScanner(bytes.NewReader(src))

	var buf bytes.Buffer
	var attributeDescription string

	for scanner.Scan() {
		writeLine := func() func() {
			wrote := false
			return func() {
				if wrote {
					return
				}
				buf.Write(scanner.Bytes())
				buf.WriteByte('\n')
				wrote = true
			}
		}()

		if m := attributeKeyPattern.FindSubmatch(scanner.Bytes()); m != nil {
			attributeDescription = spec.attributes[string(m[1])]
		}

		if attributeDescription != "" && typePattern.Match(scanner.Bytes()) {
			writeDescriptionLine(&buf, attributeDescription, false)
			attributeDescription = ""
		}

		if resourcePattern.Match(scanner.Bytes()) {
			writeLine()
			writeDescriptionLine(&buf, spec.resourceDescription, true)
		}

		if !strings.Contains(scanner.Text(), `Description:`) {
			writeLine()
		}
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("could not format source code: %w", err)
	}

	return formatted, nil
}

func writeDescriptionLine(buf *bytes.Buffer, description string, doubleNewline bool) {
	if description == "" {
		return
	}

	buf.WriteString(`Description: "`)
	var lastCh rune
	for _, ch := range description {
		if ch == '\r' {
			continue
		}

		if lastCh == '\n' && ch != '\n' {
			buf.WriteString("\" +\n\"")
		}

		switch ch {
		case '\n':
			buf.WriteString(`\n`)
		case '"':
			buf.WriteString(`\"`)
		case '\\':
			buf.WriteString(`\\`)
		default:
			buf.WriteRune(ch)
		}

		lastCh = ch
	}

	buf.WriteString("\",\n")

	if doubleNewline {
		buf.WriteByte('\n')
	}
}

func writeTerraformExamples(rootDir, providerName string, specs []*docFileSpec) error {
	for _, spec := range specs {
		exampleDir := filepath.Join(rootDir, "examples", spec.parent, fmt.Sprintf("%s_%s", providerName, spec.name))
		if err := os.MkdirAll(exampleDir, 0755); err != nil {
			return err
		}

		err := func() error {
			f, err := os.Create(filepath.Join(exampleDir, spec.parent[:len(spec.parent)-1]+".tf"))
			if err != nil {
				return err
			}
			defer f.Close()

			return writeTerraformExample(f, spec.example)
		}()
		if err != nil {
			return err
		}
	}

	return nil
}

func writeTerraformExample(f io.Writer, example string) error {
	code := false
	scanner := bufio.NewScanner(strings.NewReader(example))
	afterCodeBlock := false

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 && !code {
			continue
		}

		if strings.Contains(scanner.Text(), "```") {
			if code {
				afterCodeBlock = true
			}
			code = !code
			continue
		}

		if code {
			if afterCodeBlock {
				if _, err := f.Write([]byte{'\n'}); err != nil {
					return err
				}
				afterCodeBlock = false
			}
			if _, err := f.Write(scanner.Bytes()); err != nil {
				return err
			}
			if _, err := f.Write([]byte{'\n'}); err != nil {
				return err
			}
		} else {
			comment := bytes.TrimSpace(bytes.Trim(scanner.Bytes(), "*#"))
			if string(comment) != "Example Usage" {
				if afterCodeBlock {
					if _, err := f.Write([]byte{'\n'}); err != nil {
						return err
					}
					afterCodeBlock = false
				}
				if _, err := f.Write([]byte("# ")); err != nil {
					return err
				}
				if _, err := f.Write(comment); err != nil {
					return err
				}
				if _, err := f.Write([]byte{'\n'}); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
