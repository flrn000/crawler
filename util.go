package main

func (cfg *config) addPageVisit(normalizedURL string) (isFirstTime bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL] += 1
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}
