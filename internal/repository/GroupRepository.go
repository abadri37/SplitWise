package repository

import (
	"SplitWise/internal/domain/entity"
	"SplitWise/internal/logger"
	"sync"
)

type GroupRepository struct {
	Groups map[string]*entity.ExpenseGroup
	Mutex  sync.RWMutex
}

func (r *GroupRepository) RegisterGroup(group *entity.ExpenseGroup) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	r.Groups[group.ID] = group
	logger.GetLogger("GroupRepository").Infof("Registered group %s", group.ID)
}
