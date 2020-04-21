-- save these lines in a file called
-- setupworld.sql
 
-- turn off feedback for cleaner display
 
SET feedback off
 
-- 3 x 3 world
 
-- alive has coordinates of living cells
 
DROP TABLE alive;
 
CREATE TABLE alive (x NUMBER,y NUMBER);
 
INSERT INTO alive VALUES (2,1);
INSERT INTO alive VALUES (2,2);
INSERT INTO alive VALUES (2,3);
 
commit;
 
-- save these lines in a file called
newgeneration.SQL
 
-- adjact contains one row for each pair of 
-- coordinates that is adjacent to a living cell
 
DROP TABLE adjacent;
 
CREATE TABLE adjacent (x NUMBER,y NUMBER);
 
-- add row for each of the 8 adjacent squares
 
INSERT INTO adjacent SELECT x-1,y-1 FROM alive; 
INSERT INTO adjacent SELECT x-1,y FROM alive;
INSERT INTO adjacent SELECT x-1,y+1 FROM alive;
INSERT INTO adjacent SELECT x,y-1 FROM alive; 
INSERT INTO adjacent SELECT x,y+1 FROM alive;
INSERT INTO adjacent SELECT x+1,y-1 FROM alive; 
INSERT INTO adjacent SELECT x+1,y FROM alive;
INSERT INTO adjacent SELECT x+1,y+1 FROM alive;
 
commit;
 
-- delete rows for squares that are outside the world
 
DELETE FROM adjacent WHERE x<1 OR y<1 OR x>3 OR y>3;
 
commit;
 
-- table counts is the number of live cells
-- adjacent to that point
 
DROP TABLE counts;
 
CREATE TABLE counts AS 
SELECT x,y,COUNT(*) n
FROM adjacent a
GROUP BY x,y;
 
--    C   N                 new C
--    1   0,1             ->  0  # Lonely
--    1   4,5,6,7,8       ->  0  # Overcrowded
--    1   2,3             ->  1  # Lives
--    0   3               ->  1  # It takes three to give birth!
--    0   0,1,2,4,5,6,7,8 ->  0  # Barren
 
-- delete the ones who die
 
DELETE FROM alive a
WHERE
((a.x,a.y) NOT IN (SELECT x,y FROM counts)) OR
((SELECT c.n FROM counts c WHERE a.x=c.x AND a.y=c.y) IN 
(1,4,5,6,7,8));
 
-- insert the ones that are born
 
INSERT INTO alive a
SELECT x,y FROM counts c
WHERE c.n=3 AND
((c.x,c.y) NOT IN (SELECT x,y FROM alive));
 
commit;
 
-- create output table
 
DROP TABLE output;
 
CREATE TABLE output AS
SELECT rownum y,' ' x1,' ' x2,' ' x3
FROM dba_tables WHERE rownum < 4;
 
UPDATE output SET x1='*'
WHERE (1,y) IN
(SELECT x,y FROM alive);
 
UPDATE output SET x2='*'
WHERE (2,y) IN
(SELECT x,y FROM alive);
 
UPDATE output SET x3='*'
WHERE (3,y) IN
(SELECT x,y FROM alive);
 
commit
 
-- output configuration
 
SELECT x1||x2||x3 WLD
FROM output
ORDER BY y DESC;