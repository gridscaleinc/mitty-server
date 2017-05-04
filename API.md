# Table of contents
[Common Rules](#common-rules)
1. [Sign up](#1sign-up)
2. [Sign in](#2sign-in)
3. [Add content to gallery](#3add-content-to-gallery)
4. [Add content for event logo](#4add-content-for-event-logo)
5. [Add content for island logo](#5add-content-for-island-logo)
6. [Add content for profile icon](#6add-content-for-profile-icon)
7. [Event Searching](#7event-searching)
8. [Register New Event](#8register-new-event)
9. [Activity List](#9activity-list)
10. [Register New Activity](#10register-new-activity)
11. [Register New Activity Item](#11register-new-activity-item)
12. [Register Island](#12register-new-island)
12. [Activity Details](#13activity-details)

### [Common Rules](id:common-rules)
*表記*
```
 o:  Optional (任意)
 M:  Mandatory (必須)
```
*時間について*
```
APIでの日付と時間については常に国際標準時刻（協定世界時(UTC))として値を設定し、文字列の表記はISO8601拡張形式で表記する。
フォーマット：　　　yyyy-mm-ddThh:mm:ss
例： 　　　　　　　　　　　　　　2017-04-21T12:00:09
```
*Request*
```
X-Mitty-Access-Token: String   (O)        Access Token for Authentication
```
*Description*
```
APIによって、認証済みでないと呼び出しできない。　そういうAPIについては、該当項目は必須。
```

*Response*
```
.HTTP Status
```
```
.JSON Response
{
   result: {
      status: OK/NG,
      errorCode: String,
      errorMessage: String
   },
   xxxx: {xxx}
}

```
### 1.[Sign up](id:sign-up)
```
POST http://dev.mitty.co/api/signup
parameter:
  username: string　　（必須）　
  password: string         (必須）
  mail_address: string 　（任意？）
```

### 2.[Sign in](id:sign-in)
```
POST http://dev.mitty.co/api/signin
parameter:
  username: string
  password: string
```

### 3.[Add content to gallery](id:content-gallery)
```
POST /api/gallery/content
Header:
   access-token: String
   content-type: json/application
```
*Input parameter*
```
{
  gallery:                     (O)    -- Gallery オブジェクト自体はOptional
　　{
　　　　id:  　　　Int,          (O)     -- Gallery ID, 未設定の場合新規採番。
　　　　seq: 　　　Int,          (O)     -- 通番 未設定の場合同一ギャラリーのマックス通番＋１
　　　　caption:  string        (O)     -- ギャラリーのタイトル
　　　　briefInfo:string        (O)     -- 概要
　　　　freeText: string        (O)     -- 詳細説明
  eventId: int8           (O)     -- 設定されば場合は、Eventsテーブルに該当イベントのGallery IDを更新。
　　　　islandId: int8          (O)     -- 設定されば場合は、Islandテーブルに該当島のGallery IDを更新。
　　},
  content:　　　　　　　　　　　　（M)     -- contentオブジェクトは必須。
　{
　　　　mime      : string      (O)     -- image/gif などMIME    
　　　　name      : string      (O)     -- コンテンツの名称
　　　　link_url  : string      (M/O)   -- dataがあればNULL
　　　　data      : byte array  (M/O)   -- link_urlがあればNULL
  }
}
```
*Output response*
```
{
  result: {status: OK/NG, message:String},
  gallery: {id: Int, seq: Int, contentId: Int}
}
```
*Description*
```
考え方：
Galleryとは、イベントや島のアピールためのコンテンツの集まり、画像や、ビデオ、Youtubeなど外部リンクが考えられる。
Galleryにはコンテンツのアピール文言を格納し,コンテンツそのものはContentsテーブルに格納。
ContentsテーブルはLogoや、アイコン、写真、ビデオなどのコンテンツを一元に管理するテーブル、中の一部はGalleryに属する。

処理概要:
 ■ コンテンツ
　　　JSON形式のパラメータを読み込み、Contentsテーブルに格納し、Contents Idを採番。
　　　１）バイナリデータの場合、S3に格納し、そのS3のURLを含め情報をContentsテーブルに格納。
　　　２）外部リンクの場合、そのままcontentsテーブルに格納する。
　■　ギャラリー
　格納したContents のIDを含め、パラメータから読み込んだGallery情報とともに　Galleryテーブルに　Insertする。
　なお、Galleryにとって２件目以降のコンテンツは　Gallery IDをparameterに含めるが、全くの新規の場合はGallery Id
  がパラメータに含めず、サーバー側にて採番とする。
　■　Events / Islandの更新
　　　　Gallery　IDが新規採番された場合、Events/Islandのいずれに帰属しなければならない、EventsもしくはIslandの該当Gallery ID
  のアプデートとする。

```

### 4.[Add content for event logo](id:content-event-logo)
```
POST /api/event/logo
```
*Input parameter*
```
{}

```
*Output response*
```
{}
```
*Description*
```
JSON形式のパラメータを読み込み
```

### 5.[Add content for island logo](id:content-island-logo)
```
POST /api/island/logo
```

*Input parameter*

```
{}

```
*Output response*
```
{}
```
*Description*
```
JSON形式のパラメータを読み込み
```
### 6.[Add content for profile icon](id:content-profile-icon)
```
POST /api/profile/icon

```
*Input parameter*
```
{}

```
*Output response*
```
{}
```
*Description*
```
JSON形式のパラメータを読み込み
```

### 7.[Event Searching](id:event-search)
```
GET /api/event/list

```
*Input parameter*
```
{
 category: Enum("recommendation"/"latest"/"topRate"),
 key:   String   
}

```
*Output response*
```
{
  count: int
  [
    event1, event2,....
  ]
}

event: {
  id: int,
  title: String,
  type: enum,            (M)
  iconUrl: byte array,   (O)
  tag:                ,  (M)
  startDate: dateTime,   (M)
  endDate: dateTime,     (M)
  allDayFlag: book,      (M)
  action: string,        (M)
  image:     URL,        (O)
  island:    URL,        (O)
  placeName: string,     (O)
  address :  string,     (O)
  islandIcon: URL,       (O)
  contactMail: String,   (O)
  contactTel: string,    (O)
  infoSource: string,    (O)
  url: URL               (O)
}
```


*Description*
```
JSON形式のパラメータを読み込み、、、、
```
### 8.[Register New Event](id:event-register)
```
POST /api/new/event

```
*Input parameter*
```
{
type: string,          (M)      -- イベントの種類
tag:  string,          (M)      -- イベントについて利用者が入力したデータの分類識別。  
title: string,         (M)      -- イベントタイトル
action: string,        (M)      -- イベントの行い概要内容
startDate: dateTime,   (M)      -- イベント開始日時
endDate: dateTime,     (M)      -- イベント終了日時
allDayFlag: bool,      (M)      -- 時刻非表示フラグ。
islandId:  int,        (M)      -- 島ID
priceName1: string,    (O)      -- 価格名称１        
price1: Double ,       (O)      -- 価格額１
priceName2,            (O)      -- 価格名称2  
price2: Double ,       (O)      -- 価格額２
currency: string 　　　（O)      -- 通貨　(USD,JPY,などISO通貨３桁表記)
priceInfo: string      (O)      -- 価格について一般的な記述
description: string    (M)      -- イベントについて詳細な説明記述
contactTel: string,    (O)      -- 連絡電話番号
contactFax: string,    (O)      -- 連絡FAX
contactMail: string,   (O)      -- 連絡メール
officialUrl: URL,      (O)      -- イベント公式ページURL
organizer: string,     (O)      -- 主催者の個人や団体の名称
sourceName: string,    (M)      -- 情報源の名称
sourceUrl: URL,        (O)      -- 情報源のWebPageのURL
anticipation: string,  (O)      -- イベント参加方式、 OPEN：　自由参加、　INVITATION:招待制、PRIVATE:個人用、他の人は参加不可。
accessControl: string, (O)      -- イベント情報のアクセス制御：　PUBLIC: 全公開、　PRIVATE: 非公開、 SHARED:関係者のみ
language: string ,     (M)      -- 言語情報　(Ja_JP, en_US, en_GB) elastic　searchに使用する。
relatedActivityId: int, (O)     -- 指定された場合、Activity Itemを一件自動登録する。
asMainEvent: bool       (O)　　　-- relatedActivityIdが指定された場合のみ意味ある。
                                   trueの場合は該当activityのmainEventIdを更新する。
}
```
*Output response*
```
{
  result: {},
  eventId: int
}
```
*変更履歴*
```
4/26:  Price1 〜　PriceInfoまでの項目にデータ型を漏れてったので、追記しました。
```
*Description*
```
JSON形式のパラメータを読み込み、eventsテーブルに登録する。
イベントidを採番しレスポンスに返す。
エラーの場合、エラーコードとエラーメッセージを返す。
設定する項目は基本eventsテーブルの同名項目, parameterに指定ない項目は基本null設定するが、時間とIDについては適宜設定。
```
*Elasticについて*
```
elastic searchの登録対象項目は下記通り。
title,
action,
description,
priceInfo,
organizer,
sourceName
```
### 9.[Activity List](id:activity-list)
```
GET /api/activity/list

```
*Input parameter*
```
{
 key:   String   
}

```
*Output response*
```
{
  count: int
  [
    activity1, activity2,....
  ]
}

activity: {
 id: int,                  -- Activity のID
 eventId: int,             -- ActivityのMainEventId
 title: String,            -- ActivityのTitle
 startDateTime,            -- MainEventのstart_datetime
 eventLogoUrl              -- MainEventのLogoIDから結びつけるContentsのLinkURL
}
```
*Description*
```
活動一覧画面に表示するための検索。

SQL：
select 
         a.id,
         a.main_event_id, 
         a.title,
         e.start_datetime,
         c.link_url as event_logo_url
from activity as a 
        left outer join events as e on a.main_event_id=e.id 
        left outer join contents as c on e.logo_id=c.id
where
        owner_id=[user_id]                         (Login中のUSERID,どう取得する？）
        title like '%key%' or memo like '%key%'    (KeyがNULLの場合該当条件なし)
        
```

### 10.[Register New Activity](id:activity-register)
```
POST /api/new/activity

```
*Input parameter*
```
{
    title: string,            -- (M) タイトル
    mainEventId:Int,          -- (O) メインイベント ID   
    memo: string              -- (O)      
}

```
*Output response*
```
{ 
  result: {
    activityId: int
  }
  
  }
```
*Description*
```
活動とは一人の個人での一連の活動アイテムのまとまり。活動はいずれ一つのイベントが裏にあるが、
イベントより先に活動を登録する際は、eventIdはまだ確定していない。
このAPIは幾つの項目を登録することで、新規活動IDを採番し、結果として返す。
```

*See also*
```
  activity.sql
```

### 11.[Register New Activity Item](id:activity-register)
```
POST /api/new/activity/item

```
*Input parameter*
```
{
     activityId : int8 ,         -- (M) アクティビティ ID
     eventId: int8	,            -- (M) イベントID。
     title: string ,             -- (M) varchar(200),
     memo		 ,                  -- (O) memo
     notification: bool ,        -- (M) アラーム要否  
     notificationDateTime: timestamp,  -- (O) アラーム時刻
     asMainEvent: bool           -- (M) trueの場合、activity.mainEventId=eventIdで更新。
}

```
*Output response*
```
{
なし
}
```
*Description*
```
活動アイテムの一つを登録。
活動アイテムとはある個人が特定のイベントに参加することについての情報を管理。
```

*See also*
```
  activity.sql
  activityItem.sql
```


### 12.[Register New Island](id:island-register)
```
POST /api/new/island

```
*Input parameter*
```
{
  nickname           : string      --(O)  愛称
  name               : string      --(M)  名称
  logoId             : int         --(O)  LogoのContent Id
  description        : string      --(O)  説明
  category           : string      --(M)  カテゴリ
  mobilityType       : string      --(M)  移動性分類
  realityType        : string      --(M)  実在性分類
  ownershipType      : string      --(M)  所有者分類
  ownerName          : string      --(O)  所有者名
  ownerId            : int         --(O)  所有者のMitty User Id
  creatorId          : int         --(O)  作成者のMitty User Id 
  meetingId          : int         --(O)  会議Id
  galleryId          : int         --(O)  ギャラリーID
  tel                : string      --(O)  電話番号
  fax                : string      --(O)  FAX
  mailaddress        : string      --(O)  メールアドレス
  webpage            : string      --(O) 　WebページのURL
  likes              : string      --(O)  いいねの数
  countryCode        : string      --(O)  国コード
  countryName        : string      --(O)  国名称
  state              : string      --(O)  都道府県　
  city               : string      --(O)  市、区
  postcode           : string      --(O)  郵便番号
  thoroghfare        : string      --(O)  大通り
  subthroghfare      : string      --(O)  通り
  buildingName       : string      --(O)  建物名称
  floorNumber        : string      --(O)  フロー番号
  roomNumber         : string      --(O)  部屋番号
  address1           : string      --(O)  住所行１
  address2           : string      --(O)  住所行２
  address3           : string      --(O)  住所行３
  latitude           : double      --(O)  地理位置の緯度
  longitude          : double      --(O)  地理位置の経度
}

```
*Output response*
```
{
  result: {
    islandId: int
  }
  
}
```
*Description*
```
島とは人が集まる場所。従来的な特定な住所にある組織が入居する建物がメインだが、飛行機、タクシーなど移動体も島として登録する場合がある。また仮想的な集会場、ライブ会場なども考えられる。ゲームの世界になると、空想的なUFOなども視野にある。
こいった情報を登録するのがこの API.
```

*See also*
```
  island.sql
```

### 13.[Activity Details](id:activity-details)
```
GET /api/activity/details

```
*Input parameter*
```
{
 id:   String   
}

```
*Output response*
```
{
  activity: {
     id:
     main_event_id
     title
     memo
  }
  details:
  [
    detail, detail,....
  ]
}

detail: {
 eventId: int,             -- ActivityItemEventId
 title: String,            -- ActivityItemのTitle
 memo: String              -- ActivityItemのMemo
 notification:Bool         
 notificationTime:   date
 eventTitle:String         -- EventsのTitle
 startDateTime:Date,       -- Eventのstart_datetime
 endDateTime:Date,
 allDayFlag:Bool,          
 eventLogoUrl:String       -- EventのLogoIDから結びつけるContentsのLinkURL
}
```
*Description*
```
活動詳細のための検索。

SQL：
select 
   a.id,
   a.title,
   a.memo,
   a.main_Event_Id,
   i.event_Id,
   i.memo,
   i.notification,
   notificationdatetime,
   e.title,
   e.start_Datetime,
   e.end_Datetime,
   e.allDay_Flag,
   c.link_url as eventLogoUrl
from
   activity as a 
   left join activity_item as i on a.id=i.activity_id
   inner join events as e on i.event_id=e.id
   left outer join contents as c on e.logo_id=c.id
where 
   a.id=[id]
   and 
   a.owner_id=[loginUserId]
        
```

