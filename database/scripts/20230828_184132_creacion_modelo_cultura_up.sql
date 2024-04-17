-- object: cultura | type: SCHEMA --
-- DROP SCHEMA IF EXISTS cultura CASCADE;
-- ddl-end --

CREATE SCHEMA cultura;

CREATE TABLE cultura.actividad_cultural
(
	id serial not null,
	nombre varchar(100) not null,
	descripcion varchar(300) not null,
	estado integer not null,
	tipo_actividad integer not null,
	fecha_inicio timestamp null,
	fecha_fin timestamp null,
	lugar varchar(50) null,
	necesita_inscripcion integer null,
	enlace_inscripcion varchar(300) null,
	posee_mayor_info integer null,
	enlace_mayor_info varchar(300) null,
	imagen varchar(300) null,
	usuario_registra varchar(50) NOT null,
	activo boolean not null,
	fecha_creacion timestamp not null, 
	fecha_modificacion timestamp not null
	
)
;

CREATE TABLE cultura.actividad_grupo_cultural
(
	id serial not null,
	grupo_cultural_id integer not null,
	actividad_cultural_id integer not null,
	activo boolean not null,
	fecha_creacion timestamp not null, 
	fecha_modificacion timestamp not null
)
;

CREATE TABLE cultura.evidencia_actividad_cultural
(
	id serial not null,
	actividad_cultural_id integer not null,
	categoria_evidencia integer not null,
	contenido_evidencia varchar(300) not null,
	activo boolean not null,
	fecha_creacion timestamp not null, 
	fecha_modificacion timestamp not null
)
;

CREATE TABLE cultura.grupo_cultural
(
	id serial not null,
	nombre varchar(50) not null,
	estado integer not null,
	descripcion varchar(250) not null,
	e_mail varchar(50) not null,
	imagen varchar(300) not null,
	necesita_inscripcion integer not null,
	enlace_inscripcion varchar(300) null,
	fecha_inicio_inscripcion timestamp null,
	fecha_fin_inscripcion timestamp null,
	lider_grupo varchar(50) not null,
	activo boolean not null,
	fecha_creacion timestamp not null, 
	fecha_modificacion timestamp not null
)
;

CREATE TABLE cultura.horario_grupo_cultural
(
	id serial not null,
	grupo_cultural_id integer not null,
	lugar_reunion varchar(50) not null,
	dia_reunion varchar(25) not null,
	hora_reunion varchar(10) not null,
	activo boolean not null,
	fecha_creacion timestamp not null, 
	fecha_modificacion timestamp not null
)
;

/* Create Primary Keys, Indexes, Uniques, Checks */

ALTER TABLE cultura.actividad_cultural ADD CONSTRAINT pk_actividad_cultural
	PRIMARY KEY (id)
;

ALTER TABLE cultura.actividad_grupo_cultural ADD CONSTRAINT pk_actividad_grupo_cultural
	PRIMARY KEY (id)
;

ALTER TABLE cultura.evidencia_actividad_cultural ADD CONSTRAINT pk_evidencia_actividad_cultural
	PRIMARY KEY (id)
;

ALTER TABLE cultura.grupo_cultural ADD CONSTRAINT pk_grupo_cultural
	PRIMARY KEY (id)
;

ALTER TABLE cultura.horario_grupo_cultural ADD CONSTRAINT pk_horario_grupo_cultural
	PRIMARY KEY (id)
;

CREATE INDEX idx_actividad_grupo_cultural_actividad_cultural_id
ON cultura.actividad_grupo_cultural (actividad_cultural_id ASC)
;

CREATE INDEX idx_actividad_grupo_cultural_grupo_cultural_id
ON cultura.actividad_grupo_cultural (grupo_cultural_id ASC)
;


CREATE INDEX idx_evidencia_actividad_cultural_actividad_cultural_id 
ON cultura.evidencia_actividad_cultural (actividad_cultural_id ASC)
;

CREATE INDEX idx_horario_grupo_cultural_grupo_cultural_id
ON cultura.horario_grupo_cultural (grupo_cultural_id ASC)
;

/* Create Foreign Key Constraints */

ALTER TABLE cultura.actividad_grupo_cultural ADD CONSTRAINT fk_actividad_grupo_cultural_actividad_cultural
	FOREIGN KEY (actividad_cultural_id) REFERENCES cultura.actividad_cultural (id) 
	ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE cultura.actividad_grupo_cultural ADD CONSTRAINT fk_actividad_grupo_cultural_grupo_cultural
	FOREIGN KEY (grupo_cultural_id) REFERENCES cultura.grupo_cultural (id) 
	ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE cultura.evidencia_actividad_cultural ADD CONSTRAINT fk_evidencia_actividad_cultural_actividad_cultural
	FOREIGN KEY (actividad_cultural_id) REFERENCES cultura.actividad_cultural (id) 
	ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE cultura.horario_grupo_cultural ADD CONSTRAINT fk_horario_grupo_cultural_grupo_cultural
	FOREIGN KEY (grupo_cultural_id) REFERENCES cultura.grupo_cultural (id) 
	ON DELETE No Action ON UPDATE No Action
;
