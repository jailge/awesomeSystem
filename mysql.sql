CREATE TABLE permission (
                                id BIGINT auto_increment NOT NULL,
                                resource varchar(100) NOT NULL,
                                `action` varchar(100) NOT NULL,
                                CONSTRAINT permission_PK PRIMARY KEY (id)
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci
AUTO_INCREMENT=1;


CREATE TABLE role (
                            id BIGINT auto_increment NOT NULL,
                            role_name varchar(100) NOT NULL,
                            CONSTRAINT role_PK PRIMARY KEY (id)
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci
AUTO_INCREMENT=1;


CREATE TABLE role_permission (
                                     id BIGINT auto_increment NOT NULL,
                                     role_id BIGINT NOT NULL,
                                     permission_id BIGINT NOT NULL,
                                     CONSTRAINT role_permission_PK PRIMARY KEY (id)
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci
AUTO_INCREMENT=1;



CREATE TABLE user_role (
                               id BIGINT auto_increment NOT NULL,
                               user_id BIGINT NOT NULL,
                               role_id BIGINT NOT NULL,
                               CONSTRAINT user_role_PK PRIMARY KEY (id)
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci
AUTO_INCREMENT=1;
