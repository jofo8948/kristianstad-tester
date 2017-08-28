DROP TABLE IF EXISTS results;
DROP TABLE IF EXISTS logs;
DROP TABLE IF EXISTS resultsets;

-- Table Results
CREATE TABLE results (
    id bigserial PRIMARY KEY,
    url character varying(250) NOT NULL,
    comment character varying(250),
    start_date timestamp without time zone NOT NULL,
    duration interval NOT NULL,
    statuscode integer NOT NULL,
    size integer NOT NULL,
    iteration integer NOT NULL,
    resultset integer
);


ALTER TABLE results OWNER TO root;


-- Table Logs
CREATE TABLE logs (
    id bigserial PRIMARY KEY,
    resultset integer NOT NULL,
    message text
);


ALTER TABLE logs OWNER TO root;


-- Table ResultSets
CREATE TABLE resultsets (
    id bigserial PRIMARY KEY,
    name character varying(50) NOT NULL,
    username character varying(50) NOT NULL,
    start_date timestamp without time zone NOT NULL,
    end_date timestamp without time zone NOT NULL
);


ALTER TABLE resultsets OWNER TO root;


ALTER TABLE ONLY results
    ADD CONSTRAINT results_resultsets_fkey FOREIGN KEY (resultset) REFERENCES resultsets(id);

ALTER TABLE ONLY logs
    ADD CONSTRAINT logs_resultsets_fkey FOREIGN KEY (resultset) REFERENCES resultsets(id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM perf;
GRANT ALL ON SCHEMA public TO perf;
GRANT ALL ON SCHEMA public TO PUBLIC;

GRANT ALL PRIVILEGES ON TABLE ResultSets to perf;
GRANT ALL PRIVILEGES ON TABLE Results to perf;
GRANT ALL PRIVILEGES ON TABLE Logs to perf;
GRANT ALL PRIVILEGES ON SEQUENCE resultsets_id_seq to perf;
GRANT ALL PRIVILEGES ON SEQUENCE results_id_seq to perf;
GRANT ALL PRIVILEGES ON SEQUENCE logs_id_seq to perf;
