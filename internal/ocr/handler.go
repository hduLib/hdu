//go:build export
package ocr

import "sync"

type HandlerSync struct {
	mu *sync.RWMutex
	Handler
}

type Handler struct {
	captcha []byte
	typ     YunmaOCRType
	token   string
}

var (
	syncH *HandlerSync // for concurrent usage
	o     sync.Once
)

// InitSync a concurrent safe ocr handler and returns it. Call multi-times
// is ok as there's only one instance across the package.
func InitSync() *HandlerSync {
	o.Do(newHandlerSync)
	return syncH
}

func newHandlerSync() {
	syncH = &HandlerSync{
		mu:      &sync.RWMutex{},
		Handler: Handler{},
	}
}

func (h *HandlerSync) SetToken(token string) *HandlerSync {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.token = token
	return h
}

func (h *HandlerSync) SetType(typ YunmaOCRType) *HandlerSync {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.typ = typ
	return h
}

var h *Handler

// Init a ocr handler and returns it. Call multi-times is ok as there's
// only one instance across the package. Also see `InitSync`
func Init() *Handler {
	if h != nil {
		h = &Handler{
			captcha: make([]byte, 0),
			typ:     Common,
			token:   "",
		}
	}
	return h
}

func (h *Handler) SetToken(token string) *Handler {
	h.token = token
	return h
}

func (h *Handler) SetType(typ YunmaOCRType) *Handler {
	h.typ = typ
	return h
}
