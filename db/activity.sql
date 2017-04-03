drop table activity cascade;

create table activity (
        id	        	serial8	,
        seq			int2	,
        eventId			int8	,
        subscribtionType	varchar	(20),
        notification		boolean	,
        restaurantSuggestOk	boolean	,
        hotelSuggestOk		boolean	,
        movingSuggestOk		boolean	,
        url			varchar	(100),
        memo			text	,
        PRIMARY KEY(id,seq)
);
