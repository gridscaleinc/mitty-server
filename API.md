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
language: string       (M)      -- 言語情報　(Ja_JP, en_US, en_GB) elastic　searchに使用する。
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
 category: Enum("thisYear"/"nextYear"/"specific")
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
 id: int,
 eventId: int,
 eventTitle: String,
 memo: String
 startDateTime: dateTime,
 endDateTime: dateTime,
 notification: boolean,




}
```
*Description*
```
JSON形式のパラメータを読み込み、、、、
```

### 10.[Register New Activity](id:activity-register)
```
POST /api/activity/register

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
JSON形式のパラメータを読み込み、、、、
```
