package transmission

func (t *Transmission) currentToken() string {
	t.tokenLock.Lock()
	defer t.tokenLock.Unlock()

	return t.token
}

func (t *Transmission) updateToken(token string) {
	t.tokenLock.Lock()
	defer t.tokenLock.Unlock()

	t.token = token
}
