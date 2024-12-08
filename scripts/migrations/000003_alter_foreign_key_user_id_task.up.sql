ALTER TABLE tasks
ADD CONSTRAINT fk_user_id_tasks FOREIGN KEY (user_id) REFERENCES users(id);