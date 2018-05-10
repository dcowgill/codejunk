-- Named sequences of consecutive numbers.
create table consec_seq (
    name varchar(20) primary key,
    last bigint not null);

-- Example definition of a table that needs consecutive numbers.
-- Junk columns = coarse approximation of a "fat" table.
create table consec_thing (
    id serial primary key,
    seq_name varchar(20) not null,  -- name of sequence to use (consec_seq.name)
    seq_num bigint not null,        -- gap-free consecutive ids (from consec_seq.last)
    junk1 varchar(100),
    junk2 varchar(100),
    junk3 varchar(100),
    junk4 varchar(100),
    unique (seq_name, seq_num));

-- Sets seq_num on insert.
create or replace function consec_thing_set_seq_num() returns trigger as $$
declare
    num consec_seq.last%type;
begin
    update consec_seq set last = last+1 where name = new.seq_name returning last into new.seq_num;
    return new;
end;
$$ language plpgsql;

-- Call the function before every insert.
create trigger consec_thing_before_insert before insert on consec_thing
    for each row execute procedure consec_thing_set_seq_num();
