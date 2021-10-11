create function trigger_set_timestamp()
    returns trigger as
$$
begin
    new.update_at = now();
    return new;
end;
$$ language plpgsql;

create trigger set_timestamp
    before update on "user"
    for each row
execute procedure trigger_set_timestamp();