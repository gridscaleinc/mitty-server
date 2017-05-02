
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table activity_item (
  activity_id            int8,
  event_id	             int8,
  title	                 varchar(200),
  memo		               text,
  notification           bool,
  notificationDateTime   date,
  created                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated	               timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(activity_id, event_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table activity_item cascade;
