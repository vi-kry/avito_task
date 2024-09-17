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

