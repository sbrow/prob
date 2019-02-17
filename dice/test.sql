-- drop table if exists dice.t_2_d_d6;
-- create table dice.t_2_d_d6 (
-- 	roll_id serial primary key
-- 	,d1_side int references dice.sides(side_id)
-- 	,d2_side int references dice.sides(side_id)
-- );

-- WITH b as (
--     select side as d2_side
--     from dice.sides
--     where die_id=1
-- )
-- INSERT INTO dice.t_2_d_d6 (d1_side, d2_side)
-- SELECT a.d1_side, b.d2_side
-- FROM b 
-- CROSS JOIN dice.t_1_d_d6 as a
-- ORDER BY d1_side, d2_side ASC;

-- with a(d1) as (
--     VALUES (1), (2), (3), (4), (5), (6)
-- ),
-- b(d2) as (
--     select d1 from a
-- ),
-- c(d3) as (
--     select d1 from a
-- )
-- select a.d1, b.d2, c.d3
-- FROM a
-- CROSS JOIN b
-- CROSS JOIN c
-- ORDER BY d1, d2, d3 asc
DROP TABLE IF EXISTS dice.t_3d6;

WITH t_1(d1) AS (
        SELECT side FROM dice.sides WHERE die_id=1
), t_2(d2) AS (
        SELECT side FROM dice.sides WHERE die_id=1
), t_3(d3) AS (
        SELECT side FROM dice.sides WHERE die_id=1
)
SELECT row_number() over (ORDER BY d1, d2, d3 ASC) as roll_id, t_1.d1, t_2.d2, t_3.d3 INTO dice.t_3d6
FROM t_1
CROSS JOIN t_2
CROSS JOIN t_3;