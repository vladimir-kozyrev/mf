package parse

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/vladimir-kozyrev/mfh/dto"
)

func GetTargets(f *os.File) ([]string, error) {
	var targets []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if isTargetDeclaration(line) {
			targets = append(targets, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return targets, nil
}

const targetDeclarationRegexpPattern = "^[a-z].*:"

func isTargetDeclaration(s string) bool {
	match, err := regexp.MatchString(targetDeclarationRegexpPattern, s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return match
}

func RemoveAllAfterFirstColon(s string) string {
	return strings.Split(s, ":")[0]
}

func GetTargetsWithContent(f *os.File) ([]*dto.Target, error) {
	var inTarget bool
	var currTarget *dto.Target
	var targets []*dto.Target

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if isTargetDeclaration(line) {
			currTarget.Declaration = line
			inTarget = true
			continue
		}
		if inTarget {
			if isTargetContent(line) {
				currTarget.Content += line
			} else {
				targets = append(targets, currTarget)
				inTarget = false
				currTarget = &dto.Target{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return targets, nil
}

const targetContentRegexpPattern = "^\t"

func isTargetContent(s string) bool {
	match, err := regexp.MatchString(targetContentRegexpPattern, s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return match
}
