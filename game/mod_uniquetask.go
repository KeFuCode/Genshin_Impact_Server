package game

import _"sync"

type TaskInfo struct {
	TaskId int
	State  int
}

type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo
	// Locker     *sync.RWMutex
}

func (self *ModUniqueTask) IsTaskFinish(taskId int) bool {
	// assume uniquetask 10001 and 10002 finished
	if taskId == 10001 || taskId == 10002 {
		return true
	}

	task, ok := self.MyTaskInfo[taskId]
	if !ok {
		return false
	}

	return task.State == TASK_STATE_FINASH
}
