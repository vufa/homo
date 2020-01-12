package rule

// sinksub subscription of sink
type sinksub struct {
	id      string
	qos     uint32
	topic   string
	tqos    uint32
	ttopic  string
	channel *msgchan
}

// Newsinksub creates a new subscription of sink
func newSinkSub(subid string, subqos uint32, subtopic string, pubqos uint32, pubtopic string, channel *msgchan) *sinksub {
	return &sinksub{
		id:      subid,
		qos:     subqos,
		topic:   subtopic,
		tqos:    pubqos,
		ttopic:  pubtopic,
		channel: channel,
	}
}
