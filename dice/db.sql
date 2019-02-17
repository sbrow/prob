-- ===========
-- BEGIN SETUP
-- ===========
drop schema dice cascade;
create schema dice;

CREATE TABLE dice.dice_types (
	die_id SERIAL PRIMARY KEY
	,name TEXT NOT null
	,short text not null
	,sides INT NOT null CHECK (sides > 0)
);

create table dice.sides (
	side_id SERIAL primary key
	,die_id INT references dice.dice_types(die_id) not null
	,side int check (side > 0)
	,value INT not null default 0
);

insert into dice.dice_types (name, short, sides) values ('V.A.T.S', 'vats', 6);
insert into dice.dice_types (name, short, sides) values ('d6', 'd6', 6);

insert into dice.sides (die_id, side, value) SELECT 1, * FROM (VALUES (1,0), (2,0), (3,1), (4,1), (5,1), (6,2)) AS t;
insert into dice.sides (die_id, side, value) SELECT 2, * FROM (VALUES (1,1), (2,2), (3,3), (4,4), (5,5), (6,6)) AS t;

create table dice.t_1_d_vats (
	roll_id serial primary key
	,d1_side int references dice.sides(side_id)
);

insert into dice.t_1_d_vats (d1_side) 
    select side
    from dice.sides
    where die_id=1
    order by side_id asc;
-- =========
-- END SETUP
-- =========
    
-- Get rolls with value
with vats as (
    select * from dice.sides where die_id = 1
)
select r.*,s.value as d1_value
from dice.t_1_d_vats as r 
inner join vats as s on (r.d1_side = s.side);

-- Get rolls grouped by sum
with vats as (
    select * from dice.sides where die_id = 1
),
total as (
	select count(*) from dice.t_1_d_vats
)
select s.value, count(*), round(count(*)::numeric / (select * from total), 4) as p
from dice.t_1_d_vats as r 
inner join vats as s on (r.d1_side = s.side)
group by s.value
order by s.value asc

/*
// Rolls dice. The reroll function should determine whether or not each die needs to be re-rolled
func Roll(dice []Die, rerolls int, reroll function([]Die) []bool) {
    roll := ...
	for _, re := range reroll(roll) {
		if re {
	    	return Roll([]Die, rerolls-1, reroll)
		}
	}
	return roll
}
 */


with roll as (
	select * from t_2d6
),
reroll_1 as (
	select * from t_2d6 where d1_side=1
)
select count(*) FROM roll
 UNION ALL
 select count(*) from reroll_1;