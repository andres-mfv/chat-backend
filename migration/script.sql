CREATE TABLE users (
                       user_id SERIAL PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create a table for channels
CREATE TABLE channels (
                          channel_id SERIAL PRIMARY KEY,
                          name VARCHAR(255) UNIQUE NOT NULL,
                          is_private BOOLEAN NOT NULL DEFAULT FALSE,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create a table for user-channel subscriptions
CREATE TABLE channel_members (
                                 user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
                                 channel_id INT REFERENCES channels(channel_id) ON DELETE CASCADE,
                                 joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (user_id, channel_id)
);

-- Create a table for messages
CREATE TABLE messages (
                          message_id SERIAL PRIMARY KEY,
                          channel_id INT REFERENCES channels(channel_id) ON DELETE CASCADE,
                          user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
                          content TEXT NOT NULL,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);