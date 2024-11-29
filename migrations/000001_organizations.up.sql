CREATE TABLE organizations
(
    id          UUID PRIMARY KEY,
    owner_id    UUID         NOT NULL,
    photo_id    UUID         NOT NULL,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    phone       VARCHAR(20)  NOT NULL,
    email       VARCHAR(255) NOT NULL,
    status      VARCHAR(50)  NOT NULL CHECK (status IN ('pending', 'approved', 'rejected')),
    longtitude  FLOAT        NOT NULL,
    latitude    FLOAT        NOT NULL
);