package schedule

import "gotgpeon/logger"

func (s *peonSchedule) CleanOutdated() {
	err := s.DeletedService.DeleteOutdatedRecordList()
	if err != nil {
		logger.Errorf("CleanOutdatedRecord err: %s", err.Error())
	}
}
