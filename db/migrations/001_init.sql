CREATE TABLE
  IF NOT EXISTS process (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    command TEXT NOT NULL,
    interval INTEGER NOT NULL
  );