package libs

import (
	"sync"
	"time"

	"go.mau.fi/whatsmeow/types"
)

// SpamDetector for detect spam pattern that hits multiple times in n Seconds
type SpamDetector struct {
	// Max reply delay
	Seconds      float64
	mapHits      sync.Map
	mapIgnoreJID sync.Map
	Active       bool
}

// Create new SpamDetector
func NewSpamDetector(s float64, act bool) *SpamDetector {
	return &SpamDetector{
		mapHits: sync.Map{},
		Seconds: s,
		Active:  act,
	}
}

// Store ignored jid
func (b *SpamDetector) StoreIgnore(jid types.JID) {
	b.mapIgnoreJID.Store(jid.ToNonAD().User, time.Now())
}

// Delete jid from ignored list
func (b *SpamDetector) DeleteIgnore(jid types.JID) {
	b.mapIgnoreJID.Delete(jid.ToNonAD().User)
}

// Check if jid is in ignored list
// * it should be check separately if you need to mark
// * a number as an ignored jid
func (b *SpamDetector) Ignored(jid types.JID) bool {
	_, ok := b.mapHits.Load(jid.ToNonAD().User)
	return ok
}

// Store to Hits
func (b *SpamDetector) Store(pattern string, sent time.Time) {
	b.mapHits.Store(pattern, sent)
}

// Check if pattern exists and not expired in .Seconds also delete if it is
func (b *SpamDetector) Exists(pattern string, sent time.Time) bool {
	if tt, ok := b.mapHits.Load(pattern); ok {
		if past, ok := tt.(time.Time); ok {
			if sent.Sub(past).Seconds() <= b.Seconds {
				return true
			} else {
				b.mapHits.Delete(pattern)
			}
		}
	}

	return false
}

// Check if the pattern are hits more than 1 time in .Seconds
func (b *SpamDetector) IsSpam(pattern string, t time.Time) bool {

	if b.Exists(pattern, t) {
		return true
	}

	b.Store(pattern, t)

	return false
}


// Create Antispam
var Antispam = NewSpamDetector(3, false)
