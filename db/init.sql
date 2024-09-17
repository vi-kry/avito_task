create table public.tender (
   id uuid primary key not null default uuid_generate_v4(),
   name character varying(100) not null,
   description character varying(250) not null,
   service_type character varying(100) not null,
   status character varying(100) not null,
   organization_id uuid,
   user_id uuid,
   created_at timestamp without time zone default CURRENT_TIMESTAMP,
   updated_at timestamp without time zone default CURRENT_TIMESTAMP,
   foreign key (organization_id) references public.organization (id)
       match simple on update no action on delete cascade,
   foreign key (user_id) references public.employee (id)
       match simple on update no action on delete cascade
);

create table public.bid (
    id uuid primary key not null default uuid_generate_v4(),
    name character varying(100) not null,
    description character varying(250) not null,
    status character varying(100) not null,
    tender_id uuid,
    organization_id uuid,
    user_id uuid,
    created_at timestamp without time zone default CURRENT_TIMESTAMP,
    updated_at timestamp without time zone default CURRENT_TIMESTAMP,
    foreign key (organization_id) references public.organization (id)
        match simple on update no action on delete cascade,
    foreign key (tender_id) references public.tender (id)
        match simple on update no action on delete cascade,
    foreign key (user_id) references public.employee (id)
        match simple on update no action on delete cascade
);

create table public.employee (
     id uuid primary key not null default uuid_generate_v4(),
     username character varying(50) not null,
     first_name character varying(50),
     last_name character varying(50),
     created_at timestamp without time zone default CURRENT_TIMESTAMP,
     updated_at timestamp without time zone default CURRENT_TIMESTAMP
);
create unique index employee_username_key on employee using btree (username);

create table public.organization (
     id uuid primary key not null default uuid_generate_v4(),
     name character varying(100) not null,
     description text,
     type organization_type,
     created_at timestamp without time zone default CURRENT_TIMESTAMP,
     updated_at timestamp without time zone default CURRENT_TIMESTAMP
);

create table public.organization_responsible (
     id uuid primary key not null default uuid_generate_v4(),
     organization_id uuid,
     user_id uuid,
     foreign key (organization_id) references public.organization (id)
         match simple on update no action on delete cascade,
     foreign key (user_id) references public.employee (id)
         match simple on update no action on delete cascade
);

