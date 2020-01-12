// Package topic implements common methods to handle MQTT topics.
package topic

import (
	"errors"
	"regexp"
	"strings"
)

// ErrZeroLength is returned by Parse if a topics has a zero length.
var ErrZeroLength = errors.New("zero length topic")

// ErrWildcards is returned by Parse if a topic contains invalid wildcards.
var ErrWildcards = errors.New("invalid use of wildcards")

var multiSlashRegex = regexp.MustCompile(`/+`)

// Parse removes duplicate and trailing slashes from the supplied
// string and returns the normalized topic.
func Parse(topic string, allowWildcards bool) (string, error) {
	// check for zero length
	if topic == "" {
		return "", ErrZeroLength
	}

	// normalize topic
	topic = multiSlashRegex.ReplaceAllString(topic, "/")

	// remove trailing slashes
	topic = strings.TrimRight(topic, "/")

	// check again for zero length
	if topic == "" {
		return "", ErrZeroLength
	}

	// get first segment
	remainder := topic
	segment := topicSegment(topic, "/")

	// check all segments
	for segment != topicEnd {
		// check use of wildcards
		if (strings.Contains(segment, "+") || strings.Contains(segment, "#")) && len(segment) > 1 {
			return "", ErrWildcards
		}

		// check if wildcards are allowed
		if !allowWildcards && (segment == "#" || segment == "+") {
			return "", ErrWildcards
		}

		// check if hash is the last character
		if segment == "#" && topicShorten(remainder, "/") != topicEnd {
			return "", ErrWildcards
		}

		// get next segment
		remainder = topicShorten(remainder, "/")
		segment = topicSegment(remainder, "/")
	}

	return topic, nil
}

// ContainsWildcards tests if the supplied topic contains wildcards. The topics
// is expected to be tested and normalized using Parse beforehand.
func ContainsWildcards(topic string) bool {
	return strings.Contains(topic, "+") || strings.Contains(topic, "#")
}
