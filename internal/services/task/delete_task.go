package services

func (s *Service) DeleteTask(id int) error {
	return s.taskRepository.Delete(id)
}
