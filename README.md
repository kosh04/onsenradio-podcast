Podcast for onsen
-----------------

インターネットラジオステーション＜音泉＞のポッドキャストを生成するスクリプト

http://www.onsen.ag/

## Usage

    $ ./onsenradio-podcast [-dump] [-p PORT]
      -dump
        	Dump RSS feeds
      -p int
        	Listen port (default 8000)

## ポッドキャストを登録する

    $ ./onsenradio-podcast

このプログラムをサーバモードで起動した後
http://localhost:8000/podcast をポッドキャストURLとして登録してください。

## TODO

- [ ] 更新日時順ソート
- [ ] podcast.Item.Description の定義
- [ ] バナー画像は表示できる？
- [ ] onsen.ag への負荷対策（キャッシュ等）

## WARNING

> ※＜音泉＞アプリリニューアル後は、スマホからの視聴はアプリでのみとなります。

上記の注意書きがあるように、生成元サイトの仕様変更によってコードが
利用できなくなる可能性があります。ご利用は自己責任でお願いします。
