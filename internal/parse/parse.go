package parse

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	targetDeclarationRegexpPattern = "^[a-z_-]+:"
	targetContentRegexpPattern     = "^\t"
)

type target struct {
	Name        string
	Declaration string
	Content     string
}

func GetTargets(f *os.File) ([]string, error) {
	var targets []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if isTargetDeclaration(line) {
			targets = append(targets, removeAllAfterFirstColon(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return targets, nil
}

func isTargetDeclaration(s string) bool {
	match, err := regexp.MatchString(targetDeclarationRegexpPattern, s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return match
}

func removeAllAfterFirstColon(s string) string {
	return strings.Split(s, ":")[0]
}

func GetTargetsWithContent(f *os.File) ([]*target, error) {
	var inTarget bool
	var targets []*target
	var currTarget *target

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if isTargetDeclaration(line) {
			currTarget = &target{}
			currTarget.Declaration = line
			currTarget.Name = removeAllAfterFirstColon(line)
			inTarget = true
			continue
		}
		if inTarget {
			if isTargetContent(line) {
				currTarget.Content += line
			} else {
				targets = append(targets, currTarget)
				inTarget = false
				currTarget = nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if currTarget != nil {
		targets = append(targets, currTarget)
	}

	return targets, nil
}

func isTargetContent(s string) bool {
	match, err := regexp.MatchString(targetContentRegexpPattern, s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return match
}
