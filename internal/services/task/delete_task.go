package services

import "context"

func (s *taskService) DeleteTask(ctx context.Context, userID, id int64) error {
	_, err := s.GetTaskByID(ctx, userID, id)
	if err != nil {
		return err
	}

	// Delete the task
	return s.taskRepository.Delete(ctx, id)
}
