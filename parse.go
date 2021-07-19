package main

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type docFileSpec struct {
	name                string
	parent              string
	resourceDescription string
	attributes          map[string]string
	ImportComment       string
	ImportCommand       string
	example             string
}

func parse(rootDir string) (string, []*docFileSpec, error) {
	providerName, err := getProviderName(rootDir)
	if err != nil {
		return "", nil, err
	}

	specs, err := readDocFiles(rootDir)
	if err != nil {
		return "", nil, err
	}

	return providerName, specs, nil
}

func getProviderName(rootDir string) (string, error) {
	b, err := os.ReadFile(filepath.Join(rootDir, "main.go"))
	if err != nil {
		return "", err
	}
	if m := regexp.MustCompile(`ProviderFunc: (\w+)\.`).FindSubmatch(b); m != nil {
		return string(m[1]), nil
	} else {
		return "", errors.New("no provider")
	}
}

func readDocFiles(rootDir string) (specs []*docFileSpec, err error) {
	descriptionPattern := regexp.MustCompile(`#\s+.*\n((?:.|\n)*?)## Example`)
	attributePattern := regexp.MustCompile("[*-]\\s*`([a-z_]+)`\\s*-\\s*(?:\\([^)]*\\))?\\s*(.*)")
	importCommentPattern := regexp.MustCompile("# Import\\n((?:.*|\\n)*?)```")
	importCommandPattern := regexp.MustCompile("```.*\\n\\$ (terraform import.*)")
	examplePattern := regexp.MustCompile(`\n(## Example.*\n(?:.|\n)+?)## (?:Attribute|Argument)`)

	return specs, filepath.WalkDir(filepath.Join(rootDir, "docs"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if filepath.Base(filepath.Dir(path)) == "docs" {
			return nil
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		spec := &docFileSpec{
			name:       strings.TrimSuffix(filepath.Base(path), ".md"),
			parent:     filepath.Base(filepath.Dir(path)),
			attributes: make(map[string]string),
		}

		seenAttributes := make(map[string]bool)

		if m := descriptionPattern.FindSubmatch(b); m != nil {
			spec.resourceDescription = string(bytes.TrimSpace(m[1]))
		}

		if m := attributePattern.FindAllSubmatch(b, -1); m != nil {
			for _, sm := range m {
				name := string(sm[1])
				description := string(bytes.TrimSpace(sm[2]))

				if seenAttributes[name] {
					if spec.attributes[name] != description {
						delete(spec.attributes, name)
					}
					continue
				}

				seenAttributes[name] = true
				spec.attributes[name] = description
			}
		}

		if m := importCommentPattern.FindSubmatch(b); m != nil {
			spec.ImportComment = string(bytes.TrimSpace(m[1]))
		}

		if m := importCommandPattern.FindSubmatch(b); m != nil {
			spec.ImportCommand = string(bytes.TrimSpace(m[1]))
		}

		if m := examplePattern.FindSubmatch(b); m != nil {
			spec.example = string(bytes.TrimSpace(m[1]))
		}

		specs = append(specs, spec)

		return nil
	})
}
