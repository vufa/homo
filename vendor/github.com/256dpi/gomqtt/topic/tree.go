package topic

import (
	"fmt"
	"strings"
	"sync"
)

type node struct {
	children map[string]*node
	values   []interface{}
}

func newNode() *node {
	return &node{
		children: make(map[string]*node),
	}
}

func (n *node) removeValue(value interface{}) {
	for i, v := range n.values {
		if v == value {
			// remove without preserving order
			n.values[i] = n.values[len(n.values)-1]
			n.values = n.values[:len(n.values)-1]
			break
		}
	}
}

func (n *node) clearValues() {
	n.values = []interface{}{}
}

func (n *node) string(i int) string {
	str := ""

	if i != 0 {
		str = fmt.Sprintf("%d", len(n.values))
	}

	for key, node := range n.children {
		str += fmt.Sprintf("\n| %s'%s' => %s", strings.Repeat(" ", i*2), key, node.string(i+1))
	}

	return str
}

// A Tree implements a thread-safe topic tree.
type Tree struct {
	separator    string
	wildcardOne  string
	wildcardSome string
	root         *node
	mutex        sync.RWMutex
}

// NewTree returns a new Tree using the specified separator and wildcards.
func NewTree(separator, wildcardOne, wildcardSome string) *Tree {
	return &Tree{
		separator:    separator,
		wildcardOne:  wildcardOne,
		wildcardSome: wildcardSome,
		root:         newNode(),
	}
}

// NewStandardTree returns a new Tree using the standard MQTT separator and
// wildcards.
func NewStandardTree() *Tree {
	return NewTree("/", "+", "#")
}

// Add registers the value for the supplied topic. This function will
// automatically grow the tree. If value already exists for the given topic it
// will not be added again.
func (t *Tree) Add(topic string, value interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// add value
	t.add(value, topic, t.root)
}

func (t *Tree) add(value interface{}, topic string, node *node) {
	// add value to leaf
	if topic == topicEnd {
		// check if duplicate
		for _, v := range node.values {
			if v == value {
				return
			}
		}

		// add value
		node.values = append(node.values, value)

		return
	}

	// get segment
	segment := topicSegment(topic, t.separator)

	// get child
	child, ok := node.children[segment]
	if !ok {
		child = newNode()
		node.children[segment] = child
	}

	// descend
	t.add(value, topicShorten(topic, t.separator), child)
}

// Set sets the supplied value as the only value for the supplied topic. This
// function will automatically grow the tree.
func (t *Tree) Set(topic string, value interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// set value
	t.set(value, topic, t.root)
}

func (t *Tree) set(value interface{}, topic string, node *node) {
	// set value on leaf
	if topic == topicEnd {
		node.values = []interface{}{value}
		return
	}

	// get segment
	segment := topicSegment(topic, t.separator)

	// get child
	child, ok := node.children[segment]
	if !ok {
		child = newNode()
		node.children[segment] = child
	}

	// descend
	t.set(value, topicShorten(topic, t.separator), child)
}

// Get gets the values from the topic that exactly matches the supplied topics.
func (t *Tree) Get(topic string) []interface{} {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// get values
	return t.get(topic, t.root)
}

func (t *Tree) get(topic string, node *node) []interface{} {
	// set value on leaf
	if topic == topicEnd {
		return node.values
	}

	// get segment
	segment := topicSegment(topic, t.separator)

	// get child
	child, ok := node.children[segment]
	if !ok {
		return nil
	}

	// descend
	return t.get(topicShorten(topic, t.separator), child)
}

// Remove un-registers the value from the supplied topic. This function will
// automatically shrink the tree.
func (t *Tree) Remove(topic string, value interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// remove value
	t.remove(value, topic, t.root)
}

// Empty will unregister all values from the supplied topic. This function will
// automatically shrink the tree.
func (t *Tree) Empty(topic string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// empty values
	t.remove(nil, topic, t.root)
}

func (t *Tree) remove(value interface{}, topic string, node *node) bool {
	// clear or remove value from leaf node
	if topic == topicEnd {
		if value == nil {
			node.clearValues()
		} else {
			node.removeValue(value)
		}

		return len(node.values) == 0 && len(node.children) == 0
	}

	// get segment
	segment := topicSegment(topic, t.separator)

	// get child
	child, ok := node.children[segment]
	if !ok {
		return false
	}

	// descend and remove node if empty
	if t.remove(value, topicShorten(topic, t.separator), child) {
		delete(node.children, segment)
	}

	return len(node.values) == 0 && len(node.children) == 0
}

// Clear will unregister the supplied value from all topics. This function will
// automatically shrink the tree.
func (t *Tree) Clear(value interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// clear value
	t.clear(value, t.root)
}

func (t *Tree) clear(value interface{}, node *node) bool {
	// remove value
	node.removeValue(value)

	// remove value from all children and remove empty nodes
	for segment, child := range node.children {
		if t.clear(value, child) {
			delete(node.children, segment)
		}
	}

	return len(node.values) == 0 && len(node.children) == 0
}

// Match will return a set of values from topics that match the supplied topic.
// The result set will be cleared from duplicate values.
//
// Note: In contrast to Search, Match does not respect wildcards in the query but
// in the stored tree.
func (t *Tree) Match(topic string) []interface{} {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	// match values
	var list []interface{}
	t.match(topic, t.root, func(values []interface{}) bool {
		list = append(list, values...)
		return true
	})

	return t.clean(list)
}

// MatchFirst behaves similar to Match but only returns the first found value.
func (t *Tree) MatchFirst(topic string) interface{} {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	// match values
	var value interface{}
	t.match(topic, t.root, func(values []interface{}) bool {
		value = values[0]
		return false
	})

	return value
}

func (t *Tree) match(topic string, node *node, fn func([]interface{}) bool) {
	// add all values to the result set that match multiple levels
	if child, ok := node.children[t.wildcardSome]; ok && len(child.values) > 0 {
		if !fn(child.values) {
			return
		}
	}

	// when finished add all values to the result set
	if topic == topicEnd {
		if len(node.values) > 0 {
			fn(node.values)
		}

		return
	}

	// advance children that match a single level
	if child, ok := node.children[t.wildcardOne]; ok {
		t.match(topicShorten(topic, t.separator), child, fn)
	}

	// get segment
	segment := topicSegment(topic, t.separator)

	// match segments and get children
	if segment != t.wildcardOne && segment != t.wildcardSome {
		if child, ok := node.children[segment]; ok {
			t.match(topicShorten(topic, t.separator), child, fn)
		}
	}
}

// Search will return a set of values from topics that match the supplied topic.
// The result set will be cleared from duplicate values.
//
// Note: In contrast to Match, Search respects wildcards in the query but not in
// the stored tree.
func (t *Tree) Search(topic string) []interface{} {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	// match values
	var list []interface{}
	t.search(topic, t.root, func(values []interface{}) bool {
		list = append(list, values...)
		return true
	})

	return t.clean(list)
}

// SearchFirst behaves similar to Search but only returns the first found value.
func (t *Tree) SearchFirst(topic string) interface{} {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	// match values
	var value interface{}
	t.search(topic, t.root, func(values []interface{}) bool {
		value = values[0]
		return false
	})

	return value
}

func (t *Tree) search(topic string, node *node, fn func([]interface{}) bool) {
	// when finished add all values to the result set
	if topic == topicEnd {
		if len(node.values) > 0 {
			fn(node.values)
		}

		return
	}

	// get segment
	segment := topicSegment(topic, t.separator)

	// add all current and further values
	if segment == t.wildcardSome {
		if len(node.values) > 0 {
			if !fn(node.values) {
				return
			}
		}

		for _, child := range node.children {
			t.search(topic, child, fn)
		}
	}

	// add all current values and continue
	if segment == t.wildcardOne {
		if len(node.values) > 0 {
			if !fn(node.values) {
				return
			}
		}

		for _, child := range node.children {
			t.search(topicShorten(topic, t.separator), child, fn)
		}
	}

	// match segments and get children
	if segment != t.wildcardOne && segment != t.wildcardSome {
		if child, ok := node.children[segment]; ok {
			t.search(topicShorten(topic, t.separator), child, fn)
		}
	}
}

// clean will remove duplicates
func (t *Tree) clean(values []interface{}) []interface{} {
	result := values[:0]

	for _, v := range values {
		if contains(result, v) {
			continue
		}

		result = append(result, v)
	}

	return result
}

// Count will count all stored values in the tree. It will not filter out
// duplicate values and thus might return a different result to `len(All())`.
func (t *Tree) Count() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return t.count(t.root)
}

func (t *Tree) count(node *node) int {
	// prepare total
	total := 0

	// add children to results
	for _, child := range node.children {
		total += t.count(child)
	}

	// add values to result
	return total + len(node.values)
}

// All will return all stored values in the tree.
func (t *Tree) All() []interface{} {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return t.clean(t.all([]interface{}{}, t.root))
}

func (t *Tree) all(result []interface{}, node *node) []interface{} {
	// add children to results
	for _, child := range node.children {
		result = t.all(result, child)
	}

	// add current node to results
	return append(result, node.values...)
}

// Reset will completely clear the tree.
func (t *Tree) Reset() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.root = newNode()
}

// String will return a string representation of the tree.
func (t *Tree) String() string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	return fmt.Sprintf("topic.Tree:%s", t.root.string(0))
}

func contains(list []interface{}, value interface{}) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

var topicEnd = "\x00"

func topicShorten(topic, separator string) string {
	i := strings.Index(topic, separator)
	if i >= 0 {
		return topic[i+1:]
	}

	return topicEnd
}

func topicSegment(topic, separator string) string {
	i := strings.Index(topic, separator)
	if i >= 0 {
		return topic[:i]
	}

	return topic
}
