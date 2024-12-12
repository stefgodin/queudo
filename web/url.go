package web

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type URLParamValidator func(string) error

type URLMatcher struct {
	Validators map[string]URLParamValidator
}

func (m *URLMatcher) AddValidator(slug string, validator URLParamValidator) {
	m.Validators[slug] = validator
}

var DefaultURLMatcher = &URLMatcher{
	Validators: map[string]URLParamValidator{
		"string": func(string) error { return nil },
		"int": func(tok string) error {
			_, err := strconv.Atoi(tok)
			return err
		},
	},
}

// Matches a path to a pattern and converts any UrlParameter
// UrlParameters are matched to any * parts of a pattern
//
// Ex:
//
// intParam := IntParam()
//
// err := match("/examples/109/delete", "/examples/*/delete", &intParam)
func (m *URLMatcher) match(path string, pattern string) ([]string, error) {
	pathTokens := tokenize(path)
	patternTokens := tokenize(pattern)
	if len(pathTokens) != len(patternTokens) {
		return nil, errors.New("Path does not match pattern length")
	}

	params := []string{}
	for i, pathToken := range pathTokens {
		patternToken := patternTokens[i]
		if patternToken[:1] == ":" {
			typeSlug := patternToken[1:]
			if m.Validators[typeSlug] == nil {
				return nil, errors.New(fmt.Sprintf("No url parameter validator exists with slug: %s", typeSlug))
			}

			err := m.Validators[typeSlug](pathToken)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Path token '%s' failed %s validation: %s", pathToken, typeSlug, err))
			}
			params = append(params, pathToken)
		} else if patternToken != "*" && patternToken != pathToken {
			return nil, errors.New(fmt.Sprintf("Path does not match pattern"))
		}
	}

	return params, nil
}

func tokenize(path string) []string {
	tokens := []string{}
	if len(path) == 0 {
		return tokens
	}

	for _, token := range strings.Split(path, "/") {
		if token != "" { // Remove any empty parts
			tokens = append(tokens, token)
		}
	}
	return tokens
}
