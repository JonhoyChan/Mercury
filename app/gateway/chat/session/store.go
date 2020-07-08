/******************************************************************************
 *
 *  Description:
 *
 *  Session management.
 *
 *****************************************************************************/

package session

import (
	"outgoing/app/gateway/chat/stats"
	"outgoing/x"
	"outgoing/x/ksuid"
	"outgoing/x/log"
	"outgoing/x/types"

	"outgoing/x/websocket"
)

var GlobalSessionStore = NewSessionStore()

// SessionStore holds live sessions. Long polling sessions are stored in a linked list with
// most recent sessions on top. In addition all sessions are stored in a map indexed by session ID.
type SessionStore struct {
	cache Cache
}

// NewSessionStore initializes a session store.
func NewSessionStore() *SessionStore {
	ss := &SessionStore{
		cache: NewDefaultCache(),
	}

	stats.RegisterInt("LiveSessions")
	stats.RegisterInt("TotalSessions")

	return ss
}

// NewSession creates a new session and saves it to the session store.
func (ss *SessionStore) NewSession(conn interface{}, sid string, serverID string) {
	var s Session

	if sid == "" {
		s.sid = ksuid.New().String()
	} else {
		s.sid = sid
	}

	s.serverID = serverID

	if ss.cache.Existed(s.sid) {
		// TODO: change to panic or log.Fatal
		panic(x.Sprintf("duplicate session ID", s.sid))
	}

	switch c := conn.(type) {
	case *websocket.Connection:
		s.proto = WEBSOCKET
		s.ws = c
	default:
		s.proto = NONE
	}

	if s.proto != NONE {
		//s.subs = make(map[string]*Subscription)
		s.send = make(chan []byte, sendQueueLimit+32) // buffered
		s.stop = make(chan []byte, 1)                 // Buffered by 1 just to make it non-blocking
		s.detach = make(chan string, 64)              // buffered
	}

	ss.cache.Store(s.sid, &s)

	if s.proto == WEBSOCKET {
		// Do work in goroutines to return from serveWebSocket() to release file pointers.
		// Otherwise "too many open files" will happen.
		go s.readLoop()
		go s.writeLoop()

		log.Info("ws: session started", log.Ctx{"sid": s.sid, "count": ss.cache.Length()})
	}
}

// Get fetches a session from store by session ID.
func (ss *SessionStore) Get(sid string) *Session {
	return ss.cache.Load(sid)
}

// Delete removes session from store.
func (ss *SessionStore) Delete(s *Session) {
	ss.cache.Delete(s.sid)
}

// Shutdown terminates sessionStore. No need to clean up.
// Don't send to clustered sessions, their servers are not being shut down.
func (ss *SessionStore) Shutdown() {
	ss.cache.Shutdown()
	log.Debug("SessionStore shut down.", "sessions terminated", ss.cache.Length())
}

// EvictUser terminates all sessions of a given user.
func (ss *SessionStore) EvictUser(uid types.Uid, skipSid string) {
	ss.cache.EvictUser(uid, skipSid)
}
