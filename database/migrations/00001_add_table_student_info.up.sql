CREATE TABLE IF NOT EXISTS courses (
    id UUID primary key DEFAULT gen_random_uuid(),
    course_name text not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    archived_at TIMESTAMP WITH TIME ZONE
                             );

CREATE TABLE IF NOT EXISTS program (
    id UUID primary key DEFAULT gen_random_uuid(),
    program_name text not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    archived_at TIMESTAMP WITH TIME ZONE
                             );

CREATE TABLE IF NOT EXISTS complains_type (
    id UUID primary key DEFAULT gen_random_uuid(),
    name text not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    archived_at TIMESTAMP WITH TIME ZONE
                             );

CREATE TABLE IF NOT EXISTS hostels (
    id UUID primary key DEFAULT gen_random_uuid(),
    name text NOT NULL,
    total_rooms numeric NOT NULL CHECK ( total_rooms > 0 ),
    rooms_available numeric NOT NULL CHECK (rooms_available >=0 ),
    semester_fees    numeric         NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE ,
    archived_at TIMESTAMP WITH TIME ZONE

                             );

CREATE TABLE IF NOT EXISTS uploads (
    id UUID primary key DEFAULT gen_random_uuid(),
    url text NOT NULL,
    name text NOT NULL ,
    file_path text NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE,
    archived_at TIMESTAMP WITH TIME ZONE
                             );

CREATE TABLE IF NOT EXISTS  student_details (

    id UUID primary key DEFAULT gen_random_uuid(),
    first_name text not null,
    second_name text ,
    student_image UUID REFERENCES uploads(id) NOT NULL,
    primary_mobile_no numeric not null,
    secondary_mobile_no numeric ,
    email varchar not null,
    roll_no varchar not null,
    dob varchar not null,
    hostel_id UUID REFERENCES hostels(id),
    program_name UUID REFERENCES program(id) NOT NULL,
    course_name  UUID REFERENCES courses(id) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE ,
    archived_at TIMESTAMP WITH TIME ZONE
                             );

CREATE TABLE IF NOT EXISTS hostel_rooms_details (
    id UUID primary key DEFAULT gen_random_uuid(),
    student_id UUID REFERENCES student_details(id) NOT NULL,
    floor_no int NOT NULL,
    room_no int NOT NULL,
    is_available bool DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE ,
    archived_at TIMESTAMP WITH TIME ZONE

                             );



CREATE TABLE IF NOT EXISTS complains (
    id UUID primary key DEFAULT gen_random_uuid(),
    student_id  UUID REFERENCES student_details(id) NOT NULL,
    hostel_id   UUID REFERENCES hostels(id) NOT NULL,
    complain_type UUID REFERENCES complains_type(id) NOT NULL,
    description TEXT  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE,
    archived_at TIMESTAMP WITH TIME ZONE
                             );



CREATE TABLE IF NOT EXISTS notices (
    id UUID primary key DEFAULT gen_random_uuid(),
    upload_id UUID REFERENCES uploads(id) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE,
    archived_at TIMESTAMP WITH TIME ZONE
                             );
CREATE TABLE IF NOT EXISTS staff_members (
    id UUID primary key DEFAULT gen_random_uuid(),
    first_name text not null,
    second_name text ,
    primary_mobile_no numeric not null,
    secondary_mobile_no numeric ,
    dob varchar  NOT NULL ,
    address text NOT NULL,
    hostel_id UUID REFERENCES hostels(id),
    designation varchar NOT NULL,
    salary numeric NOT NULL CHECK ( salary > 0 ),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE,
    archived_at TIMESTAMP WITH TIME ZONE

                             );




