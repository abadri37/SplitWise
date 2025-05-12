package observer

import "SplitWise/internal/logger"

type Observer interface {
	Notify(event string)
}

type UserNotifier struct {
	UserID string
}

func (n *UserNotifier) Notify(event string) {
	logger.GetLogger("Observer").Infof("Notify User %s: %s", n.UserID, event)
}
