package main

import "log"

// Für den Fall, dass wir das Logging austauschen wollen ohne
// die Abhängigkeiten zu brechen
type LogI interface {
	Printf(format string, a ...any)
}

type Log struct {
}

func (l *Log) Printf(format string, a ...any) {
	log.Printf(format, a...)
}
