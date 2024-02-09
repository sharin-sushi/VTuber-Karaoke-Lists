# V-kara (VTbuer-Karaoke-Lists)

※README作成中※

`「推し」の「歌枠」の聴きたい「歌」` </br>
`「ささっと把握」、「さくっと再生」、「ばばっと布教」`

## 作成した目的

　youtubeでは配信者がカラオケをする「歌枠」という動画がある。

しかし、`何の歌`が`どの動画`にあって、`何分何秒目`なのか分からない。
  
ファンは歌枠情報を登録し、リスト化。  
「ちょっとあの時の歌を聴き直したい」を叶える。  

Vtuber本人は自身が歌ったこと歌を把握し、活動に役立てられる。  

    サイトリンク：完成後に掲載予定

## 技術

### 使用技術
  
| 言語、サービス  | フレームワーク/ライブラリ |
| --- | --- |
| Next.js v13.5.6 | TypeScript v5.1.6 <br> React v18.2.0 <br> react-hook-form v.7.47.0 <br/> react-select 5.7.7 <br> react-table 7.8.0 <Br> react-youtube 10.1.0 <br> |
| Go v1.18  | GORM v1.25.4 <br> GIN v1.9.1 <br> go-sqlmock v1.5.1|
|MySQL v8.0.32| - |

その他：AWS, Docker, Github,GitHub Acitons(CI), Postman, Figma, draw.io
アーキテクチャ：Go:クリーンアーキテクチャ、Next.js:[reactの流儀](https://ja.react.dev/learn/thinking-in-react)を基本とし、機能ごとにフォルダ分けすることで既存コードを探し出しやすい構成にした。

### 技術選定理由

※前提：完全未経験の独学

- Go
  - 比較対象：JAVA、Ruby、Python等
  - コンパイラ言語であり実行速度が高速である
  - 静的型付けであり、コンパイル前にバグを発見しやすい
  - 静的型付けかつ記述自由度が低いことから、以下２点を利点と考えた
    - 開発を中長期まで続けた際にも、加筆・改修しやすい
    - 他人のコードを読んだ際に学びやすい
  - 比較的後発ながらメルカリ、Docker、twitch等の使用実績があり、求人件数が多く今後の需要を見込んだ

- Next.js
  - 比較対象：Vue.js等
  - より普及しているReactベースのフレームワークである
  - SPA、特にSSRを簡単に実装できる
  - 基本的にTypeScriptと共に採用されるため静的型付け言語であるGoと平衡で学びやすいと考えた

## 設計図

  ※初期案であり、変更が多いです。デプロイ時に更新予定です。
| ER図 | 画面遷移図 main  | 画面遷移図 header/footer |
| :---: | :---: | :---: |
| [原寸画像](https://user-images.githubusercontent.com/127638412/273976430-29856108-a613-493e-b024-bb9ad7ac88d9.png) | [原寸画像](https://user-images.githubusercontent.com/127638412/273976533-1d5db155-c5a6-403c-95db-f05b6fefb3f0.png) | [原寸画像](https://user-images.githubusercontent.com/127638412/273976737-17eb88c9-dfb0-4a99-a5ba-bf909c268ada.png) |
| ![80%](https://github.com/sharin-sushi/0016go_next_relation/assets/127638412/b3a57a97-41e3-42e4-882c-177a2e317127) | ![273976533-1d5db155-c5a6-403c-95db-f05b6fefb3f0](https://github.com/sharin-sushi/0016go_next_relation/assets/127638412/b4eda633-a6b4-4d30-b446-77a0062ad79a) | ![273976737-17eb88c9-dfb0-4a99-a5ba-bf909c268ada](https://github.com/sharin-sushi/0016go_next_relation/assets/127638412/71a4338f-84d3-4e7f-ab08-c5af83207308)|

## 機能
  - ログイン（ゲストログイン機能有り、JWT使用）、ログアウト、退会の機能
  - 登録データを表やドロップダウンで閲覧できる
  - 表について(react-table)
    - 項目クリックでページ内動画再生 (react-youtube)
    - 項目クリックで動画へのリンクをコピー
    - 表にはページネーションやソート機能
    - 最近登録された50曲の表をtopページに配置
  - いいね機能
  - バリデーション(react-hook-form、goでも若干)
    - 会員登録、ログイン時：メアド、パスワード
    - データ登録時：Vtuber名、kana、紹介動画URL、歌枠動画タイトル、歌枠動画URL、曲名
  - DB登録、編集、削除機能：Vtuber、動画(youtubeの歌枠)、歌の再生開始時間等(会員専用)
　  - 登録, 編集時はプルダウンを使用し、楽に入力できるように(react-select)
    - マイページにて、自分の登録した情報の一覧を確認できる
  - DB流出時の被害減少
    - パスワード：bcryptでハッシュ化
    - メールアドレス：AESで暗号化
  - レスポンシブ対応(PC推奨) 

## 機能詳細

- 機能要件
  - ログイン機能（ゲストログイン機能有り）、ログアウト機能
  - 退会機能
  - DB登録：Vtuber、動画(youtubeの歌枠)、歌の再生開始時間等の登録/編集/削除(会員専用)
  - 登録情報をを表で閲覧
  - いいね機能(会員専用)
  - バリデーション
    - 会員登録、ログイン：メアド、パスワード
    - データ登録：Vtuber名、kana、紹介動画URL、歌枠動画タイトル、歌枠動画URL、曲名
  - レスポンジシブ
  - ブラウザ対応確認：Chrome, Opera,

- 非機能要件
  - N+1問題対策：レコードが増えても発行されるSQLが増えないように設計
  - SQLインジェクション対策：ORMを導入
  - 不正なログイン対策：JWTを使用
  - DB流出時の被害減少：
    - パスワードをbcryptで暗号化
    - メールアドレスをAESで暗号化し使用時は復号
  - クリーンアーキテクチャを採用し、拡張と修正をしやすく
  - https化 

--ver. 1.5--(知人にデバッグを兼ねて使用依頼、フィードバックを貰う)

- 機能要件
  - 削除、申請依頼

--ver. 2(Xで一般公開)--

- 機能要件
  - パスワード変更機能

- 非機能要件
  - 配信者に特別：厳正な会員登録(twitterと紐づけ？)、自身が関与するtableに対しては全権保持
  - ログイン認証をセッション管理に
  - 名前ﾖﾐｶﾞﾅ、ﾊﾟｽﾜｰﾄﾞ登録時に全角を半角に自動変換

---

## マニュアル

 自分の好きな配信者とその歌情報（動画と歌い出しの時間等）を登録できます。
 歌情報については、登録情報の詳細ページで直接動画を再生することができます。  

### 会員機能
　- 歌情報登録
　- 自分が登録した歌情報の編集

※現状はゲストログイン機能を使用できます。

### 非会員/会員共通機能
- 歌情報の閲覧とページ内視聴

### 主な用途
- 自分で登録した歌を見返すこと
- 布教
- 歌った歌の把握
 
---

## 今後実装予定の機能

(一般公開しながらのバージョンアップ)

--ver. 3--

- 機能要件
  - オリ曲、カバー曲のtable追加
  - 類似サービス追加："鳴き声"、"名言/名言"、"お話"の
  - お知らせ機能１：ログインした際に通知が届く。通知：フォロー関連の新規登録情報等
  - 会員登録時に本人確認メール
  - 配信者情報に項目追加：X, Fanbox, Booth, 画像
  - 曲登録時に±30s以内のものがある場合に警告する
  - url登録時の自動変換：youtubeでは1動画に複数種類のURLが有るが、`watch`のものしか汎用性が低い
  ユーザーがどれを登録しても自動変換して``watch`へ統一
  これにより、重複回避や、埋め込み再生時のエラー回避

- 非機能要件
  - 死活監視
  - 定期的な自動バックアップ
  - ロック機能？
  - メンテナンス中はサービス停止とする

  
--ver4--

- X投稿機能
- お知らせ機能２：新規動画をメールで？初期は１通送って、通知onを促す
- 開発者のXs表示(機能追加、デバッグ等のお知らせ)
- youtube apiを利用し、url入力で動画タイトル取得

--verX--

- ブラックリスト登録(IP BAN?)
