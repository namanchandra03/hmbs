BEGIN ;
ALTER TABLE hostel_rooms_details
    ADD COLUMN IF NOT EXISTS hostel_id UUID NOT NULL default gen_random_uuid(),
    ADD FOREIGN KEY (hostel_id) REFERENCES hostels (id);

ALTER TABLE hostel_rooms_details
    ALTER COLUMN hostel_id DROP DEFAULT;

COMMIT;