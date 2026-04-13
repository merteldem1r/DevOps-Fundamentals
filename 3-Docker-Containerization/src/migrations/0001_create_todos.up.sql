CREATE TYPE todo_priority AS ENUM ('HIGH', 'MEDIUM', 'LOW');

CREATE TABLE
    todos (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
        title VARCHAR(255) NOT NULL,
        description TEXT,
        priority todo_priority NOT NULL
    );