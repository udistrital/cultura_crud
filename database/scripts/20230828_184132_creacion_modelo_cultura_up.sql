-- object: cultura | type: SCHEMA --
-- DROP SCHEMA IF EXISTS cultura CASCADE;
-- ddl-end --

CREATE SCHEMA cultura;

CREATE TABLE cultura.actividad_cultural
(
	id_actividad_cultural integer NOT NULL   DEFAULT NEXTVAL(('"actividad_cultural_id_actividad_cultural_seq"'::text)::regclass),
	nombre varchar(100) NOT NULL,
	descripcion varchar(300) NOT NULL,
	estado integer NOT NULL,
	tipo_actividad integer NOT NULL,
	fecha_inicio timestamp NOT NULL,
	fecha_fin timestamp NOT NULL,
	lugar varchar(50) NOT NULL,
	necesita_inscripcion integer NOT NULL,
	enlace_inscripcion varchar(300) NOT NULL,
	posee_mayor_info integer NOT NULL,
	enlace_mayor_info varchar(300) NOT NULL,
	imagen varchar(300) NOT NULL,
	fecha_creacion timestamp NOT NULL, 
	fecha_modificacion timestamp NULL,
	usuario_registra varchar(50) NOT NULL
)
;

CREATE TABLE cultura.actividad_grupo_cultural
(
	id_actividad_grupo_cultural integer NOT NULL   DEFAULT NEXTVAL(('"actividad_grupo_cultural_id_actividad_grupo_cultural_seq"'::text)::regclass),
	id_grupo_cultural integer NOT NULL,
	id_activdad_cultural integer NOT NULL
)
;

CREATE TABLE cultura.evidencia_actividad_cultural
(
	id_evidencia_actividad_cultural integer NOT NULL   DEFAULT NEXTVAL(('"evidencia_actividad_cultural_id_evidencia_actividad_cultural_seq"'::text)::regclass),
	id_actividad_cultural integer NOT NULL,
	categoria_evidencia integer NOT NULL,
	contenido_evidencia varchar(300) NOT NULL
)
;

CREATE TABLE cultura.grupo_cultural
(
	id_grupo_cultural integer NOT NULL   DEFAULT NEXTVAL(('"grupo_cultural_id_grupo_cultural_seq"'::text)::regclass),
	nombre varchar(50) NOT NULL,
	estado integer NOT NULL,
	descripcion varchar(250) NOT NULL,
	e_mail varchar(50) NOT NULL,
	imagen varchar(300) NOT NULL,
	necesita_inscripcion integer NOT NULL,
	enlace_inscripcion varchar(300) NULL,
	fecha_inicio_inscripcion timestamp NULL,
	fecha_fin_inscripcion timestamp NULL
)
;

CREATE TABLE cultura.horarios_grupo_cultural
(
	id_horario_grupo_cultural integer NOT NULL   DEFAULT NEXTVAL(('"horarios_grupo_cultural_id_horario_grupo_cultural_seq"'::text)::regclass),
	id_grupo_cultural integer NOT NULL,
	lugar_reunion varchar(50) NOT NULL,
	dia_reunion varchar(25) NOT NULL,
	hora_reunion varchar(10) NOT NULL
)
;

/* Create Primary Keys, Indexes, Uniques, Checks */

ALTER TABLE cultura.actividad_cultural ADD CONSTRAINT PK_ActividadCultural
	PRIMARY KEY (id_actividad_cultural)
;

ALTER TABLE cultura.actividad_grupo_cultural ADD CONSTRAINT PK_ActividadGrupoCultural
	PRIMARY KEY (id_actividad_grupo_cultural)
;

CREATE INDEX IXFK_ActividadGrupoCultural_ActividadCultural 
ON cultura.actividad_grupo_cultural (id_activdad_cultural ASC)
;

CREATE INDEX IXFK_ActividadGrupoCultural_GrupoCultural 
ON cultura.actividad_grupo_cultural (id_grupo_cultural ASC)
;

ALTER TABLE cultura.evidencia_actividad_cultural ADD CONSTRAINT PK_EvidenciaActividadCultural
	PRIMARY KEY (id_evidencia_actividad_cultural)
;

CREATE INDEX IXFK_EvidenciaActividadCultural_ActividadCultural 
ON cultura.evidencia_actividad_cultural (id_actividad_cultural ASC)
;

ALTER TABLE cultura.grupo_cultural ADD CONSTRAINT PK_GrupoCultural
	PRIMARY KEY (Id_grupo_cultural)
;

ALTER TABLE cultura.horarios_grupo_cultural ADD CONSTRAINT PK_HorarioGrupoCultural
	PRIMARY KEY (id_horario_grupo_cultural)
;

CREATE INDEX IXFK_HorarioGrupoCultural_GrupoCultural 
ON cultura.horarios_grupo_cultural (id_grupo_cultural ASC)
;

/* Create Foreign Key Constraints */

ALTER TABLE cultura.actividad_grupo_cultural ADD CONSTRAINT FK_ActividadGrupoCultural_ActividadCultural
	FOREIGN KEY (id_activdad_cultural) REFERENCES cultura.actividad_cultural (id_actividad_cultural) 
	ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE cultura.actividad_grupo_cultural ADD CONSTRAINT FK_ActividadGrupoCultural_GrupoCultural
	FOREIGN KEY (id_grupo_cultural) REFERENCES cultura.grupo_cultural (Id_grupo_cultural) 
	ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE cultura.evidencia_actividad_cultural ADD CONSTRAINT FK_EvidenciaActividadCultural_ActividadCultural
	FOREIGN KEY (id_actividad_cultural) REFERENCES cultura.actividad_cultural (id_actividad_cultural) 
	ON DELETE No Action ON UPDATE No Action
;

ALTER TABLE cultura.horarios_grupo_cultural ADD CONSTRAINT FK_HorarioGrupoCultural_GrupoCultural
	FOREIGN KEY (id_grupo_cultural) REFERENCES cultura.grupo_cultural (Id_grupo_cultural) 
	ON DELETE No Action ON UPDATE No Action
;

/* Create Table Comments, Sequences for Autonumber Columns */

CREATE SEQUENCE cultura.actividad_cultural_id_actividad_cultural_seq INCREMENT 1 START 1
;

CREATE SEQUENCE cultura.actividad_grupo_cultural_id_actividad_grupo_cultural_seq INCREMENT 1 START 1
;

CREATE SEQUENCE cultura.evidencia_actividad_cultural_id_evidencia_actividad_cultural_seq INCREMENT 1 START 1
;

CREATE SEQUENCE cultura.grupo_cultural_id_grupo_cultural_seq INCREMENT 1 START 1
;

CREATE SEQUENCE cultura.horarios_grupo_cultural_id_horario_grupo_cultural_seq INCREMENT 1 START 1
;
