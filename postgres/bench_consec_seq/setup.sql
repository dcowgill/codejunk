-- Named sequences of consecutive numbers.
create table gapless_sequence (
    name varchar(20) primary key,
    last bigint not null);

-- Example definition of a table that needs consecutive numbers.
-- Junk columns = coarse approximation of a "fat" table.
create table my_thing (
    id serial primary key,
    seq_name varchar(20) not null, -- for easier benchmarking of parallel sequences
    seq_num bigint not null,       -- gap-free consecutive ints (from gapless_sequence.last)
    junk1 varchar(100),
    junk2 varchar(100),
    junk3 varchar(100),
    junk4 varchar(100),
    unique (seq_name, seq_num));

-- Sets seq_num on insert.
create or replace function my_thing_set_seq_num() returns trigger as $$
declare
    num gapless_sequence.last%type;
begin
    -- N.B. Ordinarily, new.seq_name be a string literal (e.g. where name =
    -- 'my_thing'), but we need to support a distinct sequence per row in order
    -- to facilitate the benchmarking multiple sequences in parallel.
    update gapless_sequence set last = last + 1 where name = new.seq_name returning last into new.seq_num;
    return new;
end;
$$ language plpgsql;

-- Call the function before every insert.
create trigger my_thing_before_insert before insert on my_thing
    for each row execute procedure my_thing_set_seq_num();
