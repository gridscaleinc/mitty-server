drop table footmark cascade;
create table footmark (
	id		serial8	PRIMARY KEY,
	event_id		int8	,
	island_id		int8	,
	mitty_id		int	,
	name_card_id		int8	,
	picture_id		int8	,
	seat_or_room_info		varchar	(100),
	checkin_time		timestamp	
);