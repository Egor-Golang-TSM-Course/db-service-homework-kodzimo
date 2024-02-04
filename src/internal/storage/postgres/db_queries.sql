--- posts table creation

--- columns table creation
create table blog.comments
(
    id             integer
        constraint id
            primary key,
    post_id        integer
        constraint post_id
            references blog.posts,
    author         integer
        constraint author
            references blog.users,
    content        varchar(1000),
    creation_date date
);

--- tags table creation
create table blog.tags
(
    id            serial
        constraint tags_id
            primary key,
    tag           varchar(16),
    creation_date date
);

--- tags-posts-links table creation
create table blog."tags-posts-links"
(
    post_id integer
        constraint post_id
            references blog.posts,
    tag_id  integer
        constraint tag_id
            references blog.tags
);


