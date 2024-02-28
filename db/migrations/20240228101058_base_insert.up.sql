INSERT INTO genres (id, title) VALUES (1, 'Fiction'),
                                      (2, 'Historical fiction'),
                                      (3, 'Non-fiction'),
                                      (4, 'Science-fiction'),
                                      (5, 'Narrative'),
                                      (6, 'Historical fantasy'),
                                      (7, 'Mystery'),
                                      (8, 'Novel'),
                                      (9, 'History'),
                                      (10, 'Humor'),
                                      (11, 'Poetry'),
                                      (12, 'Science'),
                                      (13, 'Fairy tale'),
                                      (14, 'Fantasy'),
                                      (15, 'Biography');
-- to use password_hash below you should set app authorization password_salt to 'helloWorld'
INSERT INTO authors (id, username, password_hash, created_at) VALUES (1, 'author-1', '6c475874f9d182e0f9eca934edfa76d8b5763bcf61063195b76959d452a6ce47d6f3abe921ce9e3fe70c421996280c174eafffffedf24fd3f7e9ae7eb5ee923a', now())

