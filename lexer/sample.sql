SELECT
    s1.field1 as f1,
    field2 f2,
    distinct field3,
    (field4*5)/6.0-7%8\3+5.93+s3.field2 f4,
    count(*),
    avg(field1) OVER(PARTITION By field2 order by filed3 desc rows between UNBOUNDED PRECEDING and UNBOUNDED FOLLOWING)
    userfunction(f1)
from
    ~/dir/file1.xlsx#sheet1 as s1,
    ../dir2/file2.csv s2,
    left outer join ./file3.csv s3 on s3.f1 = s2.f1 and  s3 = 4
window 
into

Where
    f1 > 4
    and (s1.f3 > '2019-01-02' or s2.f4 <> 4)
    and f4 is not null
    or s1.f1 in (1,2,3,4,5)
group by

sort by
    field1,
    field2 asc,
    field3 desc
having
    f1 = true
    and f2 <> false

union
