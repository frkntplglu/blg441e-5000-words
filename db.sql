CREATE TYPE word_level AS ENUM ('A1-A2','B1-B2', 'C1-C2');
create table word
(
    id          serial not null
        constraint word_pk
            primary key,
    vocabulary  varchar(255)                                      not null,
    definition  varchar(255)                                      not null,
    sentence    text,
    translation text                                              not null,
    level       word_level,
    is_premium  boolean
);

create table quiz
(
    id          serial not null
        constraint quiz_pk
            primary key,
    title  varchar(255)                                      not null,
    level       word_level
);

create table question
(
    title    text not null,
    option_a varchar,
    option_b varchar,
    option_c varchar,
    option_d varchar,
    answer   text not null,
    quiz_id  integer,
    id       serial
        constraint question_pk
            primary key
);

INSERT INTO word (id, vocabulary, definition, sentence, translation, level, is_premium) VALUES (1, 'wizard', 'usta', 'He''s a financial wizard.', 'O bir finans ustası.', 'C1-C2', false);
INSERT INTO word (id, vocabulary, definition, sentence, translation, level, is_premium) VALUES (2, 'plumbing', 'sıhhi tesisat', 'He had an interest in a plumbing supply store his brother Ralph operated.', 'Kardeşi Ralph''in işlettiği sıhhi tesisat dükkanına ilgi duyuyordu.', 'C1-C2', false);
INSERT INTO word (id, vocabulary, definition, sentence, translation, level, is_premium) VALUES (3, 'limb', 'uzuv', 'Cold numbs his limbs.', 'Soğuk, uzuvlarını hissizleştiriyor.', 'C1-C2', false);
INSERT INTO word (id, vocabulary, definition, sentence, translation, level, is_premium) VALUES (4, 'captive', 'esir', 'He is the captive of his own fears.', 'O kendi korkularının esiridir.', 'C1-C2', false);


INSERT INTO quiz (id, title, level) VALUES (1, 'A1-A2 Quiz 1', 'A1-A2');
INSERT INTO quiz (id, title, level) VALUES (2, 'A1-A2 Quiz 2', 'A1-A2');
INSERT INTO quiz (id, title, level) VALUES (3, 'A1-A2 Quiz 3', 'A1-A2');
INSERT INTO quiz (id, title, level) VALUES (4, 'B1-B2 Quiz 1', 'B1-B2');
INSERT INTO quiz (id, title, level) VALUES (5, 'C1-C2 Quiz 1', 'C1-C2');


INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('My mother cooks ----', 'ask', 'well', 'each', 'pay', 'well', 1, 1);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('I want to be a ----.', 'teacher', 'again', 'safe', 'well', 'teacher', 1, 2);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('The Berlin ---- came down in 1989.', 'women', 'well', 'wall', 'ask', 'wall', 1, 3);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('It isn''t ---- to leave the house after dark.', 'safe', 'look', 'women', 'thing', 'safe', 1, 4);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('I''d like you to give me an honest ----.', 'answer', 'ask', 'bus', 'women', 'answer', 1, 5);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('He cut the cake into six pieces and gave ---- child a slice.', 'weekend', 'each', 'again', 'fish', 'each', 2, 6);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('They were happy because they had caught a lot of ---- that day.', 'women', 'teacher', 'fish', 'nationality', 'fish', 2, 7);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('What''s that ---- over there?', 'teacher', 'women', 'thing', 'well', 'thing', 2, 8);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('What a lovely ----!', 'women', 'weekend', 'safe', 'dress', 'dress', 2, 9);
INSERT INTO question (title, option_a, option_b, option_c, option_d, answer, quiz_id, id) VALUES ('---- over there - there''s a rainbow!', 'bus', 'safe', 'look', 'again', 'look', 2, 10);
